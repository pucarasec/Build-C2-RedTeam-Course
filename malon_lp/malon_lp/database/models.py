from datetime import datetime

from sqlalchemy import ForeignKey, Column, text
from sqlalchemy.orm import declarative_base, relationship, validates
from sqlalchemy.types import DateTime, String, Integer, JSON, BLOB

Base = declarative_base()

class Agent(Base):
    __tablename__ = 'agents'
    id = Column(String, primary_key=True)
    created_at = Column(DateTime, nullable=False, default=datetime.now)
    last_seen_at = Column(DateTime, nullable=False, default=datetime.now)
    last_task_id = Column(Integer, nullable=False, default=0)
    tasks = relationship('Task', back_populates='agent')

class Task(Base):
    __tablename__ = 'tasks'
    id = Column(Integer, primary_key=True, autoincrement=True)
    created_at = Column(DateTime, nullable=False, default=datetime.now)
    agent_id = Column(String, ForeignKey('agents.id'))
    agent = relationship('Agent', back_populates='tasks')
    type = Column(String, nullable=False)
    info = Column(JSON)
    input = Column(BLOB)
    result = relationship('TaskResult', back_populates='task')

class TaskResult(Base):
    __tablename__ = 'task_results'
    task_id = Column(Integer, ForeignKey('tasks.id'), primary_key=True)
    created_at = Column(DateTime, nullable=False, default=datetime.now)
    task = relationship('Task', back_populates='result')
    info = Column(JSON)
    output = Column(BLOB)

class Listener(Base):
    __tablename__ = 'listeners'
    id = Column(Integer, primary_key=True, autoincrement=True)
    created_at = Column(DateTime, nullable=False, default=datetime.now)
    type = Column(String, nullable=False)
    bind_host = Column(String, nullable=False, default="0.0.0.0")
    bind_port = Column(Integer, nullable=False, unique=True)
    target_host = Column(String, nullable=False)
    target_port = Column(Integer, nullable=False)
    connection_interval_ms = Column(Integer, nullable=False, default=1000)
    sym_key = Column(BLOB, nullable=False)

    @validates('sym_key')
    def validate_sym_key(self, _, sym_key):
        assert len(sym_key) == 16
        return sym_key
