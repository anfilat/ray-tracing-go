package point3

import "github.com/anfilat/ray-tracing-go.git/vec3"

type Point3 struct {
	vec vec3.Vec3
}

func New(vec *vec3.Vec3) *Point3 {
	return &Point3{vec: *vec}
}

func (p *Point3) Vec() *vec3.Vec3 {
	return &p.vec
}
