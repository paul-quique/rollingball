package rollingball

//Point stores x and y values for a point in 2d space
type Point struct {
	X, Y float64
}

//Line represents a line between point A and point B
type Line struct {
	A, B *Point
}

//CubicBezier represents a cubic Bezier curve
type CubicBezier struct {
	A, B, C, D *Point
}

//MoveTo moves the point to the given coordinates
func (p *Point) MoveTo(x, y float64) {
	p.X, p.Y = x, y
}

//NewPoint returns a pointer to a new point with the given coordinates
func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}
