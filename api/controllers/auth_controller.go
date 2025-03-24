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

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user in the system
// @Accept json
// @Produce json
// @Param body body utils.RegisterRequest true "Registration details"
// @Success 200 {object} utils.AuthResponse
// @Failure 400 {object} utils.MessageResponse
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var request utils.RegisterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Status:  false,
		})
		return
	}
	// Call the service
	_, token, err := c.authService.Register(request.Username, request.Email, request.Password)

	if err != nil {
		log.Printf("Error in registration: %v", err)
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.UserAlreadyExists,
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

// Login handles user login
// @Summary Login user
// @Description Authenticates a user and returns a JWT token
// @Accept json
// @Produce json
// @Param body body utils.LoginRequest true "Login credentials"
// @Success 200 {object} utils.AuthResponse
// @Failure 400 {object} utils.MessageResponse
// @Router /auth/enter [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request utils.LoginRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
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

// CreateMaster handles the creation of a master user
// @Summary Create a master user
// @Description Creates a user with master privileges (only works when no master exists)
// @Accept json
// @Produce json
// @Param body body utils.CreateMasterRequest true "Master user details"
// @Success 200 {object} utils.AuthResponse
// @Failure 403 {object} utils.MessageResponse
// @Router /master/create [post]
func (c *AuthController) CreateMaster(ctx *gin.Context) {
	var request utils.CreateMasterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
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

// DeleteMaster handles the deletion of a master user
// @Summary Delete a master user
// @Description Removes a user with master privileges
// @Accept json
// @Produce json
// @Param body body utils.DeleteMasterRequest true "Master user credentials"
// @Success 200 {object} utils.MessageResponse
// @Failure 403 {object} utils.MessageResponse
// @Failure 404 {object} utils.MessageResponse
// @Router /master/delete [post]
func (c *AuthController) DeleteMaster(ctx *gin.Context) {
	var request utils.DeleteMasterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
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
