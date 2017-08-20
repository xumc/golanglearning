package main

import(
	"testing"
)
func Test_findInMetrix(t *testing.T) {
	arr :=  [][]int{
		[]int{1,2,8,9},
		[]int{2,4,9,12},
		[]int{4,7,10,13},
		[]int{6,8,11,15},
	}
	ele := 7
	re := findInMetrix(arr, ele)
	if re == false {
		t.Error("find error")
	}

	ele = 100
	re = findInMetrix(arr, ele)
	if re {
		t.Error("find error")
	}

	ele = -100
	re = findInMetrix(arr, ele)
	if re {
		t.Error("find error")
	}

	ele = 14
	re = findInMetrix(arr, ele)
	if re {
		t.Error("find error")
	}
}

func Test_replaceBlank(t *testing.T) {
	var str string = "We are happy."
	str = replaceBlank(str)
	if str != "We%20are%20happy." {
		t.Error("replace err")
	}

	str = replaceBlank("")
	if str != "" {
		t.Error("replace err")
	}


}

