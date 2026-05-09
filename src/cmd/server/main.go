package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	slog.Info("starting sai-server",
		"environment", cfg.Environment,
		"port", cfg.ServerPort,
	)

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

	slog.Info("database migrated successfully")

	var llmClient port.LLMClient
	if cfg.LLMProvider == "openai" {
		llmClient = llm.NewOpenAIClient(cfg.OpenAIKey, cfg.LLMModel)
		slog.Info("llm client initialized", "provider", cfg.LLMProvider, "model", cfg.LLMModel)
	} else {
		llmClient = llm.NewOpenAIClient("", cfg.LLMModel)
	}

	var ttsClient port.TTSClient
	ttsClient = tts.NewTTSClient(cfg.OpenAIKey, "openai")
	slog.Info("tts client initialized", "provider", "openai")

	var imageGenClient port.ImageGenClient
	imageGenClient = imagegen.NewImageGenClient(cfg.GeminiAPIKey, cfg.OpenAIKey, cfg.GeminiImageModel)
	slog.Info("image gen client initialized", "primary", "gemini", "model", cfg.GeminiImageModel)

	var storageClient port.StorageClient
	storageClient = storage.NewMinIOOrNoop(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey)
	slog.Info("storage client initialized", "endpoint", cfg.MinIOEndpoint)

	var cacheClient port.CacheClient
	cacheClient = cache.NewRedisOrNoop(cfg.RedisURL)
	slog.Info("cache client initialized", "url", cfg.RedisURL)

	a2ui := a2uiEngine.NewEngine()
	slog.Info("a2ui engine initialized")

	defer a2ui.CloseAll()

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())

	h := &handlers.Handlers{
		A2UI:      handlers.NewA2UIHandler(a2ui),
		ColdStart: handlers.NewColdStartHandler(),
		Session:   handlers.NewSessionHandler(),
		Capsule:   handlers.NewCapsuleHandler(),
		Quiz:      handlers.NewQuizHandler(),
		Socratic:  handlers.NewSocraticHandler(),
		User:      handlers.NewUserHandler(cfg.JWTSecret),
	}

	api.RegisterRoutes(r, h, cfg.JWTSecret)

	_ = db
	_ = llmClient
	_ = ttsClient
	_ = imageGenClient
	_ = storageClient
	_ = cacheClient

	slog.Info("server ready", "port", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
