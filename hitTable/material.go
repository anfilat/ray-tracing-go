package hitTable

import (
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Material interface {
	Scatter(rIn ray.Ray, rec *HitRecord) (color.Color, ray.Ray, bool)
}

type Lambertian struct {
	Albedo color.Color
}

func NewLambertian(albedo color.Color) Lambertian {
	return Lambertian{albedo}
}

func (l Lambertian) Scatter(_ ray.Ray, rec *HitRecord) (color.Color, ray.Ray, bool) {
	scatterDirection := rec.Normal.Add(vec3.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}

	scattered := ray.New(rec.P, scatterDirection)
	return l.Albedo, scattered, true
}

type Metal struct {
	Albedo color.Color
}

func NewMetal(albedo color.Color) Metal {
	return Metal{albedo}
}

func (m Metal) Scatter(rIn ray.Ray, rec *HitRecord) (color.Color, ray.Ray, bool) {
	reflected := vec3.Reflect(rIn.Dir(), rec.Normal)
	scattered := ray.New(rec.P, reflected)
	return m.Albedo, scattered, true
}
