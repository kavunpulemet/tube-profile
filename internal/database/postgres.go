package database

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"tube-profile/internal/utils"
)

type ProfileRepository interface {
	Create(ctx utils.MyContext, profile Profile) error
	GetByUserID(ctx utils.MyContext, userID string) (Profile, error)
	Update(ctx utils.MyContext, input Profile) error
}

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

//go:embed sql/CreateProfile.sql
var createProfile string

func (p *ProfilePostgres) Create(ctx utils.MyContext, profile Profile) error {
	profile.ID = uuid.New().String()

	log.Println("Repository userID", profile.UserID)

	_, err := p.db.ExecContext(ctx.Ctx, createProfile, profile.ID, profile.UserID, profile.Gender, profile.Age, profile.Weight, profile.Height, profile.Goal)
	if err != nil {
		return fmt.Errorf("failed to insert profile: %w", err)
	}

	return nil
}

//go:embed sql/GetProfile.sql
var getProfile string

func (p *ProfilePostgres) GetByUserID(ctx utils.MyContext, userID string) (Profile, error) {
	var profile Profile

	err := p.db.GetContext(ctx.Ctx, &profile, getProfile, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Profile{}, fmt.Errorf("profile not found for user_id: %s", userID)
		}
		return Profile{}, err
	}

	return profile, nil
}

func (p *ProfilePostgres) Update(ctx utils.MyContext, input Profile) error {
	var (
		queryBuilder strings.Builder
		args         []interface{}
		argIndex     int
	)

	queryBuilder.WriteString("UPDATE profiles SET ")

	if input.Gender != "" {
		argIndex++
		queryBuilder.WriteString(fmt.Sprintf("gender = $%d, ", argIndex))
		args = append(args, input.Gender)
	}
	if input.Age != 0 {
		argIndex++
		queryBuilder.WriteString(fmt.Sprintf("age = $%d, ", argIndex))
		args = append(args, input.Age)
	}
	if input.Weight != 0 {
		argIndex++
		queryBuilder.WriteString(fmt.Sprintf("weight = $%d, ", argIndex))
		args = append(args, input.Weight)
	}
	if input.Height != 0 {
		argIndex++
		queryBuilder.WriteString(fmt.Sprintf("height = $%d, ", argIndex))
		args = append(args, input.Height)
	}
	if input.Goal != "" {
		argIndex++
		queryBuilder.WriteString(fmt.Sprintf("goal = $%d, ", argIndex))
		args = append(args, input.Goal)
	}

	queryStr := queryBuilder.String()
	queryStr = queryStr[:len(queryStr)-2]

	argIndex++
	queryStr += fmt.Sprintf(" WHERE user_id = $%d;", argIndex)
	args = append(args, input.UserID)

	_, err := p.db.ExecContext(ctx.Ctx, queryStr, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("profile with user_id %s not found", input.UserID)
		}
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}
