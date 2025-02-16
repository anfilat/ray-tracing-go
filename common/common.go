package common

import (
	"math"
	"math/rand/v2"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func Random() float64 {
	// Returns a random real in [0,1).
	return rand.Float64()
}

func RandomDouble(min, max float64) float64 {
	// Returns a random real in [min,max).
	return min + (max-min)*Random()
}
