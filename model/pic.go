package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pic struct {
	gorm.Model
	Uuid   uuid.UUID
	Master string
}
