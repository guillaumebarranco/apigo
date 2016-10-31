package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func MangasShow(w http.ResponseWriter, r * http.Request) {

    mangas := Mangas{
        Manga{Name: "One Piece"},
        Manga{Name: "Naruto"},
    }

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(mangas); err != nil {
        panic(err)
    }
}

func UsersShow(w http.ResponseWriter, r * http.Request) {

    con, err := sql.Open("mysql", "root"+":"+""+"@/"+"lg")
    defer con.Close()

    rows, err := con.Query("select name from roles")

    if err != nil { /* error handling */}
    items := make([]*User, 0, 10)
    var ida string

    for rows.Next() {
        err = rows.Scan(&ida)
        if err != nil { /* error handling */}
        items = append(items, &User{Name: ida})
    }

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(items); err != nil {
        panic(err)
    }
}
