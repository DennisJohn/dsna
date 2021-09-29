package algo_strings

import (
	"sort"
	"strings"
)

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
func isValid(s string) bool {
	st := stack{[]rune{}}
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			st.push(v)
		} else {
			if len(st.s) == 0 {
				return false
			}
			switch v {
			case ']':
				if st.pop() != '[' {
					return false
				}
			case ')':
				if st.pop() != '(' {
					return false
				}
			case '}':
				if st.pop() != '{' {
					return false
				}
			default:
				return false
			}
		}
	}
	return len(st.s) == 0
}

type stack struct {
	s []rune
}

func (s *stack) push(r rune) {
	s.s = append(s.s, r)
}

func (s *stack) pop() rune {
	l := len(s.s)
	rt := s.s[l-1]
	s.s = s.s[:l-1]
	return rt
}

//1047
func removeDuplicates(s string) string {
	st := stackR{s: []byte{}}
	for i := 0; i < len(s); i++ {
		if st.max() == s[i] {
			st.pop()
		} else {
			st.push(s[i])
		}
	}
	return string(st.s)
}

type stackR struct {
	s []byte
}

func (s *stackR) push(b byte) {
	s.s = append(s.s, b)
}

func (s *stackR) max() byte {
	if len(s.s) == 0 {
		return 0
	}
	return s.s[len(s.s)-1]
}

func (s *stackR) pop() {
	s.s = s.s[:len(s.s)-1]
}

//1
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; !ok {
			m[target-v] = i
		} else {
			return []int{m[v], i}
		}
	}
	return []int{-1, -1}
}

//167
func twoSumSorted(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		o := numbers[l] + numbers[r]
		if o == target {
			return []int{l + 1, r + 1}
		}
		if o > target {
			r--
		} else {
			l++
		}
	}
	return []int{-1, -1}
}

//55
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	max := nums[0]
	for i := 1; i <= max; i++ {
		cur := nums[i]
		if cur+i >= len(nums)-1 {
			return true
		}
		if cur+i > max {
			max = cur + i
		}
	}
	return false
}

//56 Merge Intervals
func merge(intervals [][]int) [][]int {
	//Sort the intervals with respect 0th index ascending
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//Initialize the stack
	stack := stackM{}
	for i, x := range intervals {
		if i == 0 {
			stack.stackIntervals = append(stack.stackIntervals, x)
			continue
		}
		y := stack.max()
		if x[0] >= y[0] && x[0] <= y[1] {
			z := []int{0, 0}
			z[0] = y[0]
			if y[1] >= x[1] {
				z[1] = y[1]
			} else {
				z[1] = x[1]
			}
			stack.pop()
			stack.add(z)
			continue
		}
		stack.add(x)
	}
	return stack.stackIntervals
}

type stackM struct {
	stackIntervals [][]int
}

func (s *stackM) add(interval []int) {
	s.stackIntervals = append(s.stackIntervals, interval)
}

func (s *stackM) max() []int {
	return s.stackIntervals[len(s.stackIntervals)-1]
}

func (s *stackM) pop() {
	s.stackIntervals = s.stackIntervals[:len(s.stackIntervals)-1]
}

//303 Range Sum Query - Immutable
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	s := 0
	sum := make([]int, len(nums)+1)
	sum[0] = s
	for i, v := range nums {
		s = s + v
		sum[i+1] = s
	}
	return NumArray{sum: sum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

//346. Moving Average from Data Stream

type MovingAverage struct {
	size  int
	count int
	queue []int
	sum   int
}

func ConstructorM(size int) MovingAverage {
	return MovingAverage{size: size, count: 0, queue: []int{}, sum: 0}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.count < this.size {
		this.count++
		this.queue = append(this.queue, val)
		this.sum = this.sum + val
		return float64(this.sum) / float64(this.count)
	} else {
		this.sum = this.sum + val
		this.sum = this.sum - this.queue[0]
		this.queue = append(this.queue[1:], val)
		return float64(this.sum) / float64(this.size)
	}
}

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
