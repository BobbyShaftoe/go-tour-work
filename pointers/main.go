package main

import (
	_ "fmt"
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i				// point to i
	fmt.Println(p, &i)	// p points to address of i (&i)
	fmt.Println(*p) 	// read actual value of i through *p pointer

	*p = 21				// set value of i through the pointer to i
	fmt.Println(i)

	p = &j				// point to j (address)
    fmt.Println(p)

	*p = *p / 37		// divide j through the *p pointer to j
	fmt.Println(j)
}
