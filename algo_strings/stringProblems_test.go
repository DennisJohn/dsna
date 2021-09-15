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
