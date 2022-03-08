package dusk

import (
	"math"
	"time"
)

/*
  GetLunarMeanEclipticLongitude()

  @returns the mean lunar ecliptic longitude as measured from the moment of perigee
  @see p.165 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetLunarMeanEclipticLongitude(datetime time.Time) float64 {
	// the number of days since the standard epoch J2000:
	var De = GetFractionalJulianDaysSinceStandardEpoch(datetime)

	// the Moon's ecliptic longitude at tge epoch J2OOO:
	var λ0 = 218.316433

	// eq. 7.3.1 p.165 of Lawrence, J.L. 2015. Celestial Calculations. Cambridge, Ma: The MIT Press
	var λ = math.Mod((13.176339686*De)+λ0, 360)

	return λ
}

/*
  GetSolarMeanAnomalyLawrence()

  @returns the mean solar anomaly as measured from the moment of perigee
  @see p.136 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetSolarMeanAnomalyLawrence(datetime time.Time) float64 {
	// the number of days since the standard epoch J2000:
	var De = GetFractionalJulianDaysSinceStandardEpoch(datetime)

	// the Sun's ecliptic longitude at the epoch J2000 (given by the The Astronomical Almanac, 2000):
	var εg = 280.466069

	// the Sun's ecliptic lonitude at perigee of the epoch J2000 (given by the The Astronomical Almanac, 2000):
	var ϖg = 282.938346

	// eq. 6.2.3 p.134 of Lawrence, J.L. 2015. Celestial Calculations. Cambridge, Ma: The MIT Press
	var M = math.Mod(((360*De)/365.242191)+εg-ϖg, 360)

	return M
}

/*
  GetSolarEquationOfCenterLawrence()

  @returns the approximate equation of center for the Sun
  @see p.136 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetSolarEquationOfCenterLawrence(M float64) float64 {
	// the eccentricity of the Earth-Sun orbit at the epoch J2000 (given by the The Astronomical Almanac, 2000):
	var e = 0.016708

	// eq. 6.2.4 p.135 of Lawrence, J.L. 2015. Celestial Calculations. Cambridge, Ma: The MIT Press
	var Ec = (360 / math.Pi) * e * sinx(M)

	return Ec
}

/*
  GetSolarEclipticLongitudeLawrence()

  @param M - the mean solar anomaly as measured from the moment of perigee
  @param C - the approximate equation of center for the Sun
  @returns the ecliptic longitude of the Sun
  @see p.136 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetSolarEclipticLongitudeLawrence(M float64, C float64) float64 {
	// the Sun's ecliptic lonitude at perigee of the epoch J2000 (given by the The Astronomical Almanac, 2000):
	var ϖg = 282.938346

	// eq. 6.2.5 & 6.2.6 p.134 of Lawrence, J.L. 2015. Celestial Calculations. Cambridge, Ma: The MIT Press
	var λ = math.Mod(M+C+ϖg, 360)

	// adjust for angles larger than 360°:
	if λ > 360 {
		λ -= 360
	}

	return λ
}
