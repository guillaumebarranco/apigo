package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "reflect"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/davecgh/go-spew/spew"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
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

    user := "root"
    password := ""
    host := "127.0.0.1:3306"
    database := "cinema"

    con, err := sql.Open("mysql", user+":"+password+"@("+host+")"+database)
    defer con.Close()

    rows, err := con.Query("select * from films")

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

// func debugType(element interface) {
//     spew.Dump(reflect.TypeOf(element))
// }

func getDbUtil() *sql.DB {

    user := "root"
    password := ""
    host := ""
    database := "webarranco8"

    con, err := sql.Open("mysql", user+":"+password+"@"+host+"/"+database)

    if err != nil {  }

    return con
}

// func getItemsFromRows(w http.ResponseWriter, rows sql.Rows) interface {

//     var element string
//     items := make([]*Project, 0, 10)

//     for rows.Next() {
//         err := rows.Scan(&element)
//         if err != nil { spew.Dump(err) }

//         items = append(items, &Project{Name: element})
//     }

//     jsonItems := json.NewEncoder(w).Encode(items)

//     spew.Dump(reflect.TypeOf(jsonItems))

//     return jsonItems
// }

func ProjectsShow(w http.ResponseWriter, r * http.Request) {

    con := getDbUtil()
    defer con.Close()

    rows, err := con.Query("select name from project")

    if err != nil { spew.Dump(err) }

    // items := getItemsFromRows(w, rows)

    items := make([]*Project, 0, 10)

    var element string

    for rows.Next() {
        err := rows.Scan(&element)
        if err != nil { spew.Dump(err) }

        items = append(items, &Project{Name: element})
    }

    jsonItems := json.NewEncoder(w).Encode(items)

    // debugType(jsonItems)
    spew.Dump(reflect.TypeOf(jsonItems))

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := jsonItems; err != nil {
        panic(err)
    }
}
