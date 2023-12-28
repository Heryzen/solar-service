package config

import (
	"fmt"
	"os"
	"solar-service/models"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConfig() *models.Config {
	_, err := os.Stat(".env")

	if !os.IsNotExist(err) {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Println("Error while reading the env file", err)
			panic(err)
		}
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	cachePort, _ := strconv.Atoi(os.Getenv("CACHE_PORT"))
	migrateRun, err := strconv.ParseBool("MIGRATION_RUN")
	migrateVersion, _ := strconv.Atoi(os.Getenv("MIGRATION_VERSION"))


	config := &models.Config{
		Service: models.ServiceConfig{
			Name: os.Getenv("SERVICE_NAME"),
			Port: 3000,
		},
		Database: models.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
			Port:     dbPort,
		},
		Migration: models.MigrationConfig{
			Run: migrateRun,
			Version: migrateVersion,
		},
		Cache: models.CacheConfig{
			Address:  os.Getenv("CACHE_ADDRESS"),
			Driver:   os.Getenv("CACHE_DRIVER"),
			Username: os.Getenv("CACHE_USERNAME"),
			Password: os.Getenv("CACHE_PASSWORD"),
			Port: 	  cachePort,
		},
		Cron: models.CronConfig{
			PlanTimer: 		 os.Getenv("CRON_PLAN_TIMER"),
			ReportDataTimer: os.Getenv("CRON_REPORT_DATA_TIMER"),
			Location: 		 os.Getenv("CRON_LOCATION"),
		},
		Auth: models.AuthConfig{
			Email: os.Getenv("AUTH_EMAIL"),
			Password: os.Getenv("AUTH_PASSWORD"),
			BaseUrl: os.Getenv("AUTH_BASE_URL"),
		},
	}

	return config
}
