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
            en el Handler siguiente.

            - se debe verificar y descifrar el mensaje
            - delegar el mensaje al Handler siguiente
            - finalmente se debe cifrar y firmar el mensaje, para poder devolverlo
        """
        raise NotImplementedError
