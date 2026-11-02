# System Design: Load Balancing Algorithms

Load balancing is the process of distributing network traffic across multiple servers to ensure no single server becomes a bottleneck. It improves the responsiveness and availability of applications by optimizing resource utilization, maximizing throughput, minimizing response time, and avoiding overload of any single resource.

## Key Concepts:
-   **Load Balancer:** A device or software that sits in front of a group of servers and distributes client requests across them.
-   **Backend Servers (Pool):** The group of servers that process the actual requests.
-   **Health Checks:** Mechanisms used by the load balancer to determine if a backend server is capable of handling requests. Unhealthy servers are temporarily removed from the pool.
-   **Session Persistence (Sticky Sessions):** Directing a client's requests to the same server for the duration of their session. This is important for applications that store session-specific data on the server.

## Load Balancing Algorithms:

### 1. Round Robin
-   **Concept:** Distributes client requests to servers in a sequential, rotating manner. Each server gets a turn.
-   **Pros:** Simple to implement, fair distribution when servers have similar processing capabilities and requests are of similar weight.
-   **Cons:** Does not consider server load or capacity. A heavily loaded server will still receive new requests.

### 2. Weighted Round Robin
-   **Concept:** Similar to Round Robin, but assigns a "weight" to each server. Servers with higher weights receive a larger proportion of requests.
-   **Pros:** Allows administrators to account for differences in server capacity.
-   **Cons:** Still doesn't dynamically adjust to real-time server load.

### 3. Least Connections
-   **Concept:** Directs new client requests to the server with the fewest active connections.
-   **Pros:** Effective for situations where requests vary in processing time, as it tries to keep the load balanced based on current activity.
-   **Cons:** Requires the load balancer to actively track connections, which adds overhead.

### 4. Weighted Least Connections
-   **Concept:** A combination of Weighted Round Robin and Least Connections. Directs requests to the server with the fewest active connections relative to its assigned weight.
-   **Pros:** Balances load based on both server capacity and current activity.

### 5. IP Hash
-   **Concept:** Uses a hash of the client's IP address to determine which server receives the request.
-   **Pros:** Provides session persistence without needing to store session information on the load balancer. All requests from a specific client IP will go to the same server.
-   **Cons:** If a server fails, all clients hashed to that server will be affected. Uneven distribution if client IPs are not uniformly distributed.

### 6. Least Response Time
-   **Concept:** Directs requests to the server with the fewest active connections and the lowest average response time.
-   **Pros:** Optimizes for performance by sending requests to the fastest available server.
-   **Cons:** More complex to implement, requires continuous monitoring of server response times.

### 7. Source IP Affinity (or Session Affinity)
-   **Concept:** Similar to IP Hash, but can use other client-specific information (e.g., cookies, SSL session IDs) to ensure subsequent requests from the same client go to the same server.
-   **Pros:** Guarantees session persistence.
-   **Cons:** Can lead to uneven load distribution if some clients are very active.

## Go Example: (Conceptual) Load Balancer with Algorithms

See the conceptual example code in `load-balancing-algorithms.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run load-balancing-algorithms.go
```
