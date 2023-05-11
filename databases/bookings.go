package databases

import (
	"github.com/RideService/models"
)

var bookings []models.Booking

func AddBookings(booking models.Booking) {
	bookings = append(bookings, booking)
}
