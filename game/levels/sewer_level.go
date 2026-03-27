package levels

import (
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/game/actors"
	"github.com/felipelucio/go-pixel-dungeon/game/actors/traps"
)

var sewerLevelConfig = game.LevelConfig{
	Music: []string{
		game.AssetMusic["SEWERS_1"],
	},
	TilesTexture:    game.AssetTiles["TILES_SEWERS"],
	WaterTexture:    game.AssetTiles["WATER_SEWERS"],
	MinRooms:        4,
	MaxRooms:        6,
	MinSpecialRooms: 1,
	MaxSpecialRooms: 2,
	Traps: []func() *actors.Trap{
		traps.NewChillingTrap, traps.NewShockingTrap, traps.NewToxicTrap,
		traps.NewWornDartTrap, traps.NewAlarmTrap, traps.NewOozeTrap,
		traps.NewConfusionTrap, traps.NewFlockTrap, traps.NewSummoningTrap,
		traps.NewTeleportTrap, traps.NewGatewayTrap,
	},
	TrapsChances: []float64{
		4, 4, 4,
		2, 2, 1,
		1, 1, 1,
		1,
	},
	Mobs: []func() *actors.Mob{},
}

func NewSewerLevel(sess *game.Session) *game.Level {
	level := game.NewLevel()

	level.AddItemToSpawn(food.Random())
	if sess.GameState.PosNeeded() {
		sess.GameState.LimitedDrops.StrengthPotion.Add(1)
		level.AddItemToSpawn(potions.NewPotionOfStrength())
	}

	if sess.GameState.SouNeeded() {
		sess.GameState.LimitedDrops.UpgradeScrolls.Add(1)
		level.AddItemToSpawn(scrolls.NewScrollOfUpdate())
	}

	if sess.GameState.AsNeeded() {
		sess.GameState.LimitedDrops.ArcaneStylus.Add(1)
		level.AddItemToSpawn(items.NewStylus())
	}

	if sess.GameState.EnchStoneNeeded() {
		sess.GameState.LimitedDrops.EnchantmentStone.Drop()
		level.AddItemToSpawn(items.NewEnchantmentStone())
	}

	if sess.GameState.IntStoneNeeded() {
		sess.GameState.LimitedDrops.IntelligenceStone.Drop()
		level.AddItemToSpawn(items.NewIntelligenceStone())
	}

	if sess.GameState.TrinketCatalystNeeded() {
		sess.GameState.LimitedDrops.TrinketCatalyst.Drop()
		level.AddItemToSpawn(items.NewTrinketCatalyst())
	}

	return level
}
