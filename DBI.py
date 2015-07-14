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
    SELECT p.name, s.hole, s.score
    FROM scores s
    JOIN players p on p.id = s.playerId
    ''')
    return c.fetchall()

