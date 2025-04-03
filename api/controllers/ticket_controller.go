package controllers

import (
	"hcall/api/dictionaries"
	"hcall/api/logger"
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
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
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
		logger.Error("Ticket Controller: Ticket creation failed", map[string]interface{}{
			"user_id": userID,
			"email":   userEmail,
			"name":    request.Name,
			"error":   err.Error(),
		})
		utils.SendError(ctx, utils.CodeInvalidInput, dictionaries.TicketCreationFailed, err)
		return
	}

	logger.Info("Ticket Controller: Ticket created successfully", map[string]interface{}{
		"user_id": userID,
		"email":   userEmail,
		"name":    request.Name,
	})

	// Return success
	utils.SendSuccess(ctx, dictionaries.TicketCreatedSuccess, nil)
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
			logger.Error("Ticket Controller: Invalid date format in ticket query", map[string]interface{}{
				"date":  date,
				"error": err.Error(),
			})
			utils.SendError(ctx, utils.CodeInvalidInput, dictionaries.InvalidDateFormat, err)
			return
		}

		logger.Error("Ticket Controller: Failed to get tickets", map[string]interface{}{
			"author": author,
			"status": status,
			"date":   date,
			"name":   name,
			"error":  err.Error(),
		})
		utils.SendError(ctx, utils.CodeNotFound, dictionaries.NoTicketsFound, err)
		return
	}

	logger.Info("Ticket Controller: Tickets retrieved successfully", map[string]interface{}{
		"count": len(tickets),
	})

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

	utils.SendSuccess(ctx, "Tickets found", gin.H{
		"tickets": responseTickets,
	})
}

func (c *TicketController) GetTicketDetails(ctx *gin.Context) {
	ticketID := ctx.Query("ticket_id")
	if ticketID == "" {
		utils.SendError(ctx, utils.CodeInvalidInput, "Ticket ID is required", nil)
		return
	}

	// Call the service
	ticket, err := c.ticketService.GetTicketDetails(ticketID)
	if err != nil {
		logger.Error("Ticket Controller: Failed to get ticket details", map[string]interface{}{
			"ticket_id": ticketID,
			"error":     err.Error(),
		})
		utils.SendError(ctx, utils.CodeNotFound, dictionaries.TicketNotFound, err)
		return
	}

	logger.Info("Ticket Controller: Ticket details retrieved successfully", map[string]interface{}{
		"ticket_id": ticketID,
	})

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

	utils.SendSuccess(ctx, "Ticket details found", gin.H{
		"ticket": detailedResponse,
	})
}

func (c *TicketController) UpdateTicketStatus(ctx *gin.Context) {
	var request utils.UpdateTicketStatusRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
		return
	}

	// Call the service
	err := c.ticketService.UpdateTicketStatus(request.TicketID, request.Status)
	if err != nil {
		logger.Error("Ticket Controller: Failed to update ticket status", map[string]interface{}{
			"ticket_id": request.TicketID,
			"status":    request.Status,
			"error":     err.Error(),
		})
		utils.SendError(ctx, utils.CodeInvalidInput, dictionaries.TicketStatusUpdateFailed, err)
		return
	}

	logger.Info("Ticket Controller: Ticket status updated successfully", map[string]interface{}{
		"ticket_id": request.TicketID,
		"status":    request.Status,
	})

	utils.SendSuccess(ctx, dictionaries.TicketStatusUpdated, nil)
}

func (c *TicketController) UpdateTicketHistory(ctx *gin.Context) {
	var request utils.UpdateTicketHistoryRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
		return
	}

	// Call the service
	err := c.ticketService.AddTicketHistory(request.TicketID, request.Message)
	if err != nil {
		logger.Error("Ticket Controller: Failed to update ticket history", map[string]interface{}{
			"ticket_id": request.TicketID,
			"error":     err.Error(),
		})
		utils.SendError(ctx, utils.CodeInvalidInput, dictionaries.TicketHistoryAddFailed, err)
		return
	}

	logger.Info("Ticket Controller: Ticket history updated successfully", map[string]interface{}{
		"ticket_id": request.TicketID,
	})

	utils.SendSuccess(ctx, dictionaries.TicketHistoryAdded, nil)
}

func (c *TicketController) DeleteTicket(ctx *gin.Context) {
	var request utils.RemoveTicketRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, utils.CodeInvalidInput, utils.MsgInvalidInput, err)
		return
	}

	// Get user ID and role
	userID, _ := ctx.Get("userId")
	userRole, _ := ctx.Get("userRole")

	// Call the service
	err := c.ticketService.DeleteTicket(request.TicketID, userID.(uint), userRole.(models.Role))
	if err != nil {
		logger.Error("Ticket Controller: Failed to delete ticket", map[string]interface{}{
			"ticket_id": request.TicketID,
			"user_id":   userID,
			"role":      userRole,
			"error":     err.Error(),
		})
		utils.SendError(ctx, utils.CodeForbidden, dictionaries.NoPermissionToDelete, err)
		return
	}

	logger.Info("Ticket Controller: Ticket deleted successfully", map[string]interface{}{
		"ticket_id": request.TicketID,
		"user_id":   userID,
	})

	utils.SendSuccess(ctx, dictionaries.TicketDeletedSuccess, nil)
}

func (c *TicketController) CountTicket(ctx *gin.Context) {
	count, err := c.ticketService.GetCounters()
	if err != nil {
		logger.Error("Ticket Controller: Failed to get ticket counters", map[string]interface{}{
			"error": err.Error(),
		})
		utils.SendError(ctx, utils.CodeNotFound, dictionaries.TicketNotFound, err)
		return
	}

	logger.Info("Ticket Controller: Ticket counters retrieved successfully", map[string]interface{}{
		"total":    count.Total,
		"pending":  count.Pending,
		"doing":    count.Doing,
		"conclued": count.Conclued,
	})

	utils.SendSuccess(ctx, "Ticket counters retrieved", gin.H{
		"total":    count.Total,
		"pending":  count.Pending,
		"doing":    count.Doing,
		"conclued": count.Conclued,
	})
}
