package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

const wpm = 250

func main() {
	words := strings.Split(strings.Join(readText(), string(' ')), string(' '))
	fmt.Println(averageTimeToRead(len(words)), "seconds")
}

func readText() []string {
	var text []string

	file, err := os.Open("text")

	if isError(err) {
		return text
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		text = append(text, line)

		if err == io.EOF {
			break
		}

	}
	file.Close()
	log.Println(text, reflect.TypeOf(text), len(text))
	return text
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func averageTimeToRead(w int) int {
	time := (w * 60) / wpm
	return time
}
