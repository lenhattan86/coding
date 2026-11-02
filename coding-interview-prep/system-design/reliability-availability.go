package main

import (
	"fmt"
	"sync"
	"time"
)

// Service represents a component or microservice in our system.
type Service struct {
	ID      int
	Name    string
	Healthy bool
	mu      sync.Mutex // Protects Healthy status
}

// CheckHealth simulates performing a health check on a service.
func (s *Service) CheckHealth() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Simulate some health check logic
	// For demonstration, let's say service 2 fails sometimes
	if s.ID == 2 && time.Now().Second()%5 == 0 { // Fails every 5 seconds mark
		s.Healthy = false
		fmt.Printf("Service %s (ID: %d) is now UNHEALTHY.\n", s.Name, s.ID)
	} else {
		s.Healthy = true
		// fmt.Printf("Service %s (ID: %d) is HEALTHY.\n", s.Name, s.ID)
	}
}

// MonitorServices continuously monitors the health of a list of services.
func MonitorServices(services []*Service, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("\n--- Performing Health Checks ---")
		for _, s := range services {
			s.CheckHealth()
			s.mu.Lock()
			if !s.Healthy {
				fmt.Printf("ALERT: Service %s (ID: %d) is DOWN!\n", s.Name, s.ID)
			} else {
				fmt.Printf("Service %s (ID: %d) is UP.\n", s.Name, s.ID)
			}
			s.mu.Unlock()
		}
	}
}

func main() {
	fmt.Println("--- Conceptual Health Check Example ---")

	// Define some services
	services := []*Service{
		{ID: 1, Name: "AuthService", Healthy: true},
		{ID: 2, Name: "PaymentService", Healthy: true},
		{ID: 3, Name: "UserService", Healthy: true},
	}

	// Start monitoring in a goroutine
	go MonitorServices(services, 1*time.Second)

	// Let the monitoring run for a while
	fmt.Println("Monitoring services for 10 seconds. Watch for alerts.")
	time.Sleep(10 * time.Second)

	fmt.Println("\n--- Health Check Example Finished ---")
}
