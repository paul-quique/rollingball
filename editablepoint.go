package rollingball

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

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
	if e.Type == MouseReleased {
		fmt.Println("mouse relachee")
		p.Pt.MoveTo(e.Location.X, e.Location.Y)
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		p.Focused = false
	} else if e.Type == MousePressed {
		p.Focused = true
	}
}

//ID returns the point's ID
func (p EditablePoint) ID() string {
	return p.UUID
}
