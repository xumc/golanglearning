package main

import (
	"testing"
	"fmt"
	"reflect"
	"github.com/golang-collections/collections/stack"

)

func Test_findInMetrix(t *testing.T) {
	arr := [][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
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

func Test_rebuildBinaryTree(t *testing.T) {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	rootNode := rebuildBinaryTree(preOrder, inOrder)
	travelArr := travelTree(rootNode)

	if !reflect.DeepEqual(preOrder, travelArr) {
		t.Error("rebuild tree error")
	}

	preOrder = []int{1, 2, 3 ,4 ,5 ,6 ,7}
	inOrder = []int{1, 2, 3, 4, 5, 6, 7}
	rootNode = rebuildBinaryTree(preOrder, inOrder)
	travelArr = travelTree(rootNode)
	fmt.Println(travelArr)

	if !reflect.DeepEqual(preOrder, travelArr) {
		t.Error("rebuild tree error")
	}

	preOrder = []int{1}
	inOrder = []int{1}
	rootNode = rebuildBinaryTree(preOrder, inOrder)
	travelArr = travelTree(rootNode)
	fmt.Println(travelArr)

	if !reflect.DeepEqual(preOrder, travelArr) {
		t.Error("rebuild tree error")
	}
}

func test_cqueue(t testing.T) {
	q := cQueue{stack1: stack.New(), stack2: stack.New()}
	q.appendTail(1)
	q.appendTail(2)
	q.appendTail(3)
	if q.deleteHead() != 1 {
		t.Error("deleteHead error")
	}
	if q.deleteHead() != 2 {
		t.Error("deleteHead error")
	}
	if q.deleteHead() != 3 {
		t.Error("deleteHead error")
	}
}
