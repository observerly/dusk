package dusk

import (
	"math"
	"time"
)

type Sun struct {
	rise time.Time
	noon time.Time
	set  time.Time
}

/*
	GetSolarMeanAnomaly()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@returns the non-uniform or anomalous apparent motion of the Sun along the plane of the ecliptic
	@see EQ.47.3 p.338 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell
*/
func GetSolarMeanAnomaly(J float64) float64 {
	// applies modulo correction to the angle, and ensures always positive:
	return math.Mod(357.5291092+(0.98560028*J), 360) + 360
}

/*
	GetSolarEquationOfCenter()

	@param M - the mean solar anomaly for the Ephemeris time or the number of centuries since J2000 epoch
	@returns the equation of center for the Sun
	@see p.164 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetSolarEquationOfCenter(M float64) float64 {
	// applies modulo correction to the angle, and ensures always positive:
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

/*
	GetSolarTransitJulianDate()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@param M - the mean solar anomaly for the Ephemeris time or the number of centuries since J2000 epoch
	@param λ - the ecliptic longitude of the Sun (in degrees)
	@returns the Julian date for the local true solar transit (or solar noon).
*/
func GetSolarTransitJulianDate(J float64, M float64, λ float64) float64 {
	return 2451545.0 + J + 0.0053*sinx(M) - 0.0069*sinx(2*λ)
}

/*
	GetSolarDeclination()

	The declination of the Sun, δ☉, is the angle between the rays of the Sun and the plane of the Earth's equator.

	@param λ - the ecliptic longitude of the Sun (in degrees)
	@returns the declination of the Sun (in degrees)
	@see https://gml.noaa.gov/grad/solcalc/glossary.html#solardeclination
*/
func GetSolarDeclination(λ float64) float64 {
	return asinx(sinx(λ) * sinx(23.44))
}

/*
	GetSolarMeanLongitude()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@returns the mean longitude of the Sun
*/
func GetSolarMeanLongitude(J float64) float64 {
	// applies modulo correction to the angle, and ensures always positive:
	var L = math.Mod(280.4665+36000.7698*J, 360)

	// correct for negative angles
	if L < 0 {
		L += 360
	}

	return L
}

/*
	GetSolarHourAngle()

	Observing the Sun from Earth, the solar hour angle is an expression of time, expressed in angular measurement,
	usually degrees, from solar noon. At solar noon the hour angle is zero degrees, with the time before solar noon
	expressed as negative degrees, and the local time after solar noon expressed as positive degrees.

	@param δ - the ecliptic longitude of the Sun (in degrees)
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@returns the solar hour angle for a given solar declination, of some observer on Earth
	@see https://gml.noaa.gov/grad/solcalc/glossary.html#solardeclination
*/
func GetSolarHourAngle(δ float64, degreesBelowHorizon float64, latitude float64, elevation float64) float64 {
	// observations on a sea horizon needing an elevation-of-observer correction
	// (corrects for both apparent dip and terrestrial refraction):
	var corr = -degreesBelowHorizon + -2.076*math.Sqrt(elevation)*1/60

	return acosx((sinx(-0.83-corr) - (sinx(latitude) * sinx(δ))) / cosx(latitude) * cosx(δ))
}

/*
	GetSunriseSunsetTimesInUTC()

	@param datetime - the datetime of the observer (in UTC)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@returns the rise, noon and set for the Sun, in UTC (*not local time)
*/
func GetSunriseSunsetTimesInUTC(datetime time.Time, degreesBelowHorizon float64, longitude float64, latitude float64, elevation float64) Sun {
	var J float64 = GetMeanSolarTime(datetime, longitude)

	var M float64 = GetSolarMeanAnomaly(J)

	var C float64 = GetSolarEquationOfCenter(M)

	var λ float64 = GetSolarEclipticLongitude(M, C)

	var δ float64 = GetSolarDeclination(λ)

	var ω float64 = GetSolarHourAngle(δ, degreesBelowHorizon, latitude, elevation)

	var h float64 = ω / 360

	var J_transit float64 = GetSolarTransitJulianDate(J, M, λ)

	var J_rise = J_transit - h

	var J_set = J_transit + h

	sun := Sun{
		rise: GetUniversalTime(J_rise),
		noon: GetUniversalTime(J_transit),
		set:  GetUniversalTime(J_set),
	}

	return sun
}

/*
	GetSolarEclipticPosition()

	@param datetime - the datetime of the observer (in UTC)
	@returns the geocentric ecliptic coodinate (λ - geocentric longitude, β - geocentric latidude) of the Sun.
*/
func GetSolarEclipticPosition(datetime time.Time) EclipticCoordinate {
	var M = GetSolarMeanAnomalyLawrence(datetime)

	var C = GetSolarEquationOfCenterLawrence(M)

	var λ = GetSolarEclipticLongitudeLawrence(M, C)

	return EclipticCoordinate{
		Longitude: λ,
		Latitude:  0,
	}
}
