import subprocess
from .comm.client import Client
from .protocol.app_pb2 import AgentMsg, LPMsg, CommandResult


class Agent:
    def __init__(self, client: Client):
        self._client = client
    
    def heartbeat(self):
        agent_msg = AgentMsg()
        agent_msg.GetCommandListMsg.SetInParent()

        response = self._client.send_msg(agent_msg.SerializeToString())
        lp_msg = LPMsg()
        lp_msg.ParseFromString(response)

        if lp_msg.HasField('CommandListMsg'):
            command_results = []
            for command in lp_msg.CommandListMsg.commands:
                try:
                    completed = subprocess.run(
                        command.args,
                        input=command.stdin,
                        timeout=command.timeoutMillis / 1000.0,
                        capture_output=True
                    )
                    command_result = CommandResult()
                    command_result.commandId = command.id
                    command_result.exitCode = completed.returncode
                    if completed.stdout is not None:
                        command_result.stdout = completed.stdout
                    if completed.stderr is not None:
                        command_result.stderr = completed.stderr
                    command_results.append(command_result)
                except subprocess.TimeoutExpired as e:
                    command_result = CommandResult()
                    command_result.commandId = command.id
                    command_result.exitCode = -1
                    if e.stdout is not None:
                        command_result.stdout = e.stdout
                    if e.stderr is not None:
                        command_result.stderr = e.stderr
                    command_results.append(command_result)
                except PermissionError as e:
                    command_result = CommandResult()
                    command_result.commandId = command.id
                    command_result.exitCode = -1
                    command_results.append(command_result)


            if len(command_results) > 0:
                agent_msg = AgentMsg()
                agent_msg.CommandResultListMsg.commandResults.extend(command_results)
                self._client.send_msg(agent_msg.SerializeToString())
