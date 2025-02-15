package vec3

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float64
}

func New(e0, e1, e2 float64) Vec3 {
	return Vec3{e: [3]float64{e0, e1, e2}}
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
