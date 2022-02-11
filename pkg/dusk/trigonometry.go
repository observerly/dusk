package dusk

import "math"

const (
	radToDeg = 180.0 / math.Pi
	degToRad = math.Pi / 180.0
)

func sinx(x float64) float64 {
	return math.Sin(x * degToRad)
}

func cosx(x float64) float64 {
	return math.Cos(x * degToRad)
}

func tanx(x float64) float64 {
	return math.Tan(x * degToRad)
}

func asinx(x float64) float64 {
	return radToDeg * math.Asin(x)
}

func acosx(x float64) float64 {
	return radToDeg * math.Acos(x)
}

func atanx(x float64) float64 {
	return radToDeg * math.Atan(x)
}

func atan2yx(y, x float64) float64 {
	return radToDeg * math.Atan2(y, x)
}
