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
        "{}/admin/agents/{}/commands".format(base_url, agent_id),
        json=command_args
    )
    command_d = response.json()
    command_id = command_d['ID']
    while True:
        response = requests.get(
            "{}/admin/agents/{}/commands/{}/results".format(base_url, agent_id, command_id),
        )
        command_result_list = response.json()
        for command_result in command_result_list:
            if command_result['Stdout'] is not None:
                sys.stdout.write(b64decode(command_result['Stdout']).decode('utf-8'))
            if command_result['Stderr'] is not None:
                sys.stdout.write(b64decode(command_result['Stderr']).decode('utf-8'))

        if len(command_result_list) > 0:
            break
        else:
            time.sleep(1.0)
