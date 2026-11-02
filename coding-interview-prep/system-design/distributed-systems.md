# System Design: Distributed Systems Concepts

Distributed systems are collections of independent computers that appear to the users of the system as a single coherent system. They are designed to solve problems that are too large or complex for a single machine, or to provide higher availability and fault tolerance.

## Key Challenges in Distributed Systems:
-   **Concurrency:** Multiple components operating simultaneously.
-   **Partial Failures:** Components can fail independently.
-   **Network Latency & Unreliability:** Communication takes time and can fail.
-   **Data Consistency:** Ensuring all copies of data are synchronized.
-   **Clock Skew:** Different machines having slightly different times.

## Core Concepts:

### 1. Remote Procedure Call (RPC)
-   **Concept:** A protocol that allows a program on one computer to execute a procedure (subroutine) on a remote computer as if it were a local procedure call.
-   **How it works:** The client sends a request message to the server, and the server sends a response message back to the client. This is typically synchronous.
-   **Advantages:** Simplifies distributed programming by abstracting network communication.
-   **Disadvantages:** Tightly coupled, synchronous calls can lead to blocking, error handling can be complex.
-   **Examples:** gRPC, Apache Thrift, traditional SOAP/XML-RPC.

### 2. Message Queues
-   **Concept:** A form of asynchronous service-to-service communication used in serverless and microservices architectures. Messages are stored in a queue until they are processed by a consumer.
-   **How it works:** Producers send messages to a queue, and consumers retrieve messages from the queue. The queue acts as a buffer and decouples producers from consumers.
-   **Advantages:**
    -   **Decoupling:** Producers and consumers don't need to be available at the same time.
    -   **Buffering:** Handles spikes in traffic, preventing system overload.
    -   **Asynchronous Communication:** Improves responsiveness.
    -   **Reliability:** Messages can be persisted until successfully processed.
-   **Disadvantages:** Adds complexity, potential for message ordering issues, operational overhead.
-   **Examples:** Apache Kafka, RabbitMQ, Amazon SQS, Google Cloud Pub/Sub.

### 3. Consensus Algorithms
-   **Concept:** Algorithms that allow a group of distributed processes to agree on a single data value or state. This is crucial for maintaining consistency in replicated state machines.
-   **How it works:** Processes propose values, communicate with each other, and eventually agree on one value, even in the presence of failures.
-   **Advantages:** Enables fault-tolerant distributed systems, ensures data consistency across replicas.
-   **Disadvantages:** Complex to implement and understand, can introduce latency.
-   **Examples:**
    -   **Paxos:** One of the earliest and most influential consensus algorithms, but notoriously difficult to understand and implement.
    -   **Raft:** Designed to be more understandable than Paxos, while still providing strong consistency guarantees. It achieves consensus by electing a leader and replicating a log.

### 4. Distributed Transactions
-   **Concept:** A transaction that involves multiple independent participants (e.g., databases, services) and must either commit or abort across all of them.
-   **Challenges:** Ensuring atomicity across distributed components is difficult due to partial failures and network issues.
-   **Solutions:**
    -   **Two-Phase Commit (2PC):** A common protocol, but it's blocking (if the coordinator fails, participants can be stuck).
    -   **Saga Pattern:** A sequence of local transactions, where each local transaction updates the database and publishes an event to trigger the next local transaction. If a step fails, compensating transactions are executed to undo the changes.

## Go Example: (Conceptual) RPC Client/Server

See the conceptual example code in `distributed-systems.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run distributed-systems.go
```
