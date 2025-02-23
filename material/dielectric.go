package material

import (
	"math"

	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/common"
	"github.com/anfilat/ray-tracing-go.git/hit"
	"github.com/anfilat/ray-tracing-go.git/ray"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Dielectric struct {
	RefractionIndex float64
}

func NewDielectric(ri float64) Dielectric {
	return Dielectric{ri}
}

func (d Dielectric) Scatter(rIn ray.Ray, rec *hit.Record) (color.Color, ray.Ray, bool) {
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
	if cannotRefract || reflectance(cosTheta, ri) > common.Random() {
		direction = vec3.Reflect(unitDirection, rec.Normal)
	} else {
		direction = vec3.Refract(unitDirection, rec.Normal, ri)
	}
	scattered := ray.New(rec.P, direction)

	return attenuation, scattered, true
}

func reflectance(cosine, rI float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - rI) / (1 + rI)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
