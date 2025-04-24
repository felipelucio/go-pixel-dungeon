package game

import (
	"sort"

	"github.com/felipelucio/go-pixel-dungeon/core"
)

type WorldSystemError struct {
	Msg string
}

func (e WorldSystemError) Error() string {
	return e.Msg
}

type World struct {
	entities   map[string]core.Entity
	components map[string]map[string]core.Component
	systems    map[int]func(*World) error
}

func NewWorld() World {
	return World{
		make(map[string]core.Entity, 100),
		make(map[string]map[string]core.Component, 10),
		make(map[int]func(*World) error),
	}
}

func (world *World) NewEntity() *core.Entity {
	ent := core.NewEntity()
	world.entities[ent.GetID()] = ent
	return &ent
}

func (world *World) RemoveEntity(ent *core.Entity) {
	// TODO: remove all components before removing entity
	delete(world.entities, ent.GetID())
}

func (world *World) AddComponent(ent *core.Entity, comp core.Component) {
	compName := comp.GetName()
	_, ok := world.components[compName]
	if !ok {
		world.components[compName] = make(map[string]core.Component, 100)
	}
	world.components[compName][ent.GetID()] = comp
	ent.AddComponent(comp)
}

func (world *World) RemoveComponent(ent *core.Entity, compName string) {
	_, ok := world.components[compName]
	if ok {
		delete(world.components[compName], ent.GetID())
		ent.RemoveComponent(compName)
	}
}

func (world *World) GetComponents(compName string) []core.Component {
	comps := make([]core.Component, 0)
	if compSpace, ok := world.components[compName]; ok {
		for key := range compSpace {
			comps = append(comps, compSpace[key])
		}
	}
	return comps
}

func (world *World) AddSystem(system func(*World) error, order int) error {
	_, ok := world.systems[order]
	if ok {
		return WorldSystemError{"There is another System registered into '%s' order"}
	}

	world.systems[order] = system
	return nil
}

func (world *World) Update() error {
	// Let's get a slice of all keys
	keySlice := make([]int, 0)
	for key := range world.systems {
		keySlice = append(keySlice, key)
	}

	// Now sort the slice
	sort.Ints(keySlice)

	// Iterate over all keys in a sorted order
	for _, key := range keySlice {
		if err := world.systems[key](world); err != nil {
			return err
		}
	}
	return nil
}
