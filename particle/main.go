package main

import (
	"os"
	"time"
	// "math"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/ungerik/go3d/float64/vec2"

	"github.com/jstewart7/ecs"

	"github.com/jstewart7/flow/physics"
	"github.com/jstewart7/flow/particle"
	"github.com/jstewart7/flow/render"
	"github.com/jstewart7/flow/asset"
	"github.com/jstewart7/flow/interp"
)

func check(err error) {
	if err != nil { panic(err) }
}

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
	cfg := pixelgl.WindowConfig{
		Title: "Particles",
		Bounds: pixel.R(0, 0, 1920, 1080),
		VSync: true,
		Resizable: true,
		Maximized: true,
		Monitor: pixelgl.PrimaryMonitor(),
	}

	win, err := pixelgl.NewWindow(cfg)
	check(err)
	win.SetSmooth(false)

	// center := win.Bounds().Center()

	world := ecs.NewWorld()
	camera := render.NewCamera(win, 0, 0)

	load := asset.NewLoad(os.DirFS("./"))
	particleSprite, err := load.Sprite("square.png")
	check(err)

	// lightBlue := color.NRGBA{0x8a, 0xeb, 0xf1, 0xff}
	// pink := color.NRGBA{0xcd, 0x60, 0x93, 0xff}
	// particlePrefab := ecs.BlankEntity()
	// particlePrefab["render.Sprite"] = render.NewSprite(particleSprite)
	// particlePrefab["particle.Lifetime"] = particle.NewLifetime(2 * time.Second)
	// particlePrefab["particle.Color"] = particle.NewColor(interp.Linear, lightBlue, pink)
	// particlePrefab["particle.Size"] = particle.NewSize(interp.Linear, vec2.T{5, 5}, vec2.T{1, 1})
	// emitter := particle.Emitter{
	// 	Max: 100,
	// 	Rate: 1.0,
	// 	Loop: true,
	// 	Probability: 1.0,
	// 	Prefab: particlePrefab,
	// 	Builders: []particle.PrefabBuilder{
	// 		&particle.TransformBuilder{
	// 			PosPositioner: &particle.RingPositioner{vec2.T{0, 2*math.Pi}, vec2.T{80, 80}},
	// 		},
	// 		&particle.RigidbodyBuilder{
	// 			Mass: 1,
	// 			VelPositioner: &particle.AnglePositioner{10},
	// 		},
	// 	},
	// }

	// Fire
	// startColor := color.NRGBA{0xf4, 0x7e, 0x1b, 0xff}
	// endColor := color.NRGBA{0xa9, 0x3b, 0x3b, 0xff}
	// particlePrefab := ecs.BlankEntity()
	// particlePrefab["render.Sprite"] = render.NewSprite(particleSprite)
	// particlePrefab["particle.Lifetime"] = particle.NewLifetime(2 * time.Second)
	// particlePrefab["particle.Color"] = particle.NewColor(interp.Linear, startColor, endColor)
	// particlePrefab["particle.Size"] = particle.NewSize(interp.Linear, vec2.T{100, 100}, vec2.T{10, 10})
	// emitter := particle.Emitter{
	// 	Max: 100,
	// 	Rate: 1.0,
	// 	Loop: true,
	// 	Probability: 1.0,
	// 	Prefab: particlePrefab,
	// 	Builders: []particle.PrefabBuilder{
	// 		&particle.TransformBuilder{
	// 			PosPositioner: &particle.RingPositioner{vec2.T{0, 2*math.Pi}, vec2.T{0, 50}},
	// 		},
	// 		&particle.RigidbodyBuilder{
	// 			Mass: 1,
	// 			VelPositioner: &particle.ConstantPositioner{vec2.T{0, 100}},
	// 			// VelPositioner: &particle.FirePositioner{},
	// 		},
	// 	},
	// }

	lightBlue := color.NRGBA{0x8a, 0xeb, 0xf1, 0xff}
	pink := color.NRGBA{0xcd, 0x60, 0x93, 0xff}
	particlePrefab := ecs.BlankEntity()
	particlePrefab["render.Sprite"] = render.NewSprite(particleSprite)
	particlePrefab["particle.Lifetime"] = particle.NewLifetime(10 * time.Second)
	particlePrefab["particle.Color"] = particle.NewColor(interp.Linear, lightBlue, pink)
	particlePrefab["particle.Size"] = particle.NewSize(interp.Linear, vec2.T{5, 5}, vec2.T{1, 1})
	emitter := particle.Emitter{
		Max: 100,
		Rate: 0.5,
		Loop: true,
		Probability: 1.0,
		Prefab: particlePrefab,
		Builders: []particle.PrefabBuilder{
			&particle.TransformBuilder{
				PosPositioner: &particle.RectPositioner{vec2.T{0, 1080}, vec2.T{1920, 1080}},
			},
			&particle.RigidbodyBuilder{
				Mass: 1,
				VelPositioner: &particle.ConstantPositioner{vec2.T{25, -200}},
			},
		},
	}

	dt := 15 * time.Millisecond
	for !win.Pressed(pixelgl.KeyBackspace) {
		{
			taggedForDelete := ecs.TaggedWith(world, "delete")
			for _,id := range taggedForDelete {
				ecs.Delete(world, id)
			}
		}

		// emitter.Update(world, vec2.T{center.X, center.Y}, dt)
		emitter.Update(world, vec2.T{0,0}, dt)

		physics.RigidbodyPhysics(world, dt)

		render.InterpolateParticles(world, dt)

		// Draw
		win.Clear(pixel.RGB(0, 0, 0))
		win.SetMatrix(camera.Mat())
		render.DrawSprites(win, world)
		win.Update()
	}
}
