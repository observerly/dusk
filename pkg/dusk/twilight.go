package dusk

type SunriseStatus int

const (
	AboveHorizon = SunriseStatus(1)
	AtHorizon    = SunriseStatus(0)
	BelowHorizon = SunriseStatus(-1)
)
