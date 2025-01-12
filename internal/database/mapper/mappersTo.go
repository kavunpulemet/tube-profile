package mapper

import (
	"tube-profile/internal/database"
	"tube-profile/internal/model"
)

func ToDBProfile(serviceProfile model.Profile) database.Profile {
	return database.Profile{
		ID:     "",
		UserID: serviceProfile.UserID,
		Gender: serviceProfile.Gender,
		Age:    serviceProfile.Age,
		Weight: serviceProfile.Weight,
		Height: serviceProfile.Height,
		Goal:   serviceProfile.Goal,
	}
}
