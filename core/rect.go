package core

type Rect struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func NewRect(x1, y1, x2, y2 int) *Rect {
	return &Rect{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}

func (r *Rect) Width() int {
	return r.X2 - r.X1
}

func (r *Rect) Height() int {
	return r.Y2 - r.Y1
}
