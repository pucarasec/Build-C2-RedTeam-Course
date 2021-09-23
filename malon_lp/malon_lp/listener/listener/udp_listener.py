from flask import Flask, request
from base64 import b64decode, b64encode

from malon_lp.listener import Listener
from malon_lp.listener.handler.sym import EncryptedHandler
from malon_lp.listener.handler.dh import DHHandler
from malon_lp.listener.handler.api import ApiHandler

from malon_lp.crypto.dh import KeyExchange

import socket

NUMERO_COMPLETAMENTE_ARBITRARIO = 2048

class UdpListener(Listener):
    def __init__(self, api_url: str, listener_id: int, host: str, port: int, sym_key: bytes):
        handler = ApiHandler(api_url, listener_id)
        handler = DHHandler(KeyExchange(), handler)
        handler = EncryptedHandler(sym_key, handler)

        self._handler = handler
        self._port = port
        self._host = host
    
    @classmethod
    def new(cls, api_url: str, listener_id: int, host: str, port: int, sym_key: bytes) -> 'Listener':
        return cls(api_url, listener_id, host, port, sym_key)
    
    @classmethod
    def type_name(cls) -> str:
        return 'udp'
    
    def run(self):
        """
            El listener debera adjuntar un socket que soporte comunicacion UDP a una 
            interfaz IPV4 que recibira por medio del parametro 'host'
            y un puerto, que sera el parametro 'port'.
            Debera recibir mensajes de un tamaño arbitrario.
        """
        rais NotImplementedError


