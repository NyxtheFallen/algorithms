package algorithms

import (
	"math/rand"
	"reflect"
	"testing"
)

var mergeSortCases = []struct {
	description string
	input       []int
	expected    []int
}{
	{
		description: "empty array",
		input:       []int{},
		expected:    []int{},
	},
	{
		description: "single-item array",
		input:       []int{9},
		expected:    []int{9},
	},
	{
		description: "two-item array",
		input:       []int{6, 1},
		expected:    []int{1, 6},
	},
	{
		description: "three-item array",
		input:       []int{6, 1, 7},
		expected:    []int{1, 6, 7},
	},
	{
		description: "reversed array",
		input:       []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		expected:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		description: "randomized array",
		input:       []int{5, 4, 8, 7, 6, 3, 2, 1, 9},
		expected:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		description: "array with negative numbers",
		input:       []int{5, 4, -8, 7, 6, 3, -2, 1, 9},
		expected:    []int{-8, -2, 1, 3, 4, 5, 6, 7, 9},
	},
}

func TestMergeSort(t *testing.T) {
	for _, val := range mergeSortCases {
		got, _ := MergeSort(val.input)
		if !reflect.DeepEqual(got, val.expected) {
			t.Fatalf("FAIL: %s\nHad %v, expected %v, got %v", val.description, val.input, val.expected, got)
		}

		t.Logf("PASS: %v", val.description)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	input := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		input[i] = rand.Intn(1000000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, val := range mergeSortCases {
			MergeSort(val.input)
		}
	}
}
