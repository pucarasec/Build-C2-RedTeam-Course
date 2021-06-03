package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"../crypto/dh"
	protocol "../protocol/app"
	"./comm/layer0"
	"./comm/layer1"
	"./models"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
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

	tx := h.db.Where("id > ?", msg.LastId).Order("id asc").Find(&commands)

	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, cmd := range commands {
		var cmdArgs []string
		json.Unmarshal(cmd.Args, &cmdArgs)

		commandList = append(commandList, &protocol.Command{
			Id:   uint64(cmd.ID),
			Args: cmdArgs,
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
		fmt.Printf("%s: %s", agent.ID, cmdResult.Output)
		h.db.Create(&models.CommandResult{
			CommandId: uint(cmdResult.CommandId),
			AgentId:   agent.ID,
			Output:    cmdResult.Output,
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
	if agent.FirstSeen.IsZero() {
		fmt.Printf("Agent %s reported in\n", agent.ID)
		agent.FirstSeen = time.Now()
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

	agent.LastSeen = time.Now()

	h.db.Save(&agent)

	return []byte(response), nil
}

func readFromTerminal(db *gorm.DB) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		fields := strings.Split(text[:len(text)-1], " ")
		cmdArgs, err := json.Marshal(fields)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		tx := db.Create(&models.Command{
			Args: cmdArgs,
		})
		if tx.Error != nil {
			fmt.Printf("Error: %s\n", tx.Error)
			continue
		}
	}
}

func main() {
	key := []byte("some random key!")
	db := models.GetDB()
	appHandler := NewAppHandler(db)
	dhHandler := layer1.NewDHHandler(layer1.NewBasicKeyRespository(), dh.NewKeyExchange(), appHandler)
	encryptedHandler := layer0.NewEncryptedHandler(key, dhHandler)
	handler := layer0.NewHTTPHandler(encryptedHandler)
	http.Handle("/", handler)
	go readFromTerminal(db)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
