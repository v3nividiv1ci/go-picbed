package model

import (
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pic fields marked `json:"-"` can be hidden when used in json
type Pic struct {
	gorm.Model `json:"-"`
	PicName    string
	Uuid       string
	Master     string `json:"-"`
}
