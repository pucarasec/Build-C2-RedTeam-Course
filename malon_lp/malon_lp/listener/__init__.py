from typing import Mapping
from abc import ABC, abstractmethod
import pkgutil
import os


def _load_submodules():
    paths = [os.path.join(path, 'listener') for path in __path__]
    for loader, module_name, is_pkg in  pkgutil.walk_packages(paths):
        _module = loader.find_module(module_name).load_module(module_name)

class Listener(ABC):
    @classmethod
    @abstractmethod
    def new(cls, api_url: str, listener_id: int, host: str, port: int, sym_key: bytes) -> 'Listener':
        """
            Factory
        """
        pass
    
    @classmethod
    @abstractmethod
    def type_name(cls) -> str:
        """
            Devuelve un valor caracteristico de tipo String para identificar
            las caracteristicas del listener
        """
        pass
    
    @abstractmethod
    def run(self):
        """
            Ejecuta el Listener 
        """
        pass

def get_listener_types() -> Mapping[str, type]:
    return {
        listener_type.type_name(): listener_type
        for listener_type in Listener.__subclasses__()
    }

_load_submodules()
