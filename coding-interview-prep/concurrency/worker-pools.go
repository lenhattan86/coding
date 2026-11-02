package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a task to be processed by a worker.
type Job struct {
	ID     int
	Number int
}

// Result represents the outcome of a processed job.
type Result struct {
	JobID  int
	Output int
	Error  error
}

// worker function that processes jobs from the jobs channel and sends results to the results channel.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d: Started processing job %d (Number: %d)\n", id, job.ID, job.Number)
		time.Sleep(time.Duration(job.Number) * 50 * time.Millisecond) // Simulate work
		output := job.Number * 2 // Simple processing
		fmt.Printf("Worker %d: Finished processing job %d (Output: %d)\n", id, job.ID, output)
		results <- Result{JobID: job.ID, Output: output, Error: nil}
	}
}

func main() {
	fmt.Println("--- Worker Pool Example ---")

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Number: j}
	}
	close(jobs) // Close the jobs channel to signal workers that no more jobs will be sent

	// Wait for all workers to finish
	wg.Wait()
	close(results) // Close the results channel after all workers are done

	// Collect and print results
	fmt.Println("\n--- All Jobs Processed. Results: ---")
	for res := range results {
		if res.Error != nil {
			fmt.Printf("Job %d failed: %v\n", res.JobID, res.Error)
		} else {
			fmt.Printf("Job %d: Output = %d\n", res.JobID, res.Output)
		}
	}

	fmt.Println("--- Worker Pool Example Finished ---")
}
