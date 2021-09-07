from typing import NamedTuple
from base64 import b64decode, b64encode
from malon_common.protocol.app_pb2 import AgentMsg, LPMsg, Command

import requests

class ApiHandler:
    def __init__(self, base_url: str):
        self._base_url = base_url
    
    def _report_agent(self, client_id: str):
        requests.post('{}/agents/{}/report'.format(self._base_url, client_id))

    def _handle_get_command_list_msg(self, client_id: str, _msg) -> bytes:
        lp_msg = LPMsg()
        response = requests.get('{}/agents/{}/tasks/unread/'.format(self._base_url, client_id))
        if response.ok:
            tasks = response.json()
            for task in tasks:
                if task['type'] == 'command':
                    command = Command()
                    command.id = task['id']
                    for arg in task['info']['args']:
                        command.args.append(arg)
                    command.timeoutMillis = task['info']['timeout_millis']
                    lp_msg.CommandListMsg.commands.append(command)
        return lp_msg.SerializeToString()

    def _handle_command_result_list_msg(self, client_id: str, msg) -> bytes:
        for command_result in msg.commandResults:
            task_id = command_result.commandId
            data = {
                'info': {'status': command_result.exitCode},
                'output': b64encode(command_result.stderr + command_result.stdout).decode('utf-8')
            }
            response = requests.post('{}/agents/{}/tasks/{}/result'.format(self._base_url, client_id, task_id), json=data)

        lp_msg = LPMsg()
        lp_msg.SuccessMsg.SetInParent()
        return lp_msg.SerializeToString()

    def handle_msg(self, msg: bytes, client_id: str) -> bytes:
        self._report_agent(client_id)
        print(msg)
        agent_msg = AgentMsg()
        agent_msg.ParseFromString(msg)
        if agent_msg.HasField('GetCommandListMsg'):
            return self._handle_get_command_list_msg(client_id, agent_msg.GetCommandListMsg)
        elif agent_msg.HasField('CommandResultListMsg'):
            return self._handle_command_result_list_msg(client_id, agent_msg.CommandResultListMsg)
        else:
            raise RuntimeError('Unexpected message type')