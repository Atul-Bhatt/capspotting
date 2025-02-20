package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"regexp"
)

/* -- subtitle format --
43
00:00:02,462 --> 00:00:02,493
Evan told me
you didn't get into Dartmouth.
*/

func main() {
	// regular expression to match timestamp
	timestampRegex, _ := regexp.Compile("\\d\\d:\\d\\d:\\d\\d,\\d\\d\\d --> \\d\\d:\\d\\d:\\d\\d,\\d\\d\\d") // looks horrible, change later
	latestTimestamp := ""

	file, err := os.Open("superbad.srt")
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
				break
			}
		}
	}
}
