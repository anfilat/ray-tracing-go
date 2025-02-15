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
	pixelDeltaU := viewportU.DivF(imageWidth)
	pixelDeltaV := viewportV.DivF(float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Sub(
		point.NewXYZ(0, 0, focalLength),
	).Sub(
		viewportU.DivF(2),
	).Sub(
		viewportV.DivF(2),
	)
	pixel00Loc := pixelDeltaU.Add(
		pixelDeltaV,
	).MulF(
		0.5,
	).Add(
		viewportUpperLeft,
	)

	// Render

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for y := 0; y < imageHeight; y++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", imageHeight-y)
		for x := 0; x < imageWidth; x++ {
			pixelCenter := pixel00Loc.Add(
				pixelDeltaU.MulF(float64(x)),
			).Add(
				pixelDeltaV.MulF(float64(y)),
			)
			rayDirection := pixelCenter.Sub(cameraCenter)
			r := ray.New(cameraCenter, rayDirection)

			pixelColor := rayColor(r)
			pixelColor.Write(os.Stdout)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone                          \n")
}

func rayColor(r ray.Ray) color.Color {
	unitDirection := r.Dir().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)

	return color.NewRGB(1, 1, 1).MulF(1 - a).Add(
		color.NewRGB(0.5, 0.7, 1).MulF(a),
	)
}
