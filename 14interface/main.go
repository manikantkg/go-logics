package main

import (
	"fmt"
)

// Example
// Declare an Interface Type and methods does not have a body
// type Employee interface {
// 	PrintName() string                // Method with string return type
// 	PrintAddress(id int)              // Method with int parameter
// 	PrintSalary(b int, t int) float64 // Method with parameters and return type
// }

//Define Type that Satisfies an Interface

type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

// Define Type that Satisfies Multiple Interfaces

type Ploygons interface {
	Perimeter()
}

type Object interface {
	NumOfSides()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimetor of Pentagon", 5*p)
}

func (p Pentagon) NumOfSides() {
	fmt.Println("Pentagon has 5 sides")
}

//Interfaces with common Method

type Vehicle interface {
	Structure() []string // Common Method
	Speed() string
}

type Human interface {
	Structure() []string // Common Method
	Performance() string
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}

func (m Man) Performance() string {
	return "8 Hrs/Day"
}

func main() {
	/* var e1 Employee

	   e1 = Emp(1)
	   e1.PrintName("Manikanta")
	   fmt.Println(e1.PrintSalary(50000, 5)) */

	/* var p Ploygons

	   p = Pentagon(50)

	   var o Pentagon = p.(Pentagon)
	   o.NumOfSides()

	   var obj Object = Pentagon(50)
	   obj.NumOfSides()
	   var pent Pentagon = obj.(Pentagon)
	   pent.Perimeter()

	   var bmw Vehicle
	   bmw = Car("World Top Brand")

	   var labour Human
	   labour = Man("Software Developer")

	   for i, j := range bmw.Structure() {
	   	fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	   }*/

	/* var m Methods
	   m = Find(2)
	   m.Area(5) */

	//interview quesion
	// Example usage
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	// Calculate and print areas
	printArea(circle)
	printArea(rectangle)

}

type Methods interface {
	Area(s int)
}

type Find int

func (a Find) Area(s int) {
	square := s * s
	fmt.Println("aera of a triangle is :", square)
	// return square

}

//Interview programme
//create an interface Area() that return mEthod area() and we have to calculate area of any shape

// Shape interface with an Area method
type Shape interface {
	Area() float64
}

// Circle type implementing the Shape interface
type Circle struct {
	Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle type implementing the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Example function to calculate and print the area of any shape implementing the Shape interface
func printArea(s Shape) {
	fmt.Printf("Area of the shape: %.2f\n", s.Area())
}
