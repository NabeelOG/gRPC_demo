package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cleanTranscript() {
	cc, err := os.Open("cc.txt")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	defer cc.Close()

	scanner := bufio.NewScanner(cc)
	var lines []string
	even := true
	for scanner.Scan() {
		even = !even
		if even {
			lines = append(lines, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := strings.Join(lines, " ")
	fmt.Println(result)
}
