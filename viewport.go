package rollingball

//Viewport stores the area to render
type Viewport struct {
	MinX, MinY, MaxX, MaxY float64
	Width, Height          float64
	ZoomFactorX            float64
	ZoomFactorY            float64
}

//GetZoom returns the zoom values for the viewport
func (v *Viewport) GetZoom() (zX float64, zY float64) {
	return v.ZoomFactorX, v.ZoomFactorY
}

//SetZoom sets the zoom values for the viewport
func (v *Viewport) SetZoom(zX, zY float64) {
	v.ZoomFactorX, v.ZoomFactorY = zX, zY
}

//MultiplyZoom multiply both x and y zoom factors by zX and zY
func (v *Viewport) MultiplyZoom(zX, zY float64) {
	v.ZoomFactorX *= zX
	v.ZoomFactorY *= zY
	v.MinX *= zX
	v.MinY *= zY
	v.MaxX *= zX
	v.MaxY *= zY
}

func (v *Viewport) MoveX(dx float64) {
	//zoom factor does not change
	v.MinX += dx * v.ZoomFactorX
	v.MaxX += dx * v.ZoomFactorX
}

func (v *Viewport) MoveY(dy float64) {
	//zoom factor does not change
	v.MinX += dy * v.ZoomFactorY
	v.MaxX += dy * v.ZoomFactorY
}

func (v *Viewport) SetBounds(x1, y1, x2, y2 float64) {
	v.MinX = x1
	v.MinY = y1
	v.MaxX = x2
	v.MaxY = y2
	v.ZoomFactorX = (x2 - x1) / v.Width
	v.ZoomFactorY = (y2 - y1) / v.Height
}
