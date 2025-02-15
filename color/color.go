package color

import (
	"fmt"
	"io"

	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Color struct {
	vec3.Vec3
}

func New(r, g, b float64) *Color {
	return &Color{Vec3: *vec3.New(r, g, b)}
}

func WriteColor(w io.Writer, pixelColor *Color) {
	r := pixelColor.X()
	g := pixelColor.Y()
	b := pixelColor.Z()

	rByte := int(255.999 * r)
	gByte := int(255.999 * g)
	bByte := int(255.999 * b)

	fmt.Fprintf(w, "%d %d %d\n", rByte, gByte, bByte)
}
