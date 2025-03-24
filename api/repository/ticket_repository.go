package repository

import (
	"errors"

	"hcall/api/database"
	"hcall/api/models"

	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{
		DB: database.DB,
	}
}

// CreateTicket creates a new ticket
func (r *TicketRepository) CreateTicket(ticket *models.Ticket) error {
	return r.DB.Create(ticket).Error
}

// GetTicket gets a ticket by ID
func (r *TicketRepository) GetTicket(id string) (*models.Ticket, error) {
	var ticket models.Ticket
	result := r.DB.Where("id = ?", id).First(&ticket)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ticket not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &ticket, nil
}

// FindByID finds a ticket by ID
func (r *TicketRepository) FindByID(id string) (*models.Ticket, error) {
	var ticket models.Ticket
	result := r.DB.Where("id = ?", id).First(&ticket)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ticket not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &ticket, nil
}

// GetTicketWithDetails gets a ticket with all its details (images and history)
func (r *TicketRepository) GetTicketWithDetails(id string) (*models.Ticket, error) {
	var ticket models.Ticket
	result := r.DB.Preload("Images").Preload("History").Where("id = ?", id).First(&ticket)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ticket not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &ticket, nil
}

// GetTickets gets all tickets
func (r *TicketRepository) GetTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.DB.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

// GetTicketsByAuthor gets tickets by author
func (r *TicketRepository) GetTicketsByAuthor(authorEmail string) ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.DB.Where("author_email = ?", authorEmail).Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("author doesn't have tickets")
	}

	return tickets, nil
}

// GetTicketsByStatus gets tickets by status
func (r *TicketRepository) GetTicketsByStatus(status models.TicketStatus) ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.DB.Where("status = ?", status).Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("no tickets found with specified status")
	}

	return tickets, nil
}

// GetTicketsByAuthorAndStatus gets tickets by author and status
func (r *TicketRepository) GetTicketsByAuthorAndStatus(authorEmail string, status models.TicketStatus) ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.DB.Where("author_email = ? AND status = ?", authorEmail, status).Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("no tickets found with specified author and status")
	}

	return tickets, nil
}

// UpdateTicketStatus updates a ticket's status
func (r *TicketRepository) UpdateTicketStatus(id string, status models.TicketStatus) error {
	result := r.DB.Model(&models.Ticket{}).Where("id = ?", id).Update("status", status)

	if result.RowsAffected == 0 {
		return errors.New("ticket not found")
	}

	return result.Error
}

// AddTicketHistory adds a new entry to the ticket's history
func (r *TicketRepository) AddTicketHistory(history *models.TicketHistory) error {
	return r.DB.Create(history).Error
}

// DeleteTicket deletes a ticket
func (r *TicketRepository) DeleteTicket(id string) error {
	// Start a transaction
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete images associated with the ticket
	if err := tx.Where("ticket_id = ?", id).Delete(&models.Image{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete history associated with the ticket
	if err := tx.Where("ticket_id = ?", id).Delete(&models.TicketHistory{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete the ticket
	result := tx.Where("id = ?", id).Delete(&models.Ticket{})
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("ticket not found")
	}

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error
}

// CreateImage creates a new image
func (r *TicketRepository) CreateImage(image *models.Image) error {
	return r.DB.Create(image).Error
}
