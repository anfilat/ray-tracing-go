package list

import (
	"github.com/anfilat/ray-tracing-go.git/hit"
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type List struct {
	objects []hit.Table
}

func New() *List {
	return &List{}
}

func (l *List) Add(object hit.Table) {
	l.objects = append(l.objects, object)
}

func (l *List) Clear() {
	l.objects = l.objects[:0]
}

func (l *List) Hit(r ray.Ray, rayT interval.Interval) (*hit.Record, bool) {
	result := &hit.Record{}
	hitAnything := false
	closestSoFar := rayT.Max

	for _, object := range l.objects {
		tempRec, ok := object.Hit(r, interval.New(rayT.Min, closestSoFar))
		if ok {
			result = tempRec
			hitAnything = true
			closestSoFar = tempRec.T
		}
	}

	return result, hitAnything
}
