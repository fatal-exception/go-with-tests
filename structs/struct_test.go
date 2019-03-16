package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got '%.2f', want '%.2f'", got, want)
		}
	}

	t.Run("Test perimeter of rectangle", func(t *testing.T) {
		r := Rectangle{1.4, 5.6}
		want := 2*1.4 + 2*5.6
		checkPerimeter(t, r, want)
	})
	t.Run("Test perimeter of circle", func(t *testing.T) {
		c := Circle{3}
		want := 2 * math.Pi * 3
		checkPerimeter(t, c, want)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 4, Width: 5}, want: 20},
		{name: "Circle", shape: Circle{Radius: 4}, want: math.Pi * 16},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want
			if got != want {
				t.Errorf("%#v got '%.2f', want '%.2f'", tt.shape, got, want)
			}
		})
	}
}
