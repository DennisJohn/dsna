package algo_strings

import (
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	if firstUniqChar("leetcode") != 0 {
		t.Errorf("Expected %v, got %v", 1, firstUniqChar("leetcode"))
	}
}

func BenchmarkFirstuniqchar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		firstUniqChar("leetcodejfaklsfasdfjalskdjfalksdjflkasjdflkasflkasf")
	}
}

func TestLengthOfLastWord(t *testing.T) {
	if lengthOfLastWord("This is a Test") != 4 {
		t.Errorf("Expected %v, got %v", 4, lengthOfLastWord("This is a Test"))
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lengthOfLastWord("leetcodejfaklsfasdfjalskdjfalksdjflkasjdflkasflkasf")
	}
}

func TestIsAnagram(t *testing.T) {
	if !isAnagram("anagram", "nagaram") {
		t.Errorf("Exptected %v, got %v", true, false)
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isAnagram("anagram", "nagaram")
	}
}

func TestShortestDistance(t *testing.T) {
	x := shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice")
	if x != 3 {
		t.Errorf("Exptected %v, got %v", 1, x)
	}
}

func BenchmarkShortestDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice")
	}
}

func BenchmarkShortestDistance2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		shortestDistance2([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice")
	}
}

// func TestDuplicateZeros(t *testing.T) {
// 	if duplicateZeros([]int{0,0,0,0,0,0,0}) != []int{0,0,0,0,0} {
// 		t.Errorf("Exptected %v, got %v", true, false)
// 	}
// }
