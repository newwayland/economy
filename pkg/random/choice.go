// Copyright 2023 Neil Wilson. All rights reserved.
// Use of this source code is governed by the AGPL
// licence that can be found in the LICENSE file.

package random

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// WeightedElement associates an integer weight with an item to allow
// weighted selection
type WeightedElement[T any, W constraints.Integer] struct {
	Element T
	Weight  W
}

// Choice returns a random element from the non-empty slice population
//
// Choice will panic if the population is empty
func Choice[T any](population []T) T {
	return population[ChoiceIndex(population)]
}

// ChoiceIndex returns the index to a random element from the non-empty
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

// WeightedChoice returns a random element from the non-empty slice population based upon the relative weights in the population
func WeightedChoice[T any, W constraints.Integer](population []WeightedElement[T, W]) T {
	var sum int
	for _, c := range population {
		sum += int(c.Weight)
	}
	r := rand.Intn(sum)
	for _, c := range population {
		r -= int(c.Weight)
		if r < 0 {
			return c.Element
		}
	}
	panic("Error in weighted choices algorithm. Try a bigger int size")
}

// NewWeightedElement is a convenience function to create a WeightedElement structure
func NewWeightedElement[T any, W constraints.Integer](element T, weight W) WeightedElement[T, W] {
	return WeightedElement[T, W]{
		Element: element,
		Weight:  weight,
	}
}
