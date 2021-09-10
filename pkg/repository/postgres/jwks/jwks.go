package jwks

import "github.com/jmoiron/sqlx"

type JWKSStore struct {
	db *sqlx.DB
}

func NewJWKSStore(db *sqlx.DB) *JWKSStore {
	return &JWKSStore{
		db: db,
	}
}
