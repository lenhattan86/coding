package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Server represents a backend server that can handle requests.
type Server struct {
	ID      int
	Healthy bool
	Load    int // Simulate current load
}

// HandleRequest simulates a server processing a request.
func (s *Server) HandleRequest(requestID int) {
	if !s.Healthy {
		fmt.Printf("Server %d is unhealthy, cannot handle request %d.\n", s.ID, requestID)
		return
	}
	s.Load++
	fmt.Printf("Server %d handling request %d. Current load: %d\n", s.ID, requestID, s.Load)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate work
	s.Load--
	fmt.Printf("Server %d finished request %d. Current load: %d\n", s.ID, requestID, s.Load)
}

// LoadBalancer distributes requests to available servers.
type LoadBalancer struct {
	servers []*Server
	strategy string // e.g., "round-robin", "least-connections"
	nextServerIndex int // For round-robin
}

// NewLoadBalancer creates a new LoadBalancer with a list of servers.
func NewLoadBalancer(servers []*Server, strategy string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		strategy: strategy,
		nextServerIndex: 0,
	}
}

// RouteRequest routes an incoming request to a server based on the strategy.
func (lb *LoadBalancer) RouteRequest(requestID int) {
	var selectedServer *Server

	switch lb.strategy {
	case "round-robin":
		for i := 0; i < len(lb.servers); i++ {
			server := lb.servers[lb.nextServerIndex]
			lb.nextServerIndex = (lb.nextServerIndex + 1) % len(lb.servers)
			if server.Healthy {
				selectedServer = server
				break
			}
		}
	case "least-connections":
		minLoad := int(^uint(0) >> 1) // Max int value
		for _, server := range lb.servers {
			if server.Healthy && server.Load < minLoad {
				minLoad = server.Load
				selectedServer = server
			}
		}
	default:
		fmt.Println("Unknown load balancing strategy.")
		return
	}

	if selectedServer != nil {
		go selectedServer.HandleRequest(requestID) // Handle request in a goroutine
	} else {
		fmt.Printf("No healthy server available for request %d.\n", requestID)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("--- Conceptual Load Balancer Example ---")

	// Create some backend servers
	servers := []*Server{
		{ID: 1, Healthy: true},
		{ID: 2, Healthy: true},
		{ID: 3, Healthy: true},
	}

	// Create a load balancer with round-robin strategy
	lb := NewLoadBalancer(servers, "round-robin")

	// Send some requests
	fmt.Println("\n--- Round Robin Strategy ---")
	for i := 1; i <= 10; i++ {
		lb.RouteRequest(i)
		time.Sleep(100 * time.Millisecond)
	}

	// Simulate a server going down
	fmt.Println("\n--- Server 2 goes unhealthy ---")
	servers[1].Healthy = false // Server 2 (index 1)
	time.Sleep(1 * time.Second) // Give time for previous requests to settle

	// Send more requests
	for i := 11; i <= 15; i++ {
		lb.RouteRequest(i)
		time.Sleep(100 * time.Millisecond)
	}

	// Change to least-connections strategy
	fmt.Println("\n--- Switching to Least Connections Strategy ---")
	lb = NewLoadBalancer(servers, "least-connections") // Re-initialize with new strategy

	// Send more requests
	for i := 16; i <= 20; i++ {
		lb.RouteRequest(i)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(3 * time.Second) // Wait for all goroutines to finish
	fmt.Println("\n--- Load Balancer Example Finished ---")
}
