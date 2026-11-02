# Arrays

An array is a collection of items stored at contiguous memory locations. The idea is to store multiple items of the same type together. This makes it easier to calculate the position of each element by simply adding an offset to a base value, i.e., the memory address of the first element of the array (generally denoted by the name of the array).

## Characteristics:
- Fixed size: Once an array is created, its size cannot be changed.
- Homogeneous: All elements in an array must be of the same data type.
- Indexed: Elements are accessed using an index, typically starting from 0.

## Time Complexity:
- Access: O(1) (constant time)
- Search: O(n) (linear time)
- Insertion/Deletion: O(n) (linear time, as elements may need to be shifted)

## Go Example: Sum of Array Elements

See the example code in `arrays.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run arrays.go
```

```