package models

import (
	"time"

	"github.com/google/uuid"
)

// NOTE: all repo layer models must be have the suffix of Repo.

// use the base models if possible for upserts

// if base models not useable: models are split into service/handler layer models and repo layer models

// if a service layer model directly match what would be the repo layer model
// there is no need to make a new repo layer model, you can pass the service
// layer model down to the repo layer. This is subject to change to always have
// layer specific models in the future.

type User struct {
	ID            uuid.UUID `form:"id" db:"id"`
	ExternalID    *string   `form:"external_id" db:"external_id"`
	Email         string    `form:"email" db:"email"`
	FirstName     string    `form:"first_name" db:"first_name"`
	LastName      string    `form:"last_name" db:"last_name"`
	EmailVerified bool      `form:"email_verified" db:"email_verified"`
	Banned        bool      `form:"banned" db:"banned"`
	Deleted       bool      `form:"deleted" db:"deleted"`
	CreatedAt     time.Time `form:"created_at" db:"created_at"`
	UpdatedAt     time.Time `form:"updated_at" db:"updated_at"`
}
