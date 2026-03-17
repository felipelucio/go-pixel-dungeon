package game

import (
	"github.com/felipelucio/go-pixel-dungeon/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Title     string
	GameState *core.GameState
}

func NewGame() *Game {
	return &Game{
		Title:     "Game",
		GameState: core.NewGameState("Hero"),
	}
}

func (g *Game) Update() error {
	return GetSession().SceneManager().UpdateScene()
}

func (g *Game) Draw(screen *ebiten.Image) {
	GetSession().SceneManager().DrawScene(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 768 / 2, 432 / 2
}
