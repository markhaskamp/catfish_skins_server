package data

type AllScores struct {
  Scores map[string][]int
}


func NewAllScores() AllScores {
  s := make(map[string][]int)
  return AllScores{Scores: s}
}

func ResetAllScores() AllScores {
  return NewAllScores()
}

type ScoreEntry struct {
  Golfer  string
  Hole    int
  Strokes int
}


func (as *AllScores)Update(score ScoreEntry) {
  scores := as.Scores[score.Golfer]
  if len(scores) == 0 {
    as.Scores[score.Golfer] = []int{0,0,0,0,0,0,0,0,0}    
  }

  var hole int
  if score.Hole > 9 {
    hole = score.Hole - 10
  } else {
    hole = score.Hole - 1
  }
    
  as.Scores[score.Golfer][hole] = score.Strokes
}


