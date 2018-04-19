package main

import (
	"fmt"
	_ "fmt"
	"math"
	_ "math"
	"strings"
	_ "strings"
)

// Go does not have classes. However, you can define "methods on types"
// A "method is a function" with a "special receiver argument"

type Vertex struct {
	X, Y float64
}

type Names struct {
	first, last string
}

type MyFloat float64

// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the 'SomeMethod()' method has a receiver of type Vertex named 'v'

func (v Vertex) SomeMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// type is 'Names struct' and method is 'thisMethod()'
// in other words 'thisMethod()' has receiver of type Vertex named 'n'

func (n Names) thisMethod() string {
	return strings.Title(n.first + " " + n.last)
}

// A method can also be declared on non-struct types
// This example is a numeric type 'MyFloat' with an 'Absolute()' method

func (f MyFloat) Absolute() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// POINTER RECEIVERS are more common - The receiver type has the literal syntax *T for some type T
// Methods with pointer receivers directly modify what the receiver points to, and function may have no return
// In this example the Scale() method is defined on *Vertex and directly changes the Vertex value defined in main()

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y / f
}

// Method with value receiver
func (v Vertex) ScaleIt(f float64) float64 {
	v.X = v.X * f
	v.Y = v.Y / f
	return v.X / v.Y * f
}

func main() {

	// SomeMethod() method
	v := Vertex{3, 4}
	fmt.Println(v.SomeMethod())

	// 'n' is Names struct attached to thisMethod() method
	n := Names{"nick", "sinclair"}
	fmt.Println(n.thisMethod())

	// Absolute() method
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Absolute())

	// Scale() method on struct 'v' directly modifies it in function with *Vertex pointer
	v = Vertex{11, 6}
	v.Scale(10)
	fmt.Println(v.SomeMethod())

	// Works in reverse - define variable pointing to &address of value, and dereference like '*p'
	p := &Vertex{5, 6}
	fmt.Println(p.SomeMethod())
	fmt.Println(*p)

	// Methods with pointer receivers "(v *Vertex)" take value or pointer as receiver, "(&g).Scale()" by default
	g := Vertex{6, 24}
	g.Scale(2)
	fmt.Println(g)
	// (g, &h) are same things for the method Scale() with pointer receiver
	h := &g
	h.Scale(10)
	fmt.Println(g)

	// // Methods with value receivers "(v Vertex)" take value or pointer as receiver, "(*i).Scale()" by default
	// (i, &j) are same things for the method Scale() with value receiver
	i := Vertex{2, 1}
	j := &i
	fmt.Println(i.ScaleIt(10))
	fmt.Println(j.ScaleIt(0.5))

}
