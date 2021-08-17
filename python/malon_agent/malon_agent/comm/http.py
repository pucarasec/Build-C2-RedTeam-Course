import re
import requests
from typing import Optional
from base64 import b64encode, b64decode

BASE64_MATCH_REGEX = re.compile(b'(<!--)([A-Za-z0-9/+=]*|=[^=]|={3,})(-->)')


def find_encoded_msg(body: bytes) -> Optional[bytes]:
    match = BASE64_MATCH_REGEX.search(body)
    return match[0] if match is not None else None


class HttpClient:
    def __init__(self, url: str):
        self._url = url

    def send_msg(self, msg: bytes) -> bytes:
        response = requests.post(self._url, {'m': b64encode(msg)})
        if not response.ok:
            raise RuntimeError('Got HTTP status code: {}'.format(response.status_code))
        encoded_response_msg = find_encoded_msg(response.content)
        if encoded_response_msg is None:
            raise RuntimeError('No encoded message found in response')

        return b64decode(encoded_response_msg)
