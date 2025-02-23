package sphere

import (
	"math"

	"github.com/anfilat/ray-tracing-go.git/hit"
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type Sphere struct {
	center point.Point
	radius float64
	mat    hit.Material
}

func New(center point.Point, radius float64, mat hit.Material) Sphere {
	return Sphere{
		center: center,
		radius: max(0, radius),
		mat:    mat,
	}
}

func (s Sphere) Hit(r ray.Ray, rayT interval.Interval) (*hit.Record, bool) {
	oc := s.center.Sub(r.Origin())
	a := r.Dir().LengthSquared()
	h := r.Dir().Dot(oc)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return nil, false
	}

	sqrtD := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (h - sqrtD) / a
	if !rayT.Surrounds(root) {
		root = (h + sqrtD) / a
		if !rayT.Surrounds(root) {
			return nil, false
		}
	}

	result := &hit.Record{
		P:   r.At(root),
		Mat: s.mat,
		T:   root,
	}
	outwardNormal := result.P.Sub(
		s.center,
	).DivF(
		s.radius,
	)
	result.SetFaceNormal(r, outwardNormal)

	return result, true
}
