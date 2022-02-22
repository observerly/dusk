package dusk

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
