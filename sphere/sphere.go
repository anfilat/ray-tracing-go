package sphere

import (
	hitTable "github.com/anfilat/ray-tracing-go.git/hit_table"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"math"
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

func (s Sphere) Hit(r ray.Ray, rayTMin, rayTMax float64, rec *hitTable.HitRecord) bool {
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
	if root <= rayTMin || rayTMax <= root {
		root = (h + sqrtD) / a
		if root <= rayTMin || rayTMax <= root {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	rec.Normal = rec.P.Sub(
		s.center,
	).DivF(
		s.radius,
	)

	return true
}
