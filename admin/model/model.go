package model

import (
	"time"
)

type Group struct {
	Id        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
