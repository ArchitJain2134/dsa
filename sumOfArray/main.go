package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of an array")
	fmt.Scan(&n)
	arr := make([]int, n)

	// taking input in an array
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the %v element of an array: ", i+1)
		fmt.Scan(&arr[i])
	}

	// printing an array which the user is give as input to confirm
		fmt.Println(arr)

		
	}


