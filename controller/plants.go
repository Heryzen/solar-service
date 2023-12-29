package controller

import (
	"os"
	"os/signal"
	"solar-service/models"
	"solar-service/pkg/plants"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

type plantsController struct {
	config 	 *models.Config
	plantsUc plants.PlantUsecase
}

type PlantsController interface {
	PlantsScheduler()
}

// NewPlantsController func
func NewPlantsController(config *models.Config, plantsUc plants.PlantUsecase) PlantsController {
	return &plantsController{
		config,
		plantsUc,
	}
}

func (c *plantsController) PlantsScheduler() {
	loadLocation, _ := time.LoadLocation(c.config.Cron.Location) 
	scheduler := cron.New(cron.WithLocation(loadLocation))
	
	defer scheduler.Stop()
	scheduler.AddFunc(c.config.Cron.PlanTimer, c.plantsUc.PlantsCron)
	
	// add other scheduler here

	go scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}