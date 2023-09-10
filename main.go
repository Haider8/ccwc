package main

import (
	"flag"
	"fmt"
	"os"
)

var byte_count_flag = flag.Bool("c", false, "print the byte counts")
var line_count_flag = flag.Bool("l", false, "print the newline counts")
var word_count_flag = flag.Bool("w", false, "print the word counts")

func count_bytes(filename string, ch chan<- string) {
	file_info, err := os.Stat(filename)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}
	ch <- fmt.Sprintf("%d %s", file_info.Size(), filename)
}

func count_lines(filename string, ch chan<- string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}

	var count int64
	for _, d := range data {
		if d == byte('\n') {
			count++
		}
	}

	ch <- fmt.Sprintf("%d %s", count, filename)
}

func count_words(filename string, ch chan<- string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}

	var count int64
	for _, d := range data {
		if d == byte(' ') {
			count++
		}
	}

	ch <- fmt.Sprintf("%d %s", count, filename)
}

func main() {
	flag.Parse()

	ch := make(chan string)

	for _, filename := range flag.Args() {
		if *byte_count_flag {
			go count_bytes(filename, ch)
		} else if *line_count_flag {
			go count_lines(filename, ch)
		} else if *word_count_flag {
			go count_words(filename, ch)
		}
	}

	for range flag.Args() {
		fmt.Println(<-ch)
	}
}
