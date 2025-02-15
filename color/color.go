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

func (c Color) Add(c2 Color) Color {
	return New(c.vec.Add(c2.vec))
}

func (c Color) MulF(t float64) Color {
	return New(c.vec.MulF(t))
}

func (c Color) Write(w io.Writer) {
	fmt.Fprintf(w,
		"%d %d %d\n",
		int(255.999*c.R()),
		int(255.999*c.G()),
		int(255.999*c.B()),
	)
}
