package basics

import (
	"fmt"
)

//problem link : https://www.geeksforgeeks.org/dsa/sum-of-array-elements/

func sumOfElements(arr []int) int {
	sum := 0
	for i := range arr {
		sum+=arr[i]
	
	}
	return sum
}
func SumExample() {
	arr := []int{1, 2, 3, 4, 5}

	sum := sumOfElements(arr)
	fmt.Println("Sum of array elements is:", sum)
}
