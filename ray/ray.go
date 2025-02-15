package ray

import (
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Ray struct {
	origin point.Point
	dir    vec3.Vec3
}

func New(origin point.Point, dir vec3.Vec3) *Ray {
	return &Ray{origin: origin, dir: dir}
}

func (r Ray) Origin() point.Point {
	return r.origin
}

func (r Ray) Dir() vec3.Vec3 {
	return r.dir
}

func (r Ray) At(t float64) point.Point {
	return point.New(vec3.Add(r.origin.Vec(), vec3.MulF(r.dir, t)))
}
