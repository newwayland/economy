// Copyright 2023 Neil Wilson. All rights reserved.
// Use of this source code is governed by the AGPL
// licence that can be found in the LICENSE file.

package random

import (
	"reflect"
	"slices"
	"testing"
)

func TestChoiceEmptyStringSlice(t *testing.T) {
	shouldPanic(
		t,
		func() {
			_ = Choice([]string{})
		},
	)
}

func TestChoiceStringSlice(t *testing.T) {
	testList := []string{"a", "b", "c"}
	localList := make([]string, len(testList))
	copy(localList, testList)
	element := Choice(localList)
	if !reflect.DeepEqual(localList, testList) {
		t.Errorf("Choice has altered the test list")
		copy(localList, testList)
	}
	if !slices.Contains(testList, element) {
		t.Errorf("Choice didn't return an element from the list")
	}

}
