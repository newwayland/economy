// Copyright 2023 Neil Wilson. All rights reserved.
// Use of this source code is governed by the AGPL
// licence that can be found in the LICENSE file.

package random

import (
	"math/rand"
)

// Choice returns a random element from the non-empty slice population
//
// Choice will panic if the population is empty
func Choice[T any](population []T) T {
	return population[ChoiceIndex(population)]
}

// ChoiceIndex returns the index to a randome element from the non-empty
// slice population
//
// ChoiceIndex will panic if the population is empty
func ChoiceIndex[T any](population []T) int {
	return rand.Intn(len(population))
}

// WithProbability returns true if the next random number is less than chance.
//
// chance should be in the range 0.0 to 1.0
func WithProbability(chance float64) bool {
	return rand.Float64() < chance
}
