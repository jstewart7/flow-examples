package main

import (
	"fmt"
	"github.com/jstewart7/ecs"
	"github.com/jstewart7/flow/render"
	"github.com/jstewart7/flow/physics"
	"github.com/jstewart7/flow/particle"
)

func init() {
	ecs.SetRegistry(&ComponentRegistry{})
}

type ComponentRegistry struct {}
func (r *ComponentRegistry) GetArchStorageType(component interface{}) ecs.ArchComponent {
	switch component.(type) {
	case physics.Transform:
		return &physics.TransformList{}
	case *physics.Transform:
		return &physics.TransformList{}
	case physics.Input:
		return &physics.InputList{}
	case *physics.Input:
		return &physics.InputList{}
	case physics.Rigidbody:
		return &physics.RigidbodyList{}
	case *physics.Rigidbody:
		return &physics.RigidbodyList{}

	case render.Sprite:
		return &render.SpriteList{}
	case *render.Sprite:
		return &render.SpriteList{}
	case render.Keybinds:
		return &render.KeybindsList{}
	case *render.Keybinds:
		return &render.KeybindsList{}

	case particle.Lifetime:
		return &particle.LifetimeList{}
	case *particle.Lifetime:
		return &particle.LifetimeList{}
	case particle.Color:
		return &particle.ColorList{}
	case *particle.Color:
		return &particle.ColorList{}
	case particle.Size:
		return &particle.SizeList{}
	case *particle.Size:
		return &particle.SizeList{}
	default:
		panic(fmt.Sprintf("Unknown component type: %T", component))
	}
}
func (r *ComponentRegistry) GetComponentMask(component interface{}) ecs.ArchMask {
	switch component.(type) {
	case physics.Transform:
		return ecs.ArchMask(1 << 0)
	case physics.Input:
		return ecs.ArchMask(1 << 1)
	case physics.Rigidbody:
		return ecs.ArchMask(1 << 2)

		// Client Only
	case render.Sprite:
		return ecs.ArchMask(1 << 3)
	case render.Keybinds:
		return ecs.ArchMask(1 << 4)
	case particle.Lifetime:
		return ecs.ArchMask(1 << 5)
	case particle.Color:
		return ecs.ArchMask(1 << 6)
	case particle.Size:
		return ecs.ArchMask(1 << 7)

	default:
		panic(fmt.Sprintf("Unknown component type: %T", component))
	}
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// type TypeList []Type
// func (t *TypeList) ComponentSet(val interface{}) { *t = *val.(*TypeList) }
// func (t *TypeList) InternalRead(index int, val interface{}) { *val.(*Type) = (*t)[index] }
// func (t *TypeList) InternalWrite(index int, val interface{}) { (*t)[index] = *val.(*Type) }
// func (t *TypeList) InternalAppend(val interface{}) { (*t) = append((*t), val.(Type)) }
// func (t *TypeList) InternalPointer(index int) interface{} { return &(*t)[index] }
// func (t *TypeList) InternalReadVal(index int) interface{} { return (*t)[index] }
// func (t *TypeList) Delete(index int) {
// 	lastVal := (*t)[len(*t)-1]
// 	(*t)[index] = lastVal
// 	(*t) = (*t)[:len(*t)-1]
// }
// func (t *TypeList) Len() int { return len(*t) }


