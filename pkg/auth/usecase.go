package auth

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	ch "solar-service/libs/cache"
	"solar-service/models"
	"solar-service/pkg/auth/repository"
	"time"
)

type authUsecase struct {
	authRepo  repository.AuthRepo
	config 	 *models.Config
	client 	  http.Client
	cache     ch.Cache
}

// ClantUsecase ...
type AuthUsecase interface {
	Login() error
}

// NewCategoriesUsecase ...
func NewAuthUsecase(authRepo repository.AuthRepo, config *models.Config, client http.Client, cache ch.Cache) AuthUsecase {
	return &authUsecase{
		authRepo,
		config,
		client,
		cache,
	}
}

func (u *authUsecase) Login() error {
	var loginResp models.LoginResponse

	payload, err := json.Marshal(models.LoginRequest{
		Email: u.config.Auth.Email,
		Password: u.config.Auth.Password,
	})
	if err != nil {
		log.Println("[err] [auth] [Login] [json.Marshal] ", err)
		return err
	}

	req, err := http.NewRequest("POST", u.config.Auth.BaseUrl+"/login", bytes.NewBuffer(payload))
	if err != nil {
		log.Println("[err] [auth] [Login] [http.NewRequest] ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	result, err := u.client.Do(req)
	if err != nil {
		log.Println("[err] [auth] [Login] [client.Do] ", err)
		return err
	}

	defer result.Body.Close()

	if err = json.NewDecoder(result.Body).Decode(&loginResp); err != nil {
		log.Println("[err] [auth] [Login] [NewDecoder] ", err)
		return err
	}
	
	if err = u.putCache("key-cache", loginResp); err != nil {
		log.Println("[err] [auth] [Login] [putCache] ", err)
		return err
	}
	
	log.Println("[success] [auth] [Login] ")
	return nil
}

func (u *authUsecase) putCache(key string, data models.LoginResponse) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return err
	}

	duration := time.Duration(24 * 60 * int64(time.Minute))
	err = u.cache.Put(key, string(encoded), duration)
	if err != nil {
		return err
	}

	return nil
}