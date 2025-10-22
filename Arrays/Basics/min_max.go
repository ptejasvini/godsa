package basics

import (
	"fmt"
	"sort"
)

func minMax(arr []int) (int, int) {
	sort.Ints(arr)
	min, max := arr[0], arr[len(arr)-1]
	return min, max
}
func main() {
	arr := []int{8, 2, 2, 4, 6}
	min, max := minMax(arr)
	fmt.Println(min, max)
}
