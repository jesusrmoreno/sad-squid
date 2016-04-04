package main

import (
	"fmt"
	"math/rand"
)

func randomItem(list []string) string {
	return list[rand.Intn(len(list))]
}

// GenerateID generates a new Asana style ID
func GenerateID() string {
	count := rand.Intn(1000) + 2
	firstAdjective := randomItem(adjectives)
	secondAdjective := randomItem(adjectives)
	noun := randomItem(nouns)
	verb := randomItem(verbs)
	adverb := randomItem(adverbs)
	return fmt.Sprintf("%d-%s-%s-%s-%s-%s",
		count, firstAdjective, secondAdjective, noun, verb, adverb)
}
