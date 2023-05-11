package controllers

import (
	"github.com/RideService/databases"
	"github.com/RideService/models"
)

func (this *rideService) RegisterCab(userID int64, CabType uint8, loc models.Location) {
	cab := models.Cab{DriverID: userID, CabType: CabType, CurrentLocation: loc, IsAvailable: true}
	databases.RegisterCabs(cab)
}
