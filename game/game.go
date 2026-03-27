package game

import (
	"github.com/felipelucio/go-pixel-dungeon/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return core.DefaultSceneManager().UpdateScene()
}

func (g *Game) Draw(screen *ebiten.Image) {
	core.DefaultSceneManager().DrawScene(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 768 / 2, 432 / 2
}
