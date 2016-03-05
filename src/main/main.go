package main

import (
  "data"
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
)

var allScores data.AllScores
var r *mux.Router


func main() {


  allScores = data.NewAllScores()
  r = mux.NewRouter()


  r.HandleFunc("/", handleIndex)
  r.HandleFunc("/strokes/{golfer}/{hole}/{strokes}", handleStrokes)
  r.HandleFunc("/allstrokes", handleAllStrokes)


  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}


func handleIndex(w http.ResponseWriter, req *http.Request) {
  j,_ := json.Marshal(allScores)
  fmt.Fprintf(w, string(j))
}


func handleStrokes(w http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)

  hole,_    := strconv.Atoi(vars["hole"])
  strokes,_ := strconv.Atoi(vars["strokes"])

  allScores.Update(data.ScoreEntry{Golfer:  vars["golfer"],
                                   Hole:    hole,
                                   Strokes: strokes})

  j,_ := json.Marshal(allScores)
  fmt.Fprintf(w, string(j))
}


func handleAllStrokes(w http.ResponseWriter, req *http.Request) {
  j,_ := json.Marshal(allScores)
  fmt.Fprintf(w, string(j))
}

