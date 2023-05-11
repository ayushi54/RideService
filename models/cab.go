package models

const (
	SEDAN       = 1
	STANDARDSUV = 2
	COMPACTSUV  = 3
)

type Cab struct {
	DriverID        int64
	CabType         uint8
	CurrentLocation Location
	IsAvailable     bool
}

func (this *Cab) getLocation() Location {
	return this.CurrentLocation
}

func (this *Cab) setLocation(loc Location) {
	this.CurrentLocation = loc
}
