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

// CreateTicket handles ticket creation
// @Summary Create a new ticket
// @Description Creates a new ticket with the provided information
// @Accept json
// @Produce json
// @Param body body utils.CreateTicketRequest true "Ticket details"
// @Success 200 {object} utils.MessageResponse
// @Failure 400 {object} utils.MessageResponse
// @Failure 403 {object} utils.MessageResponse
// @Security Bearer
// @Router /ticket/create [post]
func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var request utils.CreateTicketRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
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

// GetTickets handles listing tickets
// @Summary List tickets
// @Description Lists tickets from a specific author or all system tickets
// @Accept json
// @Produce json
// @Param author query string false "Author's email"
// @Param status query string false "Ticket status"
// @Success 200 {object} utils.TicketsListResponse
// @Failure 404 {object} utils.MessageResponse
// @Security Bearer
// @Router /ticket/fetch [get]
func (c *TicketController) GetTickets(ctx *gin.Context) {
	author := ctx.Query("author")
	status := ctx.Query("status")

	// Call the service
	tickets, err := c.ticketService.GetTickets(author, status)
	if err != nil {
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

// GetTicketDetails handles getting ticket details
// @Summary Get ticket details
// @Description Retrieves detailed information about a specific ticket
// @Accept json
// @Produce json
// @Param ticket_id query string true "Ticket ID"
// @Success 200 {object} utils.TicketResponse
// @Failure 404 {object} utils.MessageResponse
// @Security Bearer
// @Router /ticket/info [get]
func (c *TicketController) GetTicketDetails(ctx *gin.Context) {
	ticketID := ctx.Query("ticket_id")
	if ticketID == "" {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Status:  false,
		})
		return
	}

	// Call the service
	ticket, err := c.ticketService.GetTicketDetails(ticketID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.MessageResponse{
			Message: dictionaries.TicketNotFound,
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

// UpdateTicketStatus handles updating ticket status
// @Summary Update ticket status
// @Description Updates the status of a specific ticket
// @Accept json
// @Produce json
// @Param body body utils.UpdateTicketStatusRequest true "Ticket status update details"
// @Success 200 {object} utils.MessageResponse
// @Failure 400 {object} utils.MessageResponse
// @Security Bearer
// @Router /ticket/edit [post]
func (c *TicketController) UpdateTicketStatus(ctx *gin.Context) {
	var request utils.UpdateTicketStatusRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Status:  false,
		})
		return
	}

	// Call the service
	err := c.ticketService.UpdateTicketStatus(request.TicketID, request.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.TicketStatusUpdateFailed,
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.TicketStatusUpdated,
		Status:  true,
	})
}

// UpdateTicketHistory handles updating ticket history
// @Summary Update ticket history
// @Description Adds a new entry to the ticket's history
// @Accept json
// @Produce json
// @Param body body utils.UpdateTicketHistoryRequest true "Ticket history update details"
// @Success 200 {object} utils.MessageResponse
// @Failure 400 {object} utils.MessageResponse
// @Security Bearer
// @Router /ticket/update [post]
func (c *TicketController) UpdateTicketHistory(ctx *gin.Context) {
	var request utils.UpdateTicketHistoryRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
			Status:  false,
		})
		return
	}

	// Call the service
	err := c.ticketService.AddTicketHistory(request.TicketID, request.Message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.TicketHistoryAddFailed,
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.TicketHistoryAdded,
		Status:  true,
	})
}

// DeleteTicket handles ticket deletion
// @Summary Delete ticket
// @Description Deletes a specific ticket
// @Accept json
// @Produce json
// @Param body body utils.RemoveTicketRequest true "Ticket deletion details"
// @Success 200 {object} utils.MessageResponse
// @Failure 400 {object} utils.MessageResponse
// @Security Bearer
// @Router /ticket/remove [post]
func (c *TicketController) DeleteTicket(ctx *gin.Context) {
	var request utils.RemoveTicketRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse{
			Message: dictionaries.InvalidData,
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
			Status:  false,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.MessageResponse{
		Message: dictionaries.TicketDeletedSuccess,
		Status:  true,
	})
}
