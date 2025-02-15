package main

import (
	"fmt"
	"os"
)

func writeImage() {
	width := 256
	height := 256

	fmt.Printf("P3\n%d %d\n255\n", width, height)

	for y := 0; y < height; y++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", height-y)
		for x := 0; x < width; x++ {
			r := float64(x) / float64(width-1)
			g := float64(y) / float64(height-1)
			b := 0.0

			red := int(255.999 * r)
			green := int(255.999 * g)
			blue := int(255.999 * b)
			fmt.Printf("%d %d %d\n", red, green, blue)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone                          \n")
}
