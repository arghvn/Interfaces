// we will learn about the interface and its implementation in Go programming with the help of examples.

// In Go programming, we use interfaces to store a set of methods without implementation. That is, methods of interface won't have a method body.
// Here, Shape is an interface with methods: area() and perimeter(). You can see both methods only have method signatures without any implementation.

// to use an interface, we first need to implement it by a type (struct). To implement an interface,
// a struct should provide implementations for all methods of an interface. For example,

package main

import "fmt"

// interface
type Shape interface {
	area() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println("Area:", s.area())
}

// main function
func main() {

	// assigns value to struct members
	rect := Rectangle{7, 4}

	// call calculate() with struct variable rect
	calculate(rect)

}
