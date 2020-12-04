package main

import (
    "day3"
    "io/ioutil"
    "log"
    "strings"
)

func readEntries() []string {
    data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    return strings.Split(string(data), "\n")
}

func main() {
    day3.Run(readEntries())
}