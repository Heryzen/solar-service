package main

import (
	"log"
	"net/http"
	"solar-service/controller"
	"solar-service/libs/cache"
	"solar-service/libs/config"
	"solar-service/libs/database"
	authUc "solar-service/pkg/auth"
	authRepo "solar-service/pkg/auth/repository"
	planUc "solar-service/pkg/plants"
	planRepo "solar-service/pkg/plants/repository"
)

var (
	httpCl = http.Client{}
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
	
	// Register repositories
	authRepo := authRepo.NewAuthRepo(db)
	plantRepo := planRepo.NewPlantsRepo(db)

	// Register usecases
	authUseCase := authUc.NewAuthUsecase(authRepo, configModel, httpCl, cache)
	plantsUseCase := planUc.NewPlantUsecase(db, cache, httpCl, configModel, authUseCase, plantRepo)

	// Register handlers
	plantsHnd := controller.NewPlantsController(configModel, plantsUseCase)
	authHnd := controller.NewAuthController(configModel, authUseCase)

	authHnd.Login()
	go plantsHnd.PlantsScheduler()

	log.Println("Running Solar Service On Port ", configModel.Service.Port)
}
