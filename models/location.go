package models

import "math"

const (
	Max_ALLOWED_DISTANCE = 5
)

type Location struct {
	Lattitude int64
	Longitude int64
}

func (this *Location) GetDistance(l2 Location) float64 {
	x := (this.Lattitude - l2.Lattitude)
	y := this.Longitude - l2.Longitude
	distance := math.Sqrt(float64(x*x + y*y))
	return distance
}
