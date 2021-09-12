package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	ID            uuid.UUID
	Email         string
	FirstName     string
	LastName      string
	EmailVerified bool
	Banned        bool
	Deleted       bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
