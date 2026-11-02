# System Design: API Design (REST, GraphQL, gRPC)

API (Application Programming Interface) design is crucial for how different software components or services communicate with each other. A well-designed API is easy to understand, use, and maintain, promoting interoperability and scalability.

## 1. REST (Representational State Transfer)

**Concept:** An architectural style for designing networked applications. It's based on a stateless, client-server communication model, using standard HTTP methods to perform operations on resources.

### Key Principles:
-   **Client-Server:** Separation of concerns.
-   **Stateless:** Each request from client to server must contain all the information needed to understand the request. The server does not store any client context between requests.
-   **Cacheable:** Responses must explicitly or implicitly define themselves as cacheable or non-cacheable.
-   **Uniform Interface:** Simplifies the overall system architecture.
    -   **Resource-Based:** Everything is a resource (e.g., `/users`, `/products/123`).
    -   **Standard Methods:** Uses HTTP verbs (GET, POST, PUT, DELETE, PATCH) for operations.
    -   **Self-descriptive Messages:** Messages include enough information to describe how to process them.
    -   **HATEOAS (Hypermedia as the Engine of Application State):** Resources contain links to other related resources, guiding the client through the application's state. (Often ignored in practice).
-   **Layered System:** Client cannot ordinarily tell whether it is connected directly to the end server, or to an intermediary.

### Advantages:
-   Simple to understand and implement.
-   Leverages existing HTTP infrastructure.
-   Widely adopted, good tooling and community support.
-   Human-readable.

### Disadvantages:
-   **Over-fetching/Under-fetching:** Clients often receive more data than needed (over-fetching) or need to make multiple requests to get all necessary data (under-fetching).
-   **Versioning:** Can be complex to manage API versions.
-   **Chatty:** Can require many round trips for complex data retrieval.

## 2. GraphQL

**Concept:** A query language for APIs and a runtime for fulfilling those queries with your existing data. It allows clients to request exactly the data they need, no more, no less.

### Key Principles:
-   **Declarative Data Fetching:** Clients specify the structure of the data they need.
-   **Single Endpoint:** Typically, a GraphQL API exposes a single HTTP endpoint (e.g., `/graphql`) that handles all queries and mutations.
-   **Strongly Typed Schema:** Defines the types of data and the relationships between them, providing a contract between client and server.
-   **Queries:** For reading data.
-   **Mutations:** For writing/modifying data.
-   **Subscriptions:** For real-time data updates.

### Advantages:
-   **Efficient Data Fetching:** Eliminates over-fetching and under-fetching.
-   **Reduced Round Trips:** Clients can get all necessary data in a single request.
-   **Strong Typing:** Provides validation and auto-completion for clients.
-   **Evolvable APIs:** Easy to add new fields and types without breaking existing clients.

### Disadvantages:
-   **Complexity:** Can be more complex to set up and manage than REST.
-   **Caching:** HTTP caching mechanisms are less effective with a single endpoint.
-   **Rate Limiting:** More challenging to implement than with REST.
-   **File Uploads:** Not natively supported, requires workarounds.

## 3. gRPC (Google Remote Procedure Call)

**Concept:** A modern, high-performance, open-source RPC framework that can run in any environment. It uses Protocol Buffers as its Interface Definition Language (IDL) and for message serialization.

### Key Principles:
-   **Protocol Buffers:** Language-agnostic, platform-agnostic, extensible mechanism for serializing structured data. Defines service methods and message types.
-   **HTTP/2:** Uses HTTP/2 for transport, enabling features like:
    -   **Multiplexing:** Multiple requests/responses over a single TCP connection.
    -   **Header Compression:** Reduces overhead.
    -   **Server Push:** Not directly used by gRPC, but HTTP/2 enables it.
-   **Streaming:** Supports four types of service methods:
    -   **Unary RPC:** Client sends a single request, server sends a single response.
    -   **Server Streaming RPC:** Client sends a single request, server sends a sequence of responses.
    -   **Client Streaming RPC:** Client sends a sequence of requests, server sends a single response.
    -   **Bidirectional Streaming RPC:** Both client and server send a sequence of messages using a read-write stream.

### Advantages:
-   **High Performance:** Due to Protocol Buffers (binary serialization) and HTTP/2.
-   **Strongly Typed:** Protocol Buffers enforce strict contracts.
-   **Multi-language Support:** Code generation for many languages.
-   **Streaming Capabilities:** Excellent for real-time communication.

### Disadvantages:
-   **Browser Support:** Not directly supported by browsers (requires a proxy).
-   **Learning Curve:** Steeper than REST.
-   **Human Readability:** Binary format is not human-readable.
-   **Tooling:** Ecosystem is growing but less mature than REST.

## Go Example: (Conceptual) API Interaction

See the conceptual example code in `api-design.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run api-design.go
```
