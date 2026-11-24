package main

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
		{"abc", []int{1, 1, 1}},
		{"", []int{}},
		{"aaaa", []int{4}},
		{"abab", []int{4}},
		{"abca", []int{4}}, // the trailing a pulls everything into one piece
	}
	for _, c := range cases {
		if got := PartitionLabels(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("PartitionLabels(%q) = %v, want %v", c.in, got, c.want)
		}
	}
	// Sizes must sum to the whole string.
	total := 0
	for _, n := range PartitionLabels("ababcbacadefegdehijhklij") {
		total += n
	}
	if total != len("ababcbacadefegdehijhklij") {
		t.Error("piece sizes must cover the string exactly")
	}
}
