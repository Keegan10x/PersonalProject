package viewers

import "time"

type LocationDBL interface {
	GetLocationsByTimeRange(time.Time, time.Time) ([]Location, bool, error)
	GetLocationByID(int) (Location, bool, error)
	SetLocation() error
	UpdateLocation() error
	DeleteLocation() error
}

type Location struct {
	id       int
	provider string
	lat      float64
	long     float64
}
