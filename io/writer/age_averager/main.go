package main

import (
	"fmt"
	_ "fmt"
	_ "strings"

	"github.com/pkg/errors"
)

type Person struct {
	name string
	age  uint8
}

type AgeAverager struct {
	sum   uint64
	count uint64
}

func main() {
	fmt.Println("Age averager writer")

	var c AgeAverager
	bob := Person{"Bobby", 18}
	jane := Person{"Jane Doe", 35}
	mary := Person{"Mary", 17}

	c.Write(bob.AgeToByte())
	c.Write(jane.AgeToByte())
	c.Write(mary.AgeToByte())

	var p []byte
	n, err := c.Write(p)
	fmt.Printf("%v, %v\n", n, err)
	n, err = c.Write([]byte{210})
	fmt.Printf("%v, %v\n", n, err)

	// note the print of variable 'c', see default String() function below
	fmt.Printf("Age average: %v\n", c)

}

func (p Person) AgeToByte() []byte {
	var buf = []byte{byte(p.age)}
	return buf
}

func (c *AgeAverager) Write(p []byte) (n int, err error) {
	l := len(p)
	switch {
	case p == nil:
		return 0, errors.New("Invalid parameter")
	case 127 < p[0]:
		return 0, errors.New("Age must be less than or equal to 127")
	}
	c.count++
	c.sum += uint64(p[0])
	return l, nil
}

// default string function for AgeAverager type
func (c AgeAverager) String() string {
	return fmt.Sprintf("%v", c.sum/c.count)
}
