package systems

import (
	"github.com/felipelucio/go-pixel-dungeon/components"
	"github.com/felipelucio/go-pixel-dungeon/game"
)

func MoveSystem(world *game.World) error {
	comps := world.GetComponents("Position")
	for _, comp := range comps {
		pos_comp, ok := comp.(*components.Position)
		if ok {
			pos_comp.X += 1
			pos_comp.Y += 1
		}
	}
	return nil
}
