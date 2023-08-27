// Copyright 2023 Neil Wilson. All rights reserved.
// Use of this source code is governed by the AGPL
// licence that can be found in the LICENSE file.

// Package random provides a set of random selection utility functions
package random

import (
	"math/rand"
)

// sample1Core makes a copy of the population, shuffles it
// and returns the first sampleSize elements
func sample1Core[T any](sampleSize int, population []T, length int) []T {
	source := make([]T, length)
	copy(source, population)
	rand.Shuffle(
		length,
		func(i, j int) {
			source[i], source[j] = source[j], source[i]
		},
	)
	return source[:sampleSize]
}

func sample1[T any](sampleSize int, population []T) []T {
	return validatedSample(sampleSize, population, sample1Core)
}

// sample2Core creates a sampleSize list of non-repeating random index numbers
// and then makes copies of the original population at those indexes.
func sample2Core[T any](sampleSize int, population []T, length int) []T {
	indexes := make(map[int]bool, sampleSize)
	var r int
	for i := 0; i < sampleSize; i++ {
		r = rand.Intn(length)
		for indexes[r] {
			r = rand.Intn(length)
		}
		indexes[r] = true
	}
	samples := make([]T, sampleSize)
	i := 0
	for k := range indexes {
		samples[i] = population[k]
		i++
	}
	return samples
}

func sample2[T any](sampleSize int, population []T) []T {
	return validatedSample(sampleSize, population, sample2Core)
}

// sample3Core creates a copy of the population, moves the selected component to the end and
// shrinks the random index to exclude it.
func sample3Core[T any](sampleSize int, population []T, length int) []T {
	source := make([]T, length)
	copy(source, population)
	samples := make([]T, sampleSize)
	var r int
	for i := 0; i < sampleSize; i++ {
		r = rand.Intn(length - i)
		samples[i] = source[r]
		source[r], source[length-i-1] = source[length-i-1], source[r]
	}
	return samples
}

func sample3[T any](sampleSize int, population []T) []T {
	return validatedSample(sampleSize, population, sample3Core)
}

// validatedSample validates the sampleSize against the population
// before calling the provided sample function
func validatedSample[T any](sampleSize int, population []T, fn func(int, []T, int) []T) []T {
	length := len(population)
	validateSampleSize(sampleSize, length)
	return fn(sampleSize, population, length)
}

// validateSampleSize checks the sampleSize is within the index range of the population
// and panics if not
func validateSampleSize(sampleSize int, length int) {
	if sampleSize < 0 {
		panic("sample less than zero")
	}
	if sampleSize > length {
		panic("sample larger than length of population")
	}
}

// Sample returns a sampleSize length slice of unique elements chosen
// from population. Used for random sampling without replacement.
//
// Returns a new slice containing elements from the population while
// leaving the original population unchanged.
//
// If the population contains repeats, then each occurrence is a possible
// selection in the sample.
//
// Sample will panic if the sampleSize is greater than the length of
// the population or less than zero,
func Sample[T any](sampleSize int, population []T) []T {
	length := len(population)
	validateSampleSize(sampleSize, length)
	test := float64(sampleSize*100) / float64(length)
	if test <= 30 {
		return sample2Core(sampleSize, population, length)
	}
	return sample1Core(sampleSize, population, length)
}
