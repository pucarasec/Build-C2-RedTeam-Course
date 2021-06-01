package models

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	Id        string    `json:"id" gorm:"primary_key"`
	FirstSeen time.Time `json:"first_seen"`
	LastSeen  time.Time `json:"last_seen"`
}

type Message struct {
	gorm.Model
	AgentId string
	Agent   Agent
	Text    string
}
