package core

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tileset struct {
	name       string
	image      *ebiten.Image
	tileWidth  int
	tileHeight int

	tiles []*ebiten.Image
}

func NewTileset(name string, imgPath string, tileW int, tileH int) (Tileset, error) {
	ts := Tileset{}
	img, _, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		return ts, err
	}
	ts.name = name
	ts.image = img
	ts.tileHeight = tileH
	ts.tileWidth = tileW

	imgSize := img.Bounds()
	cols := imgSize.Dx() / tileW
	rows := imgSize.Dy() / tileH
	ts.tiles = make([]*ebiten.Image, 0, cols*rows)

	for y := int(0); y < rows; y++ {
		for x := int(0); x < cols; x++ {
			dx := x * tileW
			dy := y * tileH
			rect := image.Rect(dx, dy, dx+tileW, dy+tileH)
			ts.tiles = append(ts.tiles, img.SubImage(rect).(*ebiten.Image))
		}
	}

	return ts, nil
}

func (ts *Tileset) Get(id int) *ebiten.Image {
	return ts.tiles[id]
}

func (ts *Tileset) GetTileWidth() int {
	return ts.tileWidth
}

func (ts *Tileset) GetTileHeight() int {
	return ts.tileHeight
}
