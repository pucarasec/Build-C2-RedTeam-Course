package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	protocol "../protocol/app"
	"./comm"
	"google.golang.org/protobuf/proto"
)

type Agent struct {
	client        comm.Client
	lastCommandId uint
}

type Command struct {
	Timeout time.Duration
	Env     []string
	Args    []string
	Stdin   []byte
}

type CommandResult struct {
	ExitCode int
	Stderr   []byte
	Stdout   []byte
}

func NewAgent(client comm.Client) *Agent {
	return &Agent{client: client}
}

func (a *Agent) runCommand(cmd Command) CommandResult {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	ctx, cancel := context.WithTimeout(context.Background(), cmd.Timeout)
	defer cancel()

	execCmd := exec.CommandContext(ctx, cmd.Args[0], cmd.Args[1:]...)
	execCmd.Stdin = bytes.NewReader(cmd.Stdin)
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr

	execCmd.Env = append(os.Environ(), cmd.Env...)

	err := execCmd.Run()

	cmdResult := CommandResult{
		ExitCode: 0,
		Stdout:   stdout.Bytes(),
		Stderr:   stderr.Bytes(),
	}

	if err != nil {
		switch t := err.(type) {
		case *exec.ExitError:
			fmt.Printf("ExitError: %s\n", t)
			cmdResult.ExitCode = t.ExitCode()
		case *exec.Error:
			fmt.Printf("Error: %s\n", t)
			cmdResult.ExitCode = 127
		}
	}

	return cmdResult
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
			GetCommandListMsg: &protocol.GetCommandListMsg{},
		},
	}
	lpMsg, err := a.sendMsg(agentMsg)
	if err != nil {
		return
	}
	if commandListMsg := lpMsg.GetCommandListMsg(); commandListMsg != nil {
		var commandResults []*protocol.CommandResult
		for _, cmd := range commandListMsg.GetCommands() {
			cmdResult := a.runCommand(Command{
				Timeout: time.Duration(cmd.TimeoutMillis) * time.Millisecond,
				Env:     cmd.Env,
				Args:    cmd.Args,
				Stdin:   cmd.Stdin,
			})
			commandResults = append(commandResults, &protocol.CommandResult{
				CommandId: cmd.GetId(),
				ExitCode:  int32(cmdResult.ExitCode),
				Stdout:    cmdResult.Stdout,
				Stderr:    cmdResult.Stderr,
			})
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

		a.sendMsg(agentMsg)
	}
}
