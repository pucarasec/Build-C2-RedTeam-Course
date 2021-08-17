class Protocol:
    pass

class Client(Protocol):
    def send_msg(self, msg: bytes) -> bytes:
        pass
