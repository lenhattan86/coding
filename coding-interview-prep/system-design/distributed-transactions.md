# System Design: Distributed Transactions (2PC & Saga)

In distributed systems, where operations span across multiple independent services or databases, ensuring data consistency and atomicity becomes significantly more complex than in a monolithic system with a single database. This is where **distributed transactions** come into play.

## Key Challenges:
-   **Partial Failures:** One service might succeed while another fails.
-   **Network Latency:** Communication delays between services.
-   **Concurrency:** Multiple transactions trying to modify the same data.
-   **Isolation:** Ensuring that concurrent distributed transactions don't interfere with each other.

## 1. Two-Phase Commit (2PC)

**Concept:** A distributed algorithm that coordinates all participating nodes in a distributed transaction to either commit or abort the transaction as a single atomic unit. It ensures that all participants agree on the outcome.

### Phases:

1.  **Phase 1: Commit-Request (Vote) Phase:**
    -   The **Coordinator** (transaction manager) sends a `prepare` message to all **Participants** (services/databases involved).
    -   Each Participant:
        -   Performs the transaction's operations locally.
        -   Writes all changes to a temporary log (undo/redo logs).
        -   If it can successfully commit, it replies `VOTE_COMMIT` to the Coordinator.
        -   If it cannot, it replies `VOTE_ABORT` and rolls back its local changes.
2.  **Phase 2: Commit (Decision) Phase:**
    -   **If all Participants voted `VOTE_COMMIT`:**
        -   The Coordinator sends a `GLOBAL_COMMIT` message to all Participants.
        -   Each Participant commits its local changes and replies `ACK`.
    -   **If any Participant voted `VOTE_ABORT` (or timed out):**
        -   The Coordinator sends a `GLOBAL_ABORT` message to all Participants.
        -   Each Participant rolls back its local changes and replies `ACK`.
    -   The Coordinator waits for all `ACK`s before completing the transaction.

### Advantages:
-   Provides strong consistency (ACID properties) across distributed services.
-   Relatively simple to understand conceptually.

### Disadvantages:
-   **Blocking:** If the Coordinator fails during Phase 2, Participants might be left in an uncertain (prepared) state, blocking resources until the Coordinator recovers. This is its biggest drawback.
-   **Performance:** High latency due to multiple network round-trips and synchronous nature.
-   **Single Point of Failure:** The Coordinator is a single point of failure.

## 2. Saga Pattern

**Concept:** A sequence of local transactions, where each local transaction updates the database and publishes an event to trigger the next local transaction in the saga. If a local transaction fails, the saga executes a series of **compensating transactions** to undo the changes made by previous successful local transactions.

### Types of Saga Implementations:

1.  **Choreography-based Saga:**
    -   Each service produces and consumes events, deciding when to execute its local transaction and publish the next event.
    -   **Pros:** Decoupled services, no single point of failure.
    -   **Cons:** Can be complex to manage, difficult to track the overall flow, potential for circular dependencies.
2.  **Orchestration-based Saga:**
    -   A dedicated **Saga Orchestrator** (a service) manages the entire workflow. It sends commands to participants, executes local transactions, and handles compensating transactions if needed.
    -   **Pros:** Centralized control, easier to monitor and manage, better error handling.
    -   **Cons:** Orchestrator can become a single point of failure (though it can be made highly available), increased complexity in the orchestrator.

### Advantages of Saga:
-   **Avoids Blocking:** Unlike 2PC, participants don't hold locks for the entire duration of the distributed transaction.
-   **Improved Availability:** Services can continue to operate even if other services are temporarily unavailable.
-   **Better Scalability:** More suitable for microservices architectures.

### Disadvantages of Saga:
-   **Eventual Consistency:** Data consistency is eventual, not immediate.
-   **Complexity:** More complex to implement than 2PC, especially compensating transactions.
-   **Testing:** Harder to test the entire distributed workflow.
-   **Lack of Isolation:** Requires careful design to handle concurrent updates and ensure business consistency.

## Go Example: (Conceptual) Distributed Transaction Simulation

See the conceptual example code in `distributed-transactions.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run distributed-transactions.go
```
