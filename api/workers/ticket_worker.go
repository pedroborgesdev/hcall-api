package workers

import (
	"hcall/api/config"
	"hcall/api/models"
	"hcall/api/repository"
	"log"
	"time"
)

type TicketService struct {
	ticketRepo *repository.TicketRepository
	userRepo   *repository.UserRepository
	stopChan   chan bool
}

func NewTicketService() *TicketService {
	service := &TicketService{
		ticketRepo: repository.NewTicketRepository(),
		userRepo:   repository.NewUserRepository(),
		stopChan:   make(chan bool),
	}
	// Remove the automatic start
	// go service.StartTicketWorker()  <- Remove this line
	return service
}

func (s *TicketService) StartTicketWorker() {

	Status := config.AppConfig.WorkerTicketStatus
	Looptime := config.AppConfig.WorkerTicketLooptime
	RemoveAfter := config.AppConfig.WorkerTicketRemoveAfter

	ticker := time.NewTicker(time.Duration(Looptime) * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := s.RemoveTicketsWithStatus(Status, RemoveAfter); err != nil {
				log.Printf("Error removing tickets: %v", err)
			} else {
				log.Println("Successfully removed concluded tickets")
			}
		case <-s.stopChan:
			return
		}
	}
}

func (s *TicketService) RemoveTicketsWithStatus(status string, remove_after int) error {
	err := s.ticketRepo.RemoveTicketsWithStatus(models.TicketStatus(status), remove_after)
	if err != nil {
		return err
	}
	return nil
}

// Stop the scheduler when needed (e.g., during application shutdown)
func (s *TicketService) Stop() {
	s.stopChan <- true
}
