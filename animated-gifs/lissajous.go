// Lissajous generates GIF animations of random Lissajous figures.
// To run this you need to specify an output e.g. go run lissajous.go >out.gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		oscCycles       = 5
		angularRes      = 0.001
		imageSize       = 100
		animationFrames = 64
		frameDelay      = 8
	)
	yOscFrequency := rand.Float64()
	animation := gif.GIF{LoopCount: animationFrames}
	phaseDiff := 0.0
	for i := 0; i < animationFrames; i++ {
		rect := image.Rect(0, 0, 2*imageSize+1, 2*imageSize+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < oscCycles*2*math.Pi; t += angularRes {
			x := math.Sin(t)
			y := math.Sin(t*yOscFrequency + phaseDiff)
			img.SetColorIndex(imageSize+int(x*imageSize+0.5), imageSize+int(y*imageSize+0.5), blackIndex)
		}
		phaseDiff += 0.1
		animation.Delay = append(animation.Delay, frameDelay)
		animation.Image = append(animation.Image, img)
	}
	gif.EncodeAll(out, &animation)
}
