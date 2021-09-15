package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type CommandInfo struct {
	TimeoutMs int      `json:"timeout_ms"`
	Env       []string `json:"env"`
	Args      []string `json:"args"`
}

type CommandResultInfo struct {
	ExitCode int `json:"exit_code"`
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

func runCommand(cmd Command) CommandResult {
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

func handleCommandTask(task Task) TaskResult {
	var cmdInfo CommandInfo
	json.Unmarshal(task.Info, &cmdInfo)
	cmd := Command{
		Timeout: time.Duration(cmdInfo.TimeoutMs) * time.Millisecond,
		Env:     cmdInfo.Env,
		Args:    cmdInfo.Args,
		Stdin:   task.Input,
	}
	cmdResult := runCommand(cmd)
	cmdResultInfo := CommandResultInfo{
		ExitCode: cmdResult.ExitCode,
	}
	cmdResultInfoJson, _ := json.Marshal(cmdResultInfo)
	output := append(cmdResult.Stderr, cmdResult.Stdout...)
	return TaskResult{
		TaskId: task.Id,
		Info:   cmdResultInfoJson,
		Output: output,
	}
}
