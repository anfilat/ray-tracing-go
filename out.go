package main

import (
	"fmt"
	"os"

	"github.com/anfilat/ray-tracing-go.git/color"
)

func writeImage() {
	width := 256
	height := 256

	fmt.Printf("P3\n%d %d\n255\n", width, height)

	for y := 0; y < height; y++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", height-y)
		for x := 0; x < width; x++ {
			pixelColor := color.New(float64(x)/float64(width-1), float64(y)/float64(height-1), 0)
			color.WriteColor(os.Stdout, pixelColor)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone                          \n")
}
