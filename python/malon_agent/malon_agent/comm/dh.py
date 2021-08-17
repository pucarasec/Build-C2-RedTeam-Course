from typing import Optional
from malon_common.protocol.base_pb2 import BaseMsg, ClientMsg, HandshakeMsg, ErrorType
from .client import Client
from malon_common.crypto.dh import KeyExchange, get_client_id
from malon_common.crypto.sym import SymmetricCipher

class DHClient:
    def __init__(self, ke: KeyExchange, client: Client):
        self._ke = ke
        self._client = client
        self._shared_key_cipher: Optional[SymmetricCipher] = None
    
    def get_client_id(self) -> bytes:
        return get_client_id(self._ke.get_public_key())

    def negotiate_key(self):
        msg = BaseMsg()
        msg.HandshakeMsg.PublicKey = self._ke.get_public_key()

        response = self._client.send_msg(msg.SerializeToString())
        response_msg = BaseMsg()
        response_msg.ParseFromString(response)

        shared_key = self._ke.get_shared_key(response_msg.HandshakeMsg.PublicKey)
        self._shared_key_cipher = SymmetricCipher(shared_key)

    def send_msg(self, msg: bytes) -> bytes:
        if self._shared_key_cipher is None:
            self.negotiate_key()
        
        client_msg = BaseMsg()
        client_msg.ClientMsg.ClientID = self.get_client_id()
        client_msg.ClientMsg.Payload = self._shared_key_cipher.encrypt_sign_msg(msg)

        response = self._client.send_msg(client_msg.SerializeToString())

        response_msg = BaseMsg()
        response_msg.ParseFromString(response)

        if response_msg.HasField('ServerMsg'):
            payload = self._shared_key_cipher.verify_decrypt_msg(response_msg.ServerMsg.Payload)
            if payload is not None:
                return payload
            else:
                raise RuntimeError('Message did not pass verification')
        elif response_msg.HasField('ErrorMsg'):
            if response_msg.ErrorMsg.type == ErrorType.HANDSHAKE_EXPIRED:
                self.negotiate_key()
                return self.send_msg(msg)
            else:
                raise RuntimeError('Server returned unknown error')
        else:
            raise RuntimeError('Unexpected message type')


