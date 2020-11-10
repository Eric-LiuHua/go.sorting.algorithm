package main

import (
	"fmt"
	"go_sorts/sorts"
	"math/rand"
	"sort"
	"testing"
	"time"
)

//go 测试 必须在以_test结尾的源码内以Test开头的函数会自动被执行。
//每个测试用例函数需要以Test为前缀，例如
func TestBubble(t *testing.T) {
	for _, funcc := range SortFuncs {
		start := time.Now().Unix()
		t.Log(fmt.Sprintf("sort name:%s ,start:%d", funcc.name, start))
		for _, test := range SortTestDatas {
			t.Log(len(test.input))
			actual := funcc.funcc(test.input)
			pos, sorted := sorts.CompareSlices(actual, test.expected)
			if !sorted {
				if pos == -1 {
					t.Errorf("test %s failed due to slice length changing", test.name)
				}
				t.Errorf("test %s failed at index %d", test.name, pos)
			}
		}
		t.Log(fmt.Sprintf("sort name:%s,dur:%d", funcc.name, time.Now().Unix()-start))
	}
}

//测试结构
type SortTest struct {
	input    []int
	expected []int
	name     string
}
type SortFunc struct {
	name  string
	funcc func([]int) []int
}

var arr = MakeRandArray(500_000)

var SortFuncs = []SortFunc{
	{"CountingSort", sorts.CountingSort},
	{"Merge", sorts.Merge},
}

var SortTestDatas = []SortTest{
	//Sorted slice
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted"},
	//Reversed slice
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed"},
	//Empty slice
	{[]int{}, []int{}, "Empty"},
	//Single-entry slice
	{[]int{1}, []int{1}, "Singleton"},

	//500k values sort
	{arr, getSortedVersion(arr), "Large Random"},
}

//创建随机数
func MakeRandArray(size int) []int {
	res := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		tmp := rand.Int63n(int64(size))
		res[i] = int(tmp)
	}
	return res
}

//自动的排序
func getSortedVersion(a []int) []int {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return a
}
