package main

import "time"

type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Todos []Todo

type Manga struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Mangas []Manga