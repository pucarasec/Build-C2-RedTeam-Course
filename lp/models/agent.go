package models

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	ID        string    `json:"id" gorm:"primary_key"`
	FirstSeen time.Time `json:"first_seen"`
	LastSeen  time.Time `json:"last_seen"`
}

type Command struct {
	gorm.Model
	Args  []byte `json:"args"`
	Input []byte `json:"input"`
}

type CommandResult struct {
	CommandId uint
	Output    []byte
	AgentId   string
	Command   Command
	Agent     Agent
}
