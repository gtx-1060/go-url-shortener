package app

import (
	"log"
	"url-shortener/apis"
	. "url-shortener/daos"
	"url-shortener/database/sqlite"
	"url-shortener/services"
)

func StartApp() {
	// configure db
	dao := NewDao(sqlite.Open(), sqlite.OpenNonConcurrent())
	dao.CreateTables()

	// configure router
	service := services.NewService(dao)
	router := apis.InitRouter(service)
	// start
	err := router.Run(":5361")
	if err != nil {
		log.Fatalln(err)
	}
}
