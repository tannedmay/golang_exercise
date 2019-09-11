package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func main() {
	lissajous(os.Stdout)
}

// Return random intensity of RGB
func intensity() uint8 {
	return uint8(55 + (rand.Int() % 201))
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	var palette = make([]color.Color, 0, nframes)
	palette = append(palette, color.Black)
	for i := 0; i < nframes; i++ {
		// fmt.Println(intensity)
		colr := color.RGBA{intensity(), intensity(), intensity(), 255}
		palette = append(palette, colr)
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < 2*math.Pi*cycles; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size*0.75), size+int(y*size*0.75), uint8((i%(len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
