package main

import (
	"log"

	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	s := game.NewSession(g)
	s.Audio().LoadMusic(game.AssetMusic)
	s.Audio().LoadSFX(game.AssetSounds)

	s.SceneManager().SwitchToScene(&scenes.TestScene{})

	ebiten.SetWindowSize(game.Config.WinWidth, game.Config.WinHeight)
	if game.Config.WinMode == game.Fullscreen {
		ebiten.SetFullscreen(true)
	}
	ebiten.SetVsyncEnabled(game.Config.VsyncEnabled)
	ebiten.SetWindowTitle(g.Title)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
