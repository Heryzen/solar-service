package controller

import (
	"solar-service/models"
	"solar-service/pkg/auth"
)

type authController struct {
	config 	 *models.Config
	authUc    auth.AuthUsecase
}

type AuthController interface {
	Login()
}

// NewPuthController func
func NewAuthController(config *models.Config, authUc auth.AuthUsecase) AuthController {
	return &authController{
		config,
		authUc,
	}
}

func (c *authController) Login() {
	c.authUc.Login()
}