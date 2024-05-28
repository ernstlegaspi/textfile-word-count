package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan uint8)

	fileName := "file.txt"
	file, err := os.Open(fileName)
	var counter uint8

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	wg.Add(1)
	go func() {

		for scanner.Scan() {
			var text string = scanner.Text()

			if text == "" {
				continue
			}

			counter += uint8(len(strings.Split(strings.TrimSpace(text), " ")))
		}

		ch <- counter
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalWords := <-ch

	fmt.Println("Total words: ", totalWords)
}
