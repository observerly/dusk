package dusk

import (
	"math"
	"time"
)

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

/*
	GetLunarLongitudeOfTheAscendingNode()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@returns the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date
	@see p.144 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarLongitudeOfTheAscendingNode(J float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var Ω = math.Mod(125.04452-1934.136261*J+0.0020708*math.Pow(J, 2)+math.Pow(J, 3)/450000, 360)

	// correct for negative angles
	if Ω < 0 {
		Ω += 360
	}

	return Ω
}

/*
	GetLunarLongitudeOfNutation()

	@param L - the ecliptic longitude at which the Sun could be found if its orbit were circular and free of perturbations.
	@param l - the ecliptic longitude at which the Moon could be found if its orbit were circular and free of perturbations.
	@param Ω - the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date
	@returns the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date (in arcseconds)
	@see p.144 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarLongitudeOfNutation(L float64, l float64, Ω float64) float64 {
	return -17.20*sinx(Ω) + 1.32*sinx(2*L) - 0.23*sinx(2*l) + 0.21*sinx(2*Ω)
}

/*
	GetLunarObliquityOfNutation()

	@param L - the ecliptic longitude at which the Sun could be found if its orbit were circular and free of perturbations.
	@param l - the ecliptic longitude at which the Moon could be found if its orbit were circular and free of perturbations.
	@param Ω - the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date
	@returns the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date (in arcseconds)
	@see p.144 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarObliquityOfNutation(L float64, l float64, Ω float64) float64 {
	return 9.20*cosx(Ω) + 0.57*cosx(2*L) + 0.1*cosx(2*l) - 0.09*cosx(2*Ω)
}

/*
  GetLunarEquatorialPosition()

 	@param datetime - the datetime in UTC of the observer
  @returns the Lunar equatorial position (right ascension & declination) in degrees:
*/
func GetLunarEquatorialPosition(datetime time.Time) EquatorialCoordinate {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(datetime)

	var M float64 = GetLunarMeanAnomaly(J)

	var L float64 = GetLunarMeanLongitude(J)

	var F float64 = GetLunarArgumentOfLatitude(J)

	var l float64 = GetLunarHorizontalLongitude(M, L)

	var b float64 = GetLunarHorizontalLatitude(F)

	var O float64 = GetEarthObliquity()

	// trigoneometric functions handle the correct degrees and radians conversions:
	var ra float64 = atan2yx(sinx(l)*cosx(O)-tanx(b)*sinx(O), cosx(l))

	// trigoneometric functions handle the correct degrees and radians conversions:
	var dec float64 = asinx(sinx(b)*cosx(O) + cosx(b)*sinx(O)*sinx(l))

	return EquatorialCoordinate{
		ra:  ra,
		dec: dec,
	}
}

/*
	GetLunarHourAngle()

	Observing the Moon from Earth, the lunar hour angle is an expression of time, expressed in angular measurement,
	usually degrees, from lunar noon. At lunar noon the hour angle is zero degrees, with the time before lunar noon
	expressed as negative degrees, and the local time after lunar noon expressed as positive degrees.

	@param δ - the ecliptic longitude of the Moon (in degrees)
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@returns the lunar hour angle for a given lunar declination, of some observer on Earth
*/
func GetLunarHourAngle(δ float64, latitude float64, elevation float64) float64 {
	// observations on a sea horizon needing an elevation-of-observer correction
	// (corrects for both apparent dip and terrestrial refraction):
	var corr = -2.076 * math.Sqrt(elevation) * 1 / 60

	return acosx((sinx(0.125-corr) - (sinx(latitude) * sinx(δ))) / cosx(latitude) * cosx(δ))
}
