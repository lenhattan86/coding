package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// --- Conceptual 2PC (Two-Phase Commit) Simulation ---

// Participant represents a service or database involved in a 2PC transaction.
type Participant struct {
	ID      string
	CanCommit bool // Simulates whether the participant is ready to commit
	mu      sync.Mutex
}

// Prepare simulates the prepare phase for a participant.
func (p *Participant) Prepare() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond) // Simulate work
	if p.CanCommit {
		fmt.Printf("  Participant %s: Voted COMMIT\n", p.ID)
		return true
	}
	fmt.Printf("  Participant %s: Voted ABORT\n", p.ID)
	return false
}

// Commit simulates the commit phase for a participant.
func (p *Participant) Commit() {
	p.mu.Lock()
	defer p.mu.Unlock()
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond) // Simulate work
	fmt.Printf("  Participant %s: Committed transaction\n", p.ID)
}

// Abort simulates the abort phase for a participant.
func (p *Participant) Abort() {
	p.mu.Lock()
	defer p.mu.Unlock()
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond) // Simulate work
	fmt.Printf("  Participant %s: Aborted transaction\n", p.ID)
}

// Coordinator orchestrates the 2PC transaction.
func Coordinator(participants []*Participant) {
	fmt.Println("Coordinator: Starting 2PC transaction...")

	// Phase 1: Prepare
	allVotedCommit := true
	var wg sync.WaitGroup
	for _, p := range participants {
		wg.Add(1)
		go func(p *Participant) {
			defer wg.Done()
			if !p.Prepare() {
				allVotedCommit = false // Not thread-safe, but for conceptual demo
			}
		}(p)
	}
	wg.Wait()

	// Phase 2: Commit/Abort
	if allVotedCommit {
		fmt.Println("Coordinator: All participants voted COMMIT. Sending GLOBAL_COMMIT.")
		for _, p := range participants {
			wg.Add(1)
			go func(p *Participant) {
				defer wg.Done()
				p.Commit()
			}(p)
		}
	} else {
		fmt.Println("Coordinator: Some participant voted ABORT. Sending GLOBAL_ABORT.")
		for _, p := range participants {
			wg.Add(1)
			go func(p *Participant) {
				defer wg.Done()
				p.Abort()
			}(p)
		}
	}
	wg.Wait()
	fmt.Println("Coordinator: 2PC transaction finished.")
}

// --- Conceptual Saga Pattern Simulation ---

// SagaStep represents a local transaction in a saga.
type SagaStep struct {
	Name              string
	Execute           func() error
	Compensate        func() error
	ExecutedSuccessfully bool
}

// SagaOrchestrator manages the saga workflow.
func SagaOrchestrator(steps []SagaStep) {
	fmt.Println("\nSaga Orchestrator: Starting Saga transaction...")

	executedSteps := make([]SagaStep, 0)
	sagaFailed := false

	for i, step := range steps {
		fmt.Printf("Saga Orchestrator: Executing step %d - %s\n", i+1, step.Name)
		err := step.Execute()
		if err != nil {
			fmt.Printf("Saga Orchestrator: Step %d - %s FAILED: %v. Initiating compensation.\n", i+1, step.Name, err)
			sagaFailed = true
			break
		}
		steps[i].ExecutedSuccessfully = true
		executedSteps = append(executedSteps, steps[i])
	}

	if sagaFailed {
		fmt.Println("Saga Orchestrator: Compensating for failed saga...")
		// Compensate in reverse order of successful execution
		for i := len(executedSteps) - 1; i >= 0; i-- {
			step := executedSteps[i]
			if step.ExecutedSuccessfully {
				fmt.Printf("Saga Orchestrator: Compensating step - %s\n", step.Name)
				err := step.Compensate()
				if err != nil {
					fmt.Printf("Saga Orchestrator: Compensation for %s FAILED: %v\n", step.Name, err)
					// In a real system, this would trigger alerts/manual intervention
			}
		}
		}
		fmt.Println("Saga Orchestrator: Saga ABORTED and compensated.")
	} else {
		fmt.Println("Saga Orchestrator: Saga COMPLETED successfully.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("--- Conceptual Distributed Transactions Example ---")

	// --- 2PC Example ---
	fmt.Println("\n--- Two-Phase Commit (2PC) Simulation ---")
	p1 := &Participant{ID: "ServiceA", CanCommit: true}
	p2 := &Participant{ID: "ServiceB", CanCommit: true}
	p3 := &Participant{ID: "ServiceC", CanCommit: false} // This participant will cause an abort

	participants := []*Participant{p1, p2, p3}
	Coordinator(participants)

	// Reset for another 2PC run
	p1.CanCommit = true
	p2.CanCommit = true
	p3.CanCommit = true // All will commit this time
	fmt.Println("\n--- Two-Phase Commit (2PC) Simulation (All Commit) ---")
	Coordinator(participants)

	// --- Saga Example ---
	fmt.Println("\n--- Saga Pattern Simulation ---")

	// Define saga steps
	step1 := SagaStep{
		Name: "Order Service: Create Order",
		Execute: func() error {
			fmt.Println("  Order Service: Order created.")
			return nil
		},
		Compensate: func() error {
			fmt.Println("  Order Service: Order creation rolled back.")
			return nil
		},
	}

	step2 := SagaStep{
		Name: "Payment Service: Process Payment",
		Execute: func() error {
			// Simulate payment failure 50% of the time
			if rand.Intn(100) < 50 {
				fmt.Println("  Payment Service: Payment FAILED.")
				return fmt.Errorf("payment failed")
			}
			fmt.Println("  Payment Service: Payment processed.")
			return nil
		},
		Compensate: func() error {
			fmt.Println("  Payment Service: Payment refunded.")
			return nil
		},
	}

	step3 := SagaStep{
		Name: "Inventory Service: Deduct Stock",
		Execute: func() error {
			fmt.Println("  Inventory Service: Stock deducted.")
			return nil
		},
		Compensate: func() error {
			fmt.Println("  Inventory Service: Stock added back.")
			return nil
		},
	}

	sagaSteps := []SagaStep{step1, step2, step3}
	SagaOrchestrator(sagaSteps)

	fmt.Println("\n--- Distributed Transactions Example Finished ---")
}
