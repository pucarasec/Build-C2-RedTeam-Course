package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key"`
	SharedKey []byte `json:"shared_key"`
}
