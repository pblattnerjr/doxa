package concord

import (
	"fmt"
	"sort"
	"strings"
)
type Concordance struct {
	Map map[string]ConcordanceLine
}
const (
	SortId = iota
	SortLeft = iota
	SortRight = iota
)
type Order int
var SortOrder Order
func (o *Order) String() string {
	var s string
	switch SortOrder {
	case SortId: s = "id"
	case SortLeft: s = "left"
	case SortRight: s = "Right"
	}
	return s
}
type ConcordanceLine struct {
	ID    string
	Right string
	Key   string
	Left  string
}
type SortedMap struct {
	m map[string]ConcordanceLine
	s []string
}

func (sm *SortedMap) Len() int {
	return len(sm.m)
}

func (sm *SortedMap) Less(i, j int) bool {
	var l1, l2 string
	switch SortOrder {
	case SortId:
		l1 = strings.TrimSpace(sm.m[sm.s[j]].ID)
		l2 = strings.TrimSpace(sm.m[sm.s[i]].ID)
	case SortLeft:
		l1 = strings.TrimSpace(sm.m[sm.s[j]].Left) + sm.m[sm.s[j]].Key + strings.TrimSpace(sm.m[sm.s[j]].Right)
		l2 = strings.TrimSpace(sm.m[sm.s[i]].Left) + sm.m[sm.s[i]].Key + strings.TrimSpace(sm.m[sm.s[i]].Right)
	case SortRight:
		l1 = sm.m[sm.s[j]].Key + strings.TrimSpace(sm.m[sm.s[j]].Right) + strings.TrimSpace(sm.m[sm.s[i]].Left)
		l2 = sm.m[sm.s[i]].Key + strings.TrimSpace(sm.m[sm.s[i]].Right) + strings.TrimSpace(sm.m[sm.s[j]].Left)
	}
	return l1 > l2
}

func (sm *SortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func SortedKeys(m map[string]ConcordanceLine, o Order) []string {
	SortOrder = o
	sm := new(SortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
// Find the index of a substring within []rune
// This solves the following problem:
// When we have a string of Greek, they are runes. If we take a slice,
// we will get ? output for some characters.  The solution is to convert the
// string to a []rune.  But, then, our search index won't work since the length
// of the original string is > than the length of the []rune.
func indexInRune(text []rune, what string) int {
	whatRunes := []rune(what)
	for i := range text {
		found := true
		for j := range whatRunes {
			if text[i+j] != whatRunes[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
// Centers keyword in middle, with window size = Width for left and right
// And, adds id, left, key, right to concordance map named cMap so we can Sort the lines
// by parts.
func (c *Concordance) Line(id, line, key string, width int) {
	r := []rune(line)
	rKey := []rune(key)
	keyIndex := indexInRune(r, key)
	lineLen := len(r)
	// get left context
	leftIndex := 0
	if keyIndex-width > 0 {
		leftIndex = keyIndex
		for i := keyIndex - 1; i > 0; i-- {
			if leftIndex <= (keyIndex - width) {
				break
			} else {
				leftIndex--
			}
		}
	}
	// get right context
	rightStart := keyIndex + len(rKey) + 1
	rightIndex := rightStart
	for i := rightStart; i < lineLen; i++ {
		if rightIndex > keyIndex+width {
			break
		} else {
			rightIndex++
		}
	}
	var left, right string
	if 0 <= leftIndex && leftIndex <= keyIndex && keyIndex <= len(r) {
		left = string(r[leftIndex:keyIndex])
	}
	left = fmt.Sprintf("%*s", width-len(left), left)
	keyRight := keyIndex + len(rKey)
	if keyRight < rightIndex && rightIndex <= len(r) {
		right = string(r[keyRight:rightIndex])
	} else {
		if keyRight < len(r) {
			right = string(r[keyRight:])
		}
	}
	right = fmt.Sprintf("%-*s", width-len(right), right)
	var cl ConcordanceLine
	cl.ID = id
	cl.Left = fmt.Sprintf("%*s", width+3, left)
	cl.Key = key
	cl.Right = right
	c.Map[cl.ID] = cl
}

