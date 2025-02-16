package sphere

import (
	"math"

	"github.com/anfilat/ray-tracing-go.git/hitTable"
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type Sphere struct {
	center point.Point
	radius float64
}

func New(center point.Point, radius float64) Sphere {
	return Sphere{
		center: center,
		radius: max(0, radius),
	}
}

func (s Sphere) Hit(r ray.Ray, rayT interval.Interval, rec *hitTable.HitRecord) bool {
	oc := s.center.Sub(r.Origin())
	a := r.Dir().LengthSquared()
	h := r.Dir().Dot(oc)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtD := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (h - sqrtD) / a
	if !rayT.Surrounds(root) {
		root = (h + sqrtD) / a
		if !rayT.Surrounds(root) {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.Sub(
		s.center,
	).DivF(
		s.radius,
	)
	rec.SetFaceNormal(r, outwardNormal)

	return true
}
