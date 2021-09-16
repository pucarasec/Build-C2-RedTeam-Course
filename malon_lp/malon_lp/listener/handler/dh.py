import json
from typing import Mapping, Optional
from base64 import b64encode, b64decode

from . import Handler, AuthHandler
from malon_lp.crypto.dh import KeyExchange, get_client_id
from malon_lp.crypto.sym import SymmetricCipher



KeyRespository = Mapping[str, bytes]


class DHHandler(Handler):
    def __init__(self, ke: KeyExchange, handler: AuthHandler, kr: Optional[KeyRespository] = None):
        self._ke = ke
        self._handler = handler
        self._kr = kr if kr is not None else {}
    
    def _handle_handshake_msg(self, public_key: bytes) -> bytes:
        client_id = get_client_id(public_key).hex()
        self._kr[client_id] = public_key
        response_msg = {
            'handshake_msg': {
                'public_key': b64encode(self._ke.get_public_key()).decode('utf-8')
            }
        }
        return json.dumps(response_msg).encode('utf-8')
    
    def _handle_client_msg(self, client_id: str, encrypted_payload: bytes) -> bytes:
        key = self._kr.get(client_id, None)
        if key is not None:
            cipher = SymmetricCipher(self._ke.get_shared_key(self._kr[client_id]))
            payload = cipher.verify_decrypt_msg(encrypted_payload)
            response = self._handler.handle_auth_msg(payload, client_id)
            encrypted_response = cipher.encrypt_sign_msg(response)
            return json.dumps({
                'server_msg': {
                    'payload': b64encode(encrypted_response).decode('utf-8')
                }
            }).encode('utf-8')
        else:
            return json.dumps({
                'error_msg': {
                    'type': 'HANDSHAKE_EXPIRED'
                }
            }).encode('utf-8')

    def handle_msg(self, msg: bytes) -> bytes:
        base_msg = json.loads(msg.decode('utf-8'))
        if base_msg.get('handshake_msg') is not None:
            return self._handle_handshake_msg(b64decode(base_msg['handshake_msg']['public_key']))
        elif base_msg.get('client_msg') is not None:
            return self._handle_client_msg(
                base_msg['client_msg']['client_id'],
                b64decode(base_msg['client_msg']['payload'])
            )
        else:
            raise RuntimeError('Unexpected message type')
