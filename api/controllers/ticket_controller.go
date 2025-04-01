package controllers

import (
	"log"
	"net/http"

	"hcall/api/dictionaries"
	"hcall/api/models"
	"hcall/api/services"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	ticketService *services.TicketService
}

func NewTicketController() *TicketController {
	return &TicketController{
		ticketService: services.NewTicketService(),
	}
}

func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var request utils.CreateTicketRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Get user ID and email
	userID, _ := ctx.Get("userId")
	userEmail, _ := ctx.Get("userEmail")

	// Call the service
	err := c.ticketService.CreateTicket(
		userID.(uint),
		userEmail.(string),
		request.Name,
		request.Explanation,
		request.Images,
	)

	if err != nil {
		log.Printf("Error creating ticket: %v", err)
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.TicketCreationFailed,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Return success
	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.TicketCreatedSuccess,
		Status:  true,
	})
}

func (c *TicketController) GetTickets(ctx *gin.Context) {
	author := ctx.Query("author")
	status := ctx.Query("status")
	date := ctx.Query("date")
	name := ctx.Query("name")

	// Call the service
	tickets, err := c.ticketService.GetTickets(author, status, date, name)

	if err != nil {
		if err.Error() == "Invalid date format" {
			ctx.JSON(http.StatusNotFound, utils.MessageResponse{
				Message: dictionaries.InvalidDateFormat,
				Reason:  err.Error(),
				Status:  false,
			})
			return
		}

		ctx.JSON(http.StatusNotFound, utils.MessageResponse{
			Message: dictionaries.NoTicketsFound,
			Status:  false,
		})
		return
	}

	// Convert to response format
	responseTickets := make([]models.BasicTicketResponse, len(tickets))
	for i, ticket := range tickets {
		// Get user information using AuthorID from ticket
		username, err := c.ticketService.GetUserUsername(ticket.AuthorID)
		if err != nil {
			responseTickets[i] = ticket.ToBasicResponse("Unknown User")
			continue
		}
		responseTickets[i] = ticket.ToBasicResponse(username)
	}

	ctx.JSON(http.StatusOK, utils.TicketsListResponse{
		Tickets: responseTickets,
		Status:  true,
	})
}

func (c *TicketController) GetTicketDetails(ctx *gin.Context) {
	ticketID := ctx.Query("ticket_id")
	if ticketID == "" {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Reason:  "Ticket ID is required",
			Status:  false,
		})
		return
	}

	// Call the service
	ticket, err := c.ticketService.GetTicketDetails(ticketID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.MessageResponse{
			Message: dictionaries.TicketNotFound,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Garantir explicitamente que o histórico seja um array vazio, se for nil
	if ticket.History == nil {
		ticket.History = []models.TicketHistory{}
	}

	// Garantir explicitamente que images seja um array vazio, se for nil
	if ticket.Images == nil {
		ticket.Images = []models.Image{}
	}

	// Convert to response format
	detailedResponse := ticket.ToDetailedResponse(true)

	response := utils.TicketResponse{
		DetailedTicketResponse: detailedResponse,
		Status:                 true,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *TicketController) UpdateTicketStatus(ctx *gin.Context) {
	var request utils.UpdateTicketStatusRequest

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
	err := c.ticketService.UpdateTicketStatus(request.TicketID, request.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.TicketStatusUpdateFailed,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.TicketStatusUpdated,
		Status:  true,
	})
}

func (c *TicketController) UpdateTicketHistory(ctx *gin.Context) {
	var request utils.UpdateTicketHistoryRequest

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
	err := c.ticketService.AddTicketHistory(request.TicketID, request.Message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.TicketHistoryAddFailed,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.TicketHistoryAdded,
		Status:  true,
	})
}

func (c *TicketController) DeleteTicket(ctx *gin.Context) {
	var request utils.RemoveTicketRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	// Get user ID and role
	userID, _ := ctx.Get("userId")
	userRole, _ := ctx.Get("userRole")

	// Call the service
	err := c.ticketService.DeleteTicket(request.TicketID, userID.(uint), userRole.(models.Role))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.NoPermissionToDelete,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.TicketDeletedSuccess,
		Status:  true,
	})
}

func (c *TicketController) CountTicket(ctx *gin.Context) {
	count, err := c.ticketService.GetCounters()
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.MessageResponse{
			Message: dictionaries.TicketNotFound,
			Reason:  err.Error(),
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.GetCountersResponse{
		Total:    count.Total,
		Pending:  count.Pending,
		Doing:    count.Doing,
		Conclued: count.Conclued,
		Status:   true,
	})
}
