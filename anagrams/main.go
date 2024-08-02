package main

import (
	"fmt"
	"sort"
	"strings"
)

// Helper function to check if two strings are anagrams
func areAnagrams(str1, str2 string) bool {
	// Remove spaces and convert to lower case
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Convert strings to slices of runes (characters)
	rune1 := []rune(str1)
	rune2 := []rune(str2)

	// Sort the slices
	sort.Slice(rune1, func(i, j int) bool {
		return rune1[i] < rune1[j]
	})
	sort.Slice(rune2, func(i, j int) bool {
		return rune2[i] < rune2[j]
	})

	// Convert sorted slices back to strings
	sortedStr1 := string(rune1)
	sortedStr2 := string(rune2)

	// Compare the sorted strings
	return sortedStr1 == sortedStr2
}

func main() {
	str1 := "listen"
	str2 := "silent"
	fmt.Printf("%q and %q are anagrams: %v\n", str1, str2, areAnagrams(str1, str2))

	str1 = "madam"
	str2 = "dammmam"
	fmt.Printf("%q and %q are anagrams: %v\n", str1, str2, areAnagrams(str1, str2))

	str1 = "liS ten"
	str2 = "silent"
	fmt.Printf("%q and %q are anagrams: %v\n", str1, str2, areAnagrams(str1, str2))
}
