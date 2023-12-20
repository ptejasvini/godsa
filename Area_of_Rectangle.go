package main

import "fmt"

func AreaOfRectangle(length, breadth int) int {
	return length * breadth
}

func main() {
	x := AreaOfRectangle(5, 6)
	fmt.Println(x)
}
