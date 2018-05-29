package Fundamentals

import (
	"fmt"
	"sort"
	"testing"

	"github.com/mooncaker816/go-algorithms4/base"
)

func TestRank(t *testing.T) {
	path0 := "../testdata/tinyW.txt"
	path1 := "../testdata/tinyT.txt"
	// path0 := "../testdata/largeW.txt"
	// path1 := "../testdata/largeT.txt"
	whitelist := base.ReadInts(path0)
	sort.Ints(whitelist)
	checklist := base.ReadInts(path1)
	for _, n := range checklist {
		if Rank(n, whitelist) == -1 {
			fmt.Println(n)
		}
	}
}
