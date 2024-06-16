package viewers

import (
	"fmt"
	"main/dbl/locations/types"
	"time"
)

type LocationDBL interface {
	GetLocationsByTimeRange(time.Time, time.Time) ([]*types.Location, bool, error)
	// GetLocationByID(int) ([]types.Location, bool, error)
	// SetLocation() error
	// UpdateLocation() error
	// DeleteLocation() error
}

type locationDBL struct {
	connectionPool interface{}
}

// NewLocationDBL revise this to take in a pointer to a db connection pool
func NewLocationDBL() *locationDBL {
	l := &locationDBL{
		connectionPool: "placeholder",
	}
	return l
}

func (l *locationDBL) GetLocationsByTimeRange(start time.Time, stop time.Time) ([]*types.Location, bool, error) {
	fmt.Println("logging times", start, stop)
	return nil, true, nil
}
