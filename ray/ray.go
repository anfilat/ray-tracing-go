package ray

import "github.com/anfilat/ray-tracing-go.git/point"

type Ray struct {
	origin point.Point
	dir    point.Point
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
	return point.New(point.Add(r.origin, point.MulF(r.dir, t)).Vec())
}
