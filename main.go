package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testtask/handlers"
	"testtask/storage"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Handler *handlers.RequestHandler
}

func NewService() *Service {
	reqStorage := storage.NewRequest()
	reqHandler := handlers.NewRequestHandler(reqStorage)

	service := &Service{
		Handler: reqHandler,
	}
	return service
}

func (s *Service) SetupRoutes(r *gin.Engine) {
	r.POST("requests", s.Handler.Create)
	r.GET("requests", s.Handler.GetRequests)
}

func main() {
	r := gin.Default()

	service := NewService()
	service.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, os.Interrupt, syscall.SIGTERM)
	<-quite

	// Set a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}

	log.Println("Server stopped")
}
