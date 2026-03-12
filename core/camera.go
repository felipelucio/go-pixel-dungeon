package core

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	Debug      bool
}

func NewCamera(sizeW int, sizeH int, minScale float64, maxScale float64, debug bool) Camera {
	buffW := sizeW * int(maxScale)
	buffH := sizeH * int(maxScale)
	return Camera{
		sizeH:    sizeH,
		sizeW:    sizeW,
		posX:     0,
		posY:     0,
		scale:    1.0,
		maxScale: maxScale,
		minScale: minScale,
		screen:   ebiten.NewImage(buffW, buffH),
		Debug:    debug,
	}
}

func (cam *Camera) GetBuffer() *ebiten.Image {
	return cam.screen
}

func (cam *Camera) GetDisplayRect() image.Rectangle {
	startX, startY := cam.GetWorldCoords(0, 0)
	endX, endY := cam.GetWorldCoords(cam.sizeW, cam.sizeH)
	return image.Rect(
		startX, startY,
		endX, endY,
	)
}

func (cam *Camera) ScaleTo(scale float64) {
	cam.scale = math.Max(cam.minScale, math.Min(cam.maxScale, scale))
}

func (cam *Camera) ScaleBy(scaleDt float64) {
	cam.scale = math.Max(cam.minScale, math.Min(cam.maxScale, cam.scale+scaleDt))
}

func (cam *Camera) SmoothScaleBy(scaleDt float64, time float64) {
	cam.tweenScale = NewTween(
		cam.scale,
		cam.scale+scaleDt,
		time,
		TweenEaseInOut,
	)
}

func (cam *Camera) MoveTo(x int, y int) {
	cam.tweenPos.End()
	cam.posX = x
	cam.posY = y
}

func (cam *Camera) MoveBy(dx int, dy int) {
	cam.tweenPos.End()
	cam.posX = int(cam.tweenPos.dest.X) + dx
	cam.posY = int(cam.tweenPos.dest.Y) + dy
}

func (cam *Camera) SmoothMoveBy(dx int, dy int, time float64) {
	camPosX, camPosY := cam.GetPosition()
	oldDestX := cam.tweenPos.dest.X
	oldDestY := cam.tweenPos.dest.Y

	cam.tweenPos = NewTweenVector2(
		NewVector2(float64(camPosX), float64(camPosY)),
		NewVector2(oldDestX+float64(dx), oldDestY+float64(dy)),
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
		cam.ScaleTo(newScale)
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
	if cam.Debug {
		camPosX, camPosY := cam.GetPosition()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Cam Pos: (%d, %d)", camPosX, camPosY))
	}
}

func (cam *Camera) GetPosition() (int, int) {
	return cam.posX, cam.posY
}

// GetScreenCoords converts world coords into screen coords
func (cam *Camera) GetScreenCoords(x int, y int) (int, int) {
	w, h := cam.sizeW, cam.sizeH
	camPosX, camPosY := cam.GetPosition()
	realx, realy := float64(x-camPosX), float64(y-camPosY)
	return int(realx*cam.scale + float64(w)/2), int(realy*cam.scale + float64(h)/2)
}

// GetWorldCoords converts screen coords into world coords
func (cam *Camera) GetWorldCoords(x int, y int) (int, int) {
	w, h := cam.sizeW, cam.sizeH
	camPosX, camPosY := cam.GetPosition()
	realX, realY := float64((x-w)/2)/cam.scale, float64((y-h)/2)/cam.scale
	return int(realX) + camPosX, int(realY) + camPosY
}

// GetCursorCoords converts cursor/screen coords into world coords
func (cam *Camera) GetCursorCoords() (int, int) {
	cx, cy := ebiten.CursorPosition()
	return cam.GetWorldCoords(cx, cy)
}
