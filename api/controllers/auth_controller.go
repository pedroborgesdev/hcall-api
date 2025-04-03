package controllers

import (
	"hcall/api/logger"
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
func (c *AuthController) Register(ctx *gin.Context) {
	var request utils.RegisterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
		return
	}

	// Call the service
	user, token, err := c.authService.Register(request.Username, request.Email, request.Password)
	if err != nil {
		logger.Error("Auth Controller: Registration failed", map[string]interface{}{
			"email": request.Email,
			"error": err.Error(),
		})
		utils.SendError(ctx, utils.CodeInvalidInput, "Registration failed", err)
		return
	}

	logger.Info("Auth Controller: User registered successfully", map[string]interface{}{
		"email": user.Email,
		"role":  user.Role,
	})

	// Return success
	utils.SendSuccess(ctx, "Registration successful", gin.H{
		"token": token,
		"user": gin.H{
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

// Login handles user login
func (c *AuthController) Login(ctx *gin.Context) {
	var request utils.LoginRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
		return
	}

	// Call the service
	user, token, err := c.authService.Login(request.Email, request.Password)
	if err != nil {
		logger.Error("Auth Controller: Login failed", map[string]interface{}{
			"email": request.Email,
			"error": err.Error(),
		})
		utils.SendError(ctx, utils.CodeUnauthorized, utils.MsgInvalidCredentials, err)
		return
	}

	logger.Info("Auth Controller: User logged in successfully", map[string]interface{}{
		"email": user.Email,
		"role":  user.Role,
	})

	// Return success
	utils.SendSuccess(ctx, "Login successful", gin.H{
		"token": token,
		"user": gin.H{
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

// CreateMaster creates a master user
func (c *AuthController) CreateMaster(ctx *gin.Context) {
	var request utils.CreateMasterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
		return
	}

	// Call the service
	master, token, err := c.authService.CreateMaster(request.Email, request.Password)
	if err != nil {
		logger.Error("Auth Controller: Master user creation failed", map[string]interface{}{
			"email": request.Email,
			"error": err.Error(),
		})
		utils.SendError(ctx, utils.CodeDuplicateEntry, "Master user already exists", err)
		return
	}

	logger.Info("Auth Controller: Master user created successfully", map[string]interface{}{
		"email": master.Email,
		"role":  master.Role,
	})

	// Return success
	utils.SendSuccess(ctx, "Master user created successfully", gin.H{
		"token": token,
		"user": gin.H{
			"email": master.Email,
			"role":  master.Role,
		},
	})
}

// DeleteMaster deletes the master user
func (c *AuthController) DeleteMaster(ctx *gin.Context) {
	var request utils.DeleteMasterRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
		return
	}

	// Call the service
	err := c.authService.DeleteMaster(request.Email, request.Password)
	if err != nil {
		logger.Error("Auth Controller: Master user deletion failed", map[string]interface{}{
			"email": request.Email,
			"error": err.Error(),
		})
		if err.Error() == "master user not found" {
			utils.SendError(ctx, utils.CodeNotFound, "Master user not found", err)
			return
		}
		utils.SendError(ctx, utils.CodeUnauthorized, "Invalid master credentials", err)
		return
	}

	logger.Info("Auth Controller: Master user deleted successfully", map[string]interface{}{
		"email": request.Email,
	})

	// Return success
	utils.SendSuccess(ctx, "Master user deleted successfully", nil)
}
