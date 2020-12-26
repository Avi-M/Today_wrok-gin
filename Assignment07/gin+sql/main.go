package main

import (
	"log"

	"gin+sql/configs"
	"gin+sql/database"
	"gin+sql/models"
	"gin+sql/repositories"
)

func main() {
	// database configs
	dbUser, dbPassword, dbName := "root", "", "gincrud"

	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	// unable to connect to database
	if err != nil {
		log.Fatalln(err)
	}

	// ping to database
	err = db.DB().Ping()

	// error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	// migration
	db.AutoMigrate(&models.Contact{})

	defer db.Close()

	contactRepository := repositories.NewContactRepository(db)

	route := configs.SetupRoutes(contactRepository)

	route.Run(":8000")
}