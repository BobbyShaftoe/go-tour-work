package main

import (
	"fmt"
	_ "fmt"
)

type Vertex struct {
	X, Y int
}

func (v *Vertex) addValues() int {
	return v.X + v.Y
}

func (v *Vertex) multiplyValues() int {
	return v.X * v.Y
}

func main() {

	// Case 1
	v := Vertex{10, 20}

	fmt.Printf("X: %d, Y: %d\n", v.X, v.Y)
	fmt.Println(v.addValues())
	fmt.Println(v.multiplyValues())


	// Case 2
	x := &v.X
	y := &v.Y

	*x = 22
	*y = 33

	fmt.Printf("X: %d, Y: %d\n", *x, *y)
	fmt.Println(v.addValues())
	fmt.Println(v.multiplyValues())


	// Case 3
	vertexPointer := &v
	fmt.Printf("vertexPointer type: %T, value: %v\n", vertexPointer, *vertexPointer)

  x = &vertexPointer.X
  y = &vertexPointer.Y
  *x = 100
  *y = 100

	fmt.Printf("X: %d, Y: %d\n", v.X, v.Y)
}
