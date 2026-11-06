# Citadel Securities coding questions

There is no official public catalog of every coding question Citadel Securities has asked. The lists below are compiled from candidate reports (LeetCode Discuss, Glassdoor, HackerRank OA writeups) and from the **Citadel** LeetCode company tag.

Citadel and Citadel Securities share much of the intern / new-grad engineering loop. Many public “Citadel” tags do not say which firm asked the question. Treat the LeetCode section as the shared company bank, and the first section as problems that candidates specifically attributed to **Citadel Securities**.

Official intern / new-grad process (CoderPad, Python or C++ preferred, 45-minute rounds): [Engineering interview process](https://www.citadelsecurities.com/careers/career-perspectives/internship-and-new-graduates-engineering-interview-process/).

---

## 1. Citadel Securities–specific reports

These are custom or domain-flavored problems candidates said they got in a Citadel Securities OA, phone screen, or onsite.

### Online assessment (HackerRank, typically 2 problems, 75–90 minutes)

| Problem | Pattern | Notes |
| --- | --- | --- |
| Maximum overlapping employees | Sweep line | Start/end time arrays; peak concurrency. Same family as Meeting Rooms II (LC 253). `n` up to `2e5`, so `O(n^2)` fails. |
| Max overlaps for a single employee / pivot team | Intervals + binary search | For each interval, count how many others overlap it (inclusive endpoints). Variant: max team size where one pivot overlaps everyone. |
| Minimum server upgrade time | Binary search on answer | Two servers need `t1` / `t2` seconds of work; at most one upgrades per second; server `i` cannot upgrade on multiples of `req_i`. |
| Memory-block MEX after at most one increment | Hash / MEX | Array of sizes; optionally increment one value `< n-1`; return all achievable MEX values, sorted. |
| Request throttling | Greedy / sliding window | Timestamps in increasing order, cooldown `K`; count dropped requests. |
| Process scheduling count | Combinatorics / DP | `n` processes, `m` slots; no process in two consecutive slots; answer mod `1e9+7`. |
| Subarrays with `arr[i] == arr[j] == sum(arr[i+1:j-1])` | Prefix sums + hash | Reported on a 2024 intern OA (`n` up to `1e5`). |
| Min operations to zero an array with `(x, y)` subtract | Binary search | One pick subtracts `x`; all others subtract `y`. Same family as LC 2702 (premium). |
| Order-book string / parse problem | Simulation / strings | Glassdoor: first live round described as an “order book string” DSA question. |

### Live coding / onsite (CoderPad)

| Problem | Pattern | Notes |
| --- | --- | --- |
| In-memory order book (single instrument) | Design, maps, queues | Add / cancel by id, price-time priority, level aggregation, depth. C++ container and allocation discussion is common. |
| Multi-exchange BBO / NBBO | Heaps + lazy deletion | `add(exchange, price, qty, side)`, `get_exchange_bbo(exchange)`, `get_nbbo()`. |
| Matching engine | Simulation | Quant-dev reports: write a matching engine. |
| Ring buffer / lock-free queue | Concurrency, C++ | Quant-dev onsite follow-up. |
| Currency exchange / best rate | Weighted graph (LC 399) | Equations like `BTC/USD`; find best conversion; watch self-loops. |
| Word Ladder II variant (LC 126) | BFS | Interviewer may omit “shortest” until after you code DFS. |
| Merge timestamped market-data streams | K-way merge + heaps | Stateful aggregation of price deltas into absolute prices. |
| Message batching producer | Design | Cache messages; flush by max count or max hold time; thread/process sharing. |
| Length-prefixed message parser | Streaming parse | `"02bc101234567890"` → print messages; follow-up: input split across calls. |
| Linked-list cycle detection | Floyd | Tail is null or points into the list. |
| Point-in-triangle query | Geometry | Non-overlapping triangles covering a region; return containing triangle or `-1`. |
| C++ debugging | Language | Glassdoor: second round was heavy C++ debugging. |
| Optimize a given function | OA / C++ | Junior C++ quant-dev HackerRank: speed up existing code. |

Closest LeetCode warm-ups for the trading-flavored rounds: **1801** Number of Orders in the Backlog, **295** Find Median from Data Stream, **253** Meeting Rooms II, **399** Evaluate Division, **23** Merge k Sorted Lists.

---

## 2. High-frequency LeetCode (Citadel company tag)

Public aggregators disagree on exact frequency, but these titles show up at the top of almost every Citadel list. Practice these first.

| LC | Title | Difficulty |
| --- | --- | --- |
| 239 | [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) | Hard |
| 295 | [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) | Hard |
| 146 | [LRU Cache](https://leetcode.com/problems/lru-cache/) | Medium |
| 460 | [LFU Cache](https://leetcode.com/problems/lfu-cache/) | Hard |
| 56 | [Merge Intervals](https://leetcode.com/problems/merge-intervals/) | Medium |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | Easy |
| 647 | [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) | Medium |
| 2484 | [Count Palindromic Subsequences](https://leetcode.com/problems/count-palindromic-subsequences/) | Hard |
| 399 | [Evaluate Division](https://leetcode.com/problems/evaluate-division/) | Medium |
| 380 | [Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/) | Medium |
| 1801 | [Number of Orders in the Backlog](https://leetcode.com/problems/number-of-orders-in-the-backlog/) | Medium |
| 718 | [Maximum Length of Repeated Subarray](https://leetcode.com/problems/maximum-length-of-repeated-subarray/) | Medium |
| 1515 | [Best Position for a Service Centre](https://leetcode.com/problems/best-position-for-a-service-centre/) | Hard |
| 2361 | [Minimum Costs Using the Train Line](https://leetcode.com/problems/minimum-costs-using-the-train-line/) (premium) | Hard |
| 2702 | [Minimum Operations to Make Numbers Non-positive](https://leetcode.com/problems/minimum-operations-to-make-numbers-non-positive/) (premium) | Hard |
| 3186 | [Maximum Total Damage With Spell Casting](https://leetcode.com/problems/maximum-total-damage-with-spell-casting/) | Medium |
| 2661 | [First Completely Painted Row or Column](https://leetcode.com/problems/first-completely-painted-row-or-column/) | Medium |
| 759 | [Employee Free Time](https://leetcode.com/problems/employee-free-time/) (premium) | Hard |
| 1229 | [Meeting Scheduler](https://leetcode.com/problems/meeting-scheduler/) (premium) | Medium |
| 2050 | [Parallel Courses III](https://leetcode.com/problems/parallel-courses-iii/) | Hard |
| 287 | [Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/) | Medium |
| 124 | [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) | Hard |
| 2771 | [Longest Non-decreasing Subarray From Two Arrays](https://leetcode.com/problems/longest-non-decreasing-subarray-from-two-arrays/) | Medium |
| 588 | [Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/) (premium) | Hard |
| 740 | [Delete and Earn](https://leetcode.com/problems/delete-and-earn/) | Medium |
| 1197 | [Minimum Knight Moves](https://leetcode.com/problems/minimum-knight-moves/) (premium) | Medium |

---

## 3. Full compiled LeetCode bank

Merged from Compile N Run, CodeJeet, CrackedPrep, DSA Prep, and Magicsheet (Citadel tag). Duplicates removed. Not every problem has been confirmed on a Citadel **Securities** loop.

### Easy

| LC | Title |
| --- | --- |
| 69 | [Sqrt(x)](https://leetcode.com/problems/sqrtx/) |
| 70 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) |
| 346 | [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/) (premium) |
| 412 | [Fizz Buzz](https://leetcode.com/problems/fizz-buzz/) |

### Medium

| LC | Title |
| --- | --- |
| 11 | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) |
| 12 | [Integer to Roman](https://leetcode.com/problems/integer-to-roman/) |
| 15 | [3Sum](https://leetcode.com/problems/3sum/) |
| 17 | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) |
| 39 | [Combination Sum](https://leetcode.com/problems/combination-sum/) |
| 46 | [Permutations](https://leetcode.com/problems/permutations/) |
| 47 | [Permutations II](https://leetcode.com/problems/permutations-ii/) |
| 49 | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) |
| 50 | [Pow(x, n)](https://leetcode.com/problems/powx-n/) |
| 56 | [Merge Intervals](https://leetcode.com/problems/merge-intervals/) |
| 71 | [Simplify Path](https://leetcode.com/problems/simplify-path/) |
| 79 | [Word Search](https://leetcode.com/problems/word-search/) |
| 98 | [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) |
| 103 | [Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/) |
| 116 | [Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/) |
| 117 | [Populating Next Right Pointers in Each Node II](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/) |
| 122 | [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) |
| 146 | [LRU Cache](https://leetcode.com/problems/lru-cache/) |
| 150 | [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) |
| 198 | [House Robber](https://leetcode.com/problems/house-robber/) |
| 200 | [Number of Islands](https://leetcode.com/problems/number-of-islands/) |
| 204 | [Count Primes](https://leetcode.com/problems/count-primes/) |
| 207 | [Course Schedule](https://leetcode.com/problems/course-schedule/) |
| 208 | [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) |
| 210 | [Course Schedule II](https://leetcode.com/problems/course-schedule-ii/) |
| 215 | [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) |
| 221 | [Maximal Square](https://leetcode.com/problems/maximal-square/) |
| 240 | [Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/) |
| 253 | [Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/) (premium) |
| 256 | [Paint House](https://leetcode.com/problems/paint-house/) (premium) |
| 279 | [Perfect Squares](https://leetcode.com/problems/perfect-squares/) |
| 287 | [Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/) |
| 300 | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) |
| 310 | [Minimum Height Trees](https://leetcode.com/problems/minimum-height-trees/) |
| 347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) |
| 348 | [Design Tic-Tac-Toe](https://leetcode.com/problems/design-tic-tac-toe/) (premium) |
| 380 | [Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/) |
| 399 | [Evaluate Division](https://leetcode.com/problems/evaluate-division/) |
| 435 | [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) |
| 560 | [Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/) |
| 622 | [Design Circular Queue](https://leetcode.com/problems/design-circular-queue/) |
| 647 | [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) |
| 688 | [Knight Probability in Chessboard](https://leetcode.com/problems/knight-probability-in-chessboard/) |
| 714 | [Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) |
| 718 | [Maximum Length of Repeated Subarray](https://leetcode.com/problems/maximum-length-of-repeated-subarray/) |
| 735 | [Asteroid Collision](https://leetcode.com/problems/asteroid-collision/) |
| 740 | [Delete and Earn](https://leetcode.com/problems/delete-and-earn/) |
| 752 | [Open the Lock](https://leetcode.com/problems/open-the-lock/) |
| 767 | [Reorganize String](https://leetcode.com/problems/reorganize-string/) |
| 787 | [Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/) |
| 875 | [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/) |
| 974 | [Subarray Sums Divisible by K](https://leetcode.com/problems/subarray-sums-divisible-by-k/) |
| 981 | [Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/) |
| 994 | [Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) |
| 1004 | [Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/) |
| 1048 | [Longest String Chain](https://leetcode.com/problems/longest-string-chain/) |
| 1143 | [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/) |
| 1155 | [Number of Dice Rolls With Target Sum](https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/) |
| 1197 | [Minimum Knight Moves](https://leetcode.com/problems/minimum-knight-moves/) (premium) |
| 1229 | [Meeting Scheduler](https://leetcode.com/problems/meeting-scheduler/) (premium) |
| 1268 | [Search Suggestions System](https://leetcode.com/problems/search-suggestions-system/) |
| 1670 | [Design Front Middle Back Queue](https://leetcode.com/problems/design-front-middle-back-queue/) |
| 1801 | [Number of Orders in the Backlog](https://leetcode.com/problems/number-of-orders-in-the-backlog/) |
| 1884 | [Egg Drop With 2 Eggs and N Floors](https://leetcode.com/problems/egg-drop-with-2-eggs-and-n-floors/) |
| 2044 | [Count Number of Maximum Bitwise-OR Subsets](https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/) |
| 2466 | [Count Ways To Build Good Strings](https://leetcode.com/problems/count-ways-to-build-good-strings/) |
| 2533 | [Number of Good Binary Strings](https://leetcode.com/problems/number-of-good-binary-strings/) (premium) |
| 2598 | [Smallest Missing Non-negative Integer After Operations](https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/) |
| 2661 | [First Completely Painted Row or Column](https://leetcode.com/problems/first-completely-painted-row-or-column/) |
| 2771 | [Longest Non-decreasing Subarray From Two Arrays](https://leetcode.com/problems/longest-non-decreasing-subarray-from-two-arrays/) |
| 2918 | [Minimum Equal Sum of Two Arrays After Replacing Zeros](https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/) |
| 2958 | [Length of Longest Subarray With at Most K Frequency](https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/) |
| 3186 | [Maximum Total Damage With Spell Casting](https://leetcode.com/problems/maximum-total-damage-with-spell-casting/) |
| 3484 | [Design Spreadsheet](https://leetcode.com/problems/design-spreadsheet/) |

### Hard

| LC | Title |
| --- | --- |
| 4 | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) |
| 10 | [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/) |
| 23 | [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) |
| 37 | [Sudoku Solver](https://leetcode.com/problems/sudoku-solver/) |
| 42 | [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) |
| 51 | [N-Queens](https://leetcode.com/problems/n-queens/) |
| 76 | [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) |
| 123 | [Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) |
| 124 | [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) |
| 126 | [Word Ladder II](https://leetcode.com/problems/word-ladder-ii/) |
| 127 | [Word Ladder](https://leetcode.com/problems/word-ladder/) |
| 149 | [Max Points on a Line](https://leetcode.com/problems/max-points-on-a-line/) |
| 188 | [Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) |
| 218 | [The Skyline Problem](https://leetcode.com/problems/the-skyline-problem/) |
| 224 | [Basic Calculator](https://leetcode.com/problems/basic-calculator/) |
| 239 | [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) |
| 269 | [Alien Dictionary](https://leetcode.com/problems/alien-dictionary/) (premium) |
| 295 | [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) |
| 297 | [Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) |
| 329 | [Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) |
| 332 | [Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/) |
| 381 | [Insert Delete GetRandom O(1) - Duplicates allowed](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/) |
| 460 | [LFU Cache](https://leetcode.com/problems/lfu-cache/) |
| 465 | [Optimal Account Balancing](https://leetcode.com/problems/optimal-account-balancing/) (premium) |
| 489 | [Robot Room Cleaner](https://leetcode.com/problems/robot-room-cleaner/) (premium) |
| 588 | [Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/) (premium) |
| 631 | [Design Excel Sum Formula](https://leetcode.com/problems/design-excel-sum-formula/) (premium) |
| 642 | [Design Search Autocomplete System](https://leetcode.com/problems/design-search-autocomplete-system/) (premium) |
| 759 | [Employee Free Time](https://leetcode.com/problems/employee-free-time/) (premium) |
| 765 | [Couples Holding Hands](https://leetcode.com/problems/couples-holding-hands/) |
| 782 | [Transform to Chessboard](https://leetcode.com/problems/transform-to-chessboard/) |
| 815 | [Bus Routes](https://leetcode.com/problems/bus-routes/) |
| 829 | [Consecutive Numbers Sum](https://leetcode.com/problems/consecutive-numbers-sum/) |
| 1203 | [Sort Items by Groups Respecting Dependencies](https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/) |
| 1235 | [Maximum Profit in Job Scheduling](https://leetcode.com/problems/maximum-profit-in-job-scheduling/) |
| 1515 | [Best Position for a Service Centre](https://leetcode.com/problems/best-position-for-a-service-centre/) |
| 1606 | [Find Servers That Handled Most Number of Requests](https://leetcode.com/problems/find-servers-that-handled-most-number-of-requests/) |
| 1632 | [Rank Transform of a Matrix](https://leetcode.com/problems/rank-transform-of-a-matrix/) |
| 2050 | [Parallel Courses III](https://leetcode.com/problems/parallel-courses-iii/) |
| 2361 | [Minimum Costs Using the Train Line](https://leetcode.com/problems/minimum-costs-using-the-train-line/) (premium) |
| 2484 | [Count Palindromic Subsequences](https://leetcode.com/problems/count-palindromic-subsequences/) |
| 2702 | [Minimum Operations to Make Numbers Non-positive](https://leetcode.com/problems/minimum-operations-to-make-numbers-non-positive/) (premium) |

---

## 4. Patterns they actually test

In order of how often they show up in Securities-flavored reports:

1. **Intervals / sweep line** — concurrency, free time, meeting rooms.
2. **Heaps and streaming** — median, top-K, order books, merge K streams.
3. **Binary search on the answer** — huge `n` or `t`, hidden tests kill `O(n^2)`.
4. **Design + hash + linked list** — LRU/LFU, caches, circular queues, order id indexes.
5. **DP** — stocks, delete-and-earn / house robber, palindromes, job scheduling.
6. **Graphs** — currency exchange, word ladder, courses / topo sort.
7. **C++ systems** — debugging, containers, matching engine, ring buffer.

---

## Sources

- [Citadel Securities intern / new-grad engineering interview process](https://www.citadelsecurities.com/careers/career-perspectives/internship-and-new-graduates-engineering-interview-process/)
- [Compile N Run — Citadel](https://www.compilenrun.com/interview/company/citadel/)
- [CodeJeet — Citadel](https://codejeet.com/company/citadel)
- [CrackedPrep — Citadel](https://crackedprep.com/citadel-interview-questions)
- [SpaceComplexity — top 15 Citadel problems](https://spacecomplexity.ai/blog/citadel-coding-interview-problems)
- Candidate OA writeups: [HackerRank overlap intervals](https://hiya31.medium.com/what-they-asked-me-in-the-citadel-securities-hackerrank-coding-round-ed3ceded3c04), [LeetCode Discuss intern OA](https://leetcode.com/discuss/post/7698725/), [Citadel OA Questions](https://leetcode.com/discuss/interview-question/5430576/Citadel-OA-Questions/)
- Order book: [PracHub intern order book](https://prachub.com/interview-experiences/citadel-intern-software-engineer-interview-experience-a-deep-dive-on-the-classic-orderbook-problem), [In-memory order book](https://prachub.com/interview-questions/in-memory-order-book-with-price-time-order-level-aggregation-and-order-cancel)

Last compiled: 2026-11-04.
