package models

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	ID            string    `json:"ID" gorm:"primary_key"`
	LastSeenAt    time.Time `json:"LastSeenAt"`
	LastCommandId uint      `json:"LastCommandId"`
	Commands      []Command
}

type Command struct {
	gorm.Model
	AgentId string
	Agent   *Agent
	Args    []byte        `json:"args"`
	Env     []byte        `json:"env"`
	Stdin   []byte        `json:"stdin"`
	Timeout time.Duration `json:"timeout"`
}

type CommandResult struct {
	gorm.Model
	CommandId uint
	AgentId   string
	ExitCode  int
	Stdout    []byte
	Stderr    []byte
	Command   *Command
	Agent     *Agent
}
