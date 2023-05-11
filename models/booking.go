package models

const (
	PENDING   = 1
	CONFIRMED = 2
	STARTED   = 3
	Canceled  = 4
	COMPLETED = 5
)

type Booking struct {
	source Location
	dest   Location
	userID int64
	cab    *Cab
	price  int
	status uint8
}

func (this *Booking) CreateBooking(source, dest Location, userID int64, cab *Cab, price int) *Booking {
	booking := &Booking{
		source: source,
		dest:   dest,
		userID: userID,
		cab:    cab,
		price:  price,
		status: CONFIRMED,
	}
	cab.IsAvailable = false
	return booking
}

func (this *Booking) FinishedBooking(cab *Cab) {
	cab.IsAvailable = true
}
