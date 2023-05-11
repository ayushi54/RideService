package matchinalgos

import (
	"errors"

	"github.com/RideService/databases"
	"github.com/RideService/models"
)

func calculatePrice(dist int) int {
	// Prices should be calcualted based on the current, distance, demand and current traffic situation
	return dist * 4
}

func GetAvailableCab(source, destination models.Location) (map[*models.Cab]int, error) {
	retCabs := make(map[*models.Cab]int)
	cabs := databases.GetCabs()
	for _, cab := range cabs {
		if cab.IsAvailable {
			dist := cab.CurrentLocation.GetDistance(source)
			journeyDist := source.GetDistance(destination)
			if dist <= models.Max_ALLOWED_DISTANCE {
				price := calculatePrice(int(journeyDist))
				retCabs[cab] = price
			}
		}
	}
	if len(retCabs) == 0 {
		return nil, errors.New("Cabs not available for this destination")
	}
	return retCabs, nil
}
