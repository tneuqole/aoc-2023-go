package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func p1() {
    file, err := os.Open("day4.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sum float64
    for scanner.Scan() {
        s := scanner.Text()
        raw := strings.Split(s, ":")
        raw = strings.Split(raw[1], "|")
        
        winningNums := strings.Fields(raw[0])
        card := strings.Fields(raw[1])
        
        count := -1.0
        for _, n := range winningNums {
            if slices.Contains(card, n) {
                 count++
            }
        }

        if count > -1 {
            sum += math.Pow(2, count)
        }
    }

    fmt.Println("problem 1: ", sum)
}

type card struct {
    winningNums []string
    nums []string
}

var cards []card

func parseCards() {
    file, err := os.Open("day4.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s := scanner.Text()
        raw := strings.Split(s, ":")
        raw = strings.Split(raw[1], "|")

        c := card {
            winningNums: strings.Fields(raw[0]),
            nums: strings.Fields(raw[1]),
        }
        cards = append(cards, c)
    }
}

func countCopies(idx int) int {
    copies := 1
    c := cards[idx]
    fmt.Println("Processing card: ", idx+1)
    count := 0
    for _, n := range c.winningNums {
        if slices.Contains(c.nums, n) {
            count++
        }
    }

    for j := 0; j < count; j++ {
        copies += countCopies(j+1+idx)
    }
    return copies
}

func p2() {
    parseCards()
    var copies int
    for i := 0; i < len(cards); i++ {
        copies += countCopies(i)
    }

    fmt.Println("problem 2: ", copies)
}

func main() {
    p1()
    p2()
}

