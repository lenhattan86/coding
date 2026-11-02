# Concurrency Patterns: Pipelines

A **pipeline** is a common concurrency pattern in Go that involves chaining together multiple stages of goroutines, where the output of one stage becomes the input to the next. Each stage typically performs a specific transformation or processing step on the data. Channels are used to connect these stages, acting as the "pipes" through which data flows.

## Key Concepts:
-   **Stages:** Independent goroutines, each responsible for a specific part of the processing.
-   **Channels:** Used to pass data between stages. Each channel acts as a buffer and a synchronization point.
-   **Fan-out/Fan-in:**
    -   **Fan-out:** Multiple goroutines read from the same input channel, distributing the work.
    -   **Fan-in:** Multiple goroutines write to a single output channel, combining results.
-   **Cancellation:** Often managed with `context.Context` to gracefully shut down all stages when an error occurs or the pipeline is no longer needed.

## How Pipelines Work:
1.  **Source Stage:** Generates initial data and sends it to the first channel.
2.  **Processing Stages:** Each intermediate stage reads from its input channel, performs some work, and writes its output to an output channel.
3.  **Sink Stage (or Consumer):** Reads the final processed data from the last channel.
4.  **Channel Closure:** It's crucial to close channels when no more data will be sent to them. This signals to downstream stages that they can stop reading and eventually exit.

## Advantages:
-   **Modularity:** Each stage is a self-contained unit, making the system easier to understand, test, and maintain.
-   **Concurrency:** Different stages can run concurrently, improving throughput.
-   **Scalability:** Stages can be scaled independently (e.g., by using fan-out for CPU-bound tasks).
-   **Resource Efficiency:** Data flows through the pipeline, reducing the need to store large intermediate results.

## Use Cases:
-   **Data Processing:** ETL (Extract, Transform, Load) operations.
-   **Image/Video Processing:** Applying a series of filters or transformations.
-   **Log Processing:** Parsing, filtering, and analyzing log entries.
-   **Web Request Handling:** A sequence of middleware or handlers.

## Go Example: Simple Data Processing Pipeline

See the example code in `pipelines.go`.
```go
// To run this code:
// 1. Navigate to the concurrency directory: cd coding-interview-prep/concurrency
// 2. Run the file: go run pipelines.go
```
