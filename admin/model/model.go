package model

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
