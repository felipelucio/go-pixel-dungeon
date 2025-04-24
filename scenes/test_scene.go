package scenes

import (
	"errors"
	"fmt"
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
	cam    core.Camera
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

	scene.cam = core.NewCamera(ebiten.NewImage(800, 600), 768, 432)

	return nil
}

func (scene *TestScene) Pause() {

}

func (scene *TestScene) Resume() {

}

func (scene *TestScene) Update() error {
	dir := core.NewVector2[float64](0.0, 0.0)
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

	if dir.X != 0 || dir.Y != 0 {
		scene.cam.SmoothMoveBy(int(64*dir.X), int(64*dir.Y), 0.2)
	}
	scene.cam.Update(1.0 / 60.0)
	return scene.world.Update()
}

func (scene *TestScene) Draw(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	// comp := scene.player.GetComponent("Position")
	// pos_comp, ok := comp.(*components.Position)
	// if ok {
	// 	p_str := fmt.Sprintf("Player: (%d, %d)", pos_comp.X, pos_comp.Y)
	// 	ebitenutil.DebugPrintAt(screen, p_str, 0, 20)
	// }

	mapH := scene.tm.GetHeight()
	mapW := scene.tm.GetWidth()
	tileH := scene.ts.GetTileHeight()
	tileW := scene.ts.GetTileWidth()
	camX, camY := scene.cam.GetPosition()
	for z := range 3 {
		_ = z
		for y := range mapH {
			for x := range mapW {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((x*tileW)+camX), float64((y*tileH)+camY))
				screen.DrawImage(scene.tm.GetTile(x, y), op)
			}
		}
	}

	return nil
}

func (scene *TestScene) Destroy() {

}
