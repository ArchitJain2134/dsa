// in this program we will be finding the toal number of odd numbers and even numbers in an array

package main

import "fmt"

func main() {
	var n int

	// declaring the size of an array
	fmt.Println("Enter the size of an arary")
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

	// printing a result
	odd, even := check(arr)
	fmt.Printf("Total number of odd elements in an array is : %v \n", odd)
	fmt.Printf("Total number of even elements in an array is : %v \n", even)
}

func check(arr []int) (int, int) {

	// check if the array is empty or not
	if len(arr) == 0 {
		fmt.Println("Your given array is empty")
		return 0, 0
	}

	// declaring 2 variables odd and even and setting their values to 0
	odd, even := 0, 0

	// checking whether the element is odd or even and according incermenting the value 
	for _, val := range arr {
		if val%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return odd, even

}
