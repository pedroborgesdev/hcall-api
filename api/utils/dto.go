package utils

import "hcall/api/models"

// Request DTOs

type LoginRequest struct {
	Email    string `json:"user_email" binding:"required,email"`
	Password string `json:"user_password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"user_name" binding:"required"`
	Email    string `json:"user_email" binding:"required,email"`
	Password string `json:"user_password" binding:"required"`
}

type CreateUserRequest struct {
	Username string      `json:"user_name" binding:"required"`
	Email    string      `json:"user_email" binding:"required,email"`
	Password string      `json:"user_password" binding:"required"`
	Role     models.Role `json:"user_role" binding:"required,oneof=user admin"`
}

type CreateMasterRequest struct {
	Email    string `json:"master_email" binding:"required,email"`
	Password string `json:"master_password" binding:"required"`
}

type DeleteMasterRequest struct {
	Email    string `json:"master_email" binding:"required,email"`
	Password string `json:"master_password" binding:"required"`
}

type DeleteUserRequest struct {
	Email string `json:"user_email" binding:"required,email"`
}

type CreateTicketRequest struct {
	Name        string     `json:"ticket_name" binding:"required"`
	Explanation string     `json:"ticket_explain" binding:"required"`
	Images      []ImageDTO `json:"ticket_images" binding:"omitempty,dive"`
}

type ImageDTO struct {
	Name    string `json:"image_name" binding:"required"`
	Content string `json:"image_content" binding:"required"` // Base64 encoded image data (will be stored directly in database)
	Type    string `json:"image_type" binding:"required,oneof=image/jpeg image/png image/gif"`
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
	Token   string `json:"jwt_token,omitempty"`
	Status  bool   `json:"status"`
}

type MessageResponse struct {
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	Status  bool   `json:"status"`
}

type UserResponse struct {
	Username  string      `json:"user_name"`
	Email     string      `json:"user_email"`
	Password  string      `json:"user_password,omitempty"`
	CreatedAt string      `json:"user_created_at,omitempty"`
	Role      models.Role `json:"user_role"`
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

type GetCountersResponse struct {
	Total    int  `json:"tickets_total"`
	Pending  int  `json:"tickets_pending"`
	Doing    int  `json:"tickets_doing"`
	Conclued int  `json:"tickets_conclued"`
	Status   bool `json:"status"`
}
