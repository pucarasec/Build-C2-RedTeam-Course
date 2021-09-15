from typing import Optional
from abc import ABC, abstractmethod

class Handler(ABC):
    @abstractmethod
    def handle_msg(self, msg: bytes, client_id: Optional[str] = None):
        pass
