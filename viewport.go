package rollingball

//Viewport stores the area to render
type Viewport struct {
	OffsetX, OffsetY       float64
	MinX, MinY, MaxX, MaxY float64
	Width, Height          float64
	ZoomFactorX            float64
	ZoomFactorY            float64
}

//SetZoom sets the zoom values for the viewport
func (v *Viewport) SetZoom(zX, zY float64) {
	v.ZoomFactorX, v.ZoomFactorY = zX, zY
}

//MultiplyZoom multiply both x and y zoom factors by zX and zY
func (v *Viewport) MultiplyZoom(zX, zY float64) {
	v.ZoomFactorX *= zX
	v.ZoomFactorY *= zY

	v.MinX = v.Width/2 - v.Width*(1./(2*v.ZoomFactorX))
	v.MaxX = v.Width/2 + v.Width*(1./(2*v.ZoomFactorX))
	v.MinY = v.Height/2 - v.Width*(1./(2*v.ZoomFactorY))
	v.MaxY = v.Height/2 + v.Width*(1./(2*v.ZoomFactorY))
}

func (v *Viewport) MoveX(dx float64) {
	v.OffsetX += dx
}

func (v *Viewport) MoveY(dy float64) {
	v.OffsetY += dy
}

func (v *Viewport) SetBounds(x1, y1, x2, y2 float64) {
	v.MinX = x1
	v.MinY = y1
	v.MaxX = x2
	v.MaxY = y2
	v.ZoomFactorX = (x2 - x1) / v.Width
	v.ZoomFactorY = (y2 - y1) / v.Height
}

func (v *Viewport) GetMinPosX() float64 {
	return v.MinX + v.OffsetX
}

func (v *Viewport) GetMinPosY() float64 {
	return v.MinY + v.OffsetY
}

func dist(a, b float64) float64 {
	return abs(b - a)
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
