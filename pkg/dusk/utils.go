package dusk

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
