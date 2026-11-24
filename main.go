package main

import "fmt"

// PartitionLabels splits s into the most pieces such that no letter appears
// in two pieces, returning the piece sizes. A piece must extend to the last
// occurrence of every letter it contains; cut when the scan reaches that
// furthest obligation.
func PartitionLabels(s string) []int {
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	out := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if l := last[s[i]-'a']; l > end {
			end = l // this letter obligates the piece to reach further
		}
		if i == end {
			out = append(out, end-start+1)
			start = i + 1
		}
	}
	return out
}

func main() {
	fmt.Println(PartitionLabels("ababcbacadefegdehijhklij"))
}
