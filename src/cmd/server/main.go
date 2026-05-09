package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"sai-server/internal/adapter/cache"
	"sai-server/internal/adapter/imagegen"
	"sai-server/internal/adapter/llm"
	"sai-server/internal/adapter/storage"
	"sai-server/internal/adapter/tts"
	"sai-server/internal/api"
	"sai-server/internal/api/handlers"
	"sai-server/internal/api/middleware"
	"sai-server/internal/app/dualcode"
	"sai-server/internal/app/quiz"
	sessionapp "sai-server/internal/app/session"
	"sai-server/internal/app/socratic"
	"sai-server/internal/app/user"
	"sai-server/internal/config"
	"sai-server/internal/domain"
	"sai-server/internal/port"
	a2uiEngine "sai-server/pkg/a2ui"
)

func main() {
	cfg := config.Load()

	logLevel := slog.LevelInfo
	if !cfg.IsProduction() {
		logLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(logger)

	slog.Info("starting sai-server", "environment", cfg.Environment, "port", cfg.ServerPort)

	var dialector gorm.Dialector
	if cfg.IsProduction() {
		dialector = postgres.Open(cfg.DatabaseURL())
	} else {
		_ = os.MkdirAll("data", 0755)
		dialector = sqlite.Open(cfg.SQLiteDSN())
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(
		&domain.User{},
		&domain.Concept{},
		&domain.UserConceptMastery{},
		&domain.QuizItem{},
		&domain.Session{},
		&domain.Capsule{},
		&domain.Interaction{},
	); err != nil {
		slog.Error("failed to run auto migrations", "error", err)
		os.Exit(1)
	}
	slog.Info("database migrated")

	var llmClient port.LLMClient
	if cfg.LLMProvider == "openai" {
		llmClient = llm.NewOpenAIClient(cfg.OpenAIKey, cfg.LLMModel)
	} else {
		llmClient = llm.NewOpenAIClient("", cfg.LLMModel)
	}
	slog.Info("llm client ready", "provider", cfg.LLMProvider, "model", cfg.LLMModel)

	ttsClient := tts.NewTTSClient(cfg.OpenAIKey, "openai")
	slog.Info("tts client ready")

	imageGenClient := imagegen.NewImageGenClient(cfg.GeminiAPIKey, cfg.OpenAIKey, cfg.GeminiImageModel)
	slog.Info("image gen client ready", "primary", "gemini")

	storageClient := storage.NewMinIOOrNoop(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey)
	slog.Info("storage client ready")

	cacheClient := cache.NewRedisOrNoop(cfg.RedisURL)
	slog.Info("cache client ready")

	_ = ttsClient
	_ = storageClient
	_ = cacheClient

	a2ui := a2uiEngine.NewEngine()
	defer a2ui.CloseAll()
	slog.Info("a2ui engine ready")

	userSvc := user.NewService(db, cfg.JWTSecret)
	dualCodeOrch := dualcode.NewOrchestrator(llmClient, imageGenClient, ttsClient)
	quizEngine := quiz.NewEngine(llmClient)
	socraticRem := socratic.NewRemediator(llmClient)
	sessionSvc := sessionapp.NewService(db, dualCodeOrch, quizEngine, socraticRem)
	sessionPipeline := sessionapp.NewPipeline(sessionSvc)

	slog.Info("services initialized")

	seedDemoData(db)
	slog.Info("demo data seeded")

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())

	h := &handlers.Handlers{
		A2UI:      handlers.NewA2UIHandler(a2ui),
		ColdStart: handlers.NewColdStartHandler(),
		Session:   handlers.NewSessionHandler(sessionSvc, sessionPipeline),
		Capsule:   handlers.NewCapsuleHandler(dualCodeOrch),
		Quiz:      handlers.NewQuizHandler(sessionPipeline),
		Socratic:  handlers.NewSocraticHandler(sessionPipeline),
		User:      handlers.NewUserHandler(userSvc),
	}

	api.RegisterRoutes(r, h, cfg.JWTSecret)

	slog.Info("server ready", "port", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}

func seedDemoData(db *gorm.DB) {
	var userCount int64
	db.Model(&domain.User{}).Count(&userCount)
	if userCount == 0 {
		demoHash, _ := bcrypt.GenerateFromPassword([]byte("demo1234"), bcrypt.DefaultCost)
		demoUser := domain.User{
			Email:        "demo@demo.com",
			PasswordHash: string(demoHash),
			Cluster:      strPtr("intermediate"),
		}
		theta := 0.55
		demoUser.EstimatedTheta = &theta
		if err := db.Create(&demoUser).Error; err != nil {
			slog.Error("failed to seed demo user", "error", err)
		} else {
			slog.Info("demo user created", "email", demoUser.Email)
		}
	}

	var count int64
	db.Model(&domain.Concept{}).Count(&count)
	if count > 0 {
		return
	}

	concepts := []domain.Concept{
		{Name: "Método Científico", Description: strPtr("Proceso sistemático de investigación"), Difficulty: 0.3},
		{Name: "Física Cuántica", Description: strPtr("Estudio de partículas subatómicas"), Difficulty: 0.8},
		{Name: "Teoría de la Relatividad", Description: strPtr("Espacio-tiempo y gravedad"), Difficulty: 0.9},
		{Name: "Álgebra Lineal", Description: strPtr("Vectores, matrices y transformaciones"), Difficulty: 0.5},
		{Name: "Biología Celular", Description: strPtr("Estructura y función de las células"), Difficulty: 0.4},
	}

	db.Create(&concepts)

	quizItems := []domain.QuizItem{
		{ConceptID: concepts[0].ID, DifficultyIRT: 0.3, Discrimination: 1.2, Guessing: 0.25,
			Content: map[string]interface{}{
				"question": "¿Cuál es el primer paso del método científico?",
				"options":  []string{"Observación", "Hipótesis", "Experimento", "Conclusión"},
				"correct":  0,
			}},
		{ConceptID: concepts[0].ID, DifficultyIRT: 0.5, Discrimination: 1.0, Guessing: 0.25,
			Content: map[string]interface{}{
				"question": "¿Qué es una variable independiente?",
				"options":  []string{"La que se mide", "La que se manipula", "La que se controla", "La que se ignora"},
				"correct":  1,
			}},
		{ConceptID: concepts[1].ID, DifficultyIRT: 0.7, Discrimination: 1.0, Guessing: 0.25,
			Content: map[string]interface{}{
				"question": "¿Qué describe la dualidad onda-partícula?",
				"options":  []string{"Solo ondas", "Solo partículas", "Comportamiento dual", "Ninguna"},
				"correct":  2,
			}},
		{ConceptID: concepts[3].ID, DifficultyIRT: 0.5, Discrimination: 1.0, Guessing: 0.25,
			Content: map[string]interface{}{
				"question": "¿Qué es un vector propio?",
				"options":  []string{"Un escalar", "Dirección invariante", "Matriz nula", "Determinante"},
				"correct":  1,
			}},
		{ConceptID: concepts[4].ID, DifficultyIRT: 0.4, Discrimination: 1.0, Guessing: 0.25,
			Content: map[string]interface{}{
				"question": "¿Cuál es la organela responsable de la producción de energía?",
				"options":  []string{"Núcleo", "Ribosoma", "Mitocondria", "Aparato de Golgi"},
				"correct":  2,
			}},
	}

	db.Create(&quizItems)
}

func strPtr(s string) *string { return &s }
