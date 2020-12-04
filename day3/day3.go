package day3

import (
    "fmt"
)

func treesHit(topography []string, rightBy int, downBy int) int {
    treesHit := 0
    xCoordinate := 0
    yCoordinate := 0
    for yCoordinate < len(topography) {
        if (xCoordinate >= len(topography[yCoordinate])) {
            xCoordinate = xCoordinate - len(topography[yCoordinate])
        }
        if (topography[yCoordinate][xCoordinate] == '#') {
            treesHit ++
        }
        xCoordinate += rightBy
        yCoordinate += downBy
    }
    return treesHit
}

func Run(topography []string) {
    slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

    total := 1
    for _, slope := range slopes {
        treesHit := treesHit(topography, slope[0], slope[1])
        total *= treesHit
    }
    fmt.Println(total)
}
