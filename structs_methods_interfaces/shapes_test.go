package structs_methods_interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("", func(t *testing.T) {
		circle := Circle{1}
		checkArea(t, circle, 3.141592653589793)
	})

	//t.Run("rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{12.0, 6.0}
	//	got := rectangle.Area()
	//	want := 72.0

	//	if got != want {
	//		t.Errorf("got %.2f want %.2f", got, want)
	//	}
	//})

	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{1}
	//	got := circle.Area()
	//	want := 3.141592

	//	if got != want {
	//		t.Errorf("got  %g, %f\n", got, got)
	//		t.Errorf("want %g, %f\n", want, want)
	//	}
	//})
}
