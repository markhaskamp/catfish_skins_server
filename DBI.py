import sqlite3

class DBI:

  def updateScore(self, id, hole, score):
    conn = sqlite3.connect('data/catfish.db')
    c = conn.cursor()

    t = (score, id, hole)
    c.execute('UPDATE scores SET score=? WHERE playerId=? AND hole=?;', t)
    conn.commit()
    conn.close()

  def getAllScores(self):
    conn = sqlite3.connect('data/catfish.db')
    c = conn.cursor()
    c.execute('''
    SELECT p.name, s.playerId, s.hole, s.score
    FROM scores s
    JOIN players p on p.id = s.playerId
    ''')
    scores = c.fetchall()

    h = {}
    for score in scores:
      key = str(score[0])
      if not key in h:
        h[key] = []
        h[key].append(score[1])

      h[key].append(score[3])

    return h

