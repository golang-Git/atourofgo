package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	line, _ := strconv.Atoi(os.Args[1])
	count := 0
	filename := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
		if count >= line {
			break
		}
	}
}
