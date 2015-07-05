from bottle import route, run, debug

@route('/score/<name>/<hole>/<score>', method='GET')
def score(name, hole, score):
  print 'name: {}, hole: {}, score: {}'.format(name, hole, score)

@route('/scores', method='GET')
def scores():
  return 'eddie would go'


debug(True)
run(reloader=True)

