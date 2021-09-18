package algo_strings

import "strings"

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
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i++
		}
	}
}

//243
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	i1 := -1
	i2 := -1
	min := len(wordsDict)
	for i, v := range wordsDict {
		if v == word1 {
			i1 = i
		}
		if v == word2 {
			i2 = i
		}
		if i1 != -1 && i2 != -1 && min > AbsDiff(i1, i2) {
			min = AbsDiff(i1, i2)
		}
	}
	return min
}

//243 - Slow
func shortestDistance2(wordsDict []string, word1 string, word2 string) int {
	w1 := []int{}
	w2 := []int{}
	min := len(wordsDict)
	for i, v := range wordsDict {
		if v == word1 {
			w1 = append(w1, i)
		}
		if v == word2 {
			w2 = append(w2, i)
		}
	}
	for _, x := range w1 {
		for _, y := range w2 {
			d := AbsDiff(x, y)
			if d < min {
				min = d
			}
		}
	}
	return min
}

func AbsDiff(x int, y int) int {
	if x-y < 0 {
		return y - x
	}
	return x - y
}

//1313
func decompressRLElist(nums []int) []int {
	a := []int{}
	for i := 0; i < len(nums); i = i + 2 {
		for j := 0; j < nums[i]; j++ {
			a = append(a, nums[i+1])
		}
	}
	return a
}

//844
func backspaceCompare(s string, t string) bool {
	return genString(s) == genString(t)
}

func genString(s string) string {
	var o strings.Builder
	hc := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			hc++
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			if hc > 0 {
				hc--
				continue
			}
			o.WriteByte(s[i])
		}
	}
	return o.String()
}

func backspaceCompareSingleLoop(s string, t string) bool {
	ls, lt, hs, ht, ts, tt := len(s)-1, len(t)-1, 0, 0, "", ""
	for ls >= 0 || lt >= 0 {
		for ls >= 0 {
			if s[ls] == '#' {
				hs++
				ls--
				continue
			}
			if s[ls] >= 'a' && s[ls] <= 'z' {
				if hs > 0 {
					hs--
					ls--
					continue
				}
				ts = string(s[ls])
				break
			}
		}

		for lt >= 0 {
			if t[lt] == '#' {
				ht++
				lt--
				continue
			}
			if t[lt] >= 'a' && t[lt] <= 'z' {
				if ht > 0 {
					ht--
					lt--
					continue
				}
				tt = string(t[lt])
				break
			}
		}
		if ls < 0 {
			ts = ""
		}
		if lt < 0 {
			tt = ""
		}
		if ts != tt {
			return false
		}
		ls--
		lt--
	}
	return true
}

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
