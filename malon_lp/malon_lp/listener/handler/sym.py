from typing import Optional
from . import Handler

from malon_lp.crypto.sym import SymmetricCipher

class EncryptedHandler(Handler):
    KEY_LENGTH = 16

    def __init__(self, key: bytes, handler: Handler):
        if len(key) != self.KEY_LENGTH:
            raise RuntimeError('Invalid key length, should be {}'.format(self.KEY_LENGTH))
        
        self._cipher = SymmetricCipher(key)
        self._handler = handler
    
    def handle_msg(self, msg: bytes) -> bytes:
        """
            Delega el mensaje descfirado (por la primer capa de criptografia simetrica)
            en el Handler siguiente
        """
        decrypted_msg = self._cipher.verify_decrypt_msg(msg)
        response_msg = self._handler.handle_msg(decrypted_msg)
        encrypted_response_msg = self._cipher.encrypt_sign_msg(response_msg)
        return encrypted_response_msg
