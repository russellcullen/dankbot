package status

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

var statuses []string = []string{
	"Half-Life 3",
	"Your Mom",
	"Monopoly",
	"Dank Memes",
	"Discord Bot Simulator",
	"Dungeons and Dragons",
	"Farmville",
	"Kingdom Hearts 3",
	"No Man's Buy",
	"Westworld",
	"Discord: The Game",
	"Uncivilization VI",
}

// TickerStatus calls the given callback function once per duration specified
func TickerStatus(d time.Duration, callback func(string)) {
	ticker := time.Tick(d)
	go func() {
		for range ticker {
			callback(RandomStatus())
		}
	}()
}

// RandomStatus returns a random user status
func RandomStatus() string {
	return statuses[random.Intn(len(statuses))]
}
