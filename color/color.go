package color

import (
	"fmt"
	"io"

	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Color struct {
	vec vec3.Vec3
}

func New(vec vec3.Vec3) Color {
	return Color{vec: vec}
}

func NewRGB(r, g, b float64) Color {
	return Color{vec: vec3.New(r, g, b)}
}

func (c Color) Vec() vec3.Vec3 {
	return c.vec
}

func (c Color) R() float64 {
	return c.vec.X()
}

func (c Color) G() float64 {
	return c.vec.Y()
}

func (c Color) B() float64 {
	return c.vec.Z()
}

func Add(v1, v2 Color) Color {
	return New(vec3.Add(v1.vec, v2.vec))
}

func MulF(v Color, t float64) Color {
	return New(vec3.MulF(v.vec, t))
}

func WriteColor(w io.Writer, pixelColor Color) {
	rByte := int(255.999 * pixelColor.R())
	gByte := int(255.999 * pixelColor.G())
	bByte := int(255.999 * pixelColor.B())

	fmt.Fprintf(w, "%d %d %d\n", rByte, gByte, bByte)
}
