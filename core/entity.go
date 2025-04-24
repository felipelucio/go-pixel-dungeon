package core

import (
	"github.com/google/uuid"
)

type Entity struct {
	id         string
	components map[string]Component
}

func NewEntity() Entity {
	return Entity{
		uuid.New().String(),
		make(map[string]Component, 5),
	}
}

func (ent *Entity) GetID() string {
	return ent.id
}

func (ent *Entity) AddComponent(component Component) {
	c_name := component.GetName()
	ent.components[c_name] = component
}

func (ent *Entity) RemoveComponent(componentName string) {
	_, ok := ent.components[componentName]
	if ok {
		delete(ent.components, componentName)
	}
}

func (ent *Entity) GetComponent(componentName string) Component {
	return ent.components[componentName]
}
