package main

import (
	"flag"
	"math/rand"
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/examples/community/game_of_life/life"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var (
	size       *int
	windowSize *float64
	frameRate  *int64
)

func init() {
	rand.Seed(time.Now().UnixNano())
	size = flag.Int("size", 5, "The size of each cell")
	windowSize = flag.Float64("windowSize", 800, "The pixel size of one side of the grid")
	frameRate = flag.Int64("frameRate", 33, "The framerate in milliseconds")
	flag.Parse()
}

func main() {
	pixelgl.Run(run)
}

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, *windowSize, *windowSize),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.White)

	// since the game board is square, rows and cols will be the same
	rows := int(*windowSize) / *size

	gridDraw := imdraw.New(nil)
	game := life.NewLife(rows, *size)
	last := time.Now()
	for !win.Closed() {
		// game loop
		if time.Since(last).Nanoseconds() > *frameRate {
			gridDraw.Clear()
			game.A.Draw(gridDraw)
			gridDraw.Draw(win)
			game.Step()
			last = time.Now()
		}
		win.Update()
	}
}
