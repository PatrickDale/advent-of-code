package main

import (
    "day2"
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
    day2.Run(readEntries())
}