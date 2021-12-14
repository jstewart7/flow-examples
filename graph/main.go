package main

import (
	// "os"
	"fmt"
	"math"
	// "time"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"

	// "github.com/ungerik/go3d/float64/vec2"

	// "github.com/unitoftime/ecs"

	// "github.com/unitoftime/flow/physics"
	// "github.com/unitoftime/flow/particle"
	// "github.com/unitoftime/flow/render"
	// "github.com/unitoftime/flow/asset"
	// "github.com/unitoftime/flow/interp"
)

func check(err error) {
	if err != nil { panic(err) }
}

func main() {
	fmt.Println("Starting")
	pixelgl.Run(runGame)
}

func runGame() {
	cfg := pixelgl.WindowConfig{
		Title: "Graphing",
		Bounds: pixel.R(0, 0, 1920, 1080),
		VSync: true,
		Resizable: true,
		Maximized: true,
		Monitor: pixelgl.PrimaryMonitor(),
	}

	win, err := pixelgl.NewWindow(cfg)
	check(err)
	// win.SetSmooth(false)

	// center := win.Bounds().Center()

	lightBlue := color.NRGBA{0x8a, 0xeb, 0xf1, 0xff}
	// pink := color.NRGBA{0xcd, 0x60, 0x93, 0xff}

	pad := float64(50)
	rect := pixel.R(0 + pad, 0 + pad, 1920 - pad, 1080 - pad)

	dat := []float64{
		2, 3, 4, 5, 4, 3, 2, 1,
	}

	imd := imdraw.New(nil)
	imd.Color = lightBlue
	imd.EndShape = imdraw.RoundEndShape

	min := math.MaxFloat64
	max := -math.MaxFloat64
	for _, y := range dat {
		min = math.Min(min, y)
		max = math.Max(max, y)
	}

	dx := rect.W() / float64(len(dat))
	dy := rect.H() / (max - min)
	fmt.Println(dx, dy, rect.H(), max, min)
	for x, y := range dat {
		imd.Push(pixel.V(
			rect.Min.X + float64(x) * dx,
			rect.Min.Y + (y-min) * dy))
	}
	imd.Line(5)

	// dt := 15 * time.Millisecond
	for !win.Pressed(pixelgl.KeyBackspace) {

		win.Clear(pixel.RGB(0, 0, 0))

		imd.Draw(win)

		win.Update()
	}
}
