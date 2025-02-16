package hitTable

import (
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type HitRecord struct {
	P      point.Point
	Normal point.Point
	T      float64
}

type HitTable interface {
	Hit(r ray.Ray, rayTMin, rayTMax float64, rec *HitRecord) bool
}
