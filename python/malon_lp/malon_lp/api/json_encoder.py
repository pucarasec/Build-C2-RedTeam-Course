import json
from datetime import datetime
from base64 import b64encode
from sqlalchemy.ext.declarative import DeclarativeMeta


def to_camel_case(s: str) -> str:
    return ''.join(map(str.capitalize, s.split('_')))


class BaseEncoder(json.JSONEncoder):
    def default(self, obj):
        if type(obj) == datetime:
            return obj.isoformat()
        elif type(obj) == bytes:
            return b64encode(obj).decode('utf-8')
        else:
            return super().default(obj)


class SQLAlchemyEncoder(BaseEncoder):
    def default(self, obj):
        if isinstance(obj.__class__, DeclarativeMeta):
            fields = {}
            for column in obj.__class__.__table__.columns:
                data = obj.__getattribute__(column.name)
                field = column.name
                try:
                    json.dumps(data, cls=SQLAlchemyEncoder)
                    fields[field] = data
                except TypeError:
                    fields[field] = None

            return fields
        else:
            return super().default(obj)
