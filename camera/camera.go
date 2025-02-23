package camera

import (
	"fmt"
	"math"
	"os"

	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/common"
	"github.com/anfilat/ray-tracing-go.git/hit"
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Camera struct {
	AspectRatio     float64     // Ratio of image width over height
	ImageWidth      int         // Rendered image width in pixel count
	SamplesPerPixel int         // Count of random samples for each pixel
	MaxDepth        int         // Maximum number of ray bounces into scene
	Vfov            float64     // Vertical view angle (field of view)
	LookFrom        point.Point // Point camera is looking from
	LookAt          point.Point // Point camera is looking at
	Vup             vec3.Vec3   // Camera-relative "up" direction
	DefocusAngle    float64     // Variation angle of rays through each pixel
	FocusDist       float64     // Distance from camera lookfrom point to plane of perfect focus

	imageHeight       int         // Rendered image height
	pixelSamplesScale float64     // Color scale factor for a sum of pixel samples
	center            point.Point // Camera center
	pixel00Loc        point.Point // Location of pixel 0, 0
	pixelDeltaU       vec3.Vec3   // Offset to pixel to the right
	pixelDeltaV       vec3.Vec3   // Offset to pixel below
	u, v, w           vec3.Vec3   // Camera frame basis vectors
	defocusDiskU      vec3.Vec3   // Defocus disk horizontal radius
	defocusDiskV      vec3.Vec3   // Defocus disk vertical radius
}

func New() *Camera {
	return &Camera{
		AspectRatio:     1,
		ImageWidth:      100,
		SamplesPerPixel: 10,
		MaxDepth:        10,
		Vfov:            90,
		LookFrom:        point.NewXYZ(0, 0, 0),
		LookAt:          point.NewXYZ(0, 0, -1),
		Vup:             vec3.New(0, 1, 0),
		DefocusAngle:    0,
		FocusDist:       10,
	}
}

func (c *Camera) Render(world hit.Table) {
	c.initialize()

	fmt.Printf("P3\n%d %d\n255\n", c.ImageWidth, c.imageHeight)

	for y := 0; y < c.imageHeight; y++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", c.imageHeight-y)
		for x := 0; x < c.ImageWidth; x++ {
			pixelColor := color.NewRGB(0, 0, 0)
			for sample := 0; sample < c.SamplesPerPixel; sample++ {
				r := c.getRay(x, y)
				pixelColor = pixelColor.Add(c.rayColor(r, c.MaxDepth, world))
			}
			color.Write(os.Stdout, pixelColor.MulF(c.pixelSamplesScale))
		}
	}

	fmt.Fprint(os.Stderr, "\rDone                          \n")
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.pixelSamplesScale = 1 / float64(c.SamplesPerPixel)

	c.center = c.LookFrom

	// Determine viewport dimensions.
	theta := common.DegreesToRadians(c.Vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h * c.FocusDist
	viewportWidth := viewportHeight * float64(c.ImageWidth) / float64(c.imageHeight)

	// Calculate the u,v,w unit basis vectors for the camera coordinate frame.
	c.w = c.LookFrom.Sub(c.LookAt).UnitVector()
	c.u = c.Vup.Cross(c.w).UnitVector()
	c.v = c.w.Cross(c.u)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := c.u.MulF(viewportWidth)        // Vector across viewport horizontal edge
	viewportV := c.v.Inv().MulF(viewportHeight) // Vector down viewport vertical edge

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	c.pixelDeltaU = viewportU.DivF(float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.DivF(float64(c.imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := c.center.Sub(
		c.w.MulF(c.FocusDist),
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

	// Calculate the camera defocus disk basis vectors.
	defocusRadius := c.FocusDist * math.Tan(common.DegreesToRadians(c.DefocusAngle/2))
	c.defocusDiskU = c.u.MulF(defocusRadius)
	c.defocusDiskV = c.v.MulF(defocusRadius)
}

func (c *Camera) getRay(x, y int) ray.Ray {
	// Construct a camera ray originating from the defocus disk and directed at a randomly
	// sampled point around the pixel location i, j.

	offset := c.sampleSquare()
	pixelSample := c.pixel00Loc.Add(
		c.pixelDeltaU.MulF(float64(x) + offset.X()),
	).Add(
		c.pixelDeltaV.MulF(float64(y) + offset.Y()),
	)

	var rayOrigin point.Point
	if c.DefocusAngle <= 0 {
		rayOrigin = c.center
	} else {
		rayOrigin = c.defocusDiskSample()
	}
	rayDirection := pixelSample.Sub(rayOrigin)

	return ray.New(rayOrigin, rayDirection)
}

func (c *Camera) sampleSquare() vec3.Vec3 {
	// Returns the vector to a random point in the [-.5,-.5]-[+.5,+.5] unit square.
	return vec3.New(common.Random()-0.5, common.Random()-0.5, 0)
}

func (c *Camera) defocusDiskSample() point.Point {
	// Returns a random point in the camera defocus disk.
	p := vec3.RandomInUnitDisk()
	return c.center.Add(c.defocusDiskU.MulF(p.X())).Add(c.defocusDiskV.MulF(p.Y()))
}

func (c *Camera) rayColor(r ray.Ray, depth int, world hit.Table) color.Color {
	// If we've exceeded the ray bounce limit, no more light is gathered.
	if depth <= 0 {
		return color.NewRGB(0, 0, 0)
	}

	rec, isHit := world.Hit(r, interval.New(0.001, math.Inf(1)))
	if isHit {
		attenuation, scattered, ok := rec.Mat.Scatter(r, rec)
		if ok {
			return attenuation.Mul(c.rayColor(scattered, depth-1, world))
		}
		return color.NewRGB(0, 0, 0)
	}

	unitDirection := r.Dir().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)
	return color.NewRGB(1, 1, 1).MulF(1 - a).Add(
		color.NewRGB(0.5, 0.7, 1).MulF(a),
	)
}
