package shapes

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius *  3.14
}
func RectanglePerimeter(rect Rectangle) (perimeter float64){
	perimeter = 2 * (rect.Width + rect.Height)
	
	return;
}