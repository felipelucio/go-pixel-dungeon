package game

import "github.com/felipelucio/go-pixel-dungeon/core"

type RoomDef interface {
	GetRoomType() string
	GetCategories() *map[RoomSizeCategory][3]int
	GetCatProbs() *map[RoomSizeCategory]float32

	GetMinWidth() int
	GetMinHeight() int
	GetMaxWidth() int
	GetMaxHeight() int
}

type RoomSizeCategory int

const (
	RoomCategoryNormal RoomSizeCategory = iota
	RoomCategoryLarge
	RoomCategoryGiant
)

type Room struct {
	core.Rect

	Neighbours []*Room
	// connected []Door
	Distance int
	Price    int
}

func NewRoom(x, y, w, h int) *Room {
	return &Room{
		Rect: core.Rect{
			X1: x,
			Y1: y,
			X2: x + w,
			Y2: y + h,
		},
	}
}
