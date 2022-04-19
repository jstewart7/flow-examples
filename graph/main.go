package main

import (
	// "os"
	"fmt"
	"math"
	// "time"
	// "image/color"

	"github.com/unitoftime/glitch"
	"github.com/unitoftime/glitch/shaders"

	"github.com/unitoftime/flow/render"
)

func check(err error) {
	if err != nil { panic(err) }
}

func main() {
	fmt.Println("Starting")
	glitch.Run(runGame)
}

func runGame() {
	win, err := glitch.NewWindow(1920, 1080, "Graphing", glitch.WindowConfig{
		Vsync: true,
		Samples: 4,
	})
	check(err)
	// win.SetSmooth(false)

	// center := win.Bounds().Center()

	lightBlue := glitch.RGBA{0x8a, 0xeb, 0xf1, 0xff}
//	lightBlue := color.NRGBA{0x8a, 0xeb, 0xf1, 0xff}
	// pink := color.NRGBA{0xcd, 0x60, 0x93, 0xff}

	pad := float32(50)
	rect := glitch.R(0 + pad, 0 + pad, 1920 - pad, 1080 - pad)

	dat := []float64{
		2, 3, 4, 5, 4, 3, 2, 1,
	}

	shader, err := glitch.NewShader(shaders.SpriteShader)
	if err != nil { panic(err) }

	pass := glitch.NewRenderPass(shader)

	geom := glitch.NewGeomDraw()
	geom.SetColor(lightBlue)

	min := math.MaxFloat64
	max := -math.MaxFloat64
	for _, y := range dat {
		min = math.Min(min, y)
		max = math.Max(max, y)
	}

	dx := rect.W() / float32(len(dat))
	dy := rect.H() / float32(max - min)
	fmt.Println(dx, dy, rect.H(), max, min)
	points := make([]glitch.Vec3, 0)
	for x, y := range dat {
		points = append(points, glitch.Vec3{
			rect.Min[0] + float32(x) * dx,
			rect.Min[1] + float32(y-min) * dy,
			0,
		})
	}
	graphLine := geom.LineStrip(points, 5)

	// imd := imdraw.New(nil)
	// imd.Color = lightBlue
	// imd.EndShape = imdraw.RoundEndShape

	// min := math.MaxFloat64
	// max := -math.MaxFloat64
	// for _, y := range dat {
	// 	min = math.Min(min, y)
	// 	max = math.Max(max, y)
	// }

	// dx := rect.W() / float64(len(dat))
	// dy := rect.H() / (max - min)
	// fmt.Println(dx, dy, rect.H(), max, min)
	// for x, y := range dat {
	// 	imd.Push(pixel.V(
	// 		rect.Min.X + float64(x) * dx,
	// 		rect.Min.Y + (y-min) * dy))
	// }
	// imd.Line(5)

	camera := render.NewCamera(win, 0, 0)
	camera.Update()

	// dt := 15 * time.Millisecond
	for !win.Pressed(glitch.KeyBackspace) {

		pass.Clear()
		graphLine.Draw(pass, glitch.Mat4Ident)

		glitch.Clear(glitch.RGBA{0, 0, 0, 1.0})

		winBounds := win.Bounds()
		camera.Position = glitch.Vec2{winBounds.W()/2, winBounds.H()/2}
		camera.Update()
		pass.SetUniform("projection", camera.Camera.Projection)
		pass.SetUniform("view", camera.Camera.View)
		pass.Draw(win)

		win.Update()
	}
}
