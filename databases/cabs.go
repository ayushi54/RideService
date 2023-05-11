package databases

import "github.com/RideService/models"

var cabs []*models.Cab

func RegisterCabs(cab models.Cab) {
	cabs = append(cabs, &cab)
}

func GetCabs() []*models.Cab {
	return cabs
}
