from typing import NamedTuple, Optional

from . import AuthHandler

import json
import requests

class ApiHandler(AuthHandler):
    def __init__(self, base_url: str):
        self._base_url = base_url
    
    def _report_agent(self, client_id: str):
        requests.post('{}/agents/{}/report'.format(self._base_url, client_id))

    def _handle_get_tasks_msg(self, client_id: str, _msg) -> bytes:
        response = requests.get('{}/agents/{}/tasks/unread/'.format(self._base_url, client_id))
        tasks = response.json() if response.ok else []
        return json.dumps({
            'task_list_msg': {
                'tasks': tasks
            }
        }).encode('utf-8')

    def _handle_task_results_msg(self, client_id: str, msg) -> bytes:
        for result in msg['results']:
            task_id = result['task_id']
            response = requests.post('{}/agents/{}/tasks/{}/result'.format(
                self._base_url,
                client_id,
                task_id
            ), json=result)
        return json.dumps({
            'status_msg': {
                'success': True
            }
        }).encode('utf-8')

    def handle_auth_msg(self, msg: bytes, client_id: str) -> bytes:
        self._report_agent(client_id)
        agent_msg = json.loads(msg.decode('utf-8'))
        if agent_msg.get('get_tasks_msg') is not None:
            return self._handle_get_tasks_msg(client_id, agent_msg['get_tasks_msg'])
        elif agent_msg.get('task_results_msg') is not None:
            return self._handle_task_results_msg(client_id, agent_msg['task_results_msg'])
        else:
            raise RuntimeError('Unexpected message type')