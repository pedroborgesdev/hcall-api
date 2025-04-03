package utils

import (
	"hcall/api/workers"
	"sync"
)

var (
	manager *WorkerManager
	once    sync.Once
)

type WorkerManager struct {
	wg      sync.WaitGroup
	workers map[string]*workers.TicketService
	mu      sync.Mutex
}

func GetWorkerManager() *WorkerManager {
	once.Do(func() {
		manager = &WorkerManager{
			workers: make(map[string]*workers.TicketService),
		}
	})
	return manager
}

func (wm *WorkerManager) StartAllWorkers() {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	// Start ticket worker
	ticketService := workers.NewTicketService()
	wm.workers["ticket"] = ticketService
	wm.wg.Add(1)
	go func() {
		defer wm.wg.Done()
		ticketService.StartTicketWorker()
	}()
}

func (wm *WorkerManager) StopAllWorkers() {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	// Stop all workers
	for _, worker := range wm.workers {
		worker.Stop()
	}
	wm.wg.Wait()
}
