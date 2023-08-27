// Copyright 2023 Neil Wilson. All rights reserved.
// Use of this source code is governed by the AGPL
// licence that can be found in the LICENSE file.

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

func ExampleChoice() {
	// Fix the random order
	rand.Seed(1)
	words := strings.Fields("ink runs from the corners of my mouth")

	wordChoice := Choice(words)
	fmt.Println(wordChoice)
	// Output:
	// runs
}

func ExampleChoiceIndex() {
	// Fix the random order
	rand.Seed(1)
	words := strings.Fields("ink runs from the corners of my mouth")

	index := ChoiceIndex(words)
	fmt.Println(index)
	// Output:
	// 1
}

func ExampleWithProbability() {
	fmt.Println(WithProbability(1.0))
	fmt.Println(WithProbability(0.0))
	// Output:
	// true
	// false
}

func ExampleWeightedChoice() {
	// Fix the random order
	rand.Seed(1)

	var list []WeightedElement[string, uint]
	list = append(list, NewWeightedElement("first", uint(20)))
	list = append(list, NewWeightedElement("second", uint(30)))
	list = append(list, NewWeightedElement("third", uint(10)))
	fmt.Println(WeightedChoice(list))
	// Output:
	// second
}
