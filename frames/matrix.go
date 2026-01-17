package frames

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Matrix animation generator
var Matrix = DefaultFrameType(generateMatrixFrames())

func generateMatrixFrames() []string {
	rand.Seed(time.Now().UnixNano())
	frames := []string{}

	ROWS, COLS := 24, 80       // default size; ascii.live will stretch
	numFrames := 30             // smooth animation

	for f := 0; f < numFrames; f++ {
		lines := make([]string, ROWS)
		for r := 0; r < ROWS; r++ {
			line := ""
			for c := 0; c < COLS; c++ {
				if rand.Float32() < 0.1 {
					// 10% chance of character
					line += string(randomChar())
				} else {
					line += " "
				}
			}
			lines[r] = line
		}
		frames = append(frames, strings.Join(lines, "\n"))
	}

	return frames
}

func randomChar() rune {
	chars := "01ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%^&*"
	return rune(chars[rand.Intn(len(chars))])
}
