package main

import (
	"github.com/nguyenbt456/todolist-go/app"
	"github.com/nguyenbt456/todolist-go/database"
)

func main() {
	db, err := database.ConnectPostgresDB()
	if err != nil {
		panic(err)
	}
	defer database.DisconnectPostgresDB(db)

	router := app.InitRouter()

	if err := router.Run(":9991"); err != nil {
		panic(err)
	}
}
