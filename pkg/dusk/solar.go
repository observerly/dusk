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

/*
	GetSolarEquationOfCenter()

	@param M - the Mean Solar Anomaly for the Ephemeris Time or the number of centuries since J2000 epoch
	@returns the equation of center for the Sun
	@see p.164 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetSolarEquationOfCenter(M float64) float64 {
	return 1.9148*sinx(M) + 0.0200*sinx(2*M) + 0.0003*sinx(3*M)
}

/*
	GetSolarEclipticLongitude()

	@param M - the mean solar anomaly for the Ephemeris time or the number of centuries since J2000 epoch
	@param C - the equation of center for the Sun
	@returns the apparent Solar ecliptic longitude (in degrees)
*/
func GetSolarEclipticLongitude(M float64, C float64) float64 {
	// applies modulo correction to the angle, and ensures always positive:
	return math.Mod(M+C+180+102.9372, 360) + 360
}
