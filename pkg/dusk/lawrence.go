package dusk

import (
	"math"
	"time"
)

/*
  GetLunarMeanAnomalyLawrence()

  @returns the mean lunar anomaly as measured from the moment of perigee
  @see p.165 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetLunarMeanAnomalyLawrence(datetime time.Time) float64 {
	// the number of days since the standard epoch J2000:
	var De = GetFractionalJulianDaysSinceStandardEpoch(datetime)

	var λ = GetLunarMeanEclipticLongitude(datetime)

	// the Moon's ecliptic longitude at perigee of the epoch J2000 (given by the The Astronomical Almanac, 2000):
	var ϖ0 = 83.353451

	// eq. 7.3.3 p.165 of Lawrence, J.L. 2015. Celestial Calculations. Cambridge, Ma: The MIT Press
	var M = math.Mod(λ-(0.1114041*De)-ϖ0, 360)

	// correct for negative angles
	if M < 0 {
		M += 360
	}

	return M
}

/*
  GetMoonEclipticPositionLawrence()

  @param datetime - the datetime in UTC of the observer
  @returns the geocentric ecliptic coodinate (λ - geocentric longitude, β - geocentric latidude) of the Moon.
  @see ch.7 p.163-169 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetLunarEclipticPositionLawrence(datetime time.Time) EclipticCoordinate {
	var Ωprime float64 = GetLunarCorrectedEclipticLongitudeOfTheAscendingNode(datetime)

	var λt float64 = GetLunarTrueEclipticLongitude(datetime)

	// the inclination of the Moon's orbit with respect to the ecliptic
	var ι float64 = 5.1453964

	var x float64 = cosx(λt - Ωprime)

	var y float64 = sinx(λt-Ωprime) * cosx(ι)

	// utilise atan2yx to determine a quadrant adjustment for arctan
	var T float64 = math.Mod(atan2yx(y, x), 360)

	// correct for negative angles
	if T < 0 {
		T += 360
	}

	// ensure values for ecliptic longitude are corrected to between [0°, 360°]
	var λm float64 = math.Mod(Ωprime+T, 360)

	var βm float64 = asinx(sinx(λt-Ωprime) * sinx(ι))

	return EclipticCoordinate{
		λ: λm,
		β: βm,
	}
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
