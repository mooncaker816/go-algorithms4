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

func TestRankR(t *testing.T) {
	path0 := "../testdata/tinyW.txt"
	path1 := "../testdata/tinyT.txt"
	// path0 := "../testdata/largeW.txt"
	// path1 := "../testdata/largeT.txt"
	whitelist := base.ReadInts(path0)
	sort.Ints(whitelist)
	checklist := base.ReadInts(path1)
	for _, n := range checklist {
		if RankR(n, whitelist) == -1 {
			fmt.Println(n)
		}
	}
}

var funcs = []struct {
	name string
	f    func(int, []int) int
}{
	{"rank", Rank},
	{"recursive rank", RankR},
}

var files = []struct {
	w string
	t string
}{
	{"tinyW.txt", "tinyT.txt"},
	{"largeW.txt", "largeT.txt"},
}

func BenchmarkRank(b *testing.B) {
	for _, f := range funcs {
		for _, file := range files {
			w := "../testdata/" + file.w
			t := "../testdata/" + file.t
			whitelist := base.ReadInts(w)
			checklist := base.ReadInts(t)
			b.Run(fmt.Sprintf("%s/%s", f.name, file.t), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, n := range checklist {
						f.f(n, whitelist)
					}
				}
			})
		}
	}
}
