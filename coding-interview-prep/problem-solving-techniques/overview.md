# Problem Solving Techniques for Coding Interviews

Solving coding problems effectively requires more than just knowing data structures and algorithms; it involves a systematic approach and familiarity with common problem-solving patterns. This section outlines a general strategy and various techniques to tackle interview questions.

## I. General Approach to Problem Solving

A structured approach can help you break down complex problems and arrive at optimal solutions.

1.  **Understand the Problem:**
    *   **Read Carefully:** Read the problem statement multiple times.
    *   **Clarify:** Ask clarifying questions. What are the inputs, outputs, constraints, edge cases? (e.g., "Can the input array be empty?", "Are numbers always positive?", "What is the maximum size of N?").
    *   **Examples:** Work through a small example manually. If examples are provided, understand them. Create your own simple examples.
    *   **Input/Output Format:** How will the input be given? How should the output be formatted?

2.  **Break Down the Problem:**
    *   Can the problem be divided into smaller, independent subproblems?
    *   Can you solve a simpler version of the problem first?

3.  **High-Level Plan (Brainstorming & Strategy):**
    *   **Data Structures:** What data structures might be suitable? (Arrays, Linked Lists, Stacks, Queues, Hash Maps, Trees, Graphs, Heaps, Tries, etc.)
    *   **Algorithms:** What algorithmic paradigms might apply? (Brute Force, Greedy, Divide and Conquer, Dynamic Programming, Backtracking, BFS/DFS, Two Pointers, Sliding Window, etc.)
    *   **Time/Space Complexity:** What are the expected time and space constraints? This often guides your choice of algorithm.
    *   **Initial Thoughts (Brute Force):** Always start by thinking about a brute-force solution. This helps you understand the problem better and provides a baseline for optimization.

4.  **Detailed Plan (Algorithm Design):**
    *   Refine your high-level plan into a step-by-step algorithm.
    *   Consider edge cases: empty input, single element, maximum/minimum values, nulls.
    *   Walk through your algorithm with your examples. Does it work correctly?

5.  **Implement the Solution:**
    *   Write clean, readable code.
    *   Use meaningful variable names.
    *   Add comments for complex logic.
    *   Handle errors and edge cases explicitly.

6.  **Test and Debug:**
    *   **Test with Examples:** Run your code with the examples you generated or were given.
    *   **Test Edge Cases:** Test with empty inputs, single elements, very large inputs, specific values that might break your logic.
    *   **Debug:** If your code doesn't work, use print statements or a debugger to trace execution and variable values.

7.  **Optimize (if necessary):**
    *   **Analyze Complexity:** Determine the time and space complexity of your solution.
    *   **Identify Bottlenecks:** Where is the code spending most of its time or memory?
    *   **Refine:** Can you use a more efficient data structure or algorithm? Can you reduce redundant computations?

## II. Common Algorithmic Techniques

These are recurring patterns and strategies used to solve a wide range of coding problems.

1.  **Brute Force / Exhaustive Search:**
    *   **Concept:** Try every possible solution and check if it's valid.
    *   **When to use:** Simple problems, small input sizes, or as a starting point to understand the problem before optimizing.
    *   **Example:** Checking all pairs in an array to find two numbers that sum to a target.

2.  **Divide and Conquer:**
    *   **Concept:** Break a problem into two or more smaller subproblems of the same type, solve them independently, and then combine their solutions.
    *   **When to use:** Problems that can be naturally split.
    *   **Examples:** Merge Sort, Quick Sort, Binary Search.

3.  **Recursion:**
    *   **Concept:** A function calls itself to solve smaller instances of the same problem. Requires a base case to terminate.
    *   **When to use:** Problems with recursive definitions (e.g., trees, graphs, fractals), or when an iterative solution is complex.
    *   **Examples:** Factorial, Fibonacci, Tree traversals (DFS), Backtracking.

4.  **Iteration:**
    *   **Concept:** Using loops (for, while) to repeatedly execute a block of code.
    *   **When to use:** Most problems can be solved iteratively. Often more efficient than recursion for simple problems due to less overhead.
    *   **Examples:** Array traversal, calculating sums, BFS.

5.  **Greedy Approach:**
    *   **Concept:** Make the locally optimal choice at each step with the hope that this choice will lead to a globally optimal solution.
    *   **When to use:** Problems where a local optimum guarantees a global optimum (e.g., Activity Selection, Dijkstra's, Prim's, Kruskal's).
    *   **Caution:** Proving correctness can be hard; often doesn't work for all problems.

6.  **Dynamic Programming (DP):**
    *   **Concept:** Solve complex problems by breaking them into simpler overlapping subproblems and storing the results of these subproblems to avoid recomputing them.
    *   **When to use:** Problems exhibiting optimal substructure and overlapping subproblems.
    *   **Approaches:** Memoization (top-down), Tabulation (bottom-up).
    *   **Examples:** Fibonacci, Longest Common Subsequence, Knapsack.

7.  **Two Pointers:**
    *   **Concept:** Use two pointers (indices) to traverse an array or list, often from opposite ends or at different speeds.
    *   **When to use:** Problems involving sorted arrays/lists, finding pairs, or partitioning.
    *   **Examples:** Finding pairs with a given sum in a sorted array, reversing a string/array, checking for palindromes.

8.  **Sliding Window:**
    *   **Concept:** Maintain a "window" (a contiguous subarray or substring) that expands or shrinks as it moves over a collection.
    *   **When to use:** Problems involving finding subarrays/substrings with certain properties (e.g., maximum sum subarray of a given size, longest substring without repeating characters).
    *   **Example:** Finding the maximum sum of a subarray of size K.

9.  **Hashing / Hash Maps (Dictionaries):**
    *   **Concept:** Use hash tables for O(1) average time complexity for insertion, deletion, and lookup.
    *   **When to use:** Counting frequencies, checking for duplicates, storing key-value pairs, optimizing lookups.
    *   **Examples:** Two Sum problem, finding unique elements, implementing caches.

10. **Sorting:**
    *   **Concept:** Arrange elements in a specific order.
    *   **When to use:** Often a preprocessing step that simplifies subsequent logic. Many algorithms rely on sorted input.
    *   **Examples:** Finding median, binary search, problems requiring elements to be processed in a specific order.

11. **Graph Traversal (BFS / DFS):**
    *   **Concept:** Systematic ways to visit all nodes and edges in a graph.
    *   **BFS (Breadth-First Search):** Explores layer by layer. Uses a queue. Good for shortest path in unweighted graphs, finding connected components.
    *   **DFS (Depth-First Search):** Explores as far as possible along each branch. Uses a stack (or recursion). Good for cycle detection, topological sort, finding connected components.

12. **Bit Manipulation:**
    *   **Concept:** Performing operations directly on the binary representation of numbers.
    *   **When to use:** Optimizing for space or time, solving problems related to sets, flags, or specific number properties.
    *   **Examples:** Checking even/odd, setting/clearing bits, counting set bits, finding unique numbers.

13. **Backtracking:**
    *   **Concept:** A general algorithmic technique for solving problems by trying to build a solution incrementally. If a partial solution cannot be extended to a complete solution, or if it violates any constraints, the algorithm "backtracks" (undoes its last step) and tries a different path.
    *   **When to use:** Problems that involve exploring all possible configurations or permutations to find a solution (e.g., N-Queens, Sudoku Solver, Combinations/Permutations).

## III. Optimization Strategies

Once you have a working solution, consider how to make it more efficient.

-   **Time Complexity Analysis (Big O Notation):** Understand how the runtime grows with input size. Aim for the most efficient possible.
-   **Space Complexity Analysis:** Understand how much memory your algorithm uses.
-   **Trade-offs:** Sometimes, you can trade space for time (e.g., using a hash map to reduce search time) or vice-versa.
-   **Precomputation:** Can you precompute some values to speed up repeated calculations?
-   **Early Exit Conditions:** Can you stop processing early if a solution is found or if it's clear no solution exists down a certain path?
-   **Data Structure Choice:** Is there a more appropriate data structure that offers better performance for your specific operations?

## Go Example: (Conceptual) Problem Solving Techniques

See the conceptual example code in `overview.go`.
```go
// To run this code:
// 1. Navigate to the problem-solving-techniques directory: cd coding-interview-prep/problem-solving-techniques
// 2. Run the file: go run overview.go
```
