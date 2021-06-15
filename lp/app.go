package main

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	protocol "../protocol/app"
	"./models"
)

type AppHandler struct {
	db *gorm.DB
}

func NewAppHandler(db *gorm.DB) *AppHandler {
	return &AppHandler{
		db: db,
	}
}

func (h *AppHandler) handleGetCommandListMsg(
	agent *models.Agent,
	msg *protocol.GetCommandListMsg,
) (*protocol.LPMsg, error) {

	var commands []models.Command
	var commandList []*protocol.Command

	tx := h.db.
		Where("agent_id = ?", agent.ID).
		Where("id > ?", agent.LastCommandId).
		Order("id asc").
		Find(&commands)

	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, cmd := range commands {
		var cmdArgs []string
		var cmdEnv []string
		json.Unmarshal(cmd.Args, &cmdArgs)
		json.Unmarshal(cmd.Env, &cmdEnv)

		if cmd.ID > agent.LastCommandId {
			agent.LastCommandId = cmd.ID
		}

		commandList = append(commandList, &protocol.Command{
			Id:            uint64(cmd.ID),
			Args:          cmdArgs,
			Env:           cmdEnv,
			Stdin:         cmd.Stdin,
			TimeoutMillis: cmd.Timeout.Milliseconds(),
		})
	}

	return &protocol.LPMsg{
		MsgType: &protocol.LPMsg_CommandListMsg{
			CommandListMsg: &protocol.CommandListMsg{
				Commands: commandList,
			},
		},
	}, nil

}

func (h *AppHandler) handleCommandResultListMsg(
	agent *models.Agent,
	msg *protocol.CommandResultListMsg,
) (*protocol.LPMsg, error) {

	for _, cmdResult := range msg.GetCommandResults() {
		h.db.Create(&models.CommandResult{
			CommandId: uint(cmdResult.CommandId),
			AgentId:   agent.ID,
			ExitCode:  int(cmdResult.ExitCode),
			Stdout:    cmdResult.Stdout,
			Stderr:    cmdResult.Stderr,
		})
	}

	return &protocol.LPMsg{
		MsgType: &protocol.LPMsg_SuccessMsg{
			SuccessMsg: &protocol.SuccessMsg{},
		},
	}, nil

}

func (h *AppHandler) handleAgentMsg(
	agent *models.Agent,
	msg *protocol.AgentMsg,
) (*protocol.LPMsg, error) {

	if commandListMsg := msg.GetGetCommandListMsg(); commandListMsg != nil {
		return h.handleGetCommandListMsg(agent, commandListMsg)
	} else if commandResultListMsg := msg.GetCommandResultListMsg(); commandResultListMsg != nil {
		return h.handleCommandResultListMsg(agent, commandResultListMsg)
	}

	return nil, fmt.Errorf("unexpected agent MsgType")
}

func (h *AppHandler) HandleAuthenticatedMsg(clientId string, msg []byte) ([]byte, error) {
	var agent models.Agent
	var agentMsg protocol.AgentMsg

	h.db.FirstOrCreate(&agent, models.Agent{ID: clientId})
	if agent.LastSeenAt.IsZero() {
		fmt.Printf("Agent %s reported in\n", agent.ID)
	}

	err := proto.Unmarshal(msg, &agentMsg)
	if err != nil {
		return nil, err
	}

	lpMsg, err := h.handleAgentMsg(&agent, &agentMsg)
	if err != nil {
		return nil, err
	}

	response, err := proto.Marshal(lpMsg)
	if err != nil {
		return nil, err
	}

	agent.LastSeenAt = time.Now()
	h.db.Save(&agent)

	return []byte(response), nil
}
