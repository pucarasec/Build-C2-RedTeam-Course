from .client import Client
from malon_common.crypto.sym import SymmetricCipher

class EncryptedClient:
    def __init__(self, key: bytes, client: Client):
        self._cipher = SymmetricCipher(key)
        self._client = client
    
    def send_msg(self, msg: bytes) -> bytes:
        encrypted_msg = self._cipher.encrypt_sign_msg(msg)
        encrypted_response_msg = self._client.send_msg(encrypted_msg)
        response_msg = self._cipher.verify_decrypt_msg(encrypted_response_msg)
        if response_msg is None:
            raise RuntimeError('Message did not pass HMAC verification')

        return response_msg
