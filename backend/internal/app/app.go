package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/spyatachkov/green-api/backend/internal/api/v1"
	greenapiClient "github.com/spyatachkov/green-api/backend/internal/client/http/greenapi"
	"github.com/spyatachkov/green-api/backend/internal/config"
	greenapiService "github.com/spyatachkov/green-api/backend/internal/service/greenapi"
	"github.com/spyatachkov/green-api/backend/pkg/recovery"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type App struct {
	config     *config.Config
	httpServer *http.Server
}

func New(cfg *config.Config) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Run() error {
	defer recovery.Recover()

	client := greenapiClient.NewClient(a.config.GreenAPIBaseURL, a.config.HTTPClientTimeout)
	service := greenapiService.New(client)
	handler := v1.NewHandler(service)

	router := a.setupRouter(handler)

	a.httpServer = &http.Server{
		Addr:         ":" + a.config.ServerPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	errChan := make(chan error, 1)

	recovery.SafeGo(func() {
		log.Printf("Server starting on %s", a.httpServer.Addr)
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		return err
	case <-quit:
		log.Println("Shutting down server...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}

func (a *App) setupRouter(handler *v1.Handler) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	c := cors.New(cors.Options{
		AllowedOrigins: a.config.CORSAllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	router.Use(c.Handler)

	router.Route("/api", func(r chi.Router) {
		r.Get("/settings", handler.GetSettings)
		r.Get("/state", handler.GetStateInstance)
		r.Post("/message", handler.SendMessage)
		r.Post("/file", handler.SendFileByUrl)
	})

	return router
}
