//go:build !windows && !linux

package task

import (
	"encoding/json"

	"../proto"
)

type ScreenshotResultInfo struct {
	DisplayCount int    `json:"display_count"`
	Display      int    `json:"selected_display"`
	Success      bool   `json:"success"`
	ErrorDesc    string `json:"error_desc,omitempty"`
}

type ScreenshotTaskHandler struct{}

func (*ScreenshotTaskHandler) HandleTask(task proto.Task) proto.TaskResult {
	screenshotResultInfo := ScreenshotResultInfo{
		DisplayCount: 0,
		Display:      -1,
		Success:      false,
		ErrorDesc:    "screenshots not supported in this OS",
	}
	screenshotResultInfoJson, _ := json.Marshal(screenshotResultInfo)
	return proto.TaskResult{
		TaskId: task.Id,
		Info:   screenshotResultInfoJson,
		Output: nil,
	}
}
