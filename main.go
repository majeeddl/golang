package main

import (
    "encoding/json"
    "fmt"
    "log"
)

// Avenger represents a single hero
type Avenger struct {
    RealName string `json:"real_name"`
    HeroName string `json:"hero_name"`
    Planet   string `json:"planet"`
    Alive    bool   `json:"alive"`
}

func (a *Avenger) isAlive() {
    a.Alive = true
}

func main() {
    avengers := []Avenger{
        {
            RealName: "Dr. Bruce Banner",
            HeroName: "Hulk",
            Planet:   "Midgard",
        },
        {
            RealName: "Tony Stark",
            HeroName: "Iron Man",
            Planet:   "Midgard",
        },
        {
            RealName: "Thor Odinson",
            HeroName: "Thor",
            Planet:   "Midgard",
        },
    }

    avengers[1].isAlive()

    jsonBytes, err := json.Marshal(avengers)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(string(jsonBytes))
}