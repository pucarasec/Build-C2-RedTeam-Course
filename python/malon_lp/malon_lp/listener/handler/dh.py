from typing import Mapping, Optional
from . import Handler
from malon_common.crypto.dh import KeyExchange, get_client_id
from malon_common.crypto.sym import SymmetricCipher
from malon_common.protocol.base_pb2 import BaseMsg, ErrorType


KeyRespository = Mapping[str, bytes]


class DHHandler(Handler):
    def __init__(self, ke: KeyExchange, handler: Handler, kr: Optional[KeyRespository] = None):
        self._ke = ke
        self._handler = handler
        self._kr = kr if kr is not None else {}
    
    def _handle_handshake_msg(self, public_key: bytes) -> bytes:
        client_id = get_client_id(public_key)
        self._kr[client_id] = public_key
        response_msg = BaseMsg()
        response_msg.HandshakeMsg.PublicKey = self._ke.get_public_key()
        return response_msg.SerializeToString()
    
    def _handle_client_msg(self, client_id: bytes, encrypted_payload: bytes) -> bytes:
        key = self._kr.get(client_id, None)
        response_msg = BaseMsg()
        if key is not None:
            cipher = SymmetricCipher(self._ke.get_shared_key(self._kr[client_id]))
            payload = cipher.verify_decrypt_msg(encrypted_payload)
            response = self._handler.handle_msg(payload, client_id=client_id.hex())
            encrypted_response = cipher.encrypt_sign_msg(response)
            response_msg = BaseMsg()
            response_msg.ServerMsg.Payload = encrypted_response
        else:
            response_msg.ErrorMsg.type = ErrorType.HANDSHAKE_EXPIRED

        return response_msg.SerializeToString()

    def handle_msg(self, msg: bytes, client_id: Optional[str] = None) -> bytes:
        base_msg = BaseMsg()
        base_msg.ParseFromString(msg)
        if base_msg.HasField('HandshakeMsg'):
            return self._handle_handshake_msg(base_msg.HandshakeMsg.PublicKey)
        elif base_msg.HasField('ClientMsg'):
            return self._handle_client_msg(base_msg.ClientMsg.ClientID, base_msg.ClientMsg.Payload)
        else:
            raise RuntimeError('Unexpected message type')
