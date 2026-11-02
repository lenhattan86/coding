# System Design: Basics & Scalability

System design interviews assess your ability to design complex software systems. It's about understanding trade-offs, making informed decisions, and communicating your design effectively.

## What is System Design?
System design is the process of defining the architecture, modules, interfaces, and data for a system to satisfy specified requirements. It's about building a robust, scalable, reliable, and maintainable system.

## Key Pillars of System Design:

1.  **Scalability:** The ability of a system to handle a growing amount of work by adding resources.
2.  **Reliability:** The probability that a system will perform its intended function without failure for a specified period.
3.  **Availability:** The proportion of time a system is in a functioning condition.
4.  **Efficiency:** How well a system uses its resources (CPU, memory, network, disk).
5.  **Maintainability:** How easy it is to operate, monitor, fix, and upgrade the system.

## Scalability

Scalability is a critical aspect of modern systems, allowing them to handle increased load and data.

### Types of Scaling:

1.  **Vertical Scaling (Scale Up):**
    -   **Concept:** Increasing the capacity of a single server (e.g., adding more CPU, RAM, faster disk).
    -   **Pros:** Simpler to implement, less complex to manage.
    -   **Cons:** Limited by the maximum capacity of a single machine, single point of failure, eventually more expensive.

2.  **Horizontal Scaling (Scale Out):**
    -   **Concept:** Adding more servers to distribute the load.
    -   **Pros:** Highly scalable, no single point of failure (if designed well), can be more cost-effective in the long run.
    -   **Cons:** More complex to implement (distributed systems challenges), requires load balancing, data partitioning, and distributed coordination.

### Key Techniques for Horizontal Scaling:

-   **Load Balancing:** Distributes incoming network traffic across multiple servers.
    -   **Types:** DNS-based, Hardware-based, Software-based (e.g., Nginx, HAProxy).
    -   **Algorithms:** Round Robin, Least Connections, IP Hash.
-   **Database Sharding/Partitioning:** Dividing a large database into smaller, more manageable pieces (shards) across multiple database servers.
    -   **Types:** Horizontal (row-based), Vertical (column-based), Directory-based.
    -   **Challenges:** Data consistency, rebalancing, complex queries.
-   **Caching:** Storing frequently accessed data in a faster, temporary storage layer (e.g., Redis, Memcached) to reduce load on the primary data store.
    -   **Levels:** Client-side, CDN, Web Server, Database, Application.
    -   **Strategies:** Cache-aside, Write-through, Write-back.
-   **Content Delivery Networks (CDNs):** Geographically distributed network of proxy servers and their data centers. The goal is to provide high availability and performance by distributing the service spatially relative to end-users.
-   **Message Queues:** Decouple services, buffer requests, and handle asynchronous communication (e.g., Kafka, RabbitMQ, SQS).
-   **Microservices Architecture:** Breaking down a monolithic application into smaller, independent services that can be developed, deployed, and scaled independently.

## Go Example: (Conceptual) Load Balancer

See the conceptual example code in `basics-scalability.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run basics-scalability.go
```
