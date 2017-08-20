package main

import (
	_ "fmt"
	"errors"
	"fmt"
)


func main() {
	//arr := []int{2,3,1,0,2,5,3}
	//fmt.Println(findDump(arr))

	//arr := []int{2,3,5,4,3,2,6,7}
	//fmt.Println(findDump(arr))

	//arr :=  [][]int{
	//	[]int{1,2,8,9},
	//	[]int{2,4,9,12},
	//	[]int{4,7,10,13},
	//	[]int{6,8,11,15},
	//}
	//ele := 7
	//re := findInMetrix(arr, ele)
	//if re == false {
	//	t.Error("find error")
	//}

	//ele := 100
	//re := findInMetrix(arr, ele)
	//if re {
	//	fmt.Println("find error")
	//}
	//var str string = "We are happy."
	//str = replaceBlank(str)
	//fmt.Println(str)

	linkList := createLinkListFromArr([]int{1,2,3,4,5})
	printListReversingly(linkList)


}


// 面试题1
func findDump(arr []int) (int,  error){
	if arr == nil || len(arr) == 0 {
		return 0, errors.New("empty arr found")
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 || arr[i] > (len(arr) - 1) {
			return 0, errors.New("invalid element found")
		}
	}

    for  i := 0; i < len(arr); i++ {
		for arr[i] != i {
			if arr[i] == arr[arr[i]] {
				return arr[i], nil
			}

			t := arr[i]
			arr[i] = arr[t]
			arr[t] = t
		}
	}
	return 0, errors.New("no dump found")
}

// 面试题2
func findDump2(arr []int) (int, error) {
	if arr == nil || len(arr) == 0 {
		return 0, errors.New("empty arr found")
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 || arr[i] > len(arr) {
			return 0, errors.New("invalid element found")
		}
	}
	var start int = 1
	var end int = len(arr) - 1
	for end >= start {
		var middle int = ((end - start) >> 1 ) + start
		var count int =  countRange(arr, start, middle)
		if end == start {
			if count > 1 {
				return start, nil
			} else {
				break
			}
		}

		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}
	}

	return -1, errors.New("invalid arr")
}

func countRange(arr []int, start, middle int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	var count int = 0
	for i := 0; i < len(arr); i++ {
	   if arr[i] >= start && arr[i] <= middle {
		   count++
	   }
	}
	return count
}


// 面试题4
func findInMetrix(arr [][]int, number int) bool {
	width := len(arr[0])
	height := len(arr)

	maxW := width - 1
	minH := 0

	for minH < height && maxW >= 0 {
		if arr[minH][maxW] == number {
			return true
		}

		if arr[minH][maxW] > number {
			maxW--
		} else {
			minH++
		}
	}

	return false
}

// 面试题5
func replaceBlank(str string) string {
	bbstr := []byte(str)

	blankNumber := 0
	for _, char := range bbstr {
		if char == ' ' {
			blankNumber++
		}
	}

	bstr := make([]byte, len(bbstr), blankNumber * 2 + len(str))

	for i, char := range bbstr {
		bstr[i] = char
	}

	bstr = bstr[:cap(bstr)]

	for i := len(bbstr) - 1; i >= 0; i-- {
		if bstr[i] == ' ' {
			bstr[i + 2 * (blankNumber - 1)] = '%'
			bstr[i + 2 * (blankNumber - 1) + 1] = '2'
			bstr[i + 2 * (blankNumber - 1) + 2] = '0'
			blankNumber--
		} else {
			bstr[i + 2 * blankNumber] = bstr[i]
		}
	}
	return string(bstr)
}

// 面试题6
type LinkNode struct {
	value int
	next *LinkNode
}

func createLinkListFromArr(arr []int) *LinkNode {
	var pNode, nNode, headNode *LinkNode

	for i, ele := range arr {
		if i == 0 {
			pNode = &LinkNode{value: ele}
			headNode = pNode

			continue
		}

		nNode = &LinkNode{value: ele}
		pNode.next = nNode

		pNode = nNode

	}

	return headNode
}

func printListReversingly(node *LinkNode) {
	if node == nil {
		return
	}

	printListReversingly(node.next)
	fmt.Println(node.value)
}
