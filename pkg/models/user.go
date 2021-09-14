package models

import (
	"time"

	"github.com/google/uuid"
)

// NOTE: all repo layer models must be have the suffix of Repo.

// models are split into service/handler layer models and repo layer models

// if a service layer model directly match what would be the repo layer model
// there is no need to make a new repo layer model, you can pass the service
// layer model down to the repo layer. This is subject to change to always have
// layer specific models in the future.

type UserProfile struct {
	ID            uuid.UUID `db:"id"`
	Email         string    `db:"email"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	EmailVerified bool      `db:"email_verified"`
	Banned        bool      `db:"banned"`
	Deleted       bool      `db:"deleted"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type CreateUser struct {
	Email     string `form:"email" json:"email" binding:"required"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
}

type CreateUserRepo struct {
	ID         uuid.UUID `form:"id" db:"id" binding:"required"`
	ExternalID string    `form:"external_id" db:"external_id" binding:"required"`
	Email      string    `form:"email" db:"email" binding:"required"`
	FirstName  string    `form:"first_name" db:"first_name" binding:"required"`
	LastName   string    `form:"last_name" db:"last_name" binding:"required"`
}
