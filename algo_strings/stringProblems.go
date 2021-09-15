package algo_strings

//387
func firstUniqChar(s string) int {
	m := make([]int, 26)
	for i, _ := range s {
		m[s[i]-'a']++
	}
	for i, _ := range s {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

//58 
func lengthOfLastWord(s string) int {
	isFirstLetter := false
	length := 0
	for i := len(s)-1; i >= 0; i-- {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			isFirstLetter = true
			length++
		}  else {
			if isFirstLetter {
				return length
			}	
		}
	}
	return length
}

//242
func isAnagram(s string, t string) bool {
	m := make(map[rune]bool)
	for _, v := range s {
		m[v] = true
	}
	for _, v := range t {
		if _, ok := m[v]; !ok {
			return false
		}
	}
	return true
}