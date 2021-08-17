
from .handler import Handler

from malon_common.crypto.sym import SymmetricCipher

class EncryptedHandler:
    def __init__(self, key: bytes, handler: Handler):
        self._cipher = SymmetricCipher(key)
        self._handler = handler
    
    def handle_msg(self, msg: bytes) -> bytes:
        decrypted_msg = self._cipher.verify_decrypt_msg(msg)
        response_msg = self._handler.handle_msg(decrypted_msg)
        encrypted_response_msg = self._cipher.encrypt_sign_msg(response_msg)
        return encrypted_response_msg
    