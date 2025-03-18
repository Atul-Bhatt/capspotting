package main

import (
	"os"
	"bufio"
	"strings"
	"regexp"
	"fmt"
	"log"
)

func scanCaption() {
	// regular expression to match timestamp
	timestampRegex, _ := regexp.Compile("\\d{2}:\\d{2}:\\d{2},\\d{3} --> \\d{2}:\\d{2}:\\d{2},\\d{3}")
	latestTimestamp := ""

	file, err := os.Open(os.Getenv("CAPTION_FILE"))
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	
	// Read user input for caption
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		file.Seek(0, 0)
		lines := bufio.NewScanner(file)
		for lines.Scan() {
			if timestampRegex.MatchString(lines.Text()) {
				latestTimestamp = lines.Text()
				continue
			}
			if strings.Contains(lines.Text(), input.Text()) {
				fmt.Println(latestTimestamp)
				fmt.Println(lines.Text())
				clipVideo(latestTimestamp)
				runVideo()
				break
			}
		}
	}
}
