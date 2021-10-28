package model

import "github.com/google/uuid"

type Transaction struct {
	ID        uuid.UUID
	CreatedAt int64
	UpdatedAt int64
	Type      int
}
