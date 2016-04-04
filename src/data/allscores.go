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
    as.Scores[score.Golfer] = []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}    
  }
  as.Scores[score.Golfer][score.Hole] = score.Strokes
}


