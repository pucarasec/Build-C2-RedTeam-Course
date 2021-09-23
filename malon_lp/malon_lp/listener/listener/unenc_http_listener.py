from flask import Flask, request
from base64 import b64decode, b64encode

from malon_lp.listener import Listener
from malon_lp.listener.handler.dummy_auth import DummyAuthenticationHandler
from malon_lp.listener.handler.api import ApiHandler

from malon_lp.crypto.dh import KeyExchange

class UnencryptedHttpListener(Listener):
    def __init__(self, api_url: str, listener_id: int, host: str, port: int):
        handler = ApiHandler(api_url, listener_id)
        handler = DummyAuthenticationHandler(handler)

        self._handler = handler
        self._host = host
        self._port = port
    
    @classmethod
    def new(cls, api_url: str, listener_id: int, host: str, port: int, _sym_key: bytes) -> 'Listener':
        return cls(api_url, listener_id, host, port)
    
    @classmethod
    def type_name(cls) -> str:
        return 'unenc-http'
    
    def run(self):
        """
            - Es necesario crear una aplicacion de flask la cual exponga unicamente el endpoint "/"
            y reciba el metodo "POST".
            - El mismo recibira un mensaje encodeado en base64 dentro de un form bajo la key "m" 
            - Este mensaje debe ser desencodeado para poder ser procesado por los handlers del Listener.
            - Finalmente debera devolver la respuesta recibida por los handlers como un string encodeado
            en base64, dentro de un comentario de HTML
            - Al terminar la funcion se debe correr la aplicacion en el puerto y host definidos por el Listener
        """ 
        raise NotImplementedError