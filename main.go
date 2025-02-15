package main

import (
	"fmt"
	"os"

	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

func main() {
	// Image
	const aspectRatio = 16.0 / 9.0
	const imageWidth = 400

	// Calculate the image height, and ensure that it's at least 1.
	imageHeight := int(imageWidth / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera
	const focalLength = 1.0
	const viewportHeight = 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := point.NewXYZ(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := point.NewXYZ(viewportWidth, 0, 0)
	viewportV := point.NewXYZ(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := point.DivF(viewportU, imageWidth)
	pixelDeltaV := point.DivF(viewportV, float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := point.Sub(
		point.Sub(
			point.Sub(
				cameraCenter,
				point.NewXYZ(0, 0, focalLength),
			),
			point.DivF(viewportU, 2),
		),
		point.DivF(viewportV, 2),
	)
	pixel00Loc := point.Add(
		viewportUpperLeft,
		point.MulF(point.Add(pixelDeltaU, pixelDeltaV), 0.5),
	)

	// Render

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for y := 0; y < imageHeight; y++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", imageHeight-y)
		for x := 0; x < imageWidth; x++ {
			pixelCenter := point.Add(
				pixel00Loc,
				point.Add(
					point.MulF(pixelDeltaU, float64(x)),
					point.MulF(pixelDeltaV, float64(y)),
				),
			)
			rayDirection := point.Sub(pixelCenter, cameraCenter)
			r := ray.New(cameraCenter, rayDirection)

			pixelColor := rayColor(r)
			color.WriteColor(os.Stdout, pixelColor)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone                          \n")
}

func rayColor(r ray.Ray) color.Color {
	unitDirection := r.Dir().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)

	return color.Add(
		color.MulF(color.NewRGB(1, 1, 1), 1-a),
		color.MulF(color.NewRGB(0.5, 0.7, 1), a),
	)
}
