# System Design: Databases (SQL vs NoSQL)

Databases are fundamental components of almost every software system, responsible for storing, managing, and retrieving data. Choosing the right database is a critical system design decision, heavily influenced by the application's requirements for data structure, scalability, consistency, and availability.

## SQL Databases (Relational Databases)

**Concept:** SQL (Structured Query Language) databases are based on the relational model, where data is organized into tables with predefined schemas. Each table consists of rows and columns, and relationships between tables are established using foreign keys.

### Characteristics:
-   **Structured Data:** Data is stored in a highly organized, tabular format with a fixed schema.
-   **ACID Properties:** Typically adhere to ACID (Atomicity, Consistency, Isolation, Durability) properties, ensuring strong transactional consistency.
-   **Vertical Scalability:** Traditionally scale up (more powerful server) rather than out (more servers). Horizontal scaling is possible but often more complex (e.g., sharding).
-   **Joins:** Support complex queries involving joins across multiple tables.
-   **Mature Ecosystem:** Well-established, widely used, and robust.

### Use Cases:
-   Financial systems (banking, accounting).
-   E-commerce platforms (orders, inventory).
-   Applications requiring complex transactions and strong data integrity.
-   Systems with clearly defined relationships between data entities.

### Examples:
-   MySQL
-   PostgreSQL
-   Oracle
-   SQL Server

## NoSQL Databases (Non-relational Databases)

**Concept:** NoSQL (Not only SQL) databases provide a mechanism for storage and retrieval of data that is modeled in means other than the tabular relations used in relational databases. They are designed for specific data models and have flexible schemas.

### Characteristics:
-   **Flexible Schema:** Can handle unstructured, semi-structured, and polymorphic data.
-   **BASE Properties:** Often prioritize BASE (Basically Available, Soft state, Eventual consistency) properties over strict ACID, favoring availability and partition tolerance.
-   **Horizontal Scalability:** Designed to scale out easily across many servers, handling large volumes of data and high traffic.
-   **No Joins (or limited):** Typically do not support complex joins across different data collections; relationships are often handled at the application level.
-   **Diverse Models:** Come in various types, each optimized for different use cases.

### Types of NoSQL Databases:

1.  **Key-Value Stores:**
    -   **Concept:** Simplest NoSQL model. Data is stored as a collection of key-value pairs.
    -   **Use Cases:** Caching, session management, user profiles.
    -   **Examples:** Redis, Memcached, DynamoDB (can also be document).

2.  **Document Databases:**
    -   **Concept:** Store data in flexible, semi-structured documents (e.g., JSON, BSON, XML). Documents can be nested and don't require a predefined schema.
    -   **Use Cases:** Content management, blogging platforms, e-commerce catalogs.
    -   **Examples:** MongoDB, Couchbase, Firestore.

3.  **Column-Family Stores (Wide-Column Stores):**
    -   **Concept:** Store data in tables, rows, and dynamic columns. Optimized for large datasets and high write throughput.
    -   **Use Cases:** Big data analytics, time-series data, event logging.
    -   **Examples:** Cassandra, HBase.

4.  **Graph Databases:**
    -   **Concept:** Store data in terms of nodes (entities) and edges (relationships) between them. Optimized for traversing relationships.
    -   **Use Cases:** Social networks, recommendation engines, fraud detection.
    -   **Examples:** Neo4j, Amazon Neptune.

## Choosing Between SQL and NoSQL:

| Feature           | SQL Databases                               | NoSQL Databases                                     |
| :---------------- | :------------------------------------------ | :-------------------------------------------------- |
| **Schema**        | Fixed, predefined, rigid                    | Dynamic, flexible, schema-less                      |
| **Scalability**   | Primarily vertical, horizontal (complex)    | Primarily horizontal (easy)                         |
| **Transactions**  | ACID compliant (strong consistency)         | BASE compliant (eventual consistency)               |
| **Data Model**    | Relational (tables, rows, columns)          | Key-value, Document, Column-family, Graph           |
| **Relationships** | Strong relationships (joins)                | Weak relationships (application-managed)            |
| **Use Cases**     | Complex transactions, strong data integrity | Big data, real-time web apps, flexible data needs   |

## Go Example: (Conceptual) Database Interaction

See the conceptual example code in `databases.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run databases.go
```
