package main

import (
	"fmt"
	"log"

	"github.com/felipelucio/go-pixel-dungeon/core"
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Title string
}

func (g *Game) Update() error {

	ebiten.SetWindowTitle(fmt.Sprintf("%s (FPS: %.2f)", g.Title, ebiten.ActualFPS()))
	return core.UpdateScene()
}

func (g *Game) Draw(screen *ebiten.Image) {
	core.DrawScene(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 768 / 2, 432 / 2
}

func main() {
	g := Game{"Go Pixel Dungeon"}

	gameState := core.GameState{}
	core.SceneManagerRegisterGameState(&gameState)
	core.SwitchToScene(&scenes.TestScene{})

	ebiten.SetWindowSize(game.Config.WinWidth, game.Config.WinHeight)
	if game.Config.WinMode == game.Fullscreen {
		ebiten.SetFullscreen(true)
	}
	ebiten.SetVsyncEnabled(game.Config.VsyncEnabled)
	ebiten.SetWindowTitle(g.Title)
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
