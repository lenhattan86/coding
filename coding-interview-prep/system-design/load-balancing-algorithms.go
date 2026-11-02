package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// BackendServer represents a backend server that can handle requests.
type BackendServer struct {
	ID      int
	Healthy bool
	ActiveConnections int32 // Using int32 for atomic operations
	mu      sync.Mutex // Protects other fields if needed
	Weight  int // For weighted algorithms
}

// HandleRequest simulates a server processing a request.
func (s *BackendServer) HandleRequest(requestID int) {
	if !s.Healthy {
		fmt.Printf("Server %d is unhealthy, cannot handle request %d.\n", s.ID, requestID)
		return
	}
	atomic.AddInt32(&s.ActiveConnections, 1)
	fmt.Printf("Server %d: Handling request %d. Connections: %d\n", s.ID, requestID, atomic.LoadInt32(&s.ActiveConnections))
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate work
	atomic.AddInt32(&s.ActiveConnections, -1)
	fmt.Printf("Server %d: Finished request %d. Connections: %d\n", s.ID, requestID, atomic.LoadInt32(&s.ActiveConnections))
}

// LoadBalancer distributes requests to available servers.
type LoadBalancer struct {
	servers []*BackendServer
	strategy string
	rrIndex int // For round-robin
	mu      sync.Mutex // Protects rrIndex and server list modifications
}

// NewLoadBalancer creates a new LoadBalancer.
func NewLoadBalancer(servers []*BackendServer, strategy string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		strategy: strategy,
		rrIndex: 0,
	}
}

// RouteRequest routes an incoming request to a server based on the strategy.
func (lb *LoadBalancer) RouteRequest(requestID int) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	var selectedServer *BackendServer
	healthyServers := []*BackendServer{}
	for _, s := range lb.servers {
		if s.Healthy {
			healthyServers = append(healthyServers, s)
		}
	}

	if len(healthyServers) == 0 {
		fmt.Printf("No healthy server available for request %d.\n", requestID)
		return
	}

	switch lb.strategy {
	case "round-robin":
		selectedServer = healthyServers[lb.rrIndex%len(healthyServers)]
		lb.rrIndex = (lb.rrIndex + 1) % len(healthyServers)
	case "least-connections":
		minConnections := int32(^uint32(0) >> 1) // Max int32 value
		for _, s := range healthyServers {
			currentConnections := atomic.LoadInt32(&s.ActiveConnections)
			if currentConnections < minConnections {
				minConnections = currentConnections
				selectedServer = s
			}
		}
	case "weighted-round-robin":
		// Simplified weighted round robin for demonstration
		// In a real scenario, this would involve a more complex stateful algorithm
		// to ensure weights are respected over time.
		totalWeight := 0
		for _, s := range healthyServers {
			totalWeight += s.Weight
		}
		if totalWeight == 0 { // Fallback if no weights or all 0
			selectedServer = healthyServers[lb.rrIndex%len(healthyServers)]
			lb.rrIndex = (lb.rrIndex + 1) % len(healthyServers)
			break
		}

		pick := rand.Intn(totalWeight)
		current := 0
		for _, s := range healthyServers {
			current += s.Weight
			if pick < current {
				selectedServer = s
				break
			}
		}
	default:
		fmt.Printf("Unknown load balancing strategy: %s for request %d.\n", lb.strategy, requestID)
		return
	}

	if selectedServer != nil {
		go selectedServer.HandleRequest(requestID)
	} else {
		fmt.Printf("Failed to select a server for request %d.\n", requestID)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("--- Conceptual Load Balancing Algorithms Example ---")

	// Create some backend servers with weights
	servers := []*BackendServer{
		{ID: 1, Healthy: true, Weight: 3}, // More powerful
		{ID: 2, Healthy: true, Weight: 1}, // Less powerful
		{ID: 3, Healthy: true, Weight: 2},
	}

	// --- Round Robin ---
	fmt.Println("\n--- Round Robin Strategy ---")
	lbRR := NewLoadBalancer(servers, "round-robin")
	for i := 1; i <= 6; i++ {
		lbRR.RouteRequest(i)
		time.Sleep(50 * time.Millisecond)
	}

	// --- Least Connections ---
	fmt.Println("\n--- Least Connections Strategy ---")
	lbLC := NewLoadBalancer(servers, "least-connections")
	// Reset server connections for a clean test
	for _, s := range servers {
		atomic.StoreInt32(&s.ActiveConnections, 0)
	}
	for i := 7; i <= 12; i++ {
		lbLC.RouteRequest(i)
		time.Sleep(50 * time.Millisecond)
	}

	// --- Weighted Round Robin ---
	fmt.Println("\n--- Weighted Round Robin Strategy ---")
	lbWRR := NewLoadBalancer(servers, "weighted-round-robin")
	// Reset server connections for a clean test
	for _, s := range servers {
		atomic.StoreInt32(&s.ActiveConnections, 0)
	}
	for i := 13; i <= 20; i++ {
		lbWRR.RouteRequest(i)
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(2 * time.Second) // Wait for all goroutines to finish
	fmt.Println("\n--- Load Balancing Algorithms Example Finished ---")
}
