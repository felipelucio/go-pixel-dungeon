package builder

import (
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/game/levels"
)

func BuildLevel(sess *game.Session) *game.Level {
	var level *game.Level

	switch sess.GameState.Depth {
	case 1:
	case 2:
	case 3:
	case 4:
		level = levels.NewSewerLevel(sess)
	case 5:
		level = levels.NewSewerBossLevel(sess)
	case 6:
	case 7:
	case 8:
	case 9:
		level = levels.NewPrisonLevel(sess)
	case 10:
		level = levels.NewPrisonBossLevel(sess)
	case 11:
	case 12:
	case 13:
	case 14:
		level = levels.NewCavesLevel(sess)
	case 15:
		level = levels.NewCavesBossLevel(sess)
	case 16:
	case 17:
	case 18:
	case 19:
		level = levels.NewCityLevel(sess)
	case 20:
		level = levels.NewCityBossLevel(sess)
	case 21:
	case 22:
	case 23:
	case 24:
		level = levels.NewHallLevel(sess)
	case 25:
		level = levels.NewHallBossLevel(sess)
	case 26:
		level = levels.NewLastLevel(sess)
	default:
		level = levels.DeadEndLevel(sess)
	}

	level

	return level
}
