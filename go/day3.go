package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	err := run(os.Stdout, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type slope struct {
	nextx int
	nexty int
}

type Point struct {
	x     int
	y     int
	slope slope
	trees int
}

var (
	width int = 0
)

func (p *Point) NextMove() {
	// right 3, down 1
	p.x = (p.x + p.slope.nextx) % width
	p.y = p.y + p.slope.nexty
	// fmt.Printf("\tNextMove: %v\n", p)
}

func run(w io.Writer, args []string) error {
	// if len(args) != 0 {
	// 	return errors.New("filename is hardcoded")
	// }
	file, err := os.Open("../day3/input")
	if err != nil {
		return err
	}
	defer file.Close()

	lineNo := 0
	routes := []*Point{
		&Point{x: 0, y: lineNo, slope: slope{nextx: 1, nexty: 1}},
		&Point{x: 0, y: lineNo, slope: slope{nextx: 3, nexty: 1}},
		&Point{x: 0, y: lineNo, slope: slope{nextx: 5, nexty: 1}},
		&Point{x: 0, y: lineNo, slope: slope{nextx: 7, nexty: 1}},
		&Point{x: 0, y: lineNo, slope: slope{nextx: 1, nexty: 2}},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		plots := []rune(scanner.Text())
		width = len(plots)
		for _, p := range routes {
			pos := p
			fmt.Printf("pos: %v\n", pos)
			if pos.y != lineNo {
				continue
			}
			if plots[pos.x] == '#' {
				pos.trees++

			}
			// get ready for next
			pos.NextMove()
		}
		lineNo++
	}

	result := 1
	for i, pos := range routes {
		fmt.Printf("route[%d]: trees = %d\n", i, pos)
		result = result * pos.trees
	}
	fmt.Printf("trees: %d\n", result)
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
