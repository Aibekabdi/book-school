package main

import (
	"book-school/internal/http"
	"book-school/internal/models"
	"book-school/internal/repository"
	"book-school/internal/server"
	"book-school/internal/service"
	"book-school/pkg/config"
	"book-school/pkg/postgres"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Book-School Swagger
// @version 1.0
// @description help for aidar

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// preparation configs
	conf, err := config.NewConfig("./configs/config.json")
	if err != nil {
		log.Fatalf("error occured while trying to parse config: %s", err.Error())
	}

	// preparation db
	db, err := postgres.NewPostgresDB(conf.Database)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal("can't close db, err:", err)
		} else {
			log.Println("db closed")
		}
	}()

	// preparation server
	repository := repository.NewRepository(db)
	service := service.NewService(repository, conf)

	service.School.Create(context.Background(), models.School{
		ClassCount: 10,
		Name:       "private",
		Password:   "private",
	})

	service.Admin.Create(context.Background(), models.Admin{
		Username: "admin",
		Password: "admin",
	})

	handler := http.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(conf.Api.Port, handler.InitRoutes()); err != nil {
			log.Printf("error occured while running http server: %s", err.Error())
			return
		}
	}()

	// graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println()
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	if err = srv.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
