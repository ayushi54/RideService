package main

import (
	"fmt"

	"github.com/RideService/controllers"
	"github.com/RideService/models"
)

// Testing the RideService
func main() {
	rideService := controllers.InitRideService()
	user1 := models.User{ID: 1, Name: "user1"}
	rideService.RegisterUser(user1)

	user2 := models.User{ID: 4, Name: "user2"}
	rideService.RegisterUser(user2)

	driver1 := models.User{ID: 2, Name: "driver1"}
	driver2 := models.User{ID: 3, Name: "driver2"}
	rideService.RegisterUser(driver1)
	rideService.RegisterUser(driver2)
	rideService.RegisterCab(driver2.ID, models.SEDAN, models.Location{Lattitude: 500, Longitude: 500})
	src := models.Location{Lattitude: 67, Longitude: 100}
	dest := models.Location{Lattitude: 69, Longitude: 103}
	rideService.RegisterCab(driver1.ID, models.COMPACTSUV, dest)

	cabs, err := rideService.GetAvailableCab(src, dest, user1.ID)
	if err == nil {
		for cab := range cabs {
			// Selecting first available cab
			err := rideService.CreateBooking(src, dest, cab, cabs[cab], user1.ID)
			if err != nil {
				fmt.Println("Booking failed")
				return
			}
			fmt.Println("Booking Confirmed")
			// User finished the trip
			rideService.FinishedTrip(user1.ID, cab)
			break
		}
	} else {
		fmt.Println("Booking failed")
	}
}
