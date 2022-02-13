package dusk

import "math"

/*
	GetSolarMeanAnomaly()

	@param J - the Ephemeris Time or the number of centuries since J2000 epoch
	@returns the non-uniform or anomalous apparent motion of the Sun along the plane of the ecliptic
	@see EQ.47.3 p.338 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell
*/
func GetSolarMeanAnomaly(J float64) float64 {
	// applies modulo correction to the angle, and ensures always positive:
	return math.Mod(357.5291092+(0.98560028*J), 360) + 360
}
