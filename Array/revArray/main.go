package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter the size of an array")
	fmt.Scan(&n)

	arr:= make([]int,n)

	

	for i:=0;i<n;i++{
		fmt.Printf("Enter the %v elements of an array: ",i+1)
		fmt.Scan(&arr[i])
	}
	// arr := []int{10, 20, 30, 40, 50}

	fmt.Println("Reversed array",reverseArray(arr))
	
	
}

func reverseArray (arr []int) []int{
	x:=len(arr)
	for i:=0;i<x/2;i++{
		arr[i],arr[x-i-1]=arr[x-i-1],arr[i]
	}
	return arr

}
