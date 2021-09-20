import os
import sys
from ..database.models import Listener as ListenerModel
from ..listener import Listener, get_listener_types
from typing import Dict, List
from multiprocessing import Process

def run_listener(listener: Listener):
    # sys.stdout = open(os.devnull, 'w')
    # sys.stderr = open(os.devnull, 'w')
    listener.run()

class ListenerManager:
    def __init__(self):
        self.listener_processes: Dict[int, Process] = {}
    
    def _create_listener(self, listener_model: ListenerModel) -> Listener:
        listener_class = get_listener_types()[listener_model.type]
        return listener_class.new(
            'http://127.0.0.1:5000',
            listener_model.id,
            listener_model.bind_host,
            listener_model.bind_port,
            listener_model.sym_key
        )
    
    def create_listener(self, listener_model: ListenerModel):
        self.delete_listener(listener_model.id)
        process = Process(
            target=run_listener,
            args=[self._create_listener(listener_model)],
            daemon=True
        )
        process.start()
        self.listener_processes[listener_model.id] = process
    
    def delete_listener(self, listener_id: int):
        process = self.listener_processes.get(listener_id)
        if process is not None:
            process.terminate()
