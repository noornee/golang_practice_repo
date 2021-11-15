package main

import (
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Area of a Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.00
		got := rectangle.Area()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("Area of a circle", func(t *testing.T) {

		circle := Circle{10.0}
		want := 314.1592653589793
		got := circle.Area()

		if want != got {
			t.Errorf("want %g got %g", want, got)
		}

	})

}

func TestPerimeter(t *testing.T) {

	t.Run("Perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{8.0, 5.0}
		want := 26.00
		got := rectangle.Perimeter()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("Perimeter of a circle", func(t *testing.T) {
		circle := Circle{15.05}
		want := 94.56193887305278
		got := circle.Perimeter()

		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})

}



// Table driven test

// func TestVal(t *testing.T) {

// 	areaTest := []struct {
// 		name string
// 		shape Shape
// 		want  float64
// 	}{
// 		{name: "Rectangle",shape: Rectangle{10.0, 10.0}, want: 100.0},
// 		{name: "Circle", shape: Circle{10.0}, want: 314.159265358979},
// 	}

// 	for _, ar := range areaTest {
// 		t.Run(ar.name, func(t *testing.T) {
// 			got := ar.shape.Area()
// 			want := ar.want
	
// 			if got != want {
// 				t.Errorf("want %g got %g", want, got)
// 			}
// 		})

// 	}

// }
