package main

import (
	"fmt"
	_ "fmt"
)

var globalI = 10

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	// Naked return
}

func main() {
	// var or := can be used in function
	i := globalI
	var j, k = 20, 30
	fmt.Println(i, j, k)

	fmt.Println(add(10, 20))

	a, b := swap("first", "second")
	fmt.Println(a, b)

	fmt.Println(split(17))

}
