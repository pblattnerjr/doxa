package concord

import (
	"sort"
	"strings"
)

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

