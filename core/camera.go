package core

import "github.com/hajimehoshi/ebiten/v2"

type Camera struct {
	sizeH      int
	sizeW      int
	posX       int
	posY       int
	scale      float64
	screen     *ebiten.Image
	tweenPos   TweenVector2
	tweenScale Tween
}

func NewCamera(screen *ebiten.Image, sizeW int, sizeH int) Camera {
	return Camera{
		sizeH:  sizeH,
		sizeW:  sizeW,
		screen: screen,
		posX:   0,
		posY:   0,
		scale:  1.0,
	}
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
		TweenEaseInOut,
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

func (cam *Camera) GetPosition() (int, int) {
	return cam.posX, cam.posY
}
