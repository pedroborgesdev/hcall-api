package controllers

import (
	"net/http"

	"hcall/api/dictionaries"
	"hcall/api/models"
	"hcall/api/services"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// GetUsers handles getting user information
// @Summary Get user information
// @Description Retrieves information about a specific user or lists all users
// @Accept json
// @Produce json
// @Param email query string false "User's email"
// @Param role query string false "User's role"
// @Success 200 {object} utils.UserResponse
// @Success 200 {object} utils.UsersListResponse
// @Failure 404 {object} utils.MessageResponse
// @Security Bearer
// @Router /user/fetch [get]
func (c *UserController) GetUsers(ctx *gin.Context) {
	email := ctx.Query("email")
	role := ctx.Query("role")

	// If email is provided, get specific user
	if email != "" {
		// If role is also provided, get user with specific role
		if role != "" {
			userRole := models.Role(role)
			user, err := c.userService.GetUserByEmailAndRole(email, userRole)
			if err != nil {
				ctx.JSON(http.StatusNotFound, utils.MessageResponse{
					Message: "Email aren't registered",
					Reason:  err.Error(),
					Status:  false,
				})
				return
			}

			// Return user details
			ctx.JSON(http.StatusOK, utils.UserResponse{
				Username:  user.Username,
				Email:     user.Email,
				Password:  "********",
				CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
				Role:      user.Role,
				Status:    true,
			})
			return
		}

		// Get user by email
		user, err := c.userService.GetUserByEmail(email)
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.MessageResponse{
				Message: "Email aren't registered",
				Reason:  err.Error(),
				Status:  false,
			})
			return
		}

		// Return user details
		ctx.JSON(http.StatusOK, utils.UserResponse{
			Username:  user.Username,
			Email:     user.Email,
			Password:  "********",
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			Role:      user.Role,
			Status:    true,
		})
		return
	}

	// If role is provided, get users with specific role
	if role != "" {
		userRole := models.Role(role)
		users, err := c.userService.GetUsersByRole(userRole)
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.MessageResponse{
				Message: "No users found with specified role",
				Reason:  err.Error(),
				Status:  false,
			})
			return
		}

		// Convert to response format
		responseUsers := make([]models.ResponseUser, len(users))
		for i, user := range users {
			responseUsers[i] = user.ToResponse(false)
		}

		ctx.JSON(http.StatusOK, utils.UsersListResponse{
			Users:  responseUsers,
			Status: true,
		})
		return
	}

	// Get all users
	users, err := c.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.MessageResponse{
			Message: dictionaries.InternalServerError,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Convert to response format
	responseUsers := make([]models.ResponseUser, len(users))
	for i, user := range users {
		responseUsers[i] = user.ToResponse(false)
	}

	ctx.JSON(http.StatusOK, utils.UsersListResponse{
		Users:  responseUsers,
		Status: true,
	})
}

// CreateUser handles user creation
// @Summary Create a new user
// @Description Creates a new user in the system
// @Accept json
// @Produce json
// @Param body body utils.CreateUserRequest true "User details"
// @Success 200 {object} utils.MessageResponse
// @Failure 400 {object} utils.MessageResponse
// @Router /users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var request utils.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	err := c.userService.CreateUser(request.Username, request.Email, request.Password, request.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.UserCreationFailed,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.UserCreatedSuccess,
		Status:  true,
	})
}

// DeleteUser handles user deletion
// @Summary Delete a user
// @Description Deletes a user from the system
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} utils.MessageResponse
// @Failure 404 {object} utils.MessageResponse
// @Failure 500 {object} utils.MessageResponse
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.userService.DeleteUser(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		message := dictionaries.InternalServerError
		if err.Error() == "user not found" {
			statusCode = http.StatusNotFound
			message = dictionaries.UserNotFound
		}
		ctx.JSON(statusCode, utils.MessageResponse{
			Message: message,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.UserDeletedSuccess,
		Status:  true,
	})
}
