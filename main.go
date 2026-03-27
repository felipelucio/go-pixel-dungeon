package main

import (
	"log"
	"log/slog"

	"github.com/felipelucio/go-pixel-dungeon/core"
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

func loadAssets() {
	asm := core.DefaultAssetManager()
	asm.LoadFontList(game.AssetFonts)
	asm.LoadImageList(game.AssetInterface)
	asm.LoadImageList(game.AssetEffects)
	asm.LoadImageList(game.AssetTiles)
	asm.LoadAudioList(game.AssetMusic)
	asm.LoadAudioList(game.AssetSounds)

	aum := core.DefaultAudioManager()
	for n := range game.AssetMusic {
		a, ok := asm.GetAudio(n)
		if !ok {
			slog.Warn("Music not found: %s", "name", n)
		}
		aum.RegisterMusic(n, a)
	}

	for n := range game.AssetSounds {
		a, ok := asm.GetAudio(n)
		if !ok {
			slog.Warn("SFX not found: %s", "name", n)
		}
		aum.RegisterSFX(n, a)
	}

}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	asm := core.NewAssetManager()
	core.SetDefaultAssetManager(asm)

	aum := core.NewAudioManager()
	core.SetDefaultAudioManager(aum)

	loadAssets()

	sm := core.NewSceneManager()
	sm.SwitchToScene(&scenes.TestScene{})
	core.SetDefaultSceneManager(sm)

	sess := game.NewSession(nil)
	game.SetSession(sess)

	g := game.NewGame()

	ebiten.SetWindowSize(game.Config.WinWidth, game.Config.WinHeight)
	if game.Config.WinMode == game.Fullscreen {
		ebiten.SetFullscreen(true)
	}
	ebiten.SetVsyncEnabled(game.Config.VsyncEnabled)
	ebiten.SetWindowTitle(game.GAME_TITLE)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
