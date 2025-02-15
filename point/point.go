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

func (p Point) UnitVector() vec3.Vec3 {
	return p.vec.UnitVector()
}

func Add(v1, v2 Point) Point {
	return New(vec3.Add(v1.vec, v2.vec))
}

func Sub(v1, v2 Point) Point {
	return New(vec3.Sub(v1.vec, v2.vec))
}

func MulF(v Point, t float64) Point {
	return New(vec3.MulF(v.vec, t))
}

func DivF(v Point, t float64) Point {
	return MulF(v, 1/t)
}
