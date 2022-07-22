package rollingball

var (
	CurrentTool int
	Points      []*EditablePoint
	Beziers     []CubicBezier
	Lines       []Line
	Arcs        []Arc
)

const (
	LineTool int = iota
	ArcTool
	CubicBezierTool
	EraseTool
)

func UpdateEditor() {
	switch CurrentTool {
	case LineTool:
	case ArcTool:
	case CubicBezierTool:
	case EraseTool:
	}
}
