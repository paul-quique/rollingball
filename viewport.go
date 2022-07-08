package rollingball

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//Viewport stores the area to render
type Viewport struct {
	OffsetX, OffsetY         float64
	MinX, MinY, MaxX, MaxY   float64
	Width, Height            float64
	ZoomFactorX, ZoomFactorY float64
}

//SetZoom sets the zoom values for the viewport
func (v *Viewport) SetZoom(zX, zY float64) {
	v.ZoomFactorX, v.ZoomFactorY = zX, zY

	v.MinX = v.Width/2 - v.Width*(1./(2*v.ZoomFactorX))
	v.MaxX = v.Width/2 + v.Width*(1./(2*v.ZoomFactorX))
	v.MinY = v.Height/2 - v.Width*(1./(2*v.ZoomFactorY))
	v.MaxY = v.Height/2 + v.Width*(1./(2*v.ZoomFactorY))
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

//MoveX moves the viewport along the x-axis
func (v *Viewport) MoveX(dx float64) {
	v.OffsetX += dx
}

//MoveY moves the viewport along the y-axis
func (v *Viewport) MoveY(dy float64) {
	v.OffsetY += dy
}

//Update updates the viewport if certain keys are pressed
func (v *Viewport) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		v.MoveX(1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		v.MoveX(-1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		v.MoveY(-1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		v.MoveY(1)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		v.MultiplyZoom(1.25, 1.25)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyBackspace) {
		v.MultiplyZoom(0.8, 0.8)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		v.SetZoom(1, 1)
	}
}

//GetMinPosX returns the minimum X value of the viewport
func (v *Viewport) GetMinPosX() float64 {
	return v.MinX + v.OffsetX
}

//GetMinPosY returns the minimum Y value of the viewport
func (v *Viewport) GetMinPosY() float64 {
	return v.MinY + v.OffsetY
}

//GetMousePosition computes and returns the mouse position in the game
func (v *Viewport) GetMousePosition() (x, y float64) {
	cx, cy := ebiten.CursorPosition()
	return (float64(cx) - v.GetMinPosX()) * v.ZoomFactorX, (float64(cy) - v.GetMinPosY()) * v.ZoomFactorY
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
