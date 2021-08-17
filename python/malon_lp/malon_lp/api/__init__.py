from base64 import b64decode
from datetime import datetime
from . import json_encoder
from flask import Flask, jsonify, request
from ..database import db_session, db_init
from ..database.models import Agent, Task, TaskResult
from sqlalchemy import asc

app = Flask(__name__)

app.json_encoder = json_encoder.SQLAlchemyEncoder

db_init()


@app.route("/agents/", methods=['GET'])
def agents():
    return jsonify(Agent.query.all())

@app.route("/agents/<id>/report", methods=['POST'])
def agents_ping(id: str):
    agent = Agent.query.get(id)

    if agent is None:
        print('New agent reported in: {}'.format(id))
        agent = Agent(id=id)

    agent.last_seen_at = datetime.now()

    db_session.add(agent)
    db_session.commit()

    return jsonify(agent)

@app.route("/agents/<id>", methods=['GET'])
def agent(id: str):
    agent = Agent.query.get(id)
    return jsonify(agent)

@app.route("/agents/<agent_id>/tasks/", methods=['GET'])
def agent_tasks(agent_id: str):
    tasks = Task.query.filter_by(agent_id=agent_id).all()
    return jsonify(tasks)

@app.route("/agents/<agent_id>/tasks/unread/", methods=['GET'])
def agent_tasks_unread(agent_id: str):
    agent = Agent.query.get(agent_id)
    tasks = Task.query \
                .filter_by(agent_id=agent_id) \
                .filter(Task.id > agent.last_task_id) \
                .order_by(asc(Task.id)) \
                .all()

    if len(tasks) > 0:
        agent.last_task_id = tasks[-1].id
        db_session.add(agent)
        db_session.commit()

    return jsonify(tasks)

@app.route("/agents/<agent_id>/tasks/", methods=['POST'])
def agent_tasks_post(agent_id: str):
    task_d = request.json
    task = Task(
        agent_id=agent_id,
        type=task_d.get('type'),
        info=task_d.get('info'),
        input=b64decode(task_d.get('input')) if 'input' in task_d else None
    )
    db_session.add(task)
    db_session.commit()
    return jsonify(task)

@app.route("/agents/<agent_id>/tasks/<int:task_id>/result", methods=['GET'])
def agent_task_result(agent_id: str, task_id: int):
    task_result = TaskResult.query.filter_by(task_id=task_id).all()
    return jsonify(task_result)

@app.route("/agents/<agent_id>/tasks/<int:task_id>/result", methods=['POST'])
def agent_task_result_post(agent_id: str, task_id: int):
    task_result_d = request.json
    task_result = TaskResult(
        task_id=task_id,
        info=task_result_d.get('info'),
        output=b64decode(task_result_d.get('output')) if 'output' in task_result_d else None
    )
    db_session.add(task_result)
    db_session.commit()
    return jsonify(task_result)

@app.teardown_appcontext
def shutdown_session(exception=None):
    db_session.remove()