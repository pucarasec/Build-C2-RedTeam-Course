import requests
import time
import sys
from base64 import b64decode

if len(sys.argv) < 3:
    print('Usage: {} BASE_URL AGENT_ID'.format(sys.argv[0]))
    exit(0)

base_url, agent_id = sys.argv[1:]

while True:
    command = input(">")
    command_args = command.split(' ')
    response = requests.post(
        "{}/agents/{}/tasks/".format(base_url, agent_id),
        json={
            'type': 'command',
            'info': {
                'args': command_args,
                'timeout_millis': 5000
            }
        }
    )
    task_d = response.json()
    task_id = task_d['id']
    while True:
        response = requests.get(
            "{}/agents/{}/tasks/{}/result".format(base_url, agent_id, task_id),
        )
        task_result_list = response.json()
        for task_result in task_result_list:
            sys.stdout.write(b64decode(task_result['output']).decode('utf-8'))

        if len(task_result_list) > 0:
            break
        else:
            time.sleep(1.0)
