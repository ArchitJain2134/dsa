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

	// calculating teh value by calling the function
	minval,maxval := MinAndMax(arr)
	fmt.Printf("\n Minimum value is : %v",minval)
	fmt.Printf("\n Maximum value is : %v",maxval)
}

func MinAndMax(arr[] int) (int,int) {
	
	// initialising 1st element as min and max both
	minval:=arr[0]
	maxval:=arr[0]
	if len(arr)==0{
		fmt.Println("please write the elements in an array")
		return 0,0
	}

	for _,val:=range arr{
		// checking and editting the value if any
		if val<minval{
			minval=val
		}

		if val>maxval{
			maxval=val
		}

	}
	return minval,maxval
}
