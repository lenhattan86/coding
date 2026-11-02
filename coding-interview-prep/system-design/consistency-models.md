# System Design: Consistency Models

In distributed systems, where data is replicated across multiple nodes, **consistency models** define the rules for how data updates are propagated and observed by different clients. They dictate the guarantees a system provides regarding the visibility and ordering of updates.

## Key Concepts:

-   **Replication:** Storing multiple copies of data across different nodes to improve availability and durability.
-   **Distributed System:** A system where components are located on different networked computers and communicate by passing messages.

## The CAP Theorem

The **CAP Theorem** (Consistency, Availability, Partition Tolerance) is a fundamental concept in distributed systems. It states that a distributed data store can only guarantee two out of three properties simultaneously:

1.  **Consistency (C):** All clients see the same data at the same time. After a write, any subsequent read will return that write or a more recent write. (Think of a single, up-to-date copy of the data).
2.  **Availability (A):** Every request receives a response, without guarantee that it contains the most recent write. The system remains operational even if some nodes fail.
3.  **Partition Tolerance (P):** The system continues to operate despite network partitions (communication failures between nodes).

**Implication:** In the presence of a network partition (which is inevitable in a distributed system), you must choose between Consistency and Availability.

-   **CP System:** Prioritizes Consistency and Partition Tolerance. If a partition occurs, the system will become unavailable (or return an error) to ensure data consistency. (e.g., traditional relational databases, ZooKeeper, etcd).
-   **AP System:** Prioritizes Availability and Partition Tolerance. If a partition occurs, the system remains available but might return stale data. (e.g., Cassandra, DynamoDB, CouchDB).
-   **CA System:** Prioritizes Consistency and Availability. This is only possible in the absence of network partitions, which is not realistic for truly distributed systems.

## ACID vs. BASE

These are two contrasting sets of properties for database transactions, often associated with relational (ACID) and NoSQL (BASE) databases, respectively.

### ACID (Atomicity, Consistency, Isolation, Durability)

Properties that guarantee that database transactions are processed reliably. Typically associated with traditional relational databases.

-   **Atomicity:** A transaction is treated as a single, indivisible unit of operations. Either all of its operations are completed, or none are.
-   **Consistency:** A transaction brings the database from one valid state to another. It ensures that any data written to the database must be valid according to all defined rules (constraints, triggers, cascades).
-   **Isolation:** The concurrent execution of multiple transactions results in a system state that would be achieved if the transactions were executed sequentially. Transactions appear to run in isolation from each other.
-   **Durability:** Once a transaction has been committed, it will remain committed even in the event of power loss, crashes, or errors.

### BASE (Basically Available, Soft state, Eventual consistency)

Properties that are often found in highly available, eventually consistent NoSQL databases.

-   **Basically Available:** The system guarantees availability (as per CAP). It will respond to any request, even if it cannot guarantee the most recent version of the information.
-   **Soft state:** The state of the system may change over time, even without input, due to eventual consistency.
-   **Eventual consistency:** Given enough time, provided no new updates are made, all updates will propagate through the system, and all replicas will eventually be consistent.

## Other Consistency Models:

-   **Strong Consistency:** All replicas are updated synchronously. All reads see the latest write. (e.g., Linearizability, Sequential Consistency).
-   **Eventual Consistency:** As described in BASE. Replicas eventually converge to the same state.
-   **Causal Consistency:** If process A has seen an update from process B, then process A will not see an older state of process B than what it has already seen.
-   **Read-Your-Writes Consistency:** If a process writes a value, it will always see its own write in subsequent reads.
-   **Monotonic Reads:** If a process reads a value, any subsequent reads by that process will return the same value or a more recent value.
-   **Monotonic Writes:** A write operation is guaranteed to complete before any subsequent write operation by the same process.

## Go Example: (Conceptual) Eventual Consistency Simulation

See the conceptual example code in `consistency-models.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run consistency-models.go
```
