package day1

import (
    "fmt"
    "log"
    "strconv"
)

func find(slice []string, val string) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}

func toInt(item string) int {
    int_item, err := strconv.Atoi(item)
    if err != nil {
        log.Fatal(err)
    }
    return int_item
}

func toString(item int) string {
    return strconv.Itoa(item)
}

func remove(list []string, index int) []string {
    return append(list[:index], list[index+1:]...)
}

func Run(entries []string) {
    for index, first := range entries {
        for _, second := range remove(entries, index) {
            potential_third := 2020 - (toInt(first) + toInt(second))
            found := find(entries, toString(potential_third))
            if found {
                fmt.Println(toInt(first) * toInt(second) * potential_third)
                return
            }
        }
    }
}
