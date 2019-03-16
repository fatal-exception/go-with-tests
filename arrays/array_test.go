package arrays_test

import (
	"reflect"
	"testing"

	. "github.com/fatal-exception/go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	})

	t.Run("Collection of arbitrary size", func(t *testing.T) {
		numbers := []int{1, 2, 4, 3, 1, 6, 5}
		got := Sum(numbers)
		want := 22
		if got != want {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	numbers_1 := []int{1, 1, 1}
	numbers_2 := []int{2, 2}
	got := SumAll(numbers_1, numbers_2)
	want := []int{3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestSumTail(t *testing.T) {
	numbers := []int{1, 3, 5}
	got := SumTail(numbers)
	want := 8
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	want := []int{5, 11}
	got := SumAllTails(numbers1, numbers2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestSumAllTailsEmptySlice(t *testing.T) {
	numbers1 := []int{}
	numbers2 := []int{4, 5, 6}
	want := []int{0, 11}
	got := SumAllTails(numbers1, numbers2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
