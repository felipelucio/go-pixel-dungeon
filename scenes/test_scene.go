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

func (scene *TestScene) Init() error {
	scene.gs = &game.GetSession().Game().GameState
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
	scene.tm = core.NewTilemap(&scene.ts, 512, 512)

	scene.camera = core.NewCamera(768, 432, 0.5, 2.0, true)

	return nil
}

func (scene *TestScene) Pause() {
}

func (scene *TestScene) Resume() {
}

func (scene *TestScene) Update() error {
	g := game.GetSession().Game()
	ebiten.SetWindowTitle(fmt.Sprintf("%s (FPS: %.2f | TPS: %.2f)", g.Title, ebiten.ActualFPS(), ebiten.ActualTPS()))

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
		scene.camera.SmoothScaleBy(0.25, 0.2)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		scene.camera.SmoothScaleBy(-0.25, 0.2)
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
	// comp := scene.player.GetComponent("Position")
	// pos_comp, ok := comp.(*components.Position)
	// if ok {
	// 	p_str := fmt.Sprintf("Player: (%d, %d)", pos_comp.X, pos_comp.Y)
	// 	ebitenutil.DebugPrintAt(buff, p_str, 0, 20)
	// }
	scene.drawMap(buff)
	scene.camera.Draw(screen)
	return nil
}

func (scene *TestScene) Destroy() {
}

func (scene *TestScene) drawMap(buff *ebiten.Image) {
	tileH := scene.ts.GetTileHeight()
	tileW := scene.ts.GetTileWidth()
	// buffH := buff.Bounds().Dy()
	// buffW := buff.Bounds().Dx()

	// camX, camY := scene.camera.GetPosition()
	// camHalfW := buffW / 2
	// camHalfH := buffH / 2
	// startX := math.Max(0, float64(camX-camHalfW))
	// startY := math.Max(0, float64(camY-camHalfH))
	// endX := startX + float64(camHalfW)
	// endY := startY + float64(camHalfH)

	mapRect := scene.camera.GetDisplayRect()
	startTileX := mapRect.Min.X / tileW
	startTileY := mapRect.Min.Y / tileH
	endTileX := mapRect.Max.X / tileW
	endTileY := mapRect.Max.Y / tileH
	yPos := 0
	for y := startTileY; y < endTileY; y++ {
		xPos := 0
		for x := startTileX; x < endTileX; x++ {
			t, err := scene.tm.GetTile(x, y)
			if err == nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(xPos*tileW), float64(yPos*tileH))
				buff.DrawImage(t, op)
			}
			xPos += 1
		}
		yPos += 1
	}
}
