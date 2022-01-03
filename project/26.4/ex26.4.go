package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	argsWord := os.Args[1]
	argsFile := os.Args[2]

	data, err := os.Open(argsFile)
	check(err)


	scanner := bufio.NewScanner(data)

	scanner.Split(bufio.ScanLines)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), argsWord) {
			fmt.Println(line, scanner.Text())
		}
		line++
	}

}