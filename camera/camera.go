package camera

import (
	"fmt"
	"math"
	"os"

	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/hitTable"
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Camera struct {
	AspectRatio float64 // Ratio of image width over height
	ImageWidth  int     // Rendered image width in pixel count

	imageHeight int         // Rendered image height
	center      point.Point // Camera center
	pixel00Loc  point.Point // Location of pixel 0, 0
	pixelDeltaU vec3.Vec3   // Offset to pixel to the right
	pixelDeltaV vec3.Vec3   // Offset to pixel below
}

func New() *Camera {
	return &Camera{
		AspectRatio: 1,
		ImageWidth:  100,
	}
}

func (c *Camera) Render(world hitTable.HitTable) {
	c.initialize()

	fmt.Printf("P3\n%d %d\n255\n", c.ImageWidth, c.imageHeight)

	for y := 0; y < c.imageHeight; y++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", c.imageHeight-y)
		for x := 0; x < c.ImageWidth; x++ {
			pixelCenter := c.pixel00Loc.Add(
				c.pixelDeltaU.MulF(float64(x)),
			).Add(
				c.pixelDeltaV.MulF(float64(y)),
			)
			rayDirection := pixelCenter.Sub(c.center)
			r := ray.New(c.center, rayDirection)

			pixelColor := c.rayColor(r, world)
			color.Write(os.Stdout, pixelColor)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone                          \n")
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.center = point.NewXYZ(0, 0, 0)

	// Determine viewport dimensions.
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * float64(c.ImageWidth) / float64(c.imageHeight)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := point.NewXYZ(viewportWidth, 0, 0)
	viewportV := point.NewXYZ(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	c.pixelDeltaU = viewportU.DivF(float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.DivF(float64(c.imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := c.center.Sub(
		point.NewXYZ(0, 0, focalLength),
	).Sub(
		viewportU.DivF(2),
	).Sub(
		viewportV.DivF(2),
	)

	c.pixel00Loc = c.pixelDeltaU.Add(
		c.pixelDeltaV,
	).MulF(
		0.5,
	).Add(
		viewportUpperLeft,
	)
}

func (c *Camera) rayColor(r ray.Ray, world hitTable.HitTable) color.Color {
	rec := &hitTable.HitRecord{}
	if world.Hit(r, interval.New(0, math.Inf(1)), rec) {
		return color.New(
			rec.Normal.Add(
				color.NewRGB(1, 1, 1),
			).MulF(
				0.5,
			),
		)
	}

	unitDirection := r.Dir().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)
	return color.NewRGB(1, 1, 1).MulF(1 - a).Add(
		color.NewRGB(0.5, 0.7, 1).MulF(a),
	)
}
