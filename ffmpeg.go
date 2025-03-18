package main

import (
	"os"
	"os/exec"
	"strconv"
	"bytes"
	"log"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	_ "github.com/joho/godotenv/autoload"
)

/* -- subtitle format --
43
00:00:02,462 --> 00:00:02,493
Evan told me
you didn't get into Dartmouth.
*/

func runVideo() {
	cmd := exec.Command("vlc", "out.mp4")
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

	// subtract 1 seconds for better clipping
	if seconds >= 1 {
		seconds -= 1 
	}

	_ = ffmpeg.
		Input(os.Getenv("TRAINING_DAY"), ffmpeg.KwArgs{"ss": seconds}). // -ss <time_off> start transcoding at specified time
		Output("out.mp4", ffmpeg.KwArgs{"t": 5}). // -t <duration> stop transcoding after specified duration 
		OverWriteOutput().
		Run()
}

func timestampToSeconds(timestamp string) int64 {
//00:00:02,462 --> 00:00:02,493
	hours, _ := strconv.ParseInt(timestamp[0:2], 10, 32)
	minutes, _ := strconv.ParseInt(timestamp[3:5], 10, 32)
	seconds, _ := strconv.ParseInt(timestamp[6:8], 10, 32)

	return ((hours*3600) + (minutes*60) + seconds)
}
