package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := run(os.Stdout, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(w io.Writer, args []string) error {
	// if len(args) != 0 {
	// 	return errors.New("filename is hardcoded")
	// }
	file, err := os.Open("../day2/input")
	if err != nil {
		return err
	}
	defer file.Close()

	oks := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, " ")
		lh := strings.Split(chunks[0], "-")
		lower, _ := strconv.Atoi(lh[0])
		higher, _ := strconv.Atoi(lh[1])
		target := strings.Trim(chunks[1], ":")
		letters := []rune(chunks[2])
		if (string(letters[lower-1]) == target) != (string(letters[higher-1]) == target) {
			fmt.Printf("%s\n", line)
			oks++
		}

	}

	fmt.Printf("oks: %d\n", oks)

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
