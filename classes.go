package main

import (
	"time"
)

/******************/
/*      Manga     */
/******************/

type Manga struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Mangas []Manga

/******************/
/*      User      */
/******************/

type User struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Users []User

/******************/
/*    Project     */
/******************/

type Project struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Projects []Project
