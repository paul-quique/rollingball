package rollingball

//Viewport stores the area to render
type Viewport struct {
	MinX, MinY, MaxX, MaxY float64
	Width, Height          float64 //width and height in pixels
}

//SetMinX sets the value of MinX
func (v *Viewport) SetMinX(x float64) {
	v.MinX = x
}

//SetMinY sets the value of MinX
func (v *Viewport) SetMinY(y float64) {
	v.MinY = y
}

//SetMaxX sets the value of MaxX
func (v *Viewport) SetMaxX(x float64) {
	v.MaxX = x
}

//SetMaxY sets the value of MaxY
func (v *Viewport) SetMaxY(y float64) {
	v.MaxX = y
}

//GetScale calculates and returns the scale values for the viewport
func (v Viewport) GetScale() (scaleX float64, scaleY float64) {
	scaleX = v.Width / (v.MaxX - v.MinX)
	scaleY = v.Height / (v.MaxY - v.MinY)
	return
}
