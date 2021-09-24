//go:build linux || windows

package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"

	"../proto"

	"github.com/kbinani/screenshot"
)

type ScreenshotInfo struct {
	Display int `json:"display"`
}

type ScreenshotResultInfo struct {
	DisplayCount int    `json:"display_count"`
	Display      int    `json:"selected_display"`
	Success      bool   `json:"success"`
	ErrorDesc    string `json:"error_desc,omitempty"`
}

type ScreenshotTaskHandler struct{}

func (*ScreenshotTaskHandler) HandleTask(task proto.Task) proto.TaskResult {
	var screenshotInfo ScreenshotInfo

	json.Unmarshal(task.Info, &screenshotInfo)

	displayCount := screenshot.NumActiveDisplays()

	if screenshotInfo.Display < 0 {
		screenshotResultInfo := ScreenshotResultInfo{
			DisplayCount: displayCount,
			Display:      screenshotInfo.Display,
			Success:      true,
		}
		screenshotResultInfoJson, _ := json.Marshal(screenshotResultInfo)
		return proto.TaskResult{
			TaskId: task.Id,
			Info:   screenshotResultInfoJson,
			Output: nil,
		}
	} else {
		bounds := screenshot.GetDisplayBounds(screenshotInfo.Display)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			screenshotResultInfo := ScreenshotResultInfo{
				DisplayCount: displayCount,
				Display:      screenshotInfo.Display,
				Success:      false,
				ErrorDesc:    fmt.Sprintf("%s", err),
			}
			screenshotResultInfoJson, _ := json.Marshal(screenshotResultInfo)
			return proto.TaskResult{
				TaskId: task.Id,
				Info:   screenshotResultInfoJson,
				Output: nil,
			}
		}

		var buffer bytes.Buffer
		png.Encode(&buffer, img)

		screenshotResultInfo := ScreenshotResultInfo{
			DisplayCount: displayCount,
			Display:      screenshotInfo.Display,
			Success:      true,
		}
		screenshotResultInfoJson, _ := json.Marshal(screenshotResultInfo)
		return proto.TaskResult{
			TaskId: task.Id,
			Info:   screenshotResultInfoJson,
			Output: buffer.Bytes(),
		}
	}
}
