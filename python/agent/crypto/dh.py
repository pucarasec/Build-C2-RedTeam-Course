
from typing import Optional
from hashlib import sha256
from ecdsa import ECDH, NIST256p, SigningKey, VerifyingKey

class KeyExchange:
    def __init__(self, private_key_der: Optional[bytes] = None):
        if private_key_der is not None:
            self.private_key = SigningKey.from_der(private_key_der)
        else:
            self.private_key = SigningKey.generate(curve=NIST256p)
    
    def get_shared_key(self, public_key_der: bytes) -> bytes:
        return ECDH(
            curve=NIST256p,
            private_key=self.private_key,
            public_key=VerifyingKey.from_der(public_key_der)
        ).generate_sharedsecret_bytes()
    
    def get_public_key(self) -> bytes:
        public_key: VerifyingKey = self.private_key.get_verifying_key()
        return public_key.to_der()
    
    def get_private_key(self) -> bytes:
        return self.private_key.to_der()

def get_client_id(public_key_der: bytes) -> bytes:
    return sha256(public_key_der).digest()[:16]
