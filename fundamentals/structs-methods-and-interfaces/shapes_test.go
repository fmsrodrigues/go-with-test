package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTest := []struct {
		name string
		shape Shape
		hasPerimeter float64
	} {
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasPerimeter: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{SideA: 3, Base: 12, SideC: 5, Height: 6}, hasPerimeter: 20.0},
	}

	for _, tt := range perimeterTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("got %g want %g", got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	} {
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{SideA: 3, Base: 12, SideC: 5, Height: 6}, hasArea: 36.0},
	}
	
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}