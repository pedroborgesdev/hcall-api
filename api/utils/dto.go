package utils

import "hcall/api/models"

// Request DTOs

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Username string      `json:"username" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required"`
	Role     models.Role `json:"role" binding:"required,oneof=user admin"`
}

type CreateMasterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type DeleteMasterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type DeleteUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type CreateTicketRequest struct {
	Name        string     `json:"ticket_name" binding:"required"`
	Explanation string     `json:"ticket_explain" binding:"required"`
	Images      []ImageDTO `json:"images" binding:"omitempty,dive"`
}

type ImageDTO struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"` // Base64 encoded image data (will be stored directly in database)
	Type    string `json:"type" binding:"required,oneof=image/jpeg image/png image/gif"`
}

type UpdateTicketStatusRequest struct {
	TicketID string              `json:"ticket_id" binding:"required"`
	Status   models.TicketStatus `json:"ticket_status" binding:"required,oneof=pending doing conclued"`
}

type UpdateTicketHistoryRequest struct {
	TicketID string `json:"ticket_id" binding:"required"`
	Message  string `json:"ticket_return" binding:"required"`
}

type RemoveTicketRequest struct {
	TicketID string `json:"ticket_id" binding:"required"`
}

// Response DTOs

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	Status  bool   `json:"status"`
}

type MessageResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type UserResponse struct {
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Password  string      `json:"password,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	Role      models.Role `json:"role"`
	Status    bool        `json:"status"`
}

type UsersListResponse struct {
	Users  []models.ResponseUser `json:"users"`
	Status bool                  `json:"status"`
}

type TicketsListResponse struct {
	Tickets []models.BasicTicketResponse `json:"tickets"`
	Status  bool                         `json:"status"`
}

type TicketResponse struct {
	models.DetailedTicketResponse
	Status bool `json:"status"`
}
