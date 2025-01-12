package service

import (
	"tube-profile/internal/database"
	"tube-profile/internal/database/mapper"
	"tube-profile/internal/model"
	"tube-profile/internal/utils"
)

type ProfileService interface {
	CreateProfile(ctx utils.MyContext, profile model.Profile) error
	GetProfile(ctx utils.MyContext, userID string) (model.Profile, error)
	UpdateProfile(ctx utils.MyContext, profile model.Profile) error
}

type ImplProfileService struct {
	repo      database.ProfileRepository
	jwtSecret []byte
}

func NewProfileService(repo database.ProfileRepository, jwtSecret []byte) *ImplProfileService {
	return &ImplProfileService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *ImplProfileService) CreateProfile(ctx utils.MyContext, profile model.Profile) error {
	err := s.repo.Create(ctx, mapper.ToDBProfile(profile))
	if err != nil {
		return err
	}

	return nil
}

func (s *ImplProfileService) GetProfile(ctx utils.MyContext, userID string) (model.Profile, error) {
	profile, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return model.Profile{}, err
	}

	return mapper.FromDBProfile(profile), nil
}

func (s *ImplProfileService) UpdateProfile(ctx utils.MyContext, profile model.Profile) error {
	err := s.repo.Update(ctx, mapper.ToDBProfile(profile))
	if err != nil {
		return err
	}

	return nil
}
