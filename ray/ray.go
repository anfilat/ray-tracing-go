package ray

import (
	"github.com/anfilat/ray-tracing-go.git/point3"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

type Ray struct {
	Origin *point3.Point3
	Dir    *vec3.Vec3
}

func New(origin *point3.Point3, dir *vec3.Vec3) *Ray {
	return &Ray{Origin: origin, Dir: dir}
}

func (r *Ray) At(t float64) *point3.Point3 {
	return point3.New(vec3.Add(r.Origin.Vec(), vec3.MulF(r.Dir, t)))
}
