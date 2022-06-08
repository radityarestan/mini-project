package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/radityarestan/mini-project/internal/controller"
	"github.com/radityarestan/mini-project/internal/repository"
	"github.com/radityarestan/mini-project/internal/service"
	"github.com/radityarestan/mini-project/internal/shared/config"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func dotEnv(key string) string {
	err := godotenv.Load(config.EnvPath)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	db, err := sql.Open("sqlite3", dotEnv(config.DBPath))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
	service := service.NewService(usersRepo)
	ctr := controller.NewController(service)
	router := controller.SetupRouter(ctr)

	if err := router.Run(dotEnv(config.LocalHostPort)); err != nil {
		panic(err)
	}
}
