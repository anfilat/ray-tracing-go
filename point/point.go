package point

import "github.com/anfilat/ray-tracing-go.git/vec3"

type Point struct {
	vec vec3.Vec3
}

func New(vec vec3.Vec3) Point {
	return Point{vec: vec}
}

func (p Point) Vec() vec3.Vec3 {
	return p.vec
}
