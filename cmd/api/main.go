// cmd/api/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardohass/kids-api/internal/auth"
	"github.com/eduardohass/kids-api/internal/config"
	"github.com/eduardohass/kids-api/internal/handlers"
	"github.com/eduardohass/kids-api/internal/repository"
	"github.com/eduardohass/kids-api/internal/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Carregar configurações
	cfg := config.Load()

	// Conectar ao banco de dados
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Configurar repositórios
	childRepo := repository.NewChildRepository(db)
	caretakerRepo := repository.NewCaretakerRepository(db)
	volunteerRepo := repository.NewVolunteerRepository(db)
	groupRepo := repository.NewGroupRepository(db)
	needRepo := repository.NewNeedRepository(db)
	allergyRepo := repository.NewAllergyRepository(db)

	// Configurar serviços
	childService := services.NewChildService(childRepo, needRepo, allergyRepo)
	caretakerService := services.NewCaretakerService(caretakerRepo)
	volunteerService := services.NewVolunteerService(volunteerRepo)
	groupService := services.NewGroupService(groupRepo)

	// Configurar autenticação
	authenticator := auth.NewAuthenticator(cfg.Auth0Domain, cfg.Auth0Audience)

	// Configurar router
	router := handlers.NewRouter(
		authenticator,
		childService,
		caretakerService,
		volunteerService,
		groupService,
	)

	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Iniciar servidor em uma goroutine
	go func() {
		log.Printf("Server starting on port %d", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Configurar graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Criar um deadline para o shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Shutdown do servidor
	log.Println("Server shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
