package main

import "fmt"

func sumArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	total := sumArray(numbers)
	fmt.Printf("The sum of array elements is: %d\n", total)

	// Another way to declare and initialize an array
	var fixedSizeArray [3]string
	fixedSizeArray[0] = "apple"
	fixedSizeArray[1] = "banana"
	fixedSizeArray[2] = "cherry"
	fmt.Println("Fixed size array:", fixedSizeArray)
}

