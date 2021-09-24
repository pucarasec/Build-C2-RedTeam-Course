package task

import (
	"fmt"

	"../proto"
)

type TaskHandler interface {
	HandleTask(task proto.Task) proto.TaskResult
}

func GetTaskHandler(task proto.Task) (TaskHandler, error) {
	switch task.Type {
	case "command":
		return &CommandTaskHandler{}, nil
	case "file":
		return &FileTaskHandler{}, nil
	case "screenshot":
		return &ScreenshotTaskHandler{}, nil
	}

	return nil, fmt.Errorf("unknown task type")
}
