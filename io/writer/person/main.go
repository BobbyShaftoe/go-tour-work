package main

import (
	"fmt"
	_ "fmt"
	_ "strings"

	"github.com/pkg/errors"
)

const ssnLen = 11

type Person struct {
	name string
	age  uint8
	ssn  string
}

func main() {
	fmt.Println("Person implementation of io.writer")
	/*
		Format: [
		<name-len>: byte,
		<bytes of name>: []byte,
		<age>: byte,
		<bytes of ssn>: [11]byte,
		]
	*/
	var data = []byte{
		5, 66, 111, 98, 98, 121, 35, 57, 57, 49, 45, 57, 48, 45, 49, 48, 51, 50,
	}

	fmt.Println(data)

	var p Person

	fmt.Println(p)

	_, err := p.Write(data)

	fmt.Printf("err: %v, %v\n", err, p)
}

func (r *Person) Write(p []byte) (n int, err error) {
	if r == nil {
		return 0, errors.New("Invalid receiver")
	}
	if 0 == len(p) {
		return 0, errors.New("Invalid parameter")
	}

	nameLen := int(p[0])
	if (nameLen + 1) > len(p) {
		return 0, errors.New("Insufficient data for name")
	}

	r.name = string(p[1 : nameLen+1])

	if (nameLen + 2) > len(p) {
		return 0, errors.New("Insufficient data for age")
	}

	r.age = uint8(p[nameLen+1])

	if ssnLen < len(p[nameLen+2:]) {
		return 0, errors.New("Insufficient data for ssn")
	}

	r.ssn = string(p[nameLen+2 : nameLen+2+ssnLen])

	return ssnLen + 2 + nameLen, nil
}

// default string function for AgeAverager type
func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v, ssn: %v)", p.name, p.age, p.ssn)
}
