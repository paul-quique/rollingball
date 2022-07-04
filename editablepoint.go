package rollingball

import "github.com/hajimehoshi/ebiten/v2"

//EditablePoint can be moved by the user
type EditablePoint struct {
	UUID    string
	Pt      *Point
	Focused bool
	Radius  float64
}

//IsHovered returns true if the point is hovered
func (p EditablePoint) IsHovered(e MouseEvent) bool {
	if Dist(p.Pt, e.Location) <= p.Radius {
		return true
	} else {
		return false
	}
}

//IsFocused returns true if the point is focused
func (p EditablePoint) IsFocused() bool {
	return p.Focused
}

//Update updates the point coordinates and focus state
func (p EditablePoint) Update(e MouseEvent) {
	//check if the point is already beeing moved
	if p.Focused {
		// if the mouse is released then move the point
		//to the new location
		if e.Type == MouseReleased {
			p.Pt.MoveTo(e.Location.X, e.Location.Y)
			ebiten.SetCursorShape(ebiten.CursorShapeDefault)
			p.Focused = false
		}
	} else if e.Type == MousePressed {
		//else set the point to moving
		p.Focused = true
	} else {
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	}
}

//ID returns the point's ID
func (p EditablePoint) ID() string {
	return p.UUID
}
