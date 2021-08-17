import sys

if sys.version_info.major >= 3 and sys.version_info.minor >= 8:
    from typing import Protocol

    class Handler(Protocol):
        def handle_msg(msg: bytes, *args, **kwargs):
            pass
else:
    class Handler:
        def handle_msg(msg: bytes, *args, **kwargs):
            pass
