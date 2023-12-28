package main

import (
	"fmt"
	"log"
	"solar-service/controller"
	"solar-service/libs/cache"
	"solar-service/libs/config"
	"solar-service/libs/database"
	"solar-service/pkg/plants"
	"solar-service/pkg/plants/repository"
)

var (
	dbRepoConn  database.DatabaseRepo = database.NewPostgresRepo()
)

func main() {
	configModel := config.GetConfig()
	
	// Setup database
	db, err := dbRepoConn.Connect(configModel)
	if err != nil {
		log.Fatal(err)
	}

	// Setup cache
	cache, err := cache.Connection(configModel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cache)
	// Register repositories
	plantRepo := repository.NewPlantsRepo(db)

	// Register usecases
	plantsUseCase := plants.NewPlantUsecase(plantRepo)

	// Register handlers
	plantsHnd := controller.NewPlantsController(plantsUseCase)

	fmt.Println(plantsHnd)

	log.Println("Running Solar Service On Port ", configModel.Service.Port)
}
