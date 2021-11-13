package shapes

import "testing"

func assertResult(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		assertResult(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 20.0

		assertResult(t, got, want)
	})
}

func TestArea(t *testing.T) {

	tests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range tests {
		got := tt.shape.Area()
		assertResult(t, got, tt.want)

	}
}
