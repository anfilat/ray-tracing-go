package material

import (
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/hit"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Lambertian struct {
	Albedo color.Color
}

func NewLambertian(albedo color.Color) Lambertian {
	return Lambertian{albedo}
}

func (l Lambertian) Scatter(_ ray.Ray, rec *hit.Record) (color.Color, ray.Ray, bool) {
	scatterDirection := rec.Normal.Add(vec3.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}

	scattered := ray.New(rec.P, scatterDirection)
	return l.Albedo, scattered, true
}
