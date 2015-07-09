from bottle import hook, response, route, run, debug
import json
import sqlite3

@hook('after_request')
def enable_cors():
  response.headers['Access-Control-Allow-Origin'] = '*'

@route('/score/<id>/<hole>/<score>', method='GET')
def score(id, hole, score):
  print 'id: {}, hole: {}, score: {}'.format(id, hole, score)
  conn = sqlite3.connect('data/catfish.db')
  c = conn.cursor()

  t = (id, score)
  col = 'H'+hole
  c.execute('INSERT INTO scores(playerId, ' + col + ')  VALUES(?, ?)', t)
  conn.commit()
  conn.close()

@route('/scores', method='GET')
def scores():
  conn = sqlite3.connect('data/catfish.db')
  c = conn.cursor()
  c.execute('''
  SELECT p.name, s.hole, s.score
  FROM scores s
  JOIN players p on p.id = s.playerId
  ''')
  return json.dumps(c.fetchall())


debug(True)
run(reloader=True)

