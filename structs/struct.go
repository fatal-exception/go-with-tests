package structs

import "math"

// Perimeter calculates perimeter of an object
func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

// Area return area of object
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Shape is
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle is
type Rectangle struct {
	Height float64
	Width  float64
}

// Circle is
type Circle struct {
	Radius float64
}

// Triangle is
type Triangle struct {
	Base   float64
	Height float64
}

// Area is
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter is
func (t Triangle) Perimeter() float64 {
	return 0.0 // can't actually do this without side lengths
}

// Perimeter is a type of fish
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area return area of object
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
