package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "MangaIndex",
        "GET",
        "/mangas",
        MangasShow,
    },
    Route{
        "UserIndex",
        "GET",
        "/users",
        UsersShow,
    },
    Route{
        "ProjectIndex",
        "GET",
        "/projects",
        ProjectsShow,
    },
}
