package main

import (
	"fmt"
	_ "fmt"
	"math"
	_ "math"
)

// An "interface type" is defined as a set of "method signatures"
// A value of "interface type" can hold any value that implements those "methods"

// TYPES //
// A type implements an interface by implementing its methods

type MyFloat float64
type Vertex struct {
	X, Y float64
}

// Type "MyFloat" implements the interface "MathUtil" in "Abs()" and "AreaCircle()" methods
// Type "Vertex" implements the interface "TrigUtil" in "EuclidDistance()" method

// INTERFACES //

type MathUtil interface {
	// this interface has "Abs()" method with type "MyFloat" as receiver
	Abs() float64
	// this interface has "AreaCircle()" method with type "MyFloat" as receiver
	AreaCircle() float64
}
type TrigUtil interface {
	// this interface has "EuclidDistance()" method with type "*Vertex" as receiver
	EuclidDistance() float64
}

// METHODS //

// This "EuclidDistance()" method in "MathUtil" interface requires *Vertex type
func (v *Vertex) EuclidDistance() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This "Abs()" method in "MathUtil" interface requires float64 type
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// This "AreaCircle()" method in "MathUtil" interface requires float64 type
func (f MyFloat) AreaCircle() float64 {
	area := math.Pi * math.Pow(float64(f), 2)
	return float64(area)
}

// MAIN //

func main() {
	r := float64(5.49)

	// The "MathUtil" interface is declared with value of "MyFloat" type expected by it's methods
	var a MathUtil = MyFloat(r)
	// Call "AreaCircle()" method from "MathUtil" interface
	fmt.Println("Area of circle with radius:", r, "is:", a.AreaCircle())

	// The "TrigUtil" interface is declared with value of "&Vertex" type expected by it's method
	var x, y float64 = 55, 66
	var b TrigUtil = &Vertex{x, y}
	// Call "TrigUtil" method from "TrigUtil" interface
	fmt.Println("Euclidean distance between points", x, "and", y, "is:", b.EuclidDistance())

}

// FUNDAMENTAL EXAMPLE //

// 	import "fmt"
// 	import "strings"
//
// 	type T struct {
// 		S string
// 	}
//
// 	// This method means type T implements the interface I,
// 	// but we don't need to explicitly declare that it does so.
//
// 	type I interface {
// 		M()
// 		N()
// 	}
//
// 	func (t T) M() {
// 		fmt.Println(t.S)
// 	}
// 	func (t T) N() {
// 		fmt.Println(strings.Title(t.S))
// 	}
//
//
// 	func main() {
// 		var i I = T{"hello"}
// 		i.M()
// 		i.N()
// 	}
