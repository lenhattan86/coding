# System Design: Caching Strategies

Caching is a fundamental technique in system design used to improve performance and reduce the load on backend services by storing frequently accessed data in a faster, temporary storage layer.

## Key Concepts:
-   **Cache Hit:** When requested data is found in the cache.
-   **Cache Miss:** When requested data is not found in the cache and must be fetched from the original source.
-   **Cache Invalidation:** The process of removing or updating stale data in the cache. This is one of the hardest problems in computer science.
-   **Time-To-Live (TTL):** A duration after which cached data is considered stale and should be re-fetched or removed.
-   **Eviction Policy:** Rules for deciding which data to remove from the cache when it's full.

## Where to Cache:

1.  **Client-side Cache (Browser Cache):** Stores static assets (images, CSS, JS) and API responses on the user's device.
2.  **CDN (Content Delivery Network):** Geographically distributed servers that cache static and dynamic content closer to users, reducing latency.
3.  **Web Server Cache (Reverse Proxy Cache):** Caches responses at the web server level (e.g., Nginx, Varnish) before they reach the application server.
4.  **Application Cache:** Implemented within the application code (e.g., in-memory caches like `map` in Go, or dedicated libraries).
5.  **Database Cache:** Caching layers provided by the database itself (e.g., query cache, buffer pool).
6.  **Distributed Cache:** A separate layer of caching servers (e.g., Redis, Memcached) that can be accessed by multiple application instances.

## Caching Strategies (How to manage data between cache and database):

1.  **Cache-Aside (Lazy Loading):**
    -   **Read:** Application checks cache first. If data is present (cache hit), return it. If not (cache miss), fetch from database, store in cache, then return.
    -   **Write:** Application writes directly to the database, then invalidates (deletes) the corresponding entry in the cache.
    -   **Pros:** Simple to implement, only caches requested data, resilient to cache failures.
    -   **Cons:** Cache misses can be slow (two round trips), data can be stale between write and invalidation.

2.  **Write-Through:**
    -   **Read:** Same as Cache-Aside.
    -   **Write:** Application writes data to the cache, and the cache synchronously writes the data to the database.
    -   **Pros:** Data in cache is always consistent with the database, simpler write logic.
    -   **Cons:** Higher write latency (due to synchronous write to DB), cache can contain unread data.

3.  **Write-Back (Write-Behind):**
    -   **Read:** Same as Cache-Aside.
    -   **Write:** Application writes data to the cache, and the cache asynchronously writes the data to the database.
    -   **Pros:** Very low write latency, high write throughput.
    -   **Cons:** Data loss risk if cache fails before data is written to DB, more complex to implement.

4.  **Read-Through:**
    -   **Read:** Application requests data from the cache. If data is not in cache, the cache itself is responsible for fetching it from the database, storing it, and returning it.
    -   **Write:** Typically combined with Write-Through or Write-Back.
    -   **Pros:** Application code is simpler (doesn't deal with cache misses), cache acts as a data source.
    -   **Cons:** Cache needs to know how to load data from the database.

## Cache Eviction Policies (When cache is full):

-   **LRU (Least Recently Used):** Discards the least recently used items first. (Most common)
-   **LFU (Least Frequently Used):** Discards the least frequently used items first.
-   **FIFO (First In, First Out):** Discards the oldest items first.
-   **LIFO (Last In, First Out):** Discards the newest items first.
-   **Random Replacement:** Randomly discards an item.

## Go Example: (Conceptual) Cache-Aside Implementation

See the conceptual example code in `caching-strategies.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run caching-strategies.go
```
