package main

import (
	"os"
	"os/exec"
	"bytes"
	"log"
	"bufio"
	"fmt"
	"strings"
	"strconv"
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

// -to <time_stop> stop transcoding after specified time is reached
func clipVideo(latestTimestamp string) {
	seconds := timestampToSeconds(latestTimestamp[0:8])
	fmt.Println(seconds)
	_ = ffmpeg.
		Input(os.Getenv("TRAINING_DAY"), ffmpeg.KwArgs{"ss": seconds}). // -ss <time_off> start transcoding at specified time
		Output("out2.mp4", ffmpeg.KwArgs{"t": 3}). // -t <duration> stop transcoding after specified duration 
		OverWriteOutput().
		Run()
}

func timestampToSeconds(timestamp string) int64 {
//00:00:02,462 --> 00:00:02,493
	hours, _ := strconv.ParseInt(timestamp[0:2], 10, 32)
	minutes, _ := strconv.ParseInt(timestamp[3:5], 10, 32)
	seconds, _ := strconv.ParseInt(timestamp[6:8], 10, 32)
	fmt.Printf("%d, %d, %d\n", hours, minutes, seconds)

	return ((hours*3600) + (minutes*60) + seconds)
}
