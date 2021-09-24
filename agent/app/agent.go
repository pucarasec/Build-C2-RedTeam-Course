package app

import (
	"encoding/json"

	"../comm"
	"./proto"
	"./task"
)

type Agent struct {
	client comm.Client
}

func NewAgent(client comm.Client) *Agent {
	return &Agent{client: client}
}

func (a *Agent) sendMsg(agentMsg *proto.AgentMsg) (*proto.LPMsg, error) {
	var lpMsg proto.LPMsg
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
	agentMsg := &proto.AgentMsg{
		GetTasksMsg: &proto.GetTasksMsg{},
	}
	lpMsg, err := a.sendMsg(agentMsg)
	if err != nil {
		return err
	}
	if taskLisgMsg := lpMsg.TaskListMsg; taskLisgMsg != nil {
		var results []proto.TaskResult
		for _, t := range taskLisgMsg.Tasks {
			taskHandler, _ := task.GetTaskHandler(t)
			result := taskHandler.HandleTask(t)
			results = append(results, result)
		}
		if len(results) > 0 {
			agentMsg := &proto.AgentMsg{
				TaskResultsMsg: &proto.TaskResultsMsg{
					Results: results,
				},
			}

			_, err := a.sendMsg(agentMsg)
			return err
		}
	}

	return nil
}
