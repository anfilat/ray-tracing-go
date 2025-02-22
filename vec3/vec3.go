package vec3

import (
	"fmt"
	"math"

	"github.com/anfilat/ray-tracing-go.git/common"
)

type Vec3 struct {
	e [3]float64
}

func New(e0, e1, e2 float64) Vec3 {
	return Vec3{e: [3]float64{e0, e1, e2}}
}

func Random() Vec3 {
	return Vec3{e: [3]float64{common.Random(), common.Random(), common.Random()}}
}

func RandomMM(min, max float64) Vec3 {
	return Vec3{e: [3]float64{common.RandomMM(min, max), common.RandomMM(min, max), common.RandomMM(min, max)}}
}

func RandomUnitVector() Vec3 {
	for {
		p := RandomMM(-1, 1)
		lenSQ := p.LengthSquared()
		if 1e-160 < lenSQ && lenSQ <= 1 {
			return p.DivF(math.Sqrt(lenSQ))
		}
	}
}

func RandomOnHemisphere(normal Vec3) Vec3 {
	onUnitSphere := RandomUnitVector()
	if onUnitSphere.Dot(normal) > 0 {
		// In the same hemisphere as the normal
		return onUnitSphere
	}
	return onUnitSphere.Inv()
}

func Reflect(v, n Vec3) Vec3 {
	return v.Sub(n.MulF(2 * v.Dot(n)))
}

func Refract(uv, n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := min(uv.Inv().Dot(n), 1)
	rOutPerp := uv.Add(n.MulF(cosTheta)).MulF(etaiOverEtat)
	rOutParallel := n.MulF(-math.Sqrt(math.Abs(1 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}

func (v Vec3) X() float64 {
	return v.e[0]
}

func (v Vec3) Y() float64 {
	return v.e[1]
}

func (v Vec3) Z() float64 {
	return v.e[2]
}

func (v Vec3) I(i int) float64 {
	return v.e[i]
}

func (v Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) NearZero() bool {
	// Return true if the vector is close to zero in all dimensions.
	const s = 1e-8
	return (math.Abs(v.e[0]) < s) && (math.Abs(v.e[1]) < s) && (math.Abs(v.e[2]) < s)
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivF(v.Length())
}

func (v Vec3) Print() {
	fmt.Printf("%g %g %g", v.e[0], v.e[1], v.e[2])
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{e: [3]float64{v.e[0] + v2.e[0], v.e[1] + v2.e[1], v.e[2] + v2.e[2]}}
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{e: [3]float64{v.e[0] - v2.e[0], v.e[1] - v2.e[1], v.e[2] - v2.e[2]}}
}

func (v Vec3) Mul(v2 Vec3) Vec3 {
	return Vec3{e: [3]float64{v.e[0] * v2.e[0], v.e[1] * v2.e[1], v.e[2] * v2.e[2]}}
}

func (v Vec3) MulF(t float64) Vec3 {
	return Vec3{e: [3]float64{v.e[0] * t, v.e[1] * t, v.e[2] * t}}
}

func (v Vec3) DivF(t float64) Vec3 {
	return v.MulF(1 / t)
}

func (v Vec3) Inv() Vec3 {
	return Vec3{e: [3]float64{-v.e[0], -v.e[1], -v.e[2]}}
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.e[0]*v2.e[0] + v.e[1]*v2.e[1] + v.e[2]*v2.e[2]
}

func Cross(v1, v2 Vec3) Vec3 {
	return Vec3{e: [3]float64{
		v1.e[1]*v2.e[2] - v1.e[2]*v2.e[1],
		v1.e[2]*v2.e[0] - v1.e[0]*v2.e[2],
		v1.e[0]*v2.e[1] - v1.e[1]*v2.e[0],
	}}
}
