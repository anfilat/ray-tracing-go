package point

import "github.com/anfilat/ray-tracing-go.git/vec3"

type Point struct {
	vec vec3.Vec3
}

func New(vec vec3.Vec3) Point {
	return Point{vec: vec}
}

func NewXYZ(x, y, z float64) Point {
	return Point{vec: vec3.New(x, y, z)}
}

func (p Point) Vec() vec3.Vec3 {
	return p.vec
}

func (p Point) LengthSquared() float64 {
	return p.vec.LengthSquared()
}

func (p Point) UnitVector() vec3.Vec3 {
	return p.vec.UnitVector()
}

func (p Point) Add(p2 Point) Point {
	return New(p.vec.Add(p2.vec))
}

func (p Point) Sub(p2 Point) Point {
	return New(p.vec.Sub(p2.vec))
}

func (p Point) MulF(t float64) Point {
	return New(p.vec.MulF(t))
}

func (p Point) DivF(t float64) Point {
	return p.MulF(1 / t)
}

func (p Point) Dot(p2 Point) float64 {
	return p.vec.Dot(p2.vec)
}
