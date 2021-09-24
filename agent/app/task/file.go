package task

import (
	"encoding/json"
	"fmt"
	"os"

	"../proto"
)

// FileInfo
// llega serializada en JSON en Task.Info
type FileInfo struct {
	Type     string `json:"type"`
	FilePath string `json:"file_path"`
}

// FileResultInfo
// Vuelve al listener serializada en JSON en TaskrRsult.Info
type FileResultInfo struct {
	Success   bool   `json:"success"`
	ErrorDesc string `json:"error_desc,omitempty"`
}

// FileTaskHandler maneja una tarea del tipo File
// Existen dos subtipos:
//	- "put" escribe los bytes de task.Input en el archivo especificado en Task.Info
//	- "get" lee los bytes del archivo especificado y los devuelve en TaskResult.Output
// El exito/fracaso de la tarea se informa en TaskResult.Info
type FileTaskHandler struct{}

func (*FileTaskHandler) HandleTask(task proto.Task) proto.TaskResult {
	var fileInfo FileInfo
	json.Unmarshal(task.Info, &fileInfo)

	var fileResultInfo FileResultInfo
	var fileData []byte = nil
	var err error

	switch fileInfo.Type {
	case "put":
		err = os.WriteFile(fileInfo.FilePath, task.Input, 0666)
		if err == nil {
			fileResultInfo.Success = true
		} else {
			fileResultInfo.Success = false
			fileResultInfo.ErrorDesc = fmt.Sprintf("%s", err)
		}
	case "get":
		fileData, err = os.ReadFile(fileInfo.FilePath)
		if err == nil {
			fileResultInfo.Success = true
		} else {
			fileResultInfo.Success = false
			fileResultInfo.ErrorDesc = fmt.Sprintf("%s", err)
		}
	default:
		fileResultInfo.Success = false
	}

	fileResultInfoJson, _ := json.Marshal(fileResultInfo)

	return proto.TaskResult{
		TaskId: task.Id,
		Info:   fileResultInfoJson,
		Output: fileData,
	}
}
