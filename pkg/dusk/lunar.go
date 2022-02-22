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
	GetLunarMeanElongation()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@returns the ecliptic elongation at which the Moon could be found if its orbit were circular and free of perturbations.
	@see EQ.47.2 p.338 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarMeanElongation(J float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var D = math.Mod((297.8501921 + 445267.1114034*J - 0.0018819*math.Pow(J, 2) + math.Pow(J, 3)/545868 - math.Pow(J, 4)/113065000), 360)

	// correct for negative angles
	if D < 0 {
		D += 360
	}

	return D
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
	@returns the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date (in degrees)
	@see p.144 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarLongitudeOfNutation(L float64, l float64, Ω float64) float64 {
	return (-17.20*sinx(Ω) + 1.32*sinx(2*L) - 0.23*sinx(2*l) + 0.21*sinx(2*Ω)) / 3600
}

/*
	GetLunarObliquityOfNutation()

	@param L - the ecliptic longitude at which the Sun could be found if its orbit were circular and free of perturbations.
	@param l - the ecliptic longitude at which the Moon could be found if its orbit were circular and free of perturbations.
	@param Ω - the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date
	@returns the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date (in degrees)
	@see p.144 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetLunarObliquityOfNutation(L float64, l float64, Ω float64) float64 {
	return (9.20*cosx(Ω) + 0.57*cosx(2*L) + 0.1*cosx(2*l) - 0.09*cosx(2*Ω)) / 3600
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
	GetMoonEclipticPosition()

	N.B. the ecliptic position is referenced to mean equinox of date and do not include the effect of nutation.

	@param datetime - the datetime in UTC of the observer
	@returns the geocentric ecliptic coodinate (λ - geocentric longitude, β - geocentric latidude and Δ distance between centers of the Earth and Moon, in km) of the Moon.
*/
func GetLunarEclipticPosition(datetime time.Time) EclipticCoordinate {
	var T float64 = GetCurrentJulianCenturyRelativeToJ2000(datetime)

	var D float64 = GetLunarMeanElongation(T)

	var Lʹ float64 = GetLunarMeanLongitude(T)

	var M float64 = GetSolarMeanAnomaly(T)

	var Mʹ float64 = GetLunarMeanAnomaly(T)

	var F float64 = GetLunarArgumentOfLatitude(T)

	var A1 float64 = 119.75 + 131.849*T

	var A2 float64 = 53.09 + 479264.29*T

	var A3 float64 = 313.45 + 481266.484*T

	var E float64 = 1 - 0.002516*T - 0.0000074*math.Pow(T, 2)

	var E2 float64 = math.Pow(E, 2)

	var Σl float64 = 3958*sinx(A1) + 1962*sinx(Lʹ-F) + 318*sinx(A2)

	var Σr float64 = 0

	var Σb float64 = -2235*sinx(Lʹ) + 382*sinx(A3) + 175*sinx(A1-F) + 175*sinx(A1+F) + 127*sinx(Lʹ-Mʹ) - 115*sinx(Lʹ+Mʹ)

	for i := range ta {
		r := &ta[i]

		sa, ca := sincosx(D*r.D + M*r.M + Mʹ*r.Mʹ + F*r.F)

		switch r.M {
		case 0:
			Σl += r.Σl * sa
			Σr += r.Σr * ca
		case 1, -1:
			Σl += r.Σl * sa * E
			Σr += r.Σr * ca * E
		case 2, -2:
			Σl += r.Σl * sa * E2
			Σr += r.Σr * ca * E2
		}
	}

	for i := range tb {
		r := &tb[i]

		sb := sinx(D*r.D + M*r.M + Mʹ*r.Mʹ + F*r.F)

		switch r.M {
		case 0:
			Σb += r.Σb * sb
		case 1, -1:
			Σb += r.Σb * sb * E
		case 2, -2:
			Σb += r.Σb * sb * E2
		}
	}

	return EclipticCoordinate{
		λ: Lʹ + Σl/1000000,
		β: Σb / 1000000,
		Δ: 385000.56 + Σr/1000,
	}
}

type tas struct{ D, M, Mʹ, F, Σl, Σr float64 }

var ta = [...]tas{
	{0, 0, 1, 0, 6288774, -20905355},
	{2, 0, -1, 0, 1274027, -3699111},
	{2, 0, 0, 0, 658314, -2955968},
	{0, 0, 2, 0, 213618, -569925},

	{0, 1, 0, 0, -185116, 48888},
	{0, 0, 0, 2, -114332, -3149},
	{2, 0, -2, 0, 58793, 246158},
	{2, -1, -1, 0, 57066, -152138},

	{2, 0, 1, 0, 53322, -170733},
	{2, -1, 0, 0, 45758, -204586},
	{0, 1, -1, 0, -40923, -129620},
	{1, 0, 0, 0, -34720, 108743},

	{0, 1, 1, 0, -30383, 104755},
	{2, 0, 0, -2, 15327, 10321},
	{0, 0, 1, 2, -12528, 0},
	{0, 0, 1, -2, 10980, 79661},

	{4, 0, -1, 0, 10675, -34782},
	{0, 0, 3, 0, 10034, -23210},
	{4, 0, -2, 0, 8548, -21636},
	{2, 1, -1, 0, -7888, 24208},

	{2, 1, 0, 0, -6766, 30824},
	{1, 0, -1, 0, -5163, -8379},
	{1, 1, 0, 0, 4987, -16675},
	{2, -1, 1, 0, 4036, -12831},

	{2, 0, 2, 0, 3994, -10445},
	{4, 0, 0, 0, 3861, -11650},
	{2, 0, -3, 0, 3665, 14403},
	{0, 1, -2, 0, -2689, -7003},

	{2, 0, -1, 2, -2602, 0},
	{2, -1, -2, 0, 2390, 10056},
	{1, 0, 1, 0, -2348, 6322},
	{2, -2, 0, 0, 2236, -9884},

	{0, 1, 2, 0, -2120, 5751},
	{0, 2, 0, 0, -2069, 0},
	{2, -2, -1, 0, 2048, -4950},
	{2, 0, 1, -2, -1773, 4130},

	{2, 0, 0, 2, -1595, 0},
	{4, -1, -1, 0, 1215, -3958},
	{0, 0, 2, 2, -1110, 0},
	{3, 0, -1, 0, -892, 3258},

	{2, 1, 1, 0, -810, 2616},
	{4, -1, -2, 0, 759, -1897},
	{0, 2, -1, 0, -713, -2117},
	{2, 2, -1, 0, -700, 2354},

	{2, 1, -2, 0, 691, 0},
	{2, -1, 0, -2, 596, 0},
	{4, 0, 1, 0, 549, -1423},
	{0, 0, 4, 0, 537, -1117},

	{4, -1, 0, 0, 520, -1571},
	{1, 0, -2, 0, -487, -1739},
	{2, 1, 0, -2, -399, 0},
	{0, 0, 2, -2, -381, -4421},

	{1, 1, 1, 0, 351, 0},
	{3, 0, -2, 0, -340, 0},
	{4, 0, -3, 0, 330, 0},
	{2, -1, 2, 0, 327, 0},

	{0, 2, 1, 0, -323, 1165},
	{1, 1, -1, 0, 299, 0},
	{2, 0, 3, 0, 294, 0},
	{2, 0, -1, -2, 0, 8752},
}

type tbs struct{ D, M, Mʹ, F, Σb float64 }

var tb = [...]tbs{
	{0, 0, 0, 1, 5128122},
	{0, 0, 1, 1, 280602},
	{0, 0, 1, -1, 277693},
	{2, 0, 0, -1, 173237},

	{2, 0, -1, 1, 55413},
	{2, 0, -1, -1, 46271},
	{2, 0, 0, 1, 32573},
	{0, 0, 2, 1, 17198},

	{2, 0, 1, -1, 9266},
	{0, 0, 2, -1, 8822},
	{2, -1, 0, -1, 8216},
	{2, 0, -2, -1, 4324},

	{2, 0, 1, 1, 4200},
	{2, 1, 0, -1, -3359},
	{2, -1, -1, 1, 2463},
	{2, -1, 0, 1, 2211},

	{2, -1, -1, -1, 2065},
	{0, 1, -1, -1, -1870},
	{4, 0, -1, -1, 1828},
	{0, 1, 0, 1, -1794},

	{0, 0, 0, 3, -1749},
	{0, 1, -1, 1, -1565},
	{1, 0, 0, 1, -1491},
	{0, 1, 1, 1, -1475},

	{0, 1, 1, -1, -1410},
	{0, 1, 0, -1, -1344},
	{1, 0, 0, -1, -1335},
	{0, 0, 3, 1, 1107},

	{4, 0, 0, -1, 1021},
	{4, 0, -1, 1, 833},

	{0, 0, 1, -3, 777},
	{4, 0, -2, 1, 671},
	{2, 0, 0, -3, 607},
	{2, 0, 2, -1, 596},

	{2, -1, 1, -1, 491},
	{2, 0, -2, 1, -451},
	{0, 0, 3, -1, 439},
	{2, 0, 2, 1, 422},

	{2, 0, -3, -1, 421},
	{2, 1, -1, 1, -366},
	{2, 1, 0, 1, -351},
	{4, 0, 0, 1, 331},

	{2, -1, 1, 1, 315},
	{2, -2, 0, -1, 302},
	{0, 0, 1, 3, -283},
	{2, 1, 1, -1, -229},

	{1, 1, 0, -1, 223},
	{1, 1, 0, 1, 223},
	{0, 1, -2, -1, -220},
	{2, 1, -1, -1, -220},

	{1, 0, 1, 1, -185},
	{2, -1, -2, -1, 181},
	{0, 1, 2, 1, -177},
	{4, 0, -2, -1, 176},

	{4, -1, -1, -1, 166},
	{1, 0, 1, -1, -164},
	{4, 0, 1, -1, 132},
	{1, 0, -1, -1, -119},

	{4, -1, 0, -1, 115},
	{2, -2, 0, 1, 107},
}

/*
	GetLunarHorizontalParallax()

	For the Moon, the problem of finding an accurate measurement of its standard altitude, h_0, is a little more
	complicated because h_0 is not constant over time. This equation takes into account the variations in
	 semidiamater and parallax.

	@param
	@returns the horizontal parallax of the Moon for a given distance measurement
*/
func GetLunarHorizontalParallax(Δ float64) float64 {
	return asinx(6378.14 / Δ)
}

/*
	GetLunarHourAngle()

	Observing the Moon from Earth, the lunar hour angle is an expression of time, expressed in angular measurement,
	usually degrees, from lunar noon. At lunar noon the hour angle is zero degrees, with the time before lunar noon
	expressed as negative degrees, and the local time after lunar noon expressed as positive degrees.

	@param δ - the ecliptic longitude of the Moon (in degrees)
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@param π - is the Lunar horizontal parallax
	@returns the lunar hour angle for a given lunar declination, of some observer on Earth
*/
func GetLunarHourAngle(δ float64, latitude float64, elevation float64, π float64) float64 {
	// observations on a sea horizon needing an elevation-of-observer correction
	// (corrects for both apparent dip and terrestrial refraction):
	var corr = -2.076 * math.Sqrt(elevation) * 1 / 60

	var h = 0.7275*π - 0.566667

	var H_0 = acosx((sinx(h-corr) - (sinx(latitude) * sinx(δ))) / cosx(latitude) * cosx(δ))

	return H_0
}

/*
	GetLunarTransitJulianDate()

	@param datetime - the datetime in UTC of the observer
	@param α - the right ascension position of the Moon (in degrees)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@param θ - the apparent sidereal time at Greenwhich for the desired datetime (in degrees)
	@returns the lunar transit time in Julian date format
	@see eq.15.2 p.102 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann - Bell.
*/
func GetLunarTransitJulianDate(datetime time.Time, α float64, longitude float64, θ float64) float64 {
	var d = time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, time.UTC)

	var J float64 = GetJulianDate(d)

	// correct for fractions of a day less than 0, and greater than 1.
	var m = (α + longitude - θ) / 360

	// correct for negative fractions of day less than 0.
	if m < 0 {
		m += 1
	}

	// correct for fractions of day greater than 1.
	if m > 1 {
		m -= 1
	}

	// add the days fraction to the Julian date at 0h:
	return J + m
}
