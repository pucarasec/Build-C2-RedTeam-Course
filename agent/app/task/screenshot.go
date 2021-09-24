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

// ScreenshotInfo
// JSON en Task.Info
type ScreenshotInfo struct {
	Display int `json:"display"` // display a tomar screenshot. valores negativos no tienen efecto y devuelven la cantidad de displays
}

// ScreenshotResultInfo
// JSON en TaskResult.Info
type ScreenshotResultInfo struct {
	DisplayCount int    `json:"display_count"`        // cantidad de displays del sistema
	Display      int    `json:"selected_display"`     // display seleccionado para tomar screenshot
	Success      bool   `json:"success"`              // tuvo exito?
	ErrorDesc    string `json:"error_desc,omitempty"` // descripcion del error en caso de fallar
}

// ScreenshotTaskHandler maneja una tarea del tipo screenshot
// Toma un screenshot utilizando la libreria "github.com/kbinani/screenshot"
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
