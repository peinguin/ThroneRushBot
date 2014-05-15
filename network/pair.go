package network

import (
	"sort"
)

type Pair struct {
	key string   `json:"key"`
	value string `json:"value"`
}

type ByKeys []Pair
func (a ByKeys) Len() int           { return len(a) }
func (a ByKeys) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKeys) Less(i, j int) bool { return a[i].key < a[j].key }

func sortByKey(pairs ByKeys) {
	sort.Sort(ByKeys(pairs));
}