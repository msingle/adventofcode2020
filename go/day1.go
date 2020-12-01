package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	err := run(os.Stdout, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getTwo(nums []int) (int, int) {

	// sort.Ints(nums)
	// x+ y == 2020
	fmt.Println("checking: ")
	for i, x := range nums {
		fmt.Printf("%d\n", x)
		for j := i; j < len(nums); j++ {
			y := nums[j]
			fmt.Printf("\t%d\n", y)

			if x+y == 2020 {
				return x, y
			}
		}
	}
	return 0, 0
}

func getThree(nums []int) (int, int, int) {

	// sort.Ints(nums)
	// x+ y == 2020
	fmt.Println("checking: ")
	for i, x := range nums {
		fmt.Printf("%d\n", x)
		for j := i; j < len(nums); j++ {
			y := nums[j]
			fmt.Printf("\t%d\n", y)

			for k := j; k < len(nums); k++ {
				z := nums[k]
				if x+y+z == 2020 {
					return x, y, z
				}
			}
			// if x+y == 2020 {
			// 	return x, y
			// }
		}
	}
	return 0, 0, 0
}

func run(w io.Writer, args []string) error {
	// if len(args) != 0 {
	// 	return errors.New("filename is hardcoded")
	// }
	file, err := os.Open("../day1/input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	nums := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		nums = append(nums, num)
	}

	// x, y := getTwo(nums)
	// fmt.Printf("x = %d\n", x)
	// fmt.Printf("y = %d\n", y)

	// result := x * y
	// fmt.Printf("x*y = %d\n", result)

	x, y, z := getThree(nums)
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)
	fmt.Printf("z = %d\n", z)

	result := x * y * z
	fmt.Printf("x*y*z= %d\n", result)

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
