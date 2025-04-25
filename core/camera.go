package core

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	sizeH      int
	sizeW      int
	posX       int
	posY       int
	scale      float64
	maxScale   float64
	minScale   float64
	screen     *ebiten.Image
	tweenPos   TweenVector2
	tweenScale Tween
}

func NewCamera(sizeW int, sizeH int, minScale float64, maxScale float64) Camera {
	return Camera{
		sizeH:    sizeH,
		sizeW:    sizeW,
		posX:     0,
		posY:     0,
		scale:    1.0,
		maxScale: maxScale,
		minScale: minScale,
		screen:   ebiten.NewImage(sizeW, sizeH),
	}
}

func (cam *Camera) GetBuffer() *ebiten.Image {
	return cam.screen
}

func (cam *Camera) resize() {
	newW := int(float64(cam.sizeW) * 1.0 / cam.scale)
	newH := int(float64(cam.sizeH) * 1.0 / cam.scale)
	cam.screen.Deallocate()
	cam.screen = ebiten.NewImage(newW, newH)
}

func (cam *Camera) ScaleTo(scale float64) {
	cam.scale = math.Max(cam.minScale, math.Min(cam.maxScale, cam.scale))
	cam.resize()
}

func (cam *Camera) ScaleBy(scaleDt float64) {
	cam.scale = math.Max(cam.minScale, math.Min(cam.maxScale, cam.scale+scaleDt))
	cam.resize()
}

func (cam *Camera) MoveTo(x int, y int) {
	cam.posX = x
	cam.posY = y
}

func (cam *Camera) MoveBy(dx int, dy int) {
	cam.posX += dx
	cam.posY += dy
}

func (cam *Camera) SmoothMoveBy(dx int, dy int, time float64) {
	cam.tweenPos = NewTweenVector2(
		NewVector2(float64(cam.posX), float64(cam.posY)),
		NewVector2(float64(cam.posX+dx), float64(cam.posY+dy)),
		time,
		TweenEaseOut,
	)
}

func (cam *Camera) Update(delta float64) {
	if !cam.tweenPos.IsFinished() {
		newPos, _ := cam.tweenPos.Update(delta)
		cam.posX = int(newPos.X)
		cam.posY = int(newPos.Y)
	}

	if !cam.tweenScale.IsFinished() {
		newScale, _ := cam.tweenScale.Update(delta)
		cam.scale = newScale
	}
}

func (cam *Camera) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w := cam.GetBuffer().Bounds().Dx()
	h := cam.GetBuffer().Bounds().Dy()
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0

	op.GeoM.Translate(-cx, -cy)
	op.GeoM.Scale(cam.scale, cam.scale)
	op.GeoM.Translate(cx*cam.scale, cy*cam.scale)

	screen.DrawImage(cam.GetBuffer(), op)
}

func (cam *Camera) GetPosition() (int, int) {
	return cam.posX, cam.posY
}

// GetScreenCoords converts world coords into screen coords
func (cam *Camera) GetScreenCoords(x int, y int) (int, int) {
	w, h := cam.sizeW, cam.sizeH
	realx, realy := float64(x-cam.posX), float64(y-cam.posY)
	return int(realx*cam.scale + float64(w)/2), int(realy*cam.scale + float64(h)/2)
}

// GetWorldCoords converts screen coords into world coords
func (cam *Camera) GetWorldCoords(x int, y int) (int, int) {
	w, h := cam.sizeW, cam.sizeH
	realx, realy := float64((x-w)/2)/cam.scale, float64((y-h)/2)/cam.scale
	return int(realx) + cam.posX, int(realy) + cam.posY
}

// GetCursorCoords converts cursor/screen coords into world coords
func (cam *Camera) GetCursorCoords() (int, int) {
	cx, cy := ebiten.CursorPosition()
	return cam.GetWorldCoords(cx, cy)
}
