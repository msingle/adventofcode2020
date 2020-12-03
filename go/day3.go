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

type Point struct {
	x int
	y int
}

var (
	width int = 0
)

func (p *Point) NextMove() {
	// right 3, down 1
	p.x = (p.x + 3) % width
	p.y = p.y + 1
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
	pos := Point{x: 0, y: lineNo}

	trees := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNo++
		plots := []rune(scanner.Text())
		width = len(plots)
		if plots[pos.x] == '#' {
			trees++
		}
		// get ready for next
		pos.NextMove()
	}

	fmt.Printf("trees: %d\n", trees)
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
