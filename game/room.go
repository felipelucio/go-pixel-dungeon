package game

import "github.com/felipelucio/go-pixel-dungeon/core"

type RoomSizeCategory int

const (
	RoomCategoryNormal RoomSizeCategory = iota
	RoomCategoryLarge
	RoomCategoryGiant
)

type RoomConfig struct {
	Name string
	// Min, Max, Value
	Categories   map[RoomSizeCategory][3]int
	SizeCatProbs map[RoomSizeCategory]float32
	IsEntrance   bool
	IsExit       bool
	IsSpecial    bool
}

type Room struct {
	core.Rect

	Name       string
	Neighbours []*Room
	// connected []Door
	Distance   int
	Price      int
	IsEntrance bool
	IsExit     bool
	IsSpecial  bool
}

func NewRoom() *Room {
	return &Room{
		Rect: core.Rect{
			X1: 0,
			Y1: 0,
			X2: 0,
			Y2: 0,
		},
		Name:       "Room",
		Neighbours: make([]*Room, 4),
		// connected: make([]*Door, 4),
		Distance:   0,
		Price:      0,
		IsEntrance: false,
		IsExit:     false,
		IsSpecial:  false,
	}
}
