package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "github.com/codegangsta/negroni"
)

type ScoreEntry struct {
  Golfer  string
  Hole    int
  Strokes int
}

type AllScores struct {
  Scores map[string][]int
}


func NewAllScores() AllScores {
  s := make(map[string][]int)
  return AllScores{Scores: s}
}


func main() {

  mux := http.NewServeMux()

  allScores := NewAllScores()


  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    j,_ := json.Marshal(allScores)
    fmt.Fprintf(w, string(j))
  })

  mux.HandleFunc("/score", func(w http.ResponseWriter, req *http.Request) {
    req.ParseForm()

    hole,_    := strconv.Atoi(req.FormValue("hole"))
    strokes,_ := strconv.Atoi(req.FormValue("strokes"))
    golfer    := req.FormValue("golfer")

    allScores.Update(ScoreEntry{Golfer:golfer, Hole:hole, Strokes:strokes})

    j,_ := json.Marshal(allScores)
    fmt.Fprintf(w, string(j))
  })

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":8080")
}


func (as *AllScores)Update(score ScoreEntry) {
  scores := as.Scores[score.Golfer]
  if len(scores) == 0 {
    as.Scores[score.Golfer] = []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}    
  }
  as.Scores[score.Golfer][score.Hole] = score.Strokes
}

