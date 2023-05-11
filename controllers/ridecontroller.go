package controllers

import (
	"github.com/RideService/databases"
	"github.com/RideService/matchinalgos"
	"github.com/RideService/models"
)

type rideService struct {
	userBookings map[int64]*models.Booking
}

func InitRideService() *rideService {
	obj := &rideService{}
	obj.userBookings = make(map[int64]*models.Booking)
	return obj
}

func (this *rideService) RegisterUser(user models.User) bool {
	databases.RegisterUsers(user)
	return true
}

func (this *rideService) GetAvailableCab(source, dest models.Location, userID int64) (map[*models.Cab]int, error) {
	cabs, err := matchinalgos.GetAvailableCab(source, dest)
	if err != nil {
		return nil, err
	}
	booking := &models.Booking{}
	this.userBookings[userID] = booking
	return cabs, nil
}

func (this *rideService) CreateBooking(source, dest models.Location, cab *models.Cab, price int, userID int64) error {
	this.userBookings[userID].CreateBooking(source, dest, userID, cab, price)
	databases.AddBookings(*this.userBookings[userID])
	return nil
}

func (this *rideService) FinishedTrip(userID int64, cab *models.Cab) {
	this.userBookings[userID].FinishedBooking(cab)
}
