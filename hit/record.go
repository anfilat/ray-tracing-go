package hit

import (
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type Record struct {
	P         point.Point
	Normal    point.Point
	Mat       Material
	T         float64
	FrontFace bool
}

func (h *Record) SetFaceNormal(r ray.Ray, outwardNormal point.Point) {
	// Sets the hit record normal vector.
	// NOTE: the parameter `outward_normal` is assumed to have unit length.

	h.FrontFace = r.Dir().Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Inv()
	}
}
