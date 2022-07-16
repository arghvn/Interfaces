// Implement Go Interface by Multiple Structs
// In Go, more than 1 struct can also implement a single interface. For example,

package main

import "fmt"

// interface
type Shape interface {
	area() float32
}

// Rectangle struct implements the interface
type Rectangle struct {
	length, breadth float32
}

// Rectangle provides implementation for area()
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// Triangle struct implements the interface
type Triangle struct {
	base, height float32
}

// Triangle provides implementation for area()
func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}

// access method of the interface
func calculate(s Shape) float32 {
	return s.area()
}

// main function
func main() {

	// assigns value to struct members
	r := Rectangle{7, 4}
	t := Triangle{8, 12}

	// call calculate() with struct variable rect
	rectangleArea := calculate(r)
	fmt.Println("Area of Rectangle:", rectangleArea)

	triangleArea := calculate(t)
	fmt.Println("Area of Triangle:", triangleArea)

}

// Output

// Area of Rectangle: 28
// Area of Rectangle: 48
// In the above example, we have used two structs: Rectangle and Triangle to implement the interface Shape.

// Just like before, both structs provide implementation for the area() method.

// Here, this time calculate() calls the area() using interface Shape and returns it.

// So, for the first call, calculate(r), the method will call the area() method implementation of Rectangle.
// Similarly, for calculate(t), the method will call the area() method implementation of Triangle.
