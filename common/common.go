package common

import "math"

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

type Interval struct {
	Min, Max float64
}

var (
	EmptyInterval    = NewInterval(math.Inf(1), math.Inf(-1))
	UniverseInterval = NewInterval(math.Inf(-1), math.Inf(1))
)

func NewInterval(min, max float64) Interval {
	return Interval{Min: min, Max: max}
}

func (i Interval) Size() float64 {
	return i.Max - i.Min
}

func (i Interval) Contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i Interval) Surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}
