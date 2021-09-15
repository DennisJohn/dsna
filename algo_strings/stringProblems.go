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
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			isFirstLetter = true
			length++
		} else {
			if isFirstLetter {
				return length
			}
		}
	}
	return length
}

//242
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i := range s {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

//1089
func duplicateZeros(arr []int)  {
    for i := 0; i < len(arr) - 1; i++ {
        if arr[i] == 0 {
            copy(arr[i+1:], arr[i:])
            i++
        }
    }
}

//1313
//844
//20
//1047
//1
//167
//9
//88
//234   => Linked List
//21
//160
//328
//2
//237
//83
//82
//19
//706
//382
//141
//142