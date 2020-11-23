// Package space provides a handy way to convert age from seconds on different planets.
package space

import "math"

type Planet string

// orbitInEarthYears maps each planet and its corresponding equivalent of one year
// on earth.
var orbitInEarthYears = map[string]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const earthYearInSecs float64 = 31557600

// struct space age holds values of
type spaceAge struct {
	planet    Planet
	sec       float64 // age in seconds
	age       float64
	planetSec float64 // the value of a complete orbit in seconds
}

func (s *spaceAge) toAge() {
	s.age = s.sec / s.planetSec
}

func (s spaceAge) getAge() float64 {
	s.toAge()
	return math.Round(s.age*100) / 100 //(round to nearest) 2 places.
}

func Age(secs float64, planet Planet) float64 {
	m := spaceAge{
		planet:    planet,
		sec:       secs,
		planetSec: orbitInEarthYears[string(planet)] * earthYearInSecs,
	}

	return m.getAge()
}
