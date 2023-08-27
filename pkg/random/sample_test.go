// Copyright 2023 Neil Wilson. All rights reserved.
// Use of this source code is governed by the AGPL
// licence that can be found in the LICENSE file.

package random

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

var shortStringSlice = makeStringSlice(20, 10)
var longStringSlice = makeStringSlice(100000, 10)
var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func makeStringSlice(num, length int) []string {
	result := make([]string, num)
	for i := range result {
		result[i] = randStringBytesMaskImprSrcUnsafe(length)
	}
	return result
}

func BenchmarkSampleStringSlice(b *testing.B) {
	benchData := map[string]func(int, []string) []string{
		"sample4": Sample[string],
		"sample2": sample2[string],
	}
	stringData := map[string][]string{
		"short": shortStringSlice,
		"long":  longStringSlice,
	}

	b.ResetTimer()
	for sliceName, stringSlice := range stringData {
		for j := 1; j <= 10; j++ {
			for benchName, fn := range benchData {
				b.Run(
					fmt.Sprintf("%s%s%v%%", benchName, sliceName, j*10),
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							fn(len(stringSlice)*j/10, stringSlice)
						}
					},
				)
			}
		}
	}
}

func shouldPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() { _ = recover() }()
	f()
	t.Errorf("should have panicked")
}

func TestSampleStringSlice(t *testing.T) {
	stringData := map[string][]string{
		"short": shortStringSlice,
		"long":  longStringSlice,
	}
	for sliceName, stringSlice := range stringData {
		localStringSlice := make([]string, len(stringSlice))
		copy(localStringSlice, stringSlice)
		testData := map[string]func(int, []string) []string{
			"sample1": sample1[string],
			"sample2": sample2[string],
			"sample3": sample3[string],
		}
		for testName, fn := range testData {
			shouldPanic(
				t,
				func() {
					_ = fn(len(localStringSlice)+1, localStringSlice)
				},
			)
			shouldPanic(
				t,
				func() {

					_ = fn(-1, localStringSlice)
				},
			)
			x := len(localStringSlice) / 2
			s := fn(x, localStringSlice)
			if len(s) != x {
				t.Errorf("%s has wrong length %v rather than %v", testName, len(s), x)
			}
			if !reflect.DeepEqual(localStringSlice, stringSlice) {
				t.Errorf("%s has altered %s", testName, sliceName)
				copy(localStringSlice, stringSlice)
			}
			// Add returned slice to a makeshift set
			m := make(map[string]bool)
			for _, item := range s {
				m[item] = true
			}
			set := make([]string, 0, len(s))
			for k := range m {
				set = append(set, k)
			}
			if len(set) != len(s) {
				t.Errorf("%s returned duplicate items from %s", testName, sliceName)
			}
		}
	}
}
