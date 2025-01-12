package mapper

import (
	"tube-profile/internal/database"
	"tube-profile/internal/model"
)

func FromDBProfile(dbProfile database.Profile) model.Profile {
	return model.Profile{
		UserID: dbProfile.UserID,
		Gender: dbProfile.Gender,
		Age:    dbProfile.Age,
		Weight: dbProfile.Weight,
		Height: dbProfile.Height,
		Goal:   dbProfile.Goal,
	}
}
