package common

import "math"

var Infinity = math.Inf(0)

const Pi = math.Pi

func DegreesToRadians(degrees float64) float64 {
	return degrees * Pi / 180.0
}
