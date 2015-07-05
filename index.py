from bottle import route, run, debug

@route('/score/<name>/<hole>/<score>', method='GET')
def score(name, hole, score):
  print 'name: {}, hole: {}, score: {}'.format(name, hole, score)

@route('/scores', method='GET')
def scores():
  return '''
{ "players": [
    { "id": 1,
      "name": "mark",
      "scores": [4,4,4]},
    { "id": 2,
      "name": "ric",
      "scores": [4,3,4,3,4,4]}
    ]
}
'''


debug(True)
run(reloader=True)

