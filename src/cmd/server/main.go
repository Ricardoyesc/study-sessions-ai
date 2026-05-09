package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"sai-server/internal/api"
	"sai-server/internal/api/handlers"
	"sai-server/internal/config"
	"sai-server/internal/domain"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	var dialector gorm.Dialector
	if cfg.IsProduction() {
		dialector = postgres.Open(cfg.DatabaseURL())
	} else {
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

	r := gin.Default()
	r.Use(gin.Recovery())

	h := &handlers.Handlers{
		A2UI:      handlers.NewA2UIHandler(),
		ColdStart: handlers.NewColdStartHandler(),
		Session:   handlers.NewSessionHandler(),
		Capsule:   handlers.NewCapsuleHandler(),
		Quiz:      handlers.NewQuizHandler(),
		Socratic:  handlers.NewSocraticHandler(),
		User:      handlers.NewUserHandler(cfg.JWTSecret),
	}

	api.RegisterRoutes(r, h, cfg.JWTSecret)

	slog.Info("starting server", "port", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
