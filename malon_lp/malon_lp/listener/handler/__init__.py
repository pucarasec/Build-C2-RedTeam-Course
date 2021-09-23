from typing import Optional
from abc import ABC, abstractmethod

class Handler(ABC):
    @abstractmethod
    def handle_msg(self, msg: bytes):
        """
            En cada una de las requests que 
            realiza el agente, se invoca este metodo.
            Es necesario en este lugar decidir el 
            tipo de mensaje que se esta recibiendo
            para definir si es necesario consultar 
            por tareas nuevas que deba realizar el agente
            o registrar el resultado
        """
        pass

class AuthenticatedHandler(ABC):
    @abstractmethod
    def handle_auth_msg(self, msg: bytes, client_id: str):      
        """
            Posee la  misma logica correspondiente 
            al manejo de mensajes, pero con la 
            funcionalidad adicional de poder identificar 
            al agente que se encuentra
            interactuando con el listener
        """
        pass
