package timesync

import (
	// "fmt"

	"time"
)

func main() {
	// when we want to wait till
	until, _ := time.Parse(time.RFC3339, "2023-06-22T15:04:05+02:00")

	// and now we wait
	time.Sleep(time.Until(until))

	// Do what ever we want..... 🎉
}
