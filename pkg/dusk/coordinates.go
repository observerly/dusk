package dusk

import (
	"time"
)

type Coordinate struct {
	/*
		ϕ - the latitude in degrees, e.g., altitude, latitude, declination
	*/
	Latitude float64
	/*
		θ - the longitude in degrees, e.g., azimuth, right ascension, longitude
	*/
	Longitude float64
}

type EquatorialCoordinate struct {
	/*
		Right Ascension - the right ascension in degrees
	*/
	RightAscension float64
	/*
		Declination - the declination in degrees
	*/
	Declination float64
}

type EclipticCoordinate struct {
	/*
		Longitude - the longitude in degrees
	*/
	Longitude float64
	/*
		Latitude - the latitude in degrees
	*/
	Latitude float64
	/*
		Distance - the distance in km
	*/
	Δ float64
}

type HorizontalCoordinate struct {
	/*
		altitude (a) or elevation
	*/
	Altitude float64
	/*
		azimuth (A) or elevation
	*/
	Azimuth float64
}

type TemporalHorizontalCoordinate struct {
	/*
		datetime of horizontal observation
	*/
	Datetime time.Time
	/*
		altitude (a) or elevation
	*/
	Altitude float64
	/*
		azimuth (A) or elevation
	*/
	Azimuth float64
}

type TransitHorizontalCoordinate struct {
	/*
		datetime of horizontal observation
	*/
	Datetime time.Time
	/*
		altitude (a) or elevation
	*/
	Altitude float64
	/*
		azimuth (A) or elevation
	*/
	Azimuth float64
	/*
		Is this particular a Moon rise?
	*/
	IsRise bool
	/*
		Is this particular a Moon set?
	*/
	IsSet bool
}

/*
	ConvertEclipticCoordinateToEquatorial()

	@param datetime - the datetime of the observer (in UTC)
	@param geocentric ecliptic coordinate of type EclipticCoordinate { λ, β, Λ }
	@returns the converted equatorial coordinate { ra, dec }
	@see eq13.3 & eq13.4 p.93 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func ConvertEclipticCoordinateToEquatorial(datetime time.Time, ec EclipticCoordinate) EquatorialCoordinate {
	var J = GetCurrentJulianCenturyRelativeToJ2000(datetime)

	var L float64 = GetSolarMeanLongitude(J)

	var l float64 = GetLunarMeanLongitude(J)

	var Ω float64 = GetLunarLongitudeOfTheAscendingNode(J)

	var ε float64 = GetMeanObliquityOfTheEcliptic(J) + GetNutationInObliquityOfTheEcliptic(L, l, Ω)

	var λ = ec.Longitude

	var β = ec.Latitude

	var α = atan2yx(sinx(λ)*cosx(ε)-tanx(β)*sinx(ε), cosx(λ))

	var δ = asinx(sinx(β)*cosx(ε) + cosx(β)*sinx(ε)*sinx(λ))

	return EquatorialCoordinate{
		RightAscension: α,
		Declination:    δ,
	}
}

/*
	ConvertEquatorialCoordinateToHorizontal()

	@param datetime - the datetime of the observer (in UTC)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param equatorial coordinate of type EquatorialCoordiate { ra, dec }
	@returns the equivalent horizontal coordinate for the given observers position
	@see eq13.5 and eq.6 p.93 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann-Bell.
*/
func ConvertEquatorialCoordinateToHorizontal(datetime time.Time, longitude float64, latitude float64, eq EquatorialCoordinate) HorizontalCoordinate {
	var LST float64 = GetLocalSiderealTime(datetime, longitude)

	var ra float64 = GetHourAngle(eq.RightAscension, LST)

	var dec float64 = eq.Declination

	var alt = asinx(sinx(dec)*sinx(latitude) + cosx(dec)*cosx(latitude)*cosx(ra))

	var az = acosx((sinx(dec) - sinx(alt)*sinx(latitude)) / (cosx(alt) * cosx(latitude)))

	return HorizontalCoordinate{
		Altitude: alt,
		Azimuth:  az,
	}
}
