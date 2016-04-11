package auth

import (
  "database/sql"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
)

type Auth interface {
  GetID(name string, password string) LoginResponse
}

type Sqlite struct {
}

func NewAuth() Auth {
  return Sqlite{}
}

type LoginResponse struct {
  ID int
  Msg string
}


func (s Sqlite)GetID(name string, password string) LoginResponse {

  db, err := sql.Open("sqlite3", "./data/catfish_db")
  if err != nil {
    return LoginResponse{ID: -1, Msg: fmt.Sprintf("%v", err)}
  }
  defer db.Close()
  stmt, err := db.Prepare("select id from login where name = ? and password = ?")
  if err != nil {
    return LoginResponse{ID: -1, Msg: fmt.Sprintf("%v", err)}
  }
  defer stmt.Close()

  var id int
  err = stmt.QueryRow(name, password).Scan(&id)
  if err != nil {
    return LoginResponse{ID: -1, Msg: fmt.Sprintf("%v", err)}
  }

  return LoginResponse{ID: id,
                       Msg: ""}

}

