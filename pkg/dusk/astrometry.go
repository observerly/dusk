package dusk

import "math"

/*
	GetHourAngle()

	Gets the hour angle for a particular object for a particular observer at a given datetime

	@param ra - the right ascension of type number of the observed object's equatorial coordinate (in degrees)
	@param LST - the local sidereal time of type number is defined as the hour angle of the vernal equinox (in degrees)
	@returns the calculated hour angle (in degrees)
*/
func GetHourAngle(ra float64, LST float64) float64 {
	var ha = LST*15 - ra

	// If the hour angle is less than zero, ensure we rotate by 2π radians (360 degrees)
	if ha < 0 {
		ha += 360
	}

	return ha
}

/*
	GetAngularSeparation()

	Gets the angular separation between two objects of known coordinates

	@param c1 - the first object's coordinate of type Coordinate { Latitude, Longitude }
	@param c2 - the second object's coordinate of type Coordinate { Latitude, Longitude }
	@returns the angular separation (in degrees)
*/
func GetAngularSeparation(coord1 Coordinate, coord2 Coordinate) float64 {
	var x = cosx(coord1.Latitude)*sinx(coord2.Latitude) - sinx(coord1.Latitude)*cosx(coord2.Latitude)*cosx(coord2.Longitude-coord1.Longitude)

	var y = cosx(coord2.Latitude) * sinx(coord2.Longitude-coord1.Longitude)

	var z float64 = sinx(coord1.Latitude)*sinx(coord2.Latitude) + cosx(coord1.Latitude)*cosx(coord2.Latitude)*cosx(coord2.Longitude-coord1.Longitude)

	return atan2yx(math.Sqrt(x*x+y*y), z)
}
