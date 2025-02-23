package material

import (
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/hit"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Metal struct {
	Albedo color.Color
	Fuzz   float64
}

func NewMetal(albedo color.Color, fuzz float64) Metal {
	return Metal{
		Albedo: albedo,
		Fuzz:   min(fuzz, 1),
	}
}

func (m Metal) Scatter(rIn ray.Ray, rec *hit.Record) (color.Color, ray.Ray, bool) {
	reflected := vec3.Reflect(rIn.Dir(), rec.Normal)
	reflected = reflected.UnitVector().Add(vec3.RandomUnitVector().MulF(m.Fuzz))
	scattered := ray.New(rec.P, reflected)
	return m.Albedo, scattered, scattered.Dir().Dot(rec.Normal) > 0
}
