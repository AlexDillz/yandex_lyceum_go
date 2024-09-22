// Напишите функцию TimeAgo(pastTime time.Time) string, который вычисляет время, прошедшее с момента заданного значения time.Time,
// и возвращает удобочитаемую строку, указывающую, как давно это было.

package main

import (
	"fmt"
	"time"
)

// TimeAgo возвращает строку, указывающую, сколько времени прошло с pastTime
func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	duration := now.Sub(pastTime)

	switch {
	case duration.Hours() >= 24:
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	case duration.Hours() >= 1:
		hours := int(duration.Hours())
		return fmt.Sprintf("%d hours ago", hours)
	case duration.Minutes() >= 1:
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d minutes ago", minutes)
	default:
		seconds := int(duration.Seconds())
		return fmt.Sprintf("%d seconds ago", seconds)
	}
}
