package plants

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"solar-service/models"
	"solar-service/pkg/plants/repository"

	ch "solar-service/libs/cache"

	"solar-service/pkg/auth"

	"github.com/jmoiron/sqlx"
)

type plantUsecase struct {
	db          *sqlx.DB
	cache       ch.Cache
	client 	    http.Client
	config 	    *models.Config
	authUc    	auth.AuthUsecase 
	plantsRepo  repository.PlantsRepo
}

// ClantUsecase ...
type PlantUsecase interface {
	PlantsCron()
}

// NewCategoriesUsecase ...
func NewPlantUsecase(db *sqlx.DB, cache ch.Cache, client http.Client, config *models.Config, authUc auth.AuthUsecase, plantsRepo repository.PlantsRepo) PlantUsecase {
	return &plantUsecase{
		db,
		cache,
		client,
		config,
		authUc,
		plantsRepo,
	}
}


func (u *plantUsecase) PlantsCron()  {
	data, err := u.getCache()
	if err != nil || data == (models.LoginResponse{}) {
		u.authUc.Login()
		return
	}

	tx, err := u.db.Begin()
	if err != nil {
		log.Print("[err] [Plants] [PlantsCron] [OpenTransaction] => ", err)
	}
	defer tx.Rollback()

	//get

	//insert
}

func (u *plantUsecase) getAllPlants() (models.PlantsResponse, error) {
	var plantsResp models.PlantsResponse

	req, err := http.NewRequest("POST", u.config.Auth.BaseUrl+"/login", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Println("[err] [auth] [Login] [http.NewRequest] ", err)
		return plantsResp, err
	}

	req.Header.Set("Content-Type", "application/json")
	result, err := u.client.Do(req)
	if err != nil {
		log.Println("[err] [auth] [Login] [client.Do] ", err)
		return plantsResp, err
	}

	defer result.Body.Close()

	if err = json.NewDecoder(result.Body).Decode(&plantsResp); err != nil {
		log.Println("[err] [auth] [Login] [NewDecoder] ", err)
		return plantsResp, err
	}
	
	log.Println("[success] [auth] [Login] ")
	return plantsResp, nil
}

func (u *plantUsecase) getCache() (models.LoginResponse, error) {
	resp := models.LoginResponse{}
	_, _, err := u.cache.Get("key-cache")
	if err != nil {
		return resp, err
	}

	return resp, nil
}