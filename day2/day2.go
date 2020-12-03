package day2

import (
    "fmt"
    "log"
    "regexp"
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

func parse_policy(policy_and_password string) (int, int, string, string) {
    var regex = regexp.MustCompile(`(\d+)-(\d+) ([a-zA-Z]): ([a-zA-Z]+)`)
    res := regex.FindStringSubmatch(policy_and_password)
    return toInt(res[1]) - 1, toInt(res[2]) - 1, res[3], res[4]
}

func Run(entries []string) {
    valid_passwords := 0
    for _, policy_and_password := range entries {
        first, second, char, password := parse_policy(policy_and_password)
        if (((string(password[first]) == char && string(password[second]) != char) || (string(password[first]) != char && string(password[second]) == char)) && (string(password[first]) != string(password[second]))) {
            valid_passwords++
        }
    }
    fmt.Println(valid_passwords)
}
