package structs

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("%#v got %g want %g", shape, got, want)
	// 	}
	// }

	// t.Run(("rectangles"), func(t *testing.T) {

	// 	rectangle := Rectangle{12.0, 6.0}
	// 	// got := rectangle.Area()
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run(("circles"), func(t *testing.T) {

	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
