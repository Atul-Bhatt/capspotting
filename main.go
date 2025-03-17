package main

import (
	"os"
	"os/exec"
	"bytes"
	"log"
	"bufio"
	"fmt"
	"strings"
	"regexp"
	_ "github.com/joho/godotenv/autoload"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

/* -- subtitle format --
43
00:00:02,462 --> 00:00:02,493
Evan told me
you didn't get into Dartmouth.
*/

func main() {
	//runVideo()

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
				break
			}
		}
	}
}

func runVideo() {
	cmd := exec.Command("vlc", os.Getenv("TRAINING_DAY"))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
}

func clipVideo(latestTimestamp string) {
	_ = ffmpeg.
		Input(os.Getenv("TRAINING_DAY"), ffmpeg.KwArgs{"ss": 1}).
		Output("out1.mp4", ffmpeg.KwArgs{"t": 1}).
		OverWriteOutput().
		Run()
}
