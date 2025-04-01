package repository

import (
	"errors"
	"log"
	"time"

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

// create a function that remove tickets with specified status
func (r *TicketRepository) RemoveTicketsWithStatus(status models.TicketStatus, remove_after int) error {
	// Get atual date and refator now to YYYY/MM/DD format
	now := time.Now().Format("2006/01/02")

	// subtract remove_after (days) from now
	now = time.Now().AddDate(0, 0, -remove_after).Format("2006/01/02")
	log.Println(now) // debugg

	// remove tickets with specified status and created_at before now
	result := r.DB.Where("status =? AND created_at <?", status, now).Delete(&models.Ticket{})
	return result.Error
}

// CreateTicket creates a new ticket
func (r *TicketRepository) CreateTicket(ticket *models.Ticket) error {
	return database.ExecuteInTransaction(r.DB, func(tx *gorm.DB) error {
		// First, verify if the user exists
		var user models.User
		if err := tx.Where("email = ?", ticket.AuthorEmail).First(&user).Error; err != nil {
			return errors.New("author not found")
		}

		// Set the correct author_id from the found user
		ticket.AuthorID = user.ID

		// Now create the ticket
		return tx.Create(ticket).Error
	})
}

func (r *TicketRepository) CountTicket(status string) error {
	return database.ExecuteInTransaction(r.DB, func(tx *gorm.DB) error {
		var counters models.Counters
		counters.ID = 1

		// Get the current counter or create if not exists - within transaction
		if err := tx.First(&counters).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				counters = models.Counters{Pending: 0, Doing: 0, Conclued: 0}
				if err := tx.Create(&counters).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}

		if status == "pending" {
			log.Println("pending")
			if err := tx.Model(&models.Counters{}).Where("id =?", counters.ID).
				Update("pending", gorm.Expr("pending +?", 1)).Error; err != nil {
				return err
			}
		}

		if status == "doing" {
			log.Println("doing")
			if err := tx.Model(&models.Counters{}).Where("id =?", counters.ID).
				Update("doing", gorm.Expr("doing +?", 1)).Error; err != nil {
				return err
			}
		}

		if status == "conclued" {
			log.Println("conclued")
			if err := tx.Model(&models.Counters{}).Where("id =?", counters.ID).
				Update("conclued", gorm.Expr("conclued +?", 1)).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&models.Counters{}).Where("id =?", counters.ID).
			Update("total", gorm.Expr("total +?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *TicketRepository) GetCounters() (*models.Counters, error) {
	var counters models.Counters
	result := r.DB.First(&counters)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("counters not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &counters, nil
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

func (r *TicketRepository) GetTicketsByDate(date string) ([]models.Ticket, error) {
	var tickets []models.Ticket

	// Query tickets created on or after the specified date
	if err := r.DB.Where("DATE(created_at) >= ?", date).Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("no tickets found after the specified date")
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

func (r *TicketRepository) GetTicketsByAuthorAndDate(author string, date string) ([]models.Ticket, error) {
	var tickets []models.Ticket

	// Query tickets created on or after the specified date and author
	if err := r.DB.Where("author_email =? AND DATE(created_at) >=?", author, date).Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("no tickets found after the specified date and author")
	}

	return tickets, nil
}

// GetTicketsByAuthorAndStatusAndDate gets tickets by date and status
func (r *TicketRepository) GetTicketsByStatusAndDate(status models.TicketStatus, date string) ([]models.Ticket, error) {
	var tickets []models.Ticket

	// Query tickets created on or after the specified date and status
	if err := r.DB.Where("status =? AND DATE(created_at) >=?", status, date).Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("no tickets found after the specified date and status")
	}

	return tickets, nil
}

// GetTicketsByAuthorAndStatusAndDate gets tickets by author, status and date
func (r *TicketRepository) GetTicketsByAuthorAndStatusAndDate(authorEmail string, status models.TicketStatus, date string) ([]models.Ticket, error) {
	var tickets []models.Ticket

	// Query tickets created on or after the specified date, author and status
	if err := r.DB.Where("author_email =? AND status =? AND DATE(created_at) >=?", authorEmail, status, date).Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("no tickets found after the specified date, author and status")
	}

	return tickets, nil
}

// UpdateTicketStatus updates a ticket's status
func (r *TicketRepository) UpdateTicketStatus(id string, status models.TicketStatus) error {
	return database.ExecuteInTransaction(r.DB, func(tx *gorm.DB) error {
		//get status of ticket id
		var ticket models.Ticket
		if err := tx.Where("id =?", id).First(&ticket).Error; err != nil {
			return err
		}

		// get current status of ticket
		currentStatus := ticket.Status

		if currentStatus == status {
			return errors.New("New status is the same as the current status")
		}

		result := tx.Model(&models.Ticket{}).Where("id = ?", id).Update("status", status)

		if result.RowsAffected == 0 {
			return errors.New("ticket not found")
		}

		return result.Error
	})
}

// AddTicketHistory adds a new entry to the ticket's history
func (r *TicketRepository) AddTicketHistory(history *models.TicketHistory) error {
	return r.DB.Create(history).Error
}

// DeleteTicket deletes a ticket
func (r *TicketRepository) DeleteTicket(id string) error {
	return database.ExecuteInTransaction(r.DB, func(tx *gorm.DB) error {
		// Delete images associated with the ticket
		if err := tx.Where("ticket_id = ?", id).Delete(&models.Image{}).Error; err != nil {
			return err
		}

		// Delete history associated with the ticket
		if err := tx.Where("ticket_id = ?", id).Delete(&models.TicketHistory{}).Error; err != nil {
			return err
		}

		// Delete the ticket
		result := tx.Where("id = ?", id).Delete(&models.Ticket{})
		if result.RowsAffected == 0 {
			return errors.New("ticket not found")
		}

		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

// CreateImage creates a new image
func (r *TicketRepository) CreateImage(image *models.Image) error {
	return r.DB.Create(image).Error
}

// Add these new methods
func (r *TicketRepository) GetTicketsByName(name string) ([]models.Ticket, error) {
	var tickets []models.Ticket
	result := r.DB.Where("name ILIKE ?", "%"+name+"%").Find(&tickets)
	return tickets, result.Error
}

func (r *TicketRepository) GetTicketsByAuthorAndName(authorEmail, name string) ([]models.Ticket, error) {
	var tickets []models.Ticket
	result := r.DB.Where("author_email = ? AND name ILIKE ?", authorEmail, "%"+name+"%").Find(&tickets)
	return tickets, result.Error
}

func (r *TicketRepository) GetTicketsByStatusAndName(status models.TicketStatus, name string) ([]models.Ticket, error) {
	var tickets []models.Ticket
	result := r.DB.Where("status = ? AND name ILIKE ?", status, "%"+name+"%").Find(&tickets)
	return tickets, result.Error
}

func (r *TicketRepository) GetTicketsByDateAndName(date, name string) ([]models.Ticket, error) {
	var tickets []models.Ticket
	result := r.DB.Where("DATE(created_at) = ? AND name ILIKE ?", date, "%"+name+"%").Find(&tickets)
	return tickets, result.Error
}
