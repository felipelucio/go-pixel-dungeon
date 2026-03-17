package rooms

import "github.com/felipelucio/go-pixel-dungeon/game"

type StandardRoom struct {
}

var categories = map[game.RoomSizeCategory][3]int{
	game.RoomCategoryNormal: {4, 10, 1},
	game.RoomCategoryLarge:  {10, 14, 2},
	game.RoomCategoryGiant:  {14, 18, 3},
}

var sizeCatProbs = map[game.RoomSizeCategory]float32{
	game.RoomCategoryNormal: 1,
	game.RoomCategoryLarge:  0,
	game.RoomCategoryGiant:  0,
}

func NewStandardRoom(sizeCat game.RoomSizeCategory) *StandardRoom {
	return &StandardRoom{}
}

func (r *StandardRoom) GetRoomType() string {
	return "StandardRoom"
}

func (r *StandardRoom) GetCategories() *map[game.RoomSizeCategory][3]int {
	return &categories
}

func (r *StandardRoom) GetCatProbs() *map[game.RoomSizeCategory]float32 {
	return &sizeCatProbs
}

func (r *StandardRoom) GetMinWidth() int {
	return -1
}

func (r *StandardRoom) GetMinHeight() int {
	return -1
}

func (r *StandardRoom) GetMaxWidth() int {
	return -1
}

func (r *StandardRoom) GetMaxHeight() int {
	return -1
}
