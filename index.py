from bottle import hook, response, route, run, debug
from DBI import *
import json

@hook('after_request')
def enable_cors():
  response.headers['Access-Control-Allow-Origin'] = '*'

@route('/score/<id>/<hole>/<score>', method='GET')
def score(id, hole, score):
  dbi = DBI()
  dbi.updateScore(id, hole, score)

@route('/scores', method='GET')
def scores():
  dbi = DBI()
  return dbi.getAllScores()


debug(True)
run(reloader=True)

