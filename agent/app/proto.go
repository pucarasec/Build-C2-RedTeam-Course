package app

import "encoding/json"

type TaskResult struct {
	TaskId int             `json:"task_id"`
	Info   json.RawMessage `json:"info"`
	Output []byte          `json:"output"`
}

type TaskResultsMsg struct {
	Results []TaskResult `json:"results"`
}

type Task struct {
	Id    int             `json:"id"`
	Type  string          `json:"type"`
	Info  json.RawMessage `json:"info"`
	Input []byte          `json:"input"`
}

type TaskListMsg struct {
	Tasks []Task `json:"tasks"`
}

type GetTasksMsg struct{}

type AgentMsg struct {
	TaskResultsMsg *TaskResultsMsg `json:"task_results_msg,omitempty"`
	GetTasksMsg    *GetTasksMsg    `json:"get_tasks_msg,omitempty"`
}

type StatusMsg struct {
	Success bool `json:"success"`
}

type LPMsg struct {
	TaskListMsg *TaskListMsg `json:"task_list_msg,omitempty"`
	StatusMsg   *StatusMsg   `json:"status_msg,omitempty"`
}
