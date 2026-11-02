# System Design: Database Indexing

Database indexing is a data structuring technique used to quickly locate and access data in a database. An index is a special lookup table that the database search engine can use to speed up data retrieval. Think of it like the index in a book: instead of reading the entire book to find a topic, you look up the topic in the index and go directly to the relevant page.

## Key Concepts:
-   **Index Key:** The column(s) on which the index is built.
-   **Index Structure:** How the index data is organized (e.g., B-tree, Hash).
-   **Clustered Index:** Determines the physical order of data rows in the table. A table can have only one clustered index.
-   **Non-Clustered Index:** Does not alter the physical order of table rows. It contains the index key values and pointers to the actual data rows. A table can have multiple non-clustered indexes.

## Why Use Indexes?
-   **Faster Data Retrieval:** Significantly speeds up `SELECT` queries, especially with `WHERE`, `ORDER BY`, `GROUP BY`, and `JOIN` clauses.
-   **Unique Constraints:** Indexes can enforce uniqueness on one or more columns.
-   **Primary Keys:** Automatically create a clustered index (or a unique non-clustered index if a clustered index already exists).

## Trade-offs of Indexing:
-   **Increased Storage Space:** Indexes require additional disk space.
-   **Slower Write Operations:** `INSERT`, `UPDATE`, and `DELETE` operations become slower because the database also needs to update the indexes.
-   **Overhead:** Maintaining indexes adds overhead to the database system.

## Types of Indexes:

### 1. B-Tree Index (B+ Tree is more common in practice)
-   **Concept:** A self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. Most commonly used index type in relational databases.
-   **Structure:** Nodes can have many children. All leaf nodes are at the same depth and contain pointers to the actual data or the data itself (for clustered indexes). Leaf nodes are often linked to allow efficient range scans.
-   **Use Cases:** Range queries, equality searches, `ORDER BY` clauses.

### 2. Hash Index
-   **Concept:** Uses a hash function to compute a hash code for each key, which then points to the data.
-   **Structure:** A hash table.
-   **Use Cases:** Very fast for equality searches (`WHERE column = value`).
-   **Limitations:** Not suitable for range queries (`WHERE column > value`) or `ORDER BY` clauses because the data is not stored in sorted order. Collisions can degrade performance.

### 3. Full-Text Index
-   **Concept:** Designed for efficient searching of text within large text fields (e.g., articles, documents).
-   **Use Cases:** Keyword searches, natural language queries.

### 4. Bitmap Index
-   **Concept:** Stores a bitmap (a sequence of bits) for each distinct value in a column. Each bit corresponds to a row, indicating whether that row has the value.
-   **Use Cases:** Columns with low cardinality (few distinct values), often used in data warehousing for fast querying of multiple conditions.
-   **Limitations:** Inefficient for high-cardinality columns or frequent updates.

## Indexing Best Practices:
-   **Index columns used in `WHERE`, `ORDER BY`, `GROUP BY`, `JOIN` clauses.**
-   **Avoid over-indexing:** Too many indexes can slow down writes.
-   **Consider cardinality:** Columns with high cardinality (many unique values) are good candidates for B-tree indexes.
-   **Composite indexes:** Indexes on multiple columns can be beneficial for queries involving those columns. Order of columns matters.
-   **Monitor performance:** Regularly analyze query performance and index usage.

## Go Example: (Conceptual) Indexing Impact

See the conceptual example code in `database-indexing.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run database-indexing.go
```
