import json
from typing import Mapping, Optional
from base64 import b64encode, b64decode

from . import Handler, AuthHandler
from malon_lp.crypto.dh import KeyExchange, get_client_id
from malon_lp.crypto.sym import SymmetricCipher



KeyRespository = Mapping[str, bytes]


class DummyAuthHandler(Handler):
    def __init__(self, handler: AuthHandler):
        self._handler = handler

    def handle_msg(self, msg: bytes, client_id: Optional[str] = None) -> bytes:
        """
            Desencodea el contenido del payload y delega el mensage en el AuthHandler
        """
        base_msg = json.loads(msg.decode('utf-8'))
        if base_msg.get('client_msg') is not None:
            return self._handler.handle_auth_msg(
                b64decode(base_msg['client_msg']['payload']),
                base_msg['client_msg']['client_id']
            )
        else:
            raise RuntimeError('Unexpected message type')
