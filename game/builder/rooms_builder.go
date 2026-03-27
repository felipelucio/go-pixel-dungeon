package builder

import (
	"github.com/felipelucio/go-pixel-dungeon/game"
	"github.com/felipelucio/go-pixel-dungeon/rooms"
)

func CreateRooms(level *game.Level, levelConf *game.LevelConfig) {
	sess := game.GetSession()
	rng := sess.RNG()

	if sess.GameState.Depth > 1 {
		switch rng.Int32N(0, 14) {
		case 0:
			level.Feeling = game.CHASM_FEEL
		case 1:
			level.Feeling = game.WATER_FEEL
		case 2:
			level.Feeling = game.GRASS_FEEL
		case 3:
			level.Feeling = game.DARK_FEEL
			level.ViewDistance = 5 * level.ViewDistance / 8
		case 4:
			level.Feeling = game.LARGE_FEEL
			// level.AddItemToSpawn()
		case 5:
			level.Feeling = game.TRAPS_FEEL
		case 6:
			level.Feeling = game.SECRETS_FEEL
		default:
			level.Feeling = game.NONE_FEEL
		}
	}

	levelRooms := make([]*game.Room, 10)
	entranceRoom := rooms.NewEntranceRoom()
	exitRoom := rooms.NewExitRoom()
	levelRooms = append(levelRooms, entranceRoom)
	levelRooms = append(levelRooms, exitRoom)

	// standardRooms := level.StandardRooms(level.Feeling == game.LARGE_FEEL)
	// if level.Feeling == game.LARGE_FEEL {
	// 	standardRooms = int(math.Ceil(float64(standardRooms) * 1.5))
	// }

	// for i := range standardRooms {
	// 	r := rooms.NewStandardRoom()
	// 	r.SetSize(i)
	// 	levelRooms = append(levelRooms, r)
	// }

	// painter.Paint(level, levelRooms)
}
