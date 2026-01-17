package frames

import (
	"math/rand"
	"strings"
	"time"
)

var Matrix = DefaultFrameType(generateMatrixFrames())

func generateMatrixFrames() []string {
	rand.Seed(time.Now().UnixNano())
	frames := []string{}

	ROWS, COLS := 24, 80
	numFrames := 30

	for f := 0; f < numFrames; f++ {
		lines := make([]string, ROWS)
		for r := 0; r < ROWS; r++ {
			line := ""
			for c := 0; c < COLS; c++ {
				if rand.Float32() < 0.05 { // 5% chance of a character
					line += string(randomChar())
				} else {
					line += " " // THIS is the key: real space
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
