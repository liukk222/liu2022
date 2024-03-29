package main

import (
	"fmt"
	"sort"
)

func testFloat64() {
	f := []float64{1.1, 4.4, 5.5, 3.3, 2.2}
	sort.Float64s(f)
	fmt.Printf("f: %v\n", f)
}

func testStrings() {
	ls := sort.StringSlice{
		"100",
		"200",
		"41",
		"3",
		"2",
	}
	fmt.Println(ls) //[100 42 41 3 2]
	sort.Strings(ls)
	fmt.Println(ls) //[100 2 3 41 42]

	//字符串排序，现比较高位，相同的再比较低位
	ls = sort.StringSlice{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Println(ls) //[d ac c ab e]
	sort.Strings(ls)
	fmt.Println(ls) //[ab ac c d e]

	//汉字排序，依次比较byte大小
	ls = sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
		"饿",
		"周",
	}
	fmt.Println(ls) //[啊 博 次 得 饿 周]
	sort.Strings(ls)
	fmt.Println(ls) //[博 周 啊 得 次 饿]

	for _, v := range ls {
		fmt.Println(v, []byte(v))
	}

}

/* type testSlice [][]int

func (l testSlice) Len() int           { return len(l) }
func (l testSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool { return l[i][1] < l[j][1] } */

type testSlice []map[string]float64

func (l testSlice) Len() int           { return len(l) }
func (l testSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool { return l[i]["a"] < l[j]["a"] }

func main() {
	//testStrings()
	//testFloat64()
	// ls := testSlice{
	// 	{1, 4},
	// 	{9, 3},
	// 	{7, 5},
	// }
	// fmt.Println(ls) //[[1 4] [9 3] [7 5]]
	// sort.Sort(ls)
	// fmt.Println(ls) //[[9 3] [1 4] [7 5]]

	ls := testSlice{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}

	fmt.Println(ls) //[map[a:4 b:12] map[a:3 b:11] map[a:5 b:10]]
	sort.Sort(ls)
	fmt.Println(ls) //[map[a:3 b:11] map[a:4 b:12] map[a:5 b:10]]

}
