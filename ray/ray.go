package ray

import (
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Ray struct {
	origin point.Point
	dir    vec3.Vec3
}

func New(origin, dir point.Point) Ray {
	return Ray{origin: origin, dir: dir}
}

func (r Ray) Origin() point.Point {
	return r.origin
}

func (r Ray) Dir() point.Point {
	return r.dir
}

func (r Ray) At(t float64) point.Point {
	return point.New(r.origin.Add(r.dir.MulF(t)))
}
