package controllers

import (
	"log"
	"net/http"

	"hcall/api/dictionaries"
	"hcall/api/services"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var request utils.RegisterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: err.Error(),
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}
	// Call the service
	_, token, err := c.authService.Register(request.Username, request.Email, request.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: err.Error(),
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Return success
	ctx.JSON(http.StatusOK, utils.AuthResponse{
		Message: dictionaries.UserRegisteredSuccess,
		Token:   token,
		Status:  true,
	})
}

// @Router /auth/enter [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request utils.LoginRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}
	// Call the service
	_, token, err := c.authService.Login(request.Email, request.Password)

	if err != nil {
		log.Printf("Error in login: %v", err)
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidCredentials,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Return success
	ctx.JSON(http.StatusOK, utils.AuthResponse{
		Message: dictionaries.UserLoggedSuccess,
		Token:   token,
		Status:  true,
	})
}

func (c *AuthController) CreateMaster(ctx *gin.Context) {
	var request utils.CreateMasterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}
	// Call the service
	_, token, err := c.authService.CreateMaster(request.Email, request.Password)

	if err != nil {
		log.Printf("Error creating master: %v", err)
		ctx.JSON(http.StatusForbidden, utils.MessageResponse{
			Message: dictionaries.MasterAlreadyExists,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Return success
	ctx.JSON(http.StatusOK, utils.AuthResponse{
		Message: dictionaries.MasterCreatedSuccess,
		Token:   token,
		Status:  true,
	})
}

func (c *AuthController) DeleteMaster(ctx *gin.Context) {
	var request utils.DeleteMasterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Call the service
	err := c.authService.DeleteMaster(request.Email, request.Password)

	if err != nil {
		log.Printf("Error deleting master: %v", err)
		statusCode := http.StatusForbidden
		message := dictionaries.InvalidMasterPassword
		if err.Error() == "master user not found" {
			statusCode = http.StatusNotFound
			message = dictionaries.MasterNotFound
		}
		ctx.JSON(statusCode, utils.MessageResponse{
			Message: message,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Return success
	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.MasterDeletedSuccess,
		Status:  true,
	})
}
