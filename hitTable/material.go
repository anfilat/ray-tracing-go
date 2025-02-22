package hitTable

import (
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"github.com/anfilat/ray-tracing-go.git/vec3"
	"math"
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
	Fuzz   float64
}

func NewMetal(albedo color.Color, fuzz float64) Metal {
	return Metal{
		Albedo: albedo,
		Fuzz:   min(fuzz, 1),
	}
}

func (m Metal) Scatter(rIn ray.Ray, rec *HitRecord) (color.Color, ray.Ray, bool) {
	reflected := vec3.Reflect(rIn.Dir(), rec.Normal)
	reflected = reflected.UnitVector().Add(vec3.RandomUnitVector().MulF(m.Fuzz))
	scattered := ray.New(rec.P, reflected)
	return m.Albedo, scattered, scattered.Dir().Dot(rec.Normal) > 0
}

type Dielectric struct {
	RefractionIndex float64
}

func NewDielectric(ri float64) Dielectric {
	return Dielectric{ri}
}

func (d Dielectric) Scatter(rIn ray.Ray, rec *HitRecord) (color.Color, ray.Ray, bool) {
	attenuation := color.NewRGB(1, 1, 1)

	ri := d.RefractionIndex
	if rec.FrontFace {
		ri = 1.0 / ri
	}

	unitDirection := rIn.Dir().UnitVector()
	cosTheta := min(unitDirection.Inv().Dot(rec.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	cannotRefract := ri*sinTheta > 1
	var direction vec3.Vec3
	if cannotRefract {
		direction = vec3.Reflect(unitDirection, rec.Normal)
	} else {
		direction = vec3.Refract(unitDirection, rec.Normal, ri)
	}
	scattered := ray.New(rec.P, direction)

	return attenuation, scattered, true
}
