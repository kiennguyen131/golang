package controller

import (
	"ecommerce-backend-api/init/internal/service"
	"ecommerce-backend-api/init/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewNewService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// response.SuccessResponse(c, 20001, []string{"tips", "m10"})
	response.ErrorResponse(c, 20003, "No need!!")
}
