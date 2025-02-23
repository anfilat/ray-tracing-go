package hit

import (
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type Table interface {
	Hit(r ray.Ray, rayT interval.Interval) (*Record, bool)
}
