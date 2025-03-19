// problem statement
// Rotate Array by K Positions

// Rotate an array to the right by K positions.
// Input: arr = [1, 2, 3, 4, 5], K = 2
// Output: [4, 5, 1, 2, 3]

package main

import "fmt"

func main() {
	var n int

	// declaring the size of an array
	fmt.Println("Enter the size of an array")
	fmt.Scan(&n)

	arr := make([]int, n)

	// filling each element in a declared array
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the %v element in an array : ", i+1)
		fmt.Scan(&arr[i])
	}

	// printing an array you have given as an input
	fmt.Println("array given by you is : ")
	fmt.Println(arr)

	// taking input of how many times user wants to rotate
	var k int
	fmt.Printf("How many times you want to rotate your array? \n")
	fmt.Scan(&k)

	// if rotate times is greter than size of an array
	k = k % n

	// calling and printing the result
	newarr:=rotateRight(arr,k)
	fmt.Println("Rotated array is ",newarr)

}

func rotateRight(arr []int, k int) []int {

	n:=len(arr)

	// checking if the rotation time is greater than size of array 
	k = k % n

	// returning an array and doing the rortation in return itself
	return append(arr[n-k:], arr[:n-k]...)

}
