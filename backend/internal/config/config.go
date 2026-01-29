package config

import (
	"log"
	"time"

	"github.com/spyatachkov/green-api/backend/pkg/env"
	"github.com/spyatachkov/green-api/backend/pkg/strutil"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         string
	CORSAllowedOrigins []string
	GreenAPIBaseURL    string
	HTTPClientTimeout  time.Duration
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (ok in prod)")
	}

	serverPort := env.Get("SERVER_PORT", "8080")
	corsOrigins := env.Get("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	corsOriginsList := strutil.SplitAndTrim(corsOrigins, ",")
	if len(corsOriginsList) == 0 {
		corsOriginsList = []string{"http://localhost:3000"}
	}

	greenAPIBaseURL := env.Get("GREEN_API_BASE_URL", "https://api.green-api.com")

	timeoutSeconds := env.GetInt("HTTP_CLIENT_TIMEOUT_SECONDS", 30)
	httpClientTimeout := time.Duration(timeoutSeconds) * time.Second

	return &Config{
		ServerPort:         serverPort,
		CORSAllowedOrigins: corsOriginsList,
		GreenAPIBaseURL:    greenAPIBaseURL,
		HTTPClientTimeout:  httpClientTimeout,
	}, nil
}
