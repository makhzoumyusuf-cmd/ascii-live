package frames

import (
	"math/rand"
	"strings"
	"time"
)

// Matrix animation
var Matrix = DefaultFrameType(generateFrames())

var ROWS, COLS = 40, 120

type drop struct {
	pos    int     // current head row
	length int     // length of the trail
	speed  float64 // chance to move this frame
}

// generateFrames returns a big slice of frames to make the animation smooth
func generateFrames() []string {
	rand.Seed(time.Now().UnixNano())
	numFrames := 500

	// initialize drops per column
	drops := make([]drop, COLS)
	for i := range drops {
		drops[i] = drop{
			pos:    rand.Intn(ROWS),
			length: 6 + rand.Intn(6),       // trail length 6-11
			speed:  0.9 + rand.Float64()*0.1, // chance to move per frame
		}
	}

	frames := make([]string, numFrames)

	for f := 0; f < numFrames; f++ {
		lines := make([]string, ROWS)
		for r := 0; r < ROWS; r++ {
			line := ""
			for c := 0; c < COLS; c++ {
				d := drops[c]
				if r <= d.pos && r > d.pos-d.length {
					line += string(randomChar())
				} else {
					line += " "
				}
			}
			lines[r] = line
		}

		frames[f] = strings.Join(lines, "\n")

		// move each drop down according to its speed
		for c := range drops {
			if rand.Float64() < drops[c].speed {
				drops[c].pos++
				if drops[c].pos >= ROWS+drops[c].length {
					drops[c].pos = 0
					drops[c].length = 6 + rand.Intn(6) // random new trail length
				}
			}
		}
	}

	return frames
}

func randomChar() rune {
	chars := "01ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%^&*"
	return rune(chars[rand.Intn(len(chars))])
}
