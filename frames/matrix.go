package frames

import (
	"math/rand"
	"strings"
	"time"
)

// DynamicMatrix generates a continuous Matrix rain
var Matrix = DefaultFrameType(func() []string {
	return []string{generateFrame()}
})

var ROWS, COLS = 40, 120
var drops []int

func init() {
	rand.Seed(time.Now().UnixNano())
	drops = make([]int, COLS)
	for i := 0; i < COLS; i++ {
		drops[i] = rand.Intn(ROWS)
	}
}

// generateFrame returns a single dynamic frame
func generateFrame() string {
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

	return strings.Join(lines, "\n")
}

// randomChar returns a single Matrix character
func randomChar() rune {
	chars := "01ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%^&*"
	return rune(chars[rand.Intn(len(chars))])
}

