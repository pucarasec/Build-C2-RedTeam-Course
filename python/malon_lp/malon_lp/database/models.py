from datetime import datetime

from sqlalchemy import ForeignKey, Column, text
from sqlalchemy.orm import declarative_base, relationship
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
