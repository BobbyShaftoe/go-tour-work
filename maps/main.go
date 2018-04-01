package main

import (
	_ "fmt"
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex


// Map literal - keys are required
var n = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// If the top-level type is just a type name, you can omit it from the elements of the literal
var o = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}

var p = map[int]Vertex{
	1: {40.68433, -74.39967},
	2: {37.42202, -122.08408},
}


func main()() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)

	// Mutate maps
	q := make(map[string]int)

	q["Answer"] = 42
	fmt.Println("The value:", q["Answer"])

	delete(q, "Answer")
	fmt.Println("The value:", q["Answer"])

	v, ok := q["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}






// import (
// //	"golang.org/x/tour/wc"
// "fmt"
// "strings"
// )
//
// func WordCount(s string) map[string]int {
// 	m := make(map[string]int)
// 	for i, f := range strings.Fields(s) {
// 		fmt.Println(f)
// 		m[f] = i
// 	}
// 	return m
// }
//
// func main() {
// 	x := WordCount("this is a test")
// 	fmt.Println(x)
// 	//wc.Test(WordCount)
// }