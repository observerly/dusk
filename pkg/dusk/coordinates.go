package dusk

import (
	"time"
)

type EquatorialCoordinate struct {
	ra  float64
	dec float64
}

type EclipticCoordinate struct {
	λ float64
	β float64
	Δ float64
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

	var λ = ec.λ

	var β = ec.β

	var α = atanx((sinx(λ)*cosx(ε) - tanx(β)*sinx(ε)) / cosx(λ))

	var δ = asinx(sinx(β)*cosx(ε) + cosx(β)*sinx(ε)*sinx(λ))

	if α < 0 {
		α += 180
	}

	return EquatorialCoordinate{
		ra:  α,
		dec: δ,
	}
}
