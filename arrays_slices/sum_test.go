package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertMessage := func(t *testing.T, nums []int, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d' given, %v", got, want, nums)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15
		assertMessage(t, nums, got, want)
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6
		assertMessage(t, nums, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 9})
	want := []int{6, 9}
	//want := "bob"

	//if got != want {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{0, 1, 2, 9})
	want := []int{5, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
