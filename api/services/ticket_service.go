package services

import (
	"errors"
	"time"

	"hcall/api/models"
	"hcall/api/repository"
	"hcall/api/utils"
)

type TicketService struct {
	ticketRepo *repository.TicketRepository
	userRepo   *repository.UserRepository
}

func NewTicketService() *TicketService {
	return &TicketService{
		ticketRepo: repository.NewTicketRepository(),
		userRepo:   repository.NewUserRepository(),
	}
}

// CreateTicket creates a new ticket
func (s *TicketService) CreateTicket(authorID uint, authorEmail, name, explanation string, images []utils.ImageDTO) error {
	// Create the ticket
	ticket := &models.Ticket{
		Name:        name,
		Explanation: explanation,
		Status:      models.PendingStatus,
		AuthorID:    authorID,
		AuthorEmail: authorEmail,
		CreatedAt:   time.Now(),
	}

	// Save the ticket to the database
	if err := s.ticketRepo.CreateTicket(ticket); err != nil {
		return err
	}

	// Process images if there are any
	if len(images) > 0 {
		// Save each image directly to the database with base64 content
		for _, img := range images {
			// Create image record in database
			image := &models.Image{
				TicketID:    ticket.ID,
				Name:        img.Name,
				ContentType: img.Type,
				Base64:      img.Content, // Salva o conteúdo base64 diretamente
				UploadedAt:  time.Now(),
			}

			if err := s.ticketRepo.CreateImage(image); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetTickets gets all tickets or tickets by author or status
func (s *TicketService) GetTickets(authorEmail string, status string) ([]models.Ticket, error) {
	// Get all tickets
	if authorEmail == "" && status == "" {
		return s.ticketRepo.GetTickets()
	}

	// Get tickets by author
	if authorEmail != "" && status == "" {
		return s.ticketRepo.GetTicketsByAuthor(authorEmail)
	}

	// Get tickets by status
	if authorEmail == "" && status != "" {
		statusEnum := models.TicketStatus(status)
		return s.ticketRepo.GetTicketsByStatus(statusEnum)
	}

	// Get tickets by author and status
	statusEnum := models.TicketStatus(status)
	return s.ticketRepo.GetTicketsByAuthorAndStatus(authorEmail, statusEnum)
}

// GetTicketDetails gets the details of a ticket
func (s *TicketService) GetTicketDetails(ticketID string) (*models.Ticket, error) {
	ticket, err := s.ticketRepo.GetTicketWithDetails(ticketID)
	if err != nil {
		return nil, err
	}

	// Garantir que o histórico nunca seja nulo, mas sempre um array vazio no mínimo
	if ticket.History == nil {
		ticket.History = []models.TicketHistory{}
	}

	// Garantir que o array de imagens nunca seja nulo, mas sempre um array vazio no mínimo
	if ticket.Images == nil {
		ticket.Images = []models.Image{}
	}

	return ticket, nil
}

// UpdateTicketStatus updates the status of a ticket
func (s *TicketService) UpdateTicketStatus(ticketID string, status models.TicketStatus) error {
	return s.ticketRepo.UpdateTicketStatus(ticketID, status)
}

// AddTicketHistory adds a history entry to a ticket
func (s *TicketService) AddTicketHistory(ticketID, message string) error {
	history := models.TicketHistory{
		TicketID:  ticketID,
		Message:   message,
		CreatedAt: time.Now(),
	}

	return s.ticketRepo.AddTicketHistory(&history)
}

// DeleteTicket deletes a ticket
func (s *TicketService) DeleteTicket(ticketID string, userID uint, userRole models.Role) error {
	// Get the ticket to check ownership
	ticket, err := s.ticketRepo.GetTicket(ticketID)
	if err != nil {
		return err
	}

	// Check if user is admin/master or the owner of the ticket
	if userRole != models.AdminRole && userRole != models.MasterRole && ticket.AuthorID != userID {
		return errors.New("you don't have permission to delete this ticket")
	}

	// Delete the ticket
	return s.ticketRepo.DeleteTicket(ticketID)
}

func (s *TicketService) GetUserUsername(userID uint) (string, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
