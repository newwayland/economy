package random

import (
	"fmt"
	"math/rand"
	"strings"
)

func ExampleSample() {
	// Fix the random order
	rand.Seed(1)
	words := strings.Fields("ink runs from the corners of my mouth")

	wordSample := Sample(3, words)
	fmt.Println(wordSample)
	// Output:
	// [of ink mouth]
}
