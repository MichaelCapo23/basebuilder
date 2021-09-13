package models

import (
	"time"

	"github.com/google/uuid"
)

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
	Email     string  `form:"email" binding:"required"`
	Password  string  `form:"password" binding:"required"`
	FirstName string  `form:"first_name" binding:"required"`
	LastName  string  `form:"last_name" binding:"required"`
	Phone     *string `form:"phone" binding:"omitempty,stirng"`
}
