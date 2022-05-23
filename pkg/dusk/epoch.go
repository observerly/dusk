package dusk

import (
	"math"
	"time"
)

// the epoch of Unix time start i.e., 1 January 1970 00:00:00 UTC:
var J1970 float64 = 2440587.5

// the epoch of Unix time start i.e., 1 January 2000 00:00:00 UTC:
var J2000 float64 = 2451545.0

type JulianPeriod struct {
	/*
		The current Julian Date expressed as fractions of days
	*/
	JD float64
	/*
		The current Julian Date expressed as fractions of centuries
	*/
	T float64
}

/*
	GetDatetimeZeroHour()

	@param datetime - the datetime of the observer
	@returns the datetime of the zero hour of the day
*/
func GetDatetimeZeroHour(datetime time.Time) time.Time {
	return time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, time.UTC)
}

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
	GetLocalGreenwhichSiderealTime

	@param datetime - the datetime of the observer
	@returns the local sidereal time relative to Greenwhich, UK
*/
func GetGreenwhichSiderealTime(datetime time.Time) float64 {
	JD := GetJulianDate(datetime)

	JD0 := GetJulianDate(time.Date(datetime.Year(), 1, 0, 0, 0, 0, 0, time.UTC))

	days := math.Floor(JD - JD0)

	var T = (JD0 - 2415020.0) / 36525

	var R = 6.6460656 + 2400.051262*T + 0.00002581*math.Pow(T, 2)

	var B = 24.0 - R + float64(24*(datetime.Year()-1900))

	var T0 = 0.0657098*days - B

	var hr = float64(datetime.Hour())

	var min = float64(datetime.Minute()) / 60.0

	var sec = float64(datetime.Second()) / 3600.0

	var ns = float64(datetime.Nanosecond()) / 3600000000.0

	var UT float64 = hr + min + sec + ns

	var A = float64(UT) * 1.002737909

	var GST = math.Mod(T0+A, 24)

	// correct for negative hour angles (24 hours is equivalent to 360°)
	if GST < 0 {
		GST += 24
	}

	return GST
}

/*
	GetLocalSiderealTime()

	@param datetime
	@returns returns the local sidereal time, relative to some location's longitude
*/
func GetLocalSiderealTime(datetime time.Time, longitude float64) float64 {
	var GST = GetGreenwhichSiderealTime(datetime)

	var d = (GST + longitude/15.0) / 24.0

	d = d - math.Floor(d)

	// correct for negative hour angles (24 hours is equivalent to 360°)
	if d < 0 {
		d += 1
	}

	return 24.0 * d
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
	GetFractionalJulianDayStandardEpoch()

	@returns the total number of fractional dates elapsed since the standard epoch J2000.
	@see p.136 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction Yo Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetFractionalJulianDaysSinceStandardEpoch(datetime time.Time) float64 {
	// get the Julian date:
	var JD float64 = GetJulianDate(datetime)

	// calculate the current Julian day:
	var n float64 = JD - 2451545.0

	return n
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
	GetCurrentJulianPeriod()

	@returns both the Julian date i.e., the continuous count of days and fractions of day since the beginning of the Julian period and the number of
	Julian centuries between J2000 (i.e., 1 January 2000 00:00:00 UTC) and the the datetime, rounded up the the nearest integer
	@see http://astro.vaporia.com/start/jd.html
*/
func GetCurrentJulianPeriod(datetime time.Time) JulianPeriod {
	// get the Julian date:
	var JD float64 = GetJulianDate(datetime)

	// calculate the current Julian date as fractions of centuries:
	var T float64 = (JD - 2451545.0) / 36525

	return JulianPeriod{
		JD: JD,
		T:  T,
	}
}

/*
	GetMeanGreenwhichSiderealTimeInDegrees()

	@returns the mean sidereal time at Greenwhich for the desired datetime (in degrees)
	@see eq.12.4 p.88 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann - Bell.
*/
func GetMeanGreenwhichSiderealTimeInDegrees(datetime time.Time) float64 {
	var d time.Time = time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, time.UTC)

	var julianPeriod JulianPeriod = GetCurrentJulianPeriod(d)

	// get the Julian date:
	var JD float64 = julianPeriod.JD

	// the number of Julian centuries between J2000 (i.e., 1 January 2000 00:00:00 UTC) and the the datetime:
	var T float64 = julianPeriod.T

	// applies modulo correction to the angle, and ensures always positive:
	var θ = math.Mod(280.46061837+(360.98564736629*(JD-2451545.0))+(0.000387933*math.Pow(T, 2))-(math.Pow(T, 3)/38710000), 360)

	// correct for negative angles
	if θ < 0 {
		θ += 360
	}

	return θ
}

/*
	GetApparentGreenwhichSiderealTimeInDegrees()

	@returns the apparent sidereal time at Greenwhich for the desired datetime (in degrees)
	@see eq.12.4 p.88 of Meeus, Jean. 1991. Astronomical algorithms. Richmond, Va: Willmann - Bell.
*/
func GetApparentGreenwhichSiderealTimeInDegrees(datetime time.Time) float64 {
	var θ float64 = GetMeanGreenwhichSiderealTimeInDegrees(datetime)

	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(datetime)

	var L float64 = GetSolarMeanLongitude(J)

	var l float64 = GetLunarMeanLongitude(J)

	var Ω float64 = GetLunarLongitudeOfTheAscendingNode(J)

	var ε float64 = GetMeanObliquityOfTheEcliptic(J) + GetNutationInObliquityOfTheEcliptic(L, l, Ω)

	var Δψ = GetNutationInLongitudeOfTheEcliptic(L, l, Ω)

	// applies a correction for the true vernal equinox:
	var corr = Δψ * cosx(ε)

	// applies modulo correction to the angle, and ensures always positive:
	var ϑ = math.Mod(θ+corr, 360)

	// correct for negative angles
	if ϑ < 0 {
		ϑ += 360
	}

	return ϑ
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

/*
	ConvertLocalSiderealTimeToGreenwhichSiderealTime()

	@param LST - the local sidereal time in hours (in decimal format)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@returns returns the GST in hours (in decimal format)
*/
func ConvertLocalSiderealTimeToGreenwhichSiderealTime(LST float64, longitude float64) float64 {
	var GST = LST - longitude/15

	// correct for negative hour angles
	if GST < 0 {
		GST += 24
	}

	return GST
}
