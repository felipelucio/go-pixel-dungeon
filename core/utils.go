package core

import "math"

func Clamp(val float64, min float64, max float64) float64 {
	return math.Max(min, math.Min(val, max))
}
