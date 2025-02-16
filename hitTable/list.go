package hitTable

import (
	"github.com/anfilat/ray-tracing-go.git/interval"
	"github.com/anfilat/ray-tracing-go.git/ray"
)

type List struct {
	objects []HitTable
}

func NewList() List {
	return List{}
}

func (l *List) Add(object HitTable) {
	l.objects = append(l.objects, object)
}

func (l *List) Clear() {
	l.objects = l.objects[:0]
}

func (l *List) Hit(r ray.Ray, rayT interval.Interval, rec *HitRecord) bool {
	tempRec := &HitRecord{}
	hitAnything := false
	closestSoFar := rayT.Max

	for _, object := range l.objects {
		if object.Hit(r, interval.New(rayT.Min, closestSoFar), tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			rec.Copy(tempRec)
		}
	}

	return hitAnything
}
