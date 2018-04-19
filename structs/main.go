package main

import (
	"fmt"
	_ "fmt"
)

type Vertex struct {
	X int
	Y int
	Z int
}

func main() {

	v := Vertex{1, 2, 3}
	fmt.Println(v)

	// Struct fields are accessed using a dot
	v.X, v.Y, v.Z = 4, 20, 66
	fmt.Println(v.X, v.Y, v.Z)

	// Struct fields can be accessed through a struct pointer. Simply use p.X instead of (*p).X
	p := &v
	p.X = 1e9
	fmt.Println(v)

	v1 := Vertex{X: 99}
	fmt.Println(v1)

	// Declare pointer to new Vertex struct
	p1 := &Vertex{100, 200, 300}
	fmt.Println(*p1)

}
