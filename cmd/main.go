package main

import (
	"database/sql"

	"github.com/radityarestan/mini-project/controller"
	"github.com/radityarestan/mini-project/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/mini-project.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
	ctr := controller.NewController(usersRepo)
	router := controller.SetupRouter(ctr)
	router.Run(":8080")
}
