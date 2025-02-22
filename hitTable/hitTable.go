package hitTable

import (
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type HitRecord struct {
	P         point.Point
	Normal    point.Point
	Mat       Material
	T         float64
	FrontFace bool
}

func (h *HitRecord) Copy(rec *HitRecord) {
	h.P = rec.P
	h.Normal = rec.Normal
	h.Mat = rec.Mat
	h.T = rec.T
	h.FrontFace = rec.FrontFace
}

func (h *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal point.Point) {
	// Sets the hit record normal vector.
	// NOTE: the parameter `outward_normal` is assumed to have unit length.

	h.FrontFace = r.Dir().Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Inv()
	}
}

type HitTable interface {
	Hit(r ray.Ray, rayT interval.Interval, rec *HitRecord) bool
}
