package color

import (
	"fmt"
	"io"

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
	fmt.Fprintf(w,
		"%d %d %d\n",
		int(255.999*pixelColor.X()),
		int(255.999*pixelColor.Y()),
		int(255.999*pixelColor.Z()),
	)
}
