![@observerly/dusk](https://user-images.githubusercontent.com/84131395/172596337-7499f919-3e41-48ea-a561-b88afa75b8c9.png)

![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/observerly/dusk/main?filename=go.mod&label=Go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/observerly/dusk)](https://pkg.go.dev/github.com/observerly/dusk)
[![Go Report Card](https://goreportcard.com/badge/github.com/observerly/dusk)](https://goreportcard.com/report/github.com/observerly/dusk)
[![Dusk Actions Status](https://github.com/observerly/dusk/actions/workflows/ci.yml/badge.svg)](https://github.com/observerly/celestia/actions/workflows/ci.yml)

Dusk 🌑 is a minimal dependency Go library for calculating the most opinimum time to observe various astronomical objects by utilising paramaters such as astronomical twilight, the lunar phase and the rise and set times of the moon and sun.

## Installation

Make sure you have Go installed ([download](https://golang.org/dl/)). Version `1.17` or higher is required for this package.

Initialize your project by creating a folder and then running `go mod init github.com/your/repo` ([learn more](https://blog.golang.org/using-go-modules)) inside the folder. Then install Dusk with the [`go get`](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
go get -u github.com/observerly/dusk
```

## Usage

### Get Twilight

The basic usage of this package is to use the `GetLocalTwilight()` func, this provides the Sun rise and Sun set times (in datetime format, as local time), it also provides the "duration" between these two datetimes. The local time is calculated from a UTC date for the speicif latitude and longitude coordiantes provided.

```go
package main

import (
  "time"

  "github.com/observerly/dusk/pkg/dusk"
)

func main() {
  // datetime of observation:
  datetime := time.Date(2022, 2, 17, 0, 0, 0, 0, time.UTC)

  // observer's longitude, in degrees (*west of the Greenwhich meridian is negative, east is positive):
  longitude := -155.8246

  // observer's latitude, in degrees  (*south of the equator is negative, north is positive):
  latitude := 20.0046

  // observaer's elevation above mean sea level, in meteres:
  elevation := 4207.0

  // specify the twilight to be defined as a set number of degrees *below* the horizon (e.g, civil twilight is designated as being 6 degrees below horizon):
  degreesBelowHorizon := -6.0

  twilight, location, err := dusk.GetLocalTwilight(datetime, longitude, latitude, elevation, degreesBelowHorizon)
}
```

There are three wrapper functions which allow for an easy calculation of civil, nautical and astronomical twilight.

### Get Civil Twilight

For civil twilight, the degreesBelowHorizon for the Sun needs to be -6°.

```go
twilight, location, err := dusk.GetLocalCivilTwilight(datetime, longitude, latitude, elevation)
```

### Get Nautical Twilight

For nautical twilight, the degreesBelowHorizon for the Sun needs to be -12°.

```go
twilight, location, err := dusk.GetLocalNauticalTwilight(datetime, longitude, latitude, elevation)
```

### Get Astronomical Twilight

For astronomical twilight, the degreesBelowHorizon for the Sun needs to be -18°.

```go
twilight, location, err := dusk.GetLocalAstronomicalTwilight(datetime, longitude, latitude, elevation)
```

### Get Moon Position

To calculate the rise and set of the moon, it is neccessary to calculate the equatorial position of the moon at zero HH:mm:ss, e.g., midnight, for the +/-1 day for the day you want to calculate for, e.g., d-1, d and d+1. 

This library supplies the following function to calculate the equatorial position of the moon (in degrees):

```go
package main

import (
  "fmt"
  "time"

  "github.com/observerly/dusk/pkg/dusk"
)

func main() {
  // datetime of observation:
  datetime := time.Date(2022, 2, 17, 14, 55, 0, 0, time.UTC)

  eq := dusk.GetLunarEquatorialPosition(datetime)
	
  fmt.Printf("The Moon is at the following equatorial coordinate:\n")
  fmt.Printf("Right Ascension: %e°\n", eq.ra)
  fmt.Printf("Declination: %e°\n", eq.dec)
}
```

### Get Moon Phase

To calculate the moon phase, it is neccessary to calculate the ecliptic position of the moon at the datetime required, as well as the knowing some longitude of an observer.

This library supplies the following function to calculate the phase of the moon:

```go
package main

import (
  "fmt"
  "time"

  "github.com/observerly/dusk/pkg/dusk"
)

func main() {
  // datetime of observation:
  datetime := time.Date(2022, 2, 17, 14, 55, 0, 0, time.UTC)

  // some longitude, in degrees (*west of the Greenwhich meridian is negative, east is positive):
  lonmgitude := -155.8246

  // get the ecliptic coordinate of the Moon for the datetime:
  ec := dusk.GetLunarEclipticPosition(datetime)

  // calculate the phase for the datetime, longitude and ecliptic coordinate:
  phase := dusk.GetMoonPhase(datetime, longitude, ec)
}
```

**N.B.** The equatorial coordinate system is a celestial coordinate system widely used to specify the positions of celestial objects. It may be implemented in spherical or rectangular coordinates, both defined by an origin at the centre of Earth, a fundamental plane consisting of the projection of Earth's equator onto the celestial sphere (forming the celestial equator), a primary direction towards the vernal equinox, and a right-handed convention.

## License

Dusk is free software licensed under the GNU General Public License v3.0 (GPL-3.0). See [LICENSE](./LICENSE).

The binary version of this program uses several open source libraries and components, which come with their own licensing terms. See below for an overview, and [LICENSE](./LICENSE) for details.

| Library attribution                                                   | License type |
|-----------------------------------------------------------------------|--------------|
| [zsefvlol/timezonemapper](https://github.com/zsefvlol/timezonemapper) | MIT License  |
