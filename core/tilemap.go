package core

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tilemap struct {
	width          int
	height         int
	tileset        *Tileset
	tiles          []int
	tiles_walkable []bool
	tiles_visible  []bool
	tiles_visited  []bool
}

func NewTilemap(ts *Tileset, width int, height int) Tilemap {
	tm := Tilemap{}
	tm.width = width
	tm.height = height
	tm.tileset = ts

	totalTiles := width * height
	tm.tiles = make([]int, 0, totalTiles)
	tm.tiles_walkable = make([]bool, 0, totalTiles)
	tm.tiles_visible = make([]bool, 0, totalTiles)
	tm.tiles_visited = make([]bool, 0, totalTiles)

	for range totalTiles {
		tm.tiles = append(tm.tiles, rand.Intn(8))
		tm.tiles_walkable = append(tm.tiles_walkable, true)
		tm.tiles_visible = append(tm.tiles_visible, true)
		tm.tiles_visited = append(tm.tiles_visited, true)
	}

	return tm
}

func (tm *Tilemap) getID(x int, y int) int {
	return (y * tm.width) + x
}

func (tm *Tilemap) GetWidth() int {
	return tm.width
}

func (tm *Tilemap) GetHeight() int {
	return tm.height
}

func (tm *Tilemap) SetTile(x int, y int, tileID int) {
	tid := tm.getID(x, y)
	tm.tiles[tid] = tileID
}

func (tm *Tilemap) GetTile(x int, y int) *ebiten.Image {
	tid := tm.tiles[tm.getID(x, y)]
	return tm.tileset.Get(tid)
}

func (tm *Tilemap) SetWalkable(x int, y int, walkable bool) {
	tid := tm.getID(x, y)
	tm.tiles_walkable[tid] = walkable
}

func (tm *Tilemap) IsWalkable(x int, y int) bool {
	tid := tm.getID(x, y)
	return tm.tiles_walkable[tid]
}

func (tm *Tilemap) SetVisible(x int, y int, visible bool) {
	tid := tm.getID(x, y)
	tm.tiles_visible[tid] = visible
}

func (tm *Tilemap) IsVisible(x int, y int) bool {
	tid := tm.getID(x, y)
	return tm.tiles_visible[tid]
}

func (tm *Tilemap) SetVisited(x int, y int, visited bool) {
	tid := tm.getID(x, y)
	tm.tiles_visited[tid] = visited
}

func (tm *Tilemap) IsVisited(x int, y int) bool {
	tid := tm.getID(x, y)
	return tm.tiles_visited[tid]
}
