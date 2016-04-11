package main

import (
  "data"
  "auth"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
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


  r.HandleFunc("/login", handleLogin)
  r.HandleFunc("/strokes/{golfer}/{hole}/{strokes}", handleStrokes)
  r.HandleFunc("/allstrokes", handleAllStrokes)
  r.HandleFunc("/reset", handleReset)


  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}


type login_struct struct {
  Name string
  Pwd string
}


func handleLogin(w http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    panic(fmt.Sprintf("%v", err))
  }

  var l login_struct
  err = json.Unmarshal(body, &l)
  if err != nil {
    panic(fmt.Sprintf("%v", err))
  }
  log.Println("name: ", l.Name)
  log.Println("pwd: ", l.Pwd)

  a := auth.NewAuth()
  loginResponse := a.GetID(l.Name, l.Pwd)
  j,err := json.Marshal(loginResponse)

  fmt.Fprintf(w, "%v", string(j))
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


func handleReset(w http.ResponseWriter, req *http.Request) {
  allScores = data.ResetAllScores()
}
