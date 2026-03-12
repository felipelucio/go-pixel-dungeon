package core

import (
	"errors"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tilemap struct {
	width          int
	height         int
	pixelWidth     int
	pixelHeight    int
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
	tm.pixelWidth = ts.tileWidth * tm.width
	tm.pixelHeight = ts.tileHeight * tm.height

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

func (tm *Tilemap) getID(x int, y int) (int, error) {
	if x < 0 || x >= tm.width || y < 0 || y >= tm.height {
		return -1, errors.New("out of bounds")
	}
	return (y * tm.width) + x, nil
}

func (tm *Tilemap) GetWidth() int {
	return tm.width
}

func (tm *Tilemap) GetHeight() int {
	return tm.height
}

func (tm *Tilemap) GetPixelWidth() int {
	return tm.pixelWidth
}

func (tm *Tilemap) GetPixelHeight() int {
	return tm.pixelHeight
}

func (tm *Tilemap) SetTile(x int, y int, tileID int) {
	tid, err := tm.getID(x, y)
	if err == nil {
		tm.tiles[tid] = tileID
	}
}

func (tm *Tilemap) GetTile(x int, y int) (*ebiten.Image, error) {
	tmID, err := tm.getID(x, y)
	if err != nil {
		return nil, err
	}
	tid := tm.tiles[tmID]
	return tm.tileset.Get(tid), nil
}

func (tm *Tilemap) SetWalkable(x int, y int, walkable bool) {
	tid, err := tm.getID(x, y)
	if err == nil {
		tm.tiles_walkable[tid] = walkable
	}
}

func (tm *Tilemap) IsWalkable(x int, y int) (bool, error) {
	tid, err := tm.getID(x, y)
	if err != nil {
		return false, err
	}
	return tm.tiles_walkable[tid], nil
}

func (tm *Tilemap) SetVisible(x int, y int, visible bool) {
	tid, err := tm.getID(x, y)
	if err == nil {
		tm.tiles_visible[tid] = visible
	}
}

func (tm *Tilemap) IsVisible(x int, y int) (bool, error) {
	tid, err := tm.getID(x, y)
	if err != nil {
		return false, err
	}
	return tm.tiles_visible[tid], nil
}

func (tm *Tilemap) SetVisited(x int, y int, visited bool) {
	tid, err := tm.getID(x, y)
	if err == nil {
		tm.tiles_visited[tid] = visited
	}
}

func (tm *Tilemap) IsVisited(x int, y int) (bool, error) {
	tid, err := tm.getID(x, y)
	if err != nil {
		return false, err
	}
	return tm.tiles_visited[tid], nil
}
