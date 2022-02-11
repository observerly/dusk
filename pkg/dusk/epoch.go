package dusk

import "time"

/*
	GetJulianDate()

	@returns the Julian date i.e., the continuous count of days and fractions of day since the beginning of the Julian period
	@see http://astro.vaporia.com/start/jd.html
*/
func GetJulianDate(datetime time.Time) float64 {
	// the epoch of Unix time start i.e., 1 January 1970 00:00:00 UTC:
	var J1970 float64 = 2440587.5

	// milliseconds elapsed since 1 January 1970 00:00:00 UTC up until now as an int64:
	var time int64 = datetime.UTC().UnixNano() / 1e6

	return float64(time)/86400000.0 + J1970
}
