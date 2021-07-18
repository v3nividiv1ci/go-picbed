package model

import (
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pic struct {
	gorm.Model
	PicName string
	Uuid    string
	Master  string
}
