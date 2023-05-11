package databases

import "github.com/RideService/models"

var users []models.User

func RegisterUsers(user models.User) {
	users = append(users, user)
}
