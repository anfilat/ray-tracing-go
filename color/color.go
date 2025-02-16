package color

import (
	"fmt"
	"io"
	"math"

	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Color = vec3.Vec3

func New(vec vec3.Vec3) Color {
	return vec
}

func NewRGB(r, g, b float64) Color {
	return vec3.New(r, g, b)
}

func Write(w io.Writer, pixelColor Color) {
	r := pixelColor.X()
	g := pixelColor.Y()
	b := pixelColor.Z()

	// Apply a linear to gamma transform for gamma 2
	r = linearToGamma(r)
	g = linearToGamma(g)
	b = linearToGamma(b)

	// Translate the [0,1] component values to the byte range [0,255].
	intensity := interval.New(0.000, 0.999)
	rByte := int(256 * intensity.Clamp(r))
	gByte := int(256 * intensity.Clamp(g))
	bByte := int(256 * intensity.Clamp(b))

	fmt.Fprintf(w, "%d %d %d\n", rByte, gByte, bByte)
}

func linearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}

	return 0
}
