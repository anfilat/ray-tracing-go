package hit

import (
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type Material interface {
	Scatter(rIn ray.Ray, rec *Record) (color.Color, ray.Ray, bool)
}
