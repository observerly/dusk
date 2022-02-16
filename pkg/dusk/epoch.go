package dusk

import (
	"math"
	"time"
)

// the epoch of Unix time start i.e., 1 January 1970 00:00:00 UTC:
var J1970 float64 = 2440587.5

// the epoch of Unix time start i.e., 1 January 2000 00:00:00 UTC:
var J2000 float64 = 2451545.0

/*
	GetJulianDate()

	@returns the Julian date i.e., the continuous count of days and fractions of day since the beginning of the Julian period
	@see http://astro.vaporia.com/start/jd.html
*/
func GetJulianDate(datetime time.Time) float64 {
	// milliseconds elapsed since 1 January 1970 00:00:00 UTC up until now as an int64:
	var time int64 = datetime.UTC().UnixNano() / 1e6

	return float64(time)/86400000.0 + J1970
}

/*
	GetUniversalTime()

	@returns the universal time (UTC) for a given Julian date
*/
func GetUniversalTime(JD float64) time.Time {
	return time.Unix(0, int64((JD-J1970)*86400000.0*1e6)).UTC()
}

/*
	GetCurrentJulianDayRelativeToJ2000()

	@returns the number of Julian days between J2000 (i.e., 1 January 2000 00:00:00 UTC) and the the datetime, rounded up the the nearest integer
	@see http://astro.vaporia.com/start/jd.html
*/
func GetCurrentJulianDayRelativeToJ2000(datetime time.Time) int {
	// get the Julian date:
	var JD float64 = GetJulianDate(datetime)

	// correction for the the fractional Julian Day for leap seconds and terrestrial time (TT):
	var corr float64 = 0.0008

	// calculate the current Julian day:
	var n float64 = math.Ceil(JD - 2451545.0 - corr)

	return int(n)
}

/*
	GetCurrentJulianCenturyRelativeToJ2000()

	@returns the number of Julian centuries between J2000 (i.e., 1 January 2000 00:00:00 UTC) and the the datetime, rounded up the the nearest integer
	@see http://astro.vaporia.com/start/jd.html
*/
func GetCurrentJulianCenturyRelativeToJ2000(datetime time.Time) float64 {
	// get the Julian date:
	var JD float64 = GetJulianDate(datetime)

	// calculate the current Julian century as fractions of centuries:
	var n float64 = (JD - 2451545.0) / 36525

	return n
}

/*
	GetMeanSolarTime()

	@param datetime - the datetime of the observer (in UTC)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@returns returns the mean solar time, relative to some observer's longitude on Earth
*/
func GetMeanSolarTime(datetime time.Time, longitude float64) float64 {
	// the number of Julian days between J2000 (i.e., 1 January 2000 00:00:00 UTC) and the the datetime:
	var n int = GetCurrentJulianDayRelativeToJ2000(datetime)

	return float64(n) - (longitude / 360)
}
