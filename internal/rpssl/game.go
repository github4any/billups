package rpssl

import (
	"math/rand"
	"time"
)

func GetRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(5)
}

func GetResult(userAnswer, computerAnswer string) string {
	sets := map[string][]string{
		"scissors": {"paper", "lizard"},
		"paper":    {"rock", "spock"},
		"lizard":   {"spock", "paper"},
		"spock":    {"scissors", "rock"},
		"rock":     {"scissors", "lizard"},
	}

	if computerAnswer == userAnswer {
		return "tie"
	}

	for _, value := range sets[userAnswer] {
		if value == computerAnswer {
			return "win"
		}
	}
	return "lose"
}
