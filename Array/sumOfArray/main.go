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
	    fmt.Println("Array given by you: ")
		fmt.Println(arr)


		// printing the total sum of an array 
		total:=sum(arr)
		fmt.Printf("Sum of all the elements of an array is : %v ",total)
	}

	func sum(arr[]int) (int){
		add:=0

		for _,val:= range arr{
			add=add+val
		}
		return add

	}


