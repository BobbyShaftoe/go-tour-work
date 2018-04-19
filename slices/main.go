package main

import (
	"fmt"
	_ "fmt"
	"strings"
)

// Slices are dynamically-sized. The type []T is a slice with elements of type T
// A slice is formed by specifying two indices, a low and high bound, separated by a colon. a[low : high]

// a slice which includes elements 1 through 3.  a[1:4]

var primes = [6]int{2, 3, 5, 7, 11, 13}

func main() {

	var s []int = primes[1:4]
	fmt.Println(s)

	// A slice does not store any data, it just describes a section of an underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// all elements
	c := names[0:]
	fmt.Println(c)

	fmt.Printf("len=%d cap=%d %v\n", len(c), cap(c), c)

	// create dynamically-sized arrays with make
	d := make([]int, 0, 5) // len(d)=0, cap(d)=5
	printSlice("d", d)

	d = d[:cap(d)] // change length, len(d)=5, cap(d)=5
	printSlice("d", d)

	// append to slice
	d = append(d, 9, 8, 7)
	fmt.Println("Append to slice")
	printSlice("d", d)

	// iterate over a slice
	e := d[5:9]
	for i, v := range e {
		fmt.Printf("%d: 2*%d = %d\n", i, v, 2*v)
	}

	for _, value := range e {
		fmt.Printf("%d\n", value)

	}

	fmt.Println("Tic Tac Toe - Slices of Slices")
	ticTacToe("O")
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Slices of slices (multidimensional array)
func ticTacToe(x string) {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = x
	board[2][2] = "*"
	board[1][2] = x
	board[1][0] = "*"
	board[0][2] = x

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Printf("\n%s\n%s\n%s\n", board[0], board[1], board[2])

}

// import "golang.org/x/tour/pic"
//
// func Pic(dx, dy int) [][]uint8 {
// 	p := make([][]uint8, dy)
// 	for i := range p {
// 		p[i] = make([]uint8, dx)
// 	}
//
// 	for y, row := range p {
// 		for x := range row {
// 			row[x] = uint8((x^2 - y^2)/2 - x  - 2 )
//
// 		}
// 	}
//
// 	return p
// }
//
// func main() {
// 	pic.Show(Pic)
//
// }
