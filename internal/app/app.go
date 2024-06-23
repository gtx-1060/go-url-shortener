package app

import (
	"url-shortener/internal/database"
	"url-shortener/internal/models"
	"url-shortener/internal/transport"
)

func StartApp() {
	// configure db
	db := database.ConnectDB()
	models.CreateTables(db)
	if err := db.Close(); err != nil {
		panic(err)
	}
	
	// configure router
	router := transport.InitRouter()
	// start
	err := router.Run(":5361")
	if err != nil {
		panic(err)
	}
}
