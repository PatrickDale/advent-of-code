package day3

import (
    "fmt"
)

func Run(entries []string) {
    treesHit := 0
    xCoordinate := 0
    for _, topographyLine := range entries {
        if (xCoordinate >= len(topographyLine)) {
            xCoordinate = xCoordinate - len(topographyLine)
        }
        if (topographyLine[xCoordinate] == '#') {
            treesHit ++
        }
        xCoordinate += 3
    }
    fmt.Println(treesHit)
}
