package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
)

/* -- subtitle format --
43
00:00:02,462 --> 00:00:02,493
Evan told me
you didn't get into Dartmouth.
*/

func main() {
	file, err := os.Open("superbad.srt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	
	// Read user input for caption
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lines := bufio.NewScanner(file)
		for lines.Scan() {
			if strings.Contains(lines.Text(), input.Text()) {
				fmt.Println(lines.Text())
				break
			}
		}
	}
}
