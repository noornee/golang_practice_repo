package main

import (
	"testing"
)

func TestMagnitude(t *testing.T) {

	t.Run("Length of a 2-component vector", func(t *testing.T) {
		vector := Vector2{4.0, 3.0}
		
		want := 5.00
		got := vector.Magnitude()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("Length of a 3-component vector", func(t *testing.T) {

		vector := Vector3{2.0, 4.0, 6.0}
		want := 7.483314773547883
		got := vector.Magnitude()

		if want != got {
			t.Errorf("want %g got %g", want, got)
		}

	})

}
