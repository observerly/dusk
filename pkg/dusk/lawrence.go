package dusk

import (
	"math"
	"time"
)

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
