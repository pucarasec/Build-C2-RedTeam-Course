package task

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"../proto"
)

// CommandInfo
// JSON en Task.Info
type CommandInfo struct {
	TimeoutMs int      `json:"timeout_ms"` // timeout del comando
	Env       []string `json:"env"`        // variables del torno del comando en formato ["NAME=VALUE", ...]
	Args      []string `json:"args"`       // argumentos a enviar al comando. El primeor es el path del ejecutable.
}

// CommandResultInfo
// JSON en TaskResult.Info
type CommandResultInfo struct {
	ExitCode  int    `json:"exit_code"`            // Exit code del comando ejecutado
	ErrorDesc string `json:"error_desc,omitempty"` // descripcion del error en caso de fallar
}

// CommandTaskHandler maneja una tarea del tipo command
// Ejecuta el binario en CommandInfo.Args[0], envia como stdin a Task.Input
// Almacena stderr y stdout en Task.Output
type CommandTaskHandler struct{}

func (*CommandTaskHandler) HandleTask(task proto.Task) proto.TaskResult {
	var cmdInfo CommandInfo
	json.Unmarshal(task.Info, &cmdInfo)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cmdInfo.TimeoutMs)*time.Millisecond,
	)
	defer cancel()

	execCmd := exec.CommandContext(ctx, cmdInfo.Args[0], cmdInfo.Args[1:]...)
	execCmd.Stdin = bytes.NewReader(task.Input)
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr

	execCmd.Env = append(os.Environ(), cmdInfo.Env...)

	err := execCmd.Run()

	cmdResultInfo := CommandResultInfo{
		ExitCode: 0,
	}

	if err != nil {
		switch t := err.(type) {
		case *exec.ExitError:
			cmdResultInfo.ExitCode = t.ExitCode()
		case *exec.Error:
			cmdResultInfo.ExitCode = 127
		}
		cmdResultInfo.ErrorDesc = fmt.Sprintf("Error: %s\n", err)
	}

	cmdResultInfoJson, _ := json.Marshal(cmdResultInfo)
	output := append(stderr.Bytes(), stdout.Bytes()...)
	return proto.TaskResult{
		TaskId: task.Id,
		Info:   cmdResultInfoJson,
		Output: output,
	}
}
