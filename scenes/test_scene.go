package scenes

import (
	"fmt"

	"github.com/felipelucio/go-pixel-dungeon/core"
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TestScene struct {
	level *game.Level
}

func (scene *TestScene) Init() error {
	sess := game.GetSession()
	level := game.NewLevel()
	// scene.camera = core.NewCamera(768, 432, 0.5, 2.0, true)

	return nil
}

func (scene *TestScene) Pause() {
}

func (scene *TestScene) Resume() {
}

func (scene *TestScene) Update() error {
	ebiten.SetWindowTitle(fmt.Sprintf("%s (FPS: %.2f | TPS: %.2f)", game.GAME_TITLE, ebiten.ActualFPS(), ebiten.ActualTPS()))

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
		// scene.camera.SmoothScaleBy(0.25, 0.2)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		// scene.camera.SmoothScaleBy(-0.25, 0.2)
	}

	if dir.X != 0 || dir.Y != 0 {
		// scene.camera.SmoothMoveBy(int(16*dir.X), int(16*dir.Y), 0.2)
	}

	// scene.camera.Update(1.0 / 60.0)
	return nil
}

func (scene *TestScene) Draw(screen *ebiten.Image) error {
	// buff := scene.camera.GetBuffer()
	// buff.Clear()
	// buff.Fill(color.RGBA{255, 128, 128, 255})
	// comp := scene.player.GetComponent("Position")
	// pos_comp, ok := comp.(*components.Position)
	// if ok {
	// 	p_str := fmt.Sprintf("Player: (%d, %d)", pos_comp.X, pos_comp.Y)
	// 	ebitenutil.DebugPrintAt(buff, p_str, 0, 20)
	// }
	// scene.drawMap(buff)
	// scene.camera.Draw(screen)
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

			xPos += 1
		}
		yPos += 1
	}
}
