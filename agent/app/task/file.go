package task

import (
	"encoding/json"
	"fmt"
	"os"

	"../proto"
)

type FileInfo struct {
	Type     string `json:"type"`
	FilePath string `json:"file_path"`
}

type FileResultInfo struct {
	Success   bool   `json:"success"`
	ErrorDesc string `json:"error_desc,omitempty"`
}

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
