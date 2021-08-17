from typing import Optional
import hmac
from Crypto.Cipher import AES
from Crypto.Random import get_random_bytes


class SymmetricCipher:
    def __init__(self, key: bytes) -> bytes:
        self._key = key
    
    def encrypt_msg(self, msg: bytes) -> bytes:
        iv = get_random_bytes(AES.block_size)
        cipher = AES.new(self._key, AES.MODE_OFB, iv)
        padded_msg = msg.ljust(len(msg) + ((-len(msg)) % AES.block_size), b'\x00')
        return iv + cipher.encrypt(padded_msg)[:len(msg)]
    
    def decrypt_msg(self, msg: bytes) -> bytes:
        iv, msg = msg[:AES.block_size], msg[AES.block_size:]
        cipher = AES.new(self._key, AES.MODE_OFB, iv)
        padded_msg = msg.ljust(len(msg) + ((-len(msg)) % AES.block_size), b'\x00')
        return cipher.decrypt(padded_msg)[:len(msg)]
    
    def sign_msg(self, msg: bytes) -> bytes:
        return hmac.digest(self._key, msg, 'sha1') + msg

    def verify_msg(self, msg: bytes) -> Optional[bytes]:
        sig, msg = msg[:20], msg[20:]
        verify_sig = hmac.digest(self._key, msg, 'sha1')
        if hmac.compare_digest(sig, verify_sig):
            return msg
        else:
            return None
    
    def encrypt_sign_msg(self, msg: bytes) -> bytes:
        return self.sign_msg(self.encrypt_msg(msg))
    
    def verify_decrypt_msg(self, msg: bytes) -> Optional[bytes]:
        verified_msg = self.verify_msg(msg)
        if verified_msg is not None:
            return self.decrypt_msg(verified_msg)
        else:
            return None
