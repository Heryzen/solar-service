package plants

import (
	"bytes"
	"context"
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
	ctx := context.Background()

	data, err := u.getCache()
	if err != nil || data == (models.LoginResponse{}) {
		u.authUc.Login()
		return
	}

	//get
	result, err := u.getAllPlants(data)
	if err != nil {
		log.Println("[err] [Plants] [PlantsCron] [getAllPlants] ", err)
		return
	}

	//insert
	for _, v := range result.Plants {
		err = u.plantsRepo.Create(ctx, &v)
		if err != nil {
			log.Println("[err] [Plants] [PlantsCron] [Create] ", err)
			return
		}
	}
}

func (u *plantUsecase) getAllPlants(data models.LoginResponse) (models.PlantsResponse, error) {
	var plantsResp models.PlantsResponse

	req, err := http.NewRequest("POST", u.config.Auth.BaseUrl+"/login", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Println("[err] [Plants] [getAllPlants] [http.NewRequest] ", err)
		return plantsResp, err
	}

	req.Header.Set("Content-Type", "application/json")
	result, err := u.client.Do(req)
	if err != nil {
		log.Println("[err] [Plants] [getAllPlants] [client.Do] ", err)
		return plantsResp, err
	}

	defer result.Body.Close()

	if err = json.NewDecoder(result.Body).Decode(&plantsResp); err != nil {
		log.Println("[err] [Plants] [getAllPlants] [NewDecoder] ", err)
		return plantsResp, err
	}
	
	return plantsResp, nil
}

func (u *plantUsecase) getCache() (models.LoginResponse, error) {
	resp := models.LoginResponse{}
	data, _, err := u.cache.Get("key-cache")
	if err != nil {
		log.Println("[err] [Plants] [getCache] [Get] ", err)
		return resp, err
	}

	resp, ok := data.(models.LoginResponse)
	if !ok {
		log.Println("[err] [Plants] [getCache] [NullData] ", err)
		return resp, err
	}

	return resp, nil
}