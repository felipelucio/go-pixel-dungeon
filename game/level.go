package game

import (
	"errors"

	"github.com/felipelucio/go-pixel-dungeon/game/actors"
	"github.com/felipelucio/go-pixel-dungeon/game/items"
)

type LevelFeeling int

const (
	CHASM_FEEL LevelFeeling = iota
	WATER_FEEL
	GRASS_FEEL
	DARK_FEEL
	LARGE_FEEL
	TRAPS_FEEL
	SECRETS_FEEL
	NONE_FEEL
)

type LevelConfig struct {
	Music           []string
	TilesTexture    string
	WaterTexture    string
	MinRooms        int
	MaxRooms        int
	MinSpecialRooms int
	MaxSpecialRooms int
	Traps           []func() *actors.Trap
	TrapsChances    []float64
	Mobs            []func() *actors.Mob
}

type Level struct {
	Width  int
	Height int

	Feeling      LevelFeeling
	ViewDistance int

	Tiles        []int
	Visited      []bool
	Mapped       []bool
	Discoverable []bool
	Passable     []bool
	LosBlocking  []bool
	Flamable     []bool
	Secret       []bool
	Solid        []bool
	Avoid        []bool
	Water        []bool
	Pit          []bool
	OpenSpace    []bool

	Entrance int
	Exit     int

	Mobs   []*actors.Mob
	Blobs  []*actors.Blob
	Traps  []*actors.Trap
	Heaps  []*items.Heap
	Plants []*items.Plant

	Music []string
}

func NewLevel() *Level {
	return &Level{
		Width:  0,
		Height: 0,

		Feeling:      NONE_FEEL,
		ViewDistance: 8,

		Tiles:        make([]int, 0),
		Visited:      make([]bool, 0),
		Mapped:       make([]bool, 0),
		Discoverable: make([]bool, 0),
		Passable:     make([]bool, 0),
		LosBlocking:  make([]bool, 0),
		Flamable:     make([]bool, 0),
		Secret:       make([]bool, 0),
		Solid:        make([]bool, 0),
		Avoid:        make([]bool, 0),
		Water:        make([]bool, 0),
		Pit:          make([]bool, 0),
		OpenSpace:    make([]bool, 0),

		Mobs:   make([]*actors.Mob, 10),
		Blobs:  make([]*actors.Blob, 10),
		Traps:  make([]*actors.Trap, 10),
		Heaps:  make([]*items.Heap, 10),
		Plants: make([]*items.Plant, 10),
	}
}

func (l *Level) SetSize(w, h int) {
	oldWidth := l.Width
	oldHeight := l.Height
	oldTiles := l.Tiles
	oldVisited := l.Visited
	oldMapped := l.Mapped
	oldDiscoverable := l.Discoverable
	oldPassable := l.Passable
	oldLosBlocking := l.LosBlocking
	oldFlamable := l.Flamable
	oldSecret := l.Secret
	oldSolid := l.Solid
	oldAvoid := l.Avoid
	oldWater := l.Water
	oldPit := l.Pit
	oldOpenSpace := l.OpenSpace

	l.Width = w
	l.Height = h
	l.Tiles = make([]int, w*h)
	l.Visited = make([]bool, w*h)
	l.Mapped = make([]bool, w*h)
	l.Discoverable = make([]bool, w*h)
	l.Passable = make([]bool, w*h)
	l.LosBlocking = make([]bool, w*h)
	l.Flamable = make([]bool, w*h)
	l.Secret = make([]bool, w*h)
	l.Solid = make([]bool, w*h)
	l.Avoid = make([]bool, w*h)
	l.Water = make([]bool, w*h)
	l.Pit = make([]bool, w*h)
	l.OpenSpace = make([]bool, w*h)

	for y := range oldHeight - 1 {
		for x := range oldWidth - 1 {
			idx := (y * oldWidth) + x
			if idx < l.Width*l.Height {
				l.Tiles[idx] = oldTiles[idx]
				l.Visited[idx] = oldVisited[idx]
				l.Mapped[idx] = oldMapped[idx]
				l.Discoverable[idx] = oldDiscoverable[idx]
				l.Passable[idx] = oldPassable[idx]
				l.LosBlocking[idx] = oldLosBlocking[idx]
				l.Flamable[idx] = oldFlamable[idx]
				l.Secret[idx] = oldSecret[idx]
				l.Solid[idx] = oldSolid[idx]
				l.Avoid[idx] = oldAvoid[idx]
				l.Water[idx] = oldWater[idx]
				l.Pit[idx] = oldPit[idx]
				l.OpenSpace[idx] = oldOpenSpace[idx]
			}
		}
	}
}

func (l *Level) GetWidth() int {
	return l.Width
}

func (l *Level) GetHeight() int {
	return l.Height
}

func (l *Level) SetHeight(h int) {
	l.Height = h
}

func (l *Level) checkInbound(x, y int) error {
	if x < 0 || x > l.Width || y < 0 || y >= l.Height {
		return errors.New("Out of bounds")
	}
	return nil
}

func (l *Level) GetIndex(x, y int) (int, error) {
	err := l.checkInbound(x, y)
	if err != nil {
		return -1, err
	}
	return (y * l.Width) + x, nil
}

func (l *Level) GetTile(x, y int) (int, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return -1, err
	}
	return l.Tiles[idx], nil
}

func (l *Level) SetTile(x, y, id int) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Tiles[idx] = id
	return nil
}

func (l *Level) GetEntrance() int {
	return l.Entrance
}

func (l *Level) SetEntrance(id int) error {
	l.Entrance = id
	return nil
}

func (l *Level) GetExit() int {
	return l.Entrance
}

func (l *Level) SetExit(id int) error {
	l.Exit = id
	return nil
}

func (l *Level) IsVisited(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Visited[idx], nil
}

func (l *Level) SetVisited(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Visited[idx] = val
	return nil
}

func (l *Level) IsMapped(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Mapped[idx], nil
}

func (l *Level) SetMapped(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Mapped[idx] = val
	return nil
}

func (l *Level) IsDiscoverable(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Discoverable[idx], nil
}

func (l *Level) SetDiscoverable(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Mapped[idx] = val
	return nil
}

func (l *Level) IsPassable(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Passable[idx], nil
}

func (l *Level) SetPassable(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Passable[idx] = val
	return nil
}

func (l *Level) IsLOSBlocking(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.LosBlocking[idx], nil
}

func (l *Level) SetLOSBlocking(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.LosBlocking[idx] = val
	return nil
}

func (l *Level) IsFlamable(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Flamable[idx], nil
}

func (l *Level) SetFlamable(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Flamable[idx] = val
	return nil
}

func (l *Level) IsSecret(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Secret[idx], nil
}

func (l *Level) SetSecret(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Secret[idx] = val
	return nil
}

func (l *Level) IsSolid(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Solid[idx], nil
}

func (l *Level) SetSolid(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Solid[idx] = val
	return nil
}

func (l *Level) IsAvoid(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Avoid[idx], nil
}

func (l *Level) SetAvoid(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Avoid[idx] = val
	return nil
}

func (l *Level) IsWater(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Water[idx], nil
}

func (l *Level) SetWater(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Water[idx] = val
	return nil
}

func (l *Level) IsPit(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.Pit[idx], nil
}

func (l *Level) SetPit(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.Pit[idx] = val
	return nil
}

func (l *Level) IsOpenSpace(x, y int) (bool, error) {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return false, err
	}
	return l.OpenSpace[idx], nil
}

func (l *Level) SetOpenSpace(x, y int, val bool) error {
	idx, err := l.GetIndex(x, y)
	if err != nil {
		return err
	}
	l.OpenSpace[idx] = val
	return nil
}

func (l *Level) GetMobs() []*actors.Mob {
	return l.Mobs
}

func (l *Level) AddMob(mob *actors.Mob) error {
	return errors.New("Not implemented")
}

func (l *Level) RemoveMob(mob *actors.Mob) error {
	return errors.New("Not implemented")
}

func (l *Level) GetBlobs() []*actors.Blob {
	return l.Blobs
}

func (l *Level) AddBlob(blob *actors.Blob) error {
	return errors.New("Not implemented")
}

func (l *Level) RemoveBlob(blob *actors.Blob) error {
	return errors.New("Not implemented")
}

func (l *Level) GetTraps() []*actors.Trap {
	return l.Traps
}

func (l *Level) AddTrap(trap *actors.Trap) error {
	return errors.New("Not implemented")
}

func (l *Level) RemoveTrap(trap *actors.Trap) error {
	return errors.New("Not implemented")
}

func (l *Level) GetHeaps() []*items.Heap {
	return l.Heaps
}

func (l *Level) AddHeap(heap *items.Heap) error {
	return errors.New("Not implemented")
}

func (l *Level) RemoveHeap(heap *items.Heap) error {
	return errors.New("Not implemented")
}

func (l *Level) GetPlants() []*items.Plant {
	return l.Plants
}

func (l *Level) AddPlant(plant *items.Plant) error {
	return errors.New("Not implemented")
}

func (l *Level) RemovePlant(plant *items.Plant) error {
	return errors.New("Not implemented")
}
