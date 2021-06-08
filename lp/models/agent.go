package models

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	ID         string    `json:"ID" gorm:"primary_key"`
	LastSeenAt time.Time `json:"LastSeenAt"`
	Commands   []Command
}

type Command struct {
	gorm.Model
	AgentId string
	Agent   *Agent
	Args    []byte `json:"args"`
	Input   []byte `json:"input"`
}

type CommandResult struct {
	gorm.Model
	CommandId uint
	AgentId   string
	Output    []byte
	Command   *Command
	Agent     *Agent
}
