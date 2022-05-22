package dusk

import (
	"math"
)

/*
  GetDoesObjectRiseOrSet()

  @returns a boolean which determines if the object's EquatorialCoordinate{} in question rises or sets for the given Observer's latitude
  @see p.117 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetDoesObjectRiseOrSet(eq EquatorialCoordinate, latitude float64) bool {
	// If |Ar| > 1, the object never rises above the horizon
	Ar := sinx(eq.Declination) / cosx(latitude)

	// If |H1| > 1, the object is always below the horizon
	H1 := tanx(latitude) * tanx(eq.Declination)

	return math.Abs(Ar) < 1 && math.Abs(H1) < 1
}
