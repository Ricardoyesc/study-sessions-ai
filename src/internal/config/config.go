package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	ServerPort  string
	JWTSecret   string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	RedisURL string

	MinIOEndpoint  string
	MinIOAccessKey string
	MinIOSecretKey string
	MinIOBucket    string

	LLMProvider string
	LLMModel    string
	OpenAIKey   string

	TTSProvider  string
	GeminiAPIKey string

	ImageProvider    string
	GeminiImageModel string
}

func Load() *Config {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	cfg := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "dev-secret-change-me"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "sai_learning"),
		DBUser:     getEnv("DB_USER", "sai_user"),
		DBPassword: getEnv("DB_PASSWORD", ""),

		RedisURL: getEnv("REDIS_URL", "localhost:6379"),

		MinIOEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIOAccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MinIOSecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
		MinIOBucket:    getEnv("MINIO_BUCKET", "sai-assets"),

		LLMProvider: getEnv("LLM_PROVIDER", "openai"),
		LLMModel:    getEnv("LLM_MODEL", "gpt-4o"),
		OpenAIKey:   getEnv("OPENAI_API_KEY", ""),

		TTSProvider:  getEnv("TTS_PROVIDER", "gemini"),
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),

		ImageProvider:    getEnv("IMAGE_PROVIDER", "gemini"),
		GeminiImageModel: getEnv("GEMINI_IMAGE_MODEL", "gemini-2.0-flash-exp-image-generation"),
	}

	return cfg
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) DatabaseURL() string {
	if c.IsProduction() {
		return "host=" + c.DBHost + " port=" + c.DBPort + " user=" + c.DBUser +
			" password=" + c.DBPassword + " dbname=" + c.DBName + " sslmode=disable"
	}
	return "data/sai_dev.db"
}

func (c *Config) SQLiteDSN() string {
	return "data/sai_dev.db"
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	if viper.IsSet(key) {
		return viper.GetString(key)
	}
	return fallback
}
