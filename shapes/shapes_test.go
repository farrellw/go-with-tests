package shapes

import (
	"testing"
)

func floatCompareHelper(t *testing.T, got, want float64){
	t.Helper()

	if got != want {
		t.Errorf("got '%.2f', but want '%.2f'", got, want)
	}
}

func TestRectanglePerimeter(t *testing.T){
	t.Run("Returns the perimeter of a rectangle", func(t *testing.T){
		rectangle := Rectangle{10.0, 10.0}
		got := RectanglePerimeter(rectangle)
		want := 40.0

		floatCompareHelper(t, got, want)
	})

}

func TestArea(t *testing.T){
	areaTests := []struct {
		shapeName string
		shape Shape
		want float64
	}{
		{shapeName: "Rectangle", shape: Rectangle{7.0, 8.5}, want: 59.5},
		{shapeName: "Circle", shape: Circle{4.0}, want: 50.24},
		{shapeName: "Triangle", shape: Triangle{3.0, 9.5}, want: 14.25},
	}

	for _, tt := range areaTests {
		t.Run(tt.shapeName, func(t *testing.T){
			got := tt.shape.Area()
		
			floatCompareHelper(t, got, tt.want)
		})
	}
}