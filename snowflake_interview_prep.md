# Snowflake Coding Interview Prep

Snowflake coding rounds tend to be LeetCode medium/hard, with a lean toward data-systems problems (parsing, hashing, sorting, joins, caching, concurrency). Expect to write working code, discuss complexity, and handle edge cases and scale ("what if the data doesn't fit in memory?").

Legend: 🟢 Easy · 🟡 Medium · 🔴 Hard

---

## 1. Arrays, Strings & Hashing

| # | Problem | Diff | Notes |
|---|---------|------|-------|
| 1 | Two Sum | 🟢 | Hash map, one pass |
| 2 | Group Anagrams | 🟡 | Sorted key / count tuple |
| 3 | Longest Substring Without Repeating Characters | 🟡 | Sliding window |
| 4 | Minimum Window Substring | 🔴 | Sliding window + counts |
| 5 | Subarray Sum Equals K | 🟡 | Prefix sum + hash map |
| 6 | Product of Array Except Self | 🟡 | Prefix/suffix, no division |
| 7 | Top K Frequent Elements | 🟡 | Heap vs bucket sort |
| 8 | Merge Intervals | 🟡 | Sort by start |
| 9 | Insert Interval | 🟡 | Linear scan |
| 10 | Meeting Rooms II | 🟡 | Heap or sweep line |
| 11 | Trapping Rain Water | 🔴 | Two pointers / monotonic stack |
| 12 | Longest Consecutive Sequence | 🟡 | Set, O(n) |
| 13 | 3Sum | 🟡 | Sort + two pointers, dedupe |
| 14 | Next Permutation | 🟡 | In-place |
| 15 | Find All Duplicates in an Array | 🟡 | Index marking |

## 2. Linked Lists

| # | Problem | Diff |
|---|---------|------|
| 16 | Reverse Linked List (iterative + recursive) | 🟢 |
| 17 | Merge Two Sorted Lists | 🟢 |
| 18 | Merge K Sorted Lists | 🔴 |
| 19 | Linked List Cycle II (find cycle start) | 🟡 |
| 20 | Remove Nth Node From End | 🟡 |
| 21 | Copy List with Random Pointer | 🟡 |
| 22 | Reverse Nodes in k-Group | 🔴 |

## 3. Trees & Graphs

| # | Problem | Diff | Notes |
|---|---------|------|-------|
| 23 | Validate Binary Search Tree | 🟡 | Bounds, not just parent compare |
| 24 | Lowest Common Ancestor of a Binary Tree | 🟡 | |
| 25 | Binary Tree Level Order Traversal | 🟡 | BFS |
| 26 | Serialize and Deserialize Binary Tree | 🔴 | Common follow-up topic |
| 27 | Binary Tree Maximum Path Sum | 🔴 | |
| 28 | Kth Smallest Element in a BST | 🟡 | Follow-up: frequent inserts/deletes |
| 29 | Number of Islands | 🟡 | DFS/BFS/Union-Find |
| 30 | Clone Graph | 🟡 | |
| 31 | Course Schedule I / II | 🟡 | Topological sort, cycle detection |
| 32 | Word Ladder | 🔴 | BFS, bidirectional BFS follow-up |
| 33 | Alien Dictionary | 🔴 | Topo sort from orderings |
| 34 | Network Delay Time | 🟡 | Dijkstra |
| 35 | Cheapest Flights Within K Stops | 🟡 | Bellman-Ford / BFS |
| 36 | Redundant Connection | 🟡 | Union-Find |
| 37 | Implement Trie (+ wildcard search) | 🟡 | |

## 4. Dynamic Programming

| # | Problem | Diff |
|---|---------|------|
| 38 | Climbing Stairs / House Robber | 🟢 |
| 39 | Coin Change (min coins + count ways) | 🟡 |
| 40 | Longest Increasing Subsequence (O(n log n)) | 🟡 |
| 41 | Longest Common Subsequence | 🟡 |
| 42 | Edit Distance | 🔴 |
| 43 | Word Break (I and II) | 🟡 / 🔴 |
| 44 | Partition Equal Subset Sum | 🟡 |
| 45 | Decode Ways | 🟡 |
| 46 | Unique Paths with Obstacles | 🟡 |
| 47 | Regular Expression Matching | 🔴 |
| 48 | Best Time to Buy and Sell Stock (all variants) | 🟡 |

## 5. Heaps, Binary Search & Sorting

| # | Problem | Diff | Notes |
|---|---------|------|-------|
| 49 | Kth Largest Element in an Array | 🟡 | Heap vs quickselect |
| 50 | Find Median from Data Stream | 🔴 | Two heaps |
| 51 | Search in Rotated Sorted Array | 🟡 | |
| 52 | Median of Two Sorted Arrays | 🔴 | |
| 53 | Koko Eating Bananas / Capacity to Ship Packages | 🟡 | Binary search on answer |
| 54 | Sliding Window Maximum | 🔴 | Monotonic deque |
| 55 | K Closest Points to Origin | 🟡 | |
| 56 | Merge K Sorted Arrays / External Merge Sort | 🟡 | Data larger than memory |

---

## 6. Design / "Build It" Problems (very Snowflake-flavored)

Write full, working, tested code for these.

57. **LRU Cache** — O(1) get/put; hash map + doubly linked list. Follow-up: make it thread-safe.
58. **LFU Cache** — O(1) all operations.
59. **Rate Limiter** — token bucket, fixed window, sliding window log/counter. Discuss distributed variant.
60. **Time-Based Key-Value Store** — binary search over timestamps.
61. **Design Hit Counter** — sliding window with a bounded buffer.
62. **Iterator problems** — Flatten Nested List Iterator, Peeking Iterator, Zigzag Iterator.
63. **Min Stack / Max Stack**.
64. **Design an In-Memory File System**.
65. **Design a Key-Value Store with transactions** — begin/commit/rollback, nested.
66. **Design a Task Scheduler with dependencies** — topo sort + worker pool.
67. **Insert Delete GetRandom O(1)**.
68. **Logger Rate Limiter / Dedup stream of events**.

## 7. Data-Systems Flavored Problems

These map to what a query engine actually does. Good signal to interviewers at a database company.

69. **Implement a hash join** — build/probe phases; handle duplicate keys and NULLs. Follow-up: build side doesn't fit in memory (grace hash join / partitioning).
70. **Implement sort-merge join** on two sorted inputs.
71. **Implement GROUP BY aggregation** over a stream of rows (`COUNT`, `SUM`, `AVG`, `MIN/MAX`) with hash aggregation. Follow-up: spill to disk.
72. **External sort** — sort 100 GB with 1 GB RAM (run generation + k-way merge).
73. **Top-K over a huge stream** — heap, and merging per-node results.
74. **Parse and evaluate a simple SQL-like expression** — tokenizer + recursive descent (`WHERE a > 5 AND (b = 'x' OR c < 3)`).
75. **Basic Calculator I / II / III** — stack / precedence parsing.
76. **CSV parser** — quoted fields, embedded commas/newlines, escaped quotes.
77. **JSON parser / flattener** — flatten nested JSON into dotted column paths (relevant to VARIANT).
78. **Implement a Bloom filter** — discuss false-positive rate, sizing.
79. **Implement a B-tree or B+ tree insert/search** (at least describe splitting).
80. **Consistent hashing** — add/remove node, virtual nodes.
81. **Merge overlapping time ranges across many sorted partitions** (min/max pruning, micro-partition style).
82. **Run-length / dictionary encoding** — encode and decode a column.
83. **Implement a simple query planner step** — predicate pushdown on an expression tree.
84. **Detect duplicate rows across two large files** — hashing, partitioning, bloom pre-filter.
85. **Reservoir sampling** — uniform sample from a stream of unknown size.
86. **Count-Min Sketch / HyperLogLog** — conceptual + basic implementation.

## 8. Concurrency

87. **Thread-safe LRU cache** — locks, read/write lock, lock striping.
88. **Producer-Consumer with a bounded blocking queue**.
89. **Implement a thread pool / worker queue**.
90. **Print in Order / Print FooBar Alternately** — semaphores / condition variables.
91. **Dining Philosophers** — deadlock avoidance.
92. **Concurrent counter / Atomic ops** — compare-and-swap vs lock.
93. **Read-Write Lock** from scratch.

## 9. SQL (usually appears too)

94. Second highest salary / Nth highest salary.
95. Top-N per group using `ROW_NUMBER()` / `RANK()` / `DENSE_RANK()`.
96. Running total and moving average with window frames.
97. Consecutive days / gaps and islands.
98. Self-join: employees earning more than their manager.
99. Find duplicate rows and delete all but one.
100. Sessionization — group events into sessions with a 30 min gap (`LAG` + cumulative sum).
101. Pivot / unpivot rows to columns.
102. Retention / funnel query: users active day 0 and day N.
103. Explain the difference between `INNER`, `LEFT`, `FULL`, `CROSS` joins, and anti/semi joins (`NOT EXISTS` vs `NOT IN` with NULLs).
104. Query optimization: given a slow query, what would you check? (join order, filters, clustering, partition pruning)

---

## Concept Questions to Be Ready For

- Time/space complexity of every solution; best/average/worst case.
- Hash map internals: collisions, resizing, amortized cost.
- Quicksort vs mergesort vs heapsort; stable vs unstable; when each wins.
- B-tree vs LSM tree; row store vs column store; why columnar compresses well.
- What is partition pruning / micro-partitions / clustering?
- Hash join vs sort-merge join vs nested loop join: when to choose each.
- CAP theorem, consistency models, idempotency, exactly-once vs at-least-once.
- Deadlock conditions; mutex vs semaphore vs condition variable.
- How would you scale your solution to billions of rows / multiple machines?

---

## Suggested 4-Week Plan

| Week | Focus | Target |
|------|-------|--------|
| 1 | Arrays, strings, hashing, linked lists (1–22) | 2–3 problems/day |
| 2 | Trees, graphs, DP (23–48) | 2 problems/day |
| 3 | Heaps/binary search (49–56) + design problems (57–68) | Full implementations with tests |
| 4 | Data-systems (69–86), concurrency (87–93), SQL (94–104), 2–3 mock interviews | Timed, talk out loud |

## Interview Tips

1. **Clarify first**: input size, duplicates, NULLs, sorted or not, memory limits, streaming vs batch.
2. **State brute force, then optimize**; say the complexity at each step.
3. **Talk while coding**; narrate invariants and edge cases.
4. **Test with your own examples**, including empty input, single element, duplicates, and overflow.
5. **Always offer the scale follow-up**: "If this doesn't fit in memory, I'd partition by hash and process each partition independently."
6. Practice in the language you'll use in the interview (Python/Java/C++), and know its heap, deque, and sorted-container APIs cold.
