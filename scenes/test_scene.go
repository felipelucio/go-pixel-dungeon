package scenes

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"github.com/felipelucio/go-pixel-dungeon/components"
	"github.com/felipelucio/go-pixel-dungeon/core"
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TestScene struct {
	gs     *core.GameState
	world  game.World
	player *core.Entity
	ts     core.Tileset
	tm     core.Tilemap
	camera core.Camera
}

func (scene *TestScene) Init(state *core.GameState) error {
	scene.gs = state
	scene.world = game.NewWorld()
	ok := scene.world.AddSystem(systems.MoveSystem, 10)
	if ok != nil {
		fmt.Printf("%s", ok.Error())
		return errors.New(ok.Error())
	}

	scene.player = scene.world.NewEntity()
	scene.world.AddComponent(scene.player, &components.Position{})
	tilePath := fmt.Sprintf("%s/%s", game.Config.AssetsPath, game.TILES_SEWERS)
	ts, err := core.NewTileset("ts0", tilePath, 16, 16)
	if err != nil {
		log.Fatal(err)
	}
	scene.ts = ts
	scene.tm = core.NewTilemap(&scene.ts, 768/2/16, 432/2/16)

	scene.camera = core.NewCamera(768, 432, 0.5, 2.0)

	return nil
}

func (scene *TestScene) Pause() {

}

func (scene *TestScene) Resume() {

}

func (scene *TestScene) Update() error {
	dir := core.NewVector2(0.0, 0.0)
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		dir.X = -1.0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		dir.X = 1.0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		dir.Y = -1.0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		dir.Y = 1.0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		scene.camera.ScaleBy(0.50)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		scene.camera.ScaleBy(-0.50)
	}

	if dir.X != 0 || dir.Y != 0 {
		scene.camera.SmoothMoveBy(int(16*dir.X), int(16*dir.Y), 0.2)
	}

	scene.camera.Update(1.0 / 60.0)
	return scene.world.Update()
}

func (scene *TestScene) Draw(screen *ebiten.Image) error {
	buff := scene.camera.GetBuffer()
	buff.Clear()
	buff.Fill(color.RGBA{255, 128, 128, 255})
	ebitenutil.DebugPrint(buff, "Hello, World!")
	// comp := scene.player.GetComponent("Position")
	// pos_comp, ok := comp.(*components.Position)
	// if ok {
	// 	p_str := fmt.Sprintf("Player: (%d, %d)", pos_comp.X, pos_comp.Y)
	// 	ebitenutil.DebugPrintAt(buff, p_str, 0, 20)
	// }

	mapH := scene.tm.GetHeight()
	mapW := scene.tm.GetWidth()
	tileH := scene.ts.GetTileHeight()
	tileW := scene.ts.GetTileWidth()
	camX, camY := scene.camera.GetPosition()
	for z := range 3 {
		_ = z
		for y := range mapH {
			for x := range mapW {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((x*tileW)-camX), float64((y*tileH)-camY))
				buff.DrawImage(scene.tm.GetTile(x, y), op)
			}
		}
	}

	scene.camera.Draw(screen)

	return nil
}

func (scene *TestScene) Destroy() {

}
