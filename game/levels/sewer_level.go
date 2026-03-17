package levels

import (
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/game/actors"
	"github.com/felipelucio/go-pixel-dungeon/game/actors/traps"
)

type SewerLevel struct {
	game.Level
}

func NewSewerLevel(w, h int) *SewerLevel {
	return &SewerLevel{
		Level: *game.NewLevel(w, h),
	}
}

func (l *SewerLevel) PlayLevelMusic() {
	game.GetSession().Audio().PlayMusic(game.AssetMusic["SEWERS_1"])
}

func (l *SewerLevel) StandardRooms(forceMax bool) int {
	rng := game.GetSession().RNG()
	if forceMax {
		return 6
	}
	return 4 + rng.Chances([]float64{1, 3, 1})
}

func (l *SewerLevel) SpecialRooms(forceMax bool) int {
	rng := game.GetSession().RNG()
	if forceMax {
		return 2
	}
	return 1 + rng.Chances([]float64{1, 4})
}

func (l *SewerLevel) TilesTexture() string {
	return game.AssetTiles["TILES_SEWERS"]
}

func (l *SewerLevel) WaterTexture() string {
	return game.AssetTiles["WATER_SEWERS"]
}

func (l *SewerLevel) TrapClasses() []func() *actors.Trap {
	gs := game.GetSession().Game().GameState
	if gs.Depth == 1 {
		return []func() *actors.Trap{
			traps.NewWornDartTrap,
		}
	} else {
		return []func() *actors.Trap{
			traps.NewChillingTrap, traps.NewShockingTrap, traps.NewToxicTrap,
			traps.NewWornDartTrap, traps.NewAlarmTrap, traps.NewOozeTrap,
			traps.NewConfusionTrap, traps.NewFlockTrap, traps.NewSummoningTrap,
			traps.NewTeleportTrap, traps.NewGatewayTrap,
		}
	}
}

func (l *SewerLevel) TrapChances() []float64 {
	gs := game.GetSession().Game().GameState
	if gs.Depth == 1 {
		return []float64{1}
	} else {
		return []float64{
			4, 4, 4,
			2, 2, 1,
			1, 1, 1,
			1,
		}
	}
}

func (l *SewerLevel) CreateMobs() {

}

func (l *SewerLevel) AddVisuals() {

}

func (l *SewerLevel) BuildFlagMaps() {

}
