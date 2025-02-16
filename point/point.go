package point

import "github.com/anfilat/ray-tracing-go.git/vec3"

type Point = vec3.Vec3

func New(vec vec3.Vec3) Point {
	return vec
}

func NewXYZ(x, y, z float64) Point {
	return vec3.New(x, y, z)
}
