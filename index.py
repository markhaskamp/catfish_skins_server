from bottle import hook, response, route, run, debug
import json

@hook('after_request')
def enable_cors():
  response.headers['Access-Control-Allow-Origin'] = '*'

@route('/score/<name>/<hole>/<score>', method='GET')
def score(name, hole, score):
  print 'name: {}, hole: {}, score: {}'.format(name, hole, score)

@route('/scores', method='GET')
def scores():
    return ( {
"players": [
    {"id": 1,
    "name": "mark",
    "scores": [1,2,3,4]},
    {"id": 2,
     "name": "ric",
     "scores": [5,6,7,8]}
    ]})



debug(True)
run(reloader=True)

