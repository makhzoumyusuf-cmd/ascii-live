package frames

import (
	"math/rand"
	"strings"
	"time"
)

var Matrix = DefaultFrameType(generateMatrixFrames())

func generateMatrixFrames() []string {
	rand.Seed(time.Now().UnixNano())
	ROWS, COLS := 40, 120       // bigger default size
	numFrames := 50              // more frames for smoothness
	frames := []string{}

	// each column has a drop position
	drops := make([]int, COLS)
	for i := 0; i < COLS; i++ {
		drops[i] = rand.Intn(ROWS)
	}

	for f := 0; f < numFrames; f++ {
		lines := make([]string, ROWS)
		for r := 0; r < ROWS; r++ {
			line := ""
			for c := 0; c < COLS; c++ {
				trailLength := 6
				if r >= drops[c]-trailLength && r <= drops[c] {
					line += string(randomChar())
				} else {
					line += " "
				}
			}
			lines[r] = line
		}
		frames = append(frames, strings.Join(lines, "\n"))

		// move drops down
		for c := 0; c < COLS; c++ {
			if rand.Float32() > 0.02 {
				drops[c]++
				if drops[c] >= ROWS {
					drops[c] = 0
				}
			} else {
				drops[c] = 0
			}
		}
	}

	return frames
}

func randomChar() rune {
	chars := "01ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%^&*"
	return rune(chars[rand.Intn(len(chars))])
}
