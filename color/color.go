package color

import (
	"fmt"
	"io"

	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Color struct {
	vec vec3.Vec3
}

func New(r, g, b float64) Color {
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

func WriteColor(w io.Writer, pixelColor Color) {
	rByte := int(255.999 * pixelColor.R())
	gByte := int(255.999 * pixelColor.G())
	bByte := int(255.999 * pixelColor.B())

	fmt.Fprintf(w, "%d %d %d\n", rByte, gByte, bByte)
}
