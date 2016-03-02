package main

import (
  "data"
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "github.com/codegangsta/negroni"
)


func main() {

  mux := http.NewServeMux()

  allScores := data.NewAllScores()


  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    j,_ := json.Marshal(allScores)
    fmt.Fprintf(w, string(j))
  })

  mux.HandleFunc("/score", func(w http.ResponseWriter, req *http.Request) {
    req.ParseForm()

    hole,_    := strconv.Atoi(req.FormValue("hole"))
    strokes,_ := strconv.Atoi(req.FormValue("strokes"))
    golfer    := req.FormValue("golfer")

    allScores.Update(data.ScoreEntry{Golfer:golfer,
                                Hole:hole,
                                Strokes:strokes})

    j,_ := json.Marshal(allScores)
    fmt.Fprintf(w, string(j))
  })

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":8080")
}


