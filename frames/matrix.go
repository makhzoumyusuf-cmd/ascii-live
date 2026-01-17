package frames

import (
	"math/rand"
	"strings"
	"time"
)

// Matrix animation generator
var Matrix = DefaultFrameType(generateMatrixFrames())

func generateMatrixFrames() []string {
	rand.Seed(time.Now().UnixNano())
	frames := []string{}

	ROWS, COLS := 24, 80       // default size
	numFrames := 30             // smooth animation

	// Initialize drops: each column has a vertical position
	drops := make([]int, COLS)
	for i := 0; i < COLS; i++ {
		drops[i] = rand.Intn(ROWS)
	}

	for f := 0; f < numFrames; f++ {
		lines := make([]string, ROWS)

		for r := 0; r < ROWS; r++ {
			line := ""
			for c := 0; c < COLS; c++ {
				// Only draw a character if this row is within the drop trail
				trailLength := 5
				if r >= drops[c]-trailLength && r <= drops[c] {
					line += string(randomChar())
				} else {
					line += " " // real space
				}
			}
			lines[r] = line
		}

		frames = append(frames, strings.Join(lines, "\n"))

		// Move drops down
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

