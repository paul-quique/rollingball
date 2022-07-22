package rollingball

import "fmt"

var (
	CurrentTool int
	Points      []EditablePoint
	Beziers     []CubicBezier
	Lines       []Line
	Arcs        []Arc
)

const EditablePointRadius = 10

const (
	CubicBezierTool int = iota
	ArcTool
	LineTool
	EraseTool
)

func UpdateEditor(px, py float64) {
	switch CurrentTool {
	case LineTool:
		isFocused := false
		cp := EditablePoint{
			UUID:    RandomID(),
			Focused: &isFocused,
			Radius:  EditablePointRadius,
			Pt:      NewPoint(px, py),
		}
		Points = append(Points, cp)
		AddMouseListener(cp)

		if len(Points) == 2 {
			Lines = append(Lines, Line{A: Points[0], B: Points[1]})
			Points = []EditablePoint{}
		}
	case ArcTool:
		isFocused := false
		cp := EditablePoint{
			UUID:    RandomID(),
			Focused: &isFocused,
			Radius:  EditablePointRadius,
			Pt:      NewPoint(px, py),
		}
		Points = append(Points, cp)
		AddMouseListener(cp)

		if len(Points) == 3 {
			Arcs = append(Arcs, Arc{Center: Points[0], A: Points[1], B: Points[2]})
			Points = []EditablePoint{}
		}
	case CubicBezierTool:
		fmt.Print("c")
		isFocused := false
		cp := EditablePoint{
			UUID:    RandomID(),
			Focused: &isFocused,
			Radius:  EditablePointRadius,
			Pt:      NewPoint(px, py),
		}
		Points = append(Points, cp)
		AddMouseListener(cp)

		if len(Points) == 4 {
			Beziers = append(Beziers, CubicBezier{A: Points[0], B: Points[1], C: Points[2], D: Points[3]})
			Points = []EditablePoint{}
		}
	case EraseTool:
	}
}
