package dusk

import "math"

/*
	GetLunarMeanLongitude()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@returns the ecliptic longitude at which the Moon could be found if its orbit were circular and free of perturbations.
	@see EQ.47.1 p.338 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarMeanLongitude(J float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var L = math.Mod((218.3164477 + 481267.88123421*J - 0.0015786*math.Pow(J, 2) + math.Pow(J, 3)/538841 - math.Pow(J, 4)/65194000), 360)

	// correct for negative angles
	if L < 0 {
		L += 360
	}

	return L
}

/*
	GetLunarMeanAnomaly()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@returns the non-uniform or anomalous apparent motion of the Moon along the plane of the ecliptic
	@see EQ.47.4 p.338 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarMeanAnomaly(J float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var M float64 = math.Mod((134.9633964 + 477198.8675055*J + 0.0087414*math.Pow(J, 2) + math.Pow(J, 3)/69699 - math.Pow(J, 4)/14712000), 360)

	// correct for negative angles
	if M < 0 {
		M += 360
	}

	return M
}

/*
  GetLunarArgumentOfLatitude()

  @param J - the Ephemeris time or the number of centuries since J2000 epoch
  @returns the Lunar argument of latitude
	@see EQ.47.5 p.338 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarArgumentOfLatitude(J float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var F float64 = math.Mod((93.272095 + 483202.0175233*J - 0.0036539*math.Pow(J, 2) + math.Pow(J, 3)/3526000 - math.Pow(J, 4)/863310000), 360)

	// correct for negative angles
	if F < 0 {
		F += 360
	}

	return F
}

/*
  GetLunarHorizontalLongitude()

 	@param M - the mean lunar anomaly for the Ephemeris time or the number of centuries since J2000 epoch
	@param L - the ecliptic longitude at which the Moon could be found if its orbit were circular and free of perturbations.
  @returns the Lunar horizontal longitude
*/
func GetLunarHorizontalLongitude(M float64, L float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var l = math.Mod(L+6.289*sinx(M), 360)

	// correct for negative angles
	if l < 0 {
		l += 360
	}

	return l
}

/*
  GetLunarHorizontalLatitude()

 	@param F - the Lunar argument of latitude
  @returns the Lunar horizontal latitude
*/
func GetLunarHorizontalLatitude(F float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var b = math.Mod(5.128*sinx(F), 360)

	// correct for negative angles
	if b < 0 {
		b += 360
	}

	return b
}
