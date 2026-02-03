package ui

import (
	"fmt"
	"time"
)

// ShowSpinner renders a short inline spinner message.
func ShowSpinner(message string, duration time.Duration) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	end := time.Now().Add(duration)
	i := 0
	for time.Now().Before(end) {
		fmt.Printf("\r%s %s", frames[i%len(frames)], message)
		time.Sleep(60 * time.Millisecond)
		i++
	}
	fmt.Printf("\r%s %s\n", frames[i%len(frames)], message)
}
