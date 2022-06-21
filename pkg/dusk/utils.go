package dusk

import "math"

/*
	@brief axial tilt, also known as obliquity, is the angle between an object's rotational axis and its orbital axis, which is the line perpendicular to its orbital plane.
*/
var TERRA_AXIAL_TILT float64 = 23.4397

/*
	GetEarthObliquity()

	@returns the known earth obliquity in degrees
*/
func GetEarthObliquity() float64 {
	return TERRA_AXIAL_TILT
}

/*
	GetMeanObliquityOfTheEcliptic()

	@param J - the Ephemeris time or the number of centuries since J2000 epoch
	@returns the mean obliquity of the ecliptic (in degrees), as adopted by the Internal Astronomical Union (IAU)
	@see Astronomical Almanac for the year 1984 (Washington, D.C.; 1983) p.s26
*/
func GetMeanObliquityOfTheEcliptic(J float64) float64 {
	// correct for large angles (+ive or -ive), i.e., applies modulo correction to the angle, and ensures always positive:
	var ε float64 = math.Mod(23.4392917-0.0130041667*J-0.00000016667*math.Pow(J, 2)+0.0000005027778*math.Pow(J, 3), 360)

	// correct for negative angles
	if ε < 0 {
		ε += 360
	}

	return ε
}

/*
	GetNutationInLongitudeOfTheEcliptic()

	@param L - the ecliptic longitude at which the Sun could be found if its orbit were circular and free of perturbations.
	@param l - the ecliptic longitude at which the Moon could be found if its orbit were circular and free of perturbations.
	@param Ω - the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date
	@returns the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date (in degrees)
	@see p.144 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetNutationInLongitudeOfTheEcliptic(L float64, l float64, Ω float64) float64 {
	return (-17.20*sinx(Ω) + 1.32*sinx(2*L) - 0.23*sinx(2*l) + 0.21*sinx(2*Ω)) / 3600
}

/*
	GetNutationInObliquityOfTheEcliptic()

	@param L - the ecliptic longitude at which the Sun could be found if its orbit were circular and free of perturbations.
	@param l - the ecliptic longitude at which the Moon could be found if its orbit were circular and free of perturbations.
	@param Ω - the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date
	@returns the longitude of the ascending node of the Moon's mean orbit on the ecliptic, measured from the mean equinox of date (in degrees)
	@see p.144 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetNutationInObliquityOfTheEcliptic(L float64, l float64, Ω float64) float64 {
	return (9.20*cosx(Ω) + 0.57*cosx(2*L) + 0.1*cosx(2*l) - 0.09*cosx(2*Ω)) / 3600
}

/*
	GetArgumentOfLocalSiderealTimeForTransit()

	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param δ - equatorial declination of the observed object
	@returns the argument of local sidereal time (LST) observed in degrees
	@see ch.5 p.115 eq 5.2.1 & 5.2.2 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetArgumentOfLocalSiderealTimeForTransit(latitude float64, δ float64) float64 {
	return acosx(-tanx(latitude) * tanx(δ))
}

/*
	GetAtmosphericRefraction()

	@param altitude - is the altitude of the object in degrees
	@returns the atmospheric refraction in degrees for all angles from 0° - 90°
	@see p.106 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func GetAtmosphericRefraction(altitude float64) *float64 {
	if altitude < 0 {
		return nil
	}

	R := (1.02 / tanx(altitude+(10.3/(altitude+5.11)))) / 60

	return &R
}

/*
	GetRelativeAirMass()

	@param altitude - is the altitude of the object in degrees
	@returns the relative air mass, the ratio of absolute air masses (as defined above) at oblique incidence relative to that at zenith.
*/
func GetRelativeAirMass(altitude float64) *float64 {
	if altitude < 0 {
		return nil
	}

	X := 1 / sinx(altitude+(244/(165+47*math.Pow(altitude, 1.1))))

	return &X
}

/*
	GetApparentAltitude()

	@param altitude - is the altitude of the object in degrees
	@returns the apparent altitude in degrees
*/
func GetApparentAltitude(altitude float64) *float64 {
	if altitude < 0 {
		return nil
	}

	R := GetAtmosphericRefraction(altitude)

	app := altitude + *R

	return &app
}
