package rooms

import (
	"github.com/felipelucio/go-pixel-dungeon/game"
)

var entranceRoomConfig = game.RoomConfig{
	Name: "EntranceRoom",
	Categories: map[game.RoomSizeCategory][3]int{
		game.RoomCategoryNormal: {4, 10, 1},
		game.RoomCategoryLarge:  {10, 14, 2},
		game.RoomCategoryGiant:  {14, 18, 3},
	},
	SizeCatProbs: map[game.RoomSizeCategory]float32{
		game.RoomCategoryNormal: 1,
		game.RoomCategoryLarge:  0,
		game.RoomCategoryGiant:  0,
	},
}

func NewEntranceRoom(sizeCat game.RoomSizeCategory) *game.Room {
	r := game.NewRoom()
	rconf := entranceRoomConfig
	r.Name = rconf.Name

	return r
}
