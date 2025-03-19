package main

import "fmt"

func main() {

	var x, y int

	// taking inputs for 1st array
	fmt.Printf("Enter the size of first array : ")
	fmt.Scan(&x)
	fmt.Println("Enter the elements of first sorted array")
	arr1 := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Printf("Your first array is: \n %v \n", arr1)

	// taking inputs for 2nd array
	fmt.Printf("Enter the size of second array : ")
	fmt.Scan(&y)
	fmt.Println("Enter the elements of second sorted array")

	arr2 := make([]int, y)
	for j := 0; j < y; j++ {
		fmt.Scan(&arr2[j])
	}
	fmt.Printf("Your second array is: \n %v \n", arr2)

	// calling the function and printing the result
	result:=sortedArray(arr1,arr2)
	fmt.Println("After merging successfully new array is : ",result)

}


func sortedArray (arr1[]int,arr2[]int) ([]int){

	merged:=[]int{}

	i,j:=0,0

	// Compare elements from both slices and pick the smaller one
	for i<len(arr1) && j<len(arr2){

		if arr1[i]<arr2[j]{
			merged = append(merged, arr1[i])
			i++
		}else{
			merged = append(merged, arr2[j])
			j++
		}
	}

	// adding any remaining elements
	merged=append(merged, arr1[i:]...)
	merged=append(merged, arr2[j:]...)
	return merged

	
}
