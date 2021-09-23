import json
from typing import Mapping, Optional
from base64 import b64encode, b64decode

from . import Handler, AuthenticatedHandler
from malon_lp.crypto.dh import KeyExchange, get_client_id
from malon_lp.crypto.sym import SymmetricCipher


KeyRespository = Mapping[str, bytes]


class DummyAuthenticationHandler(Handler):
    def __init__(self, handler: AuthenticatedHandler):
        self._handler = handler

    def handle_msg(self, msg: bytes, client_id: Optional[str] = None) -> bytes:
        """
            Desencodea el contenido del payload y delega el mensage en el AuthenticatedHandler

            Recibira un mensaje con la siguiente estructura:
            msg = b'{"client_msg": {"payload": "cHVjYXJhCg==", "client_id":"9927094366ecadea795bc37a32e5e140"}}'

            Y debera devolver un mensaje en bytes de la  forma 
            b'{"server_msg": {"payload": "cHVjYXJhCg==" } }'

            Pistas:
                - utilizar la libreria 'json' para deserializar msg 
                - utilizar la libreria 'b64decode' para desencodear el payload enviado al AuthenticatedHandler
        """
        raise NotImplementedError
