package core

import (
	"math"
)

type TweenType int

const (
	TweenLinear TweenType = iota
	TweenEaseIn
	TweenEaseOut
	TweenEaseInOut
)

type Tween struct {
	orig     float64
	dest     float64
	easing   TweenType
	time     float64
	currTime float64
}

func NewTween(orig float64, dest float64, time float64, easing TweenType) Tween {
	return Tween{
		orig,
		dest,
		easing,
		time,
		0.0,
	}
}

func (tween *Tween) Update(delta float64) (float64, bool) {
	tween.currTime = math.Min(tween.currTime+delta, tween.time)
	pct := tween.currTime / tween.time
	return tween.orig + (tween.dest-tween.orig)*pct, tween.IsFinished()
}

func (tween *Tween) IsFinished() bool {
	return tween.currTime == tween.time
}

type TweenVector2 struct {
	orig     Vector2[float64]
	dest     Vector2[float64]
	easing   TweenType
	time     float64
	currTime float64
}

func NewTweenVector2(orig Vector2[float64], dest Vector2[float64], time float64, easing TweenType) TweenVector2 {
	return TweenVector2{
		orig,
		dest,
		easing,
		time,
		0.0,
	}
}

func (tween *TweenVector2) Update(delta float64) (Vector2[float64], bool) {
	tween.currTime = math.Min(tween.currTime+delta, tween.time)
	pct := tween.currTime / tween.time
	return Vector2[float64]{
		tween.orig.X + (tween.dest.X-tween.orig.X)*pct,
		tween.orig.Y + (tween.dest.Y-tween.orig.Y)*pct,
	}, tween.IsFinished()
}

func (tween *TweenVector2) IsFinished() bool {
	return tween.currTime == tween.time
}

type TweenVector3 struct {
	orig     Vector3[float64]
	dest     Vector3[float64]
	easing   TweenType
	time     float64
	currTime float64
}

func NewTweenVector3(orig Vector3[float64], dest Vector3[float64], time float64, easing TweenType) TweenVector3 {
	return TweenVector3{
		orig,
		dest,
		easing,
		time,
		0.0,
	}
}

func (tween *TweenVector3) Update(delta float64) (Vector3[float64], bool) {
	tween.currTime = math.Min(tween.currTime+delta, tween.time)
	pct := tween.currTime / tween.time
	return Vector3[float64]{
		tween.orig.X + (tween.dest.X-tween.orig.X)*pct,
		tween.orig.Y + (tween.dest.Y-tween.orig.Y)*pct,
		tween.orig.Z + (tween.dest.Z-tween.orig.Z)*pct,
	}, tween.IsFinished()
}

func (tween *TweenVector3) IsFinished() bool {
	return tween.currTime == tween.time
}
