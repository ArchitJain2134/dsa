// example of palindrome array is [1,2,3,2,1]
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

	val:=palindrome(arr)

	if val == 1{
		fmt.Println("given array is palindrome")
	}else{
		fmt.Println("given array is not palindrome")
	}

}

func palindrome(arr [] int) int{
	if len(arr)==0{
		fmt.Println("array givem by you is empty")
		return 0
	}

	n:=len(arr)
	
	for i:=0;i<n/2;i++{
		if arr[i]!=arr[n-1-i]{
			return 0
		}
		
	}
	return 1
}
