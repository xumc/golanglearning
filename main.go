package main

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"math"

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

	//linkList := createLinkListFromArr([]int{1,2,3,4,5})
	//printListReversingly(linkList)

	//preOrder := []int{1,2,4,7,3,5,6,8}
	//inOrder := []int{4,7,2,1,5,3,8,6}
	//rootNode := rebuildBinaryTree(preOrder, inOrder)
	//fmt.Println(travelTree(rootNode))

	//q := cQueue{stack1: stack.New(), stack2: stack.New()}
	//q.appendTail(1)
	//q.appendTail(2)
	//q.appendTail(3)
	//fmt.Println(q.deleteHead())
	//fmt.Println(q.deleteHead())
	//fmt.Println(q.deleteHead())
	//fmt.Println(q.deleteHead())

	//fmt.Println(sum(4, 200))

	world := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		world[i] = make([]bool, 8)
	}
	queens(world, 8, 0, 0)

}

// 面试题1
func findDump(arr []int) (int, error) {
	if arr == nil || len(arr) == 0 {
		return 0, errors.New("empty arr found")
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 || arr[i] > (len(arr)-1) {
			return 0, errors.New("invalid element found")
		}
	}

	for i := 0; i < len(arr); i++ {
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
		var middle int = ((end - start) >> 1) + start
		var count int = countRange(arr, start, middle)
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

	bstr := make([]byte, len(bbstr), blankNumber*2+len(str))

	for i, char := range bbstr {
		bstr[i] = char
	}

	bstr = bstr[:cap(bstr)]

	for i := len(bbstr) - 1; i >= 0; i-- {
		if bstr[i] == ' ' {
			bstr[i+2*(blankNumber-1)] = '%'
			bstr[i+2*(blankNumber-1)+1] = '2'
			bstr[i+2*(blankNumber-1)+2] = '0'
			blankNumber--
		} else {
			bstr[i+2*blankNumber] = bstr[i]
		}
	}
	return string(bstr)
}

// 面试题6
type LinkNode struct {
	value int
	next  *LinkNode
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

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// 面试题7
func rebuildBinaryTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}

	if len(preOrder) == 1 {
		leaf := TreeNode{value: preOrder[0]}
		return &leaf
	}

	for _, ele := range preOrder {
		for i := 0; i < len(inOrder); i++ {
			if ele == inOrder[i] {
				leftSub := inOrder[:i]
				left := rebuildBinaryTree(preOrder[1:(1+len(leftSub))], leftSub)
				right := rebuildBinaryTree(preOrder[(1+len(leftSub)):], inOrder[(i+1):])
				var rootNode TreeNode = TreeNode{value: ele, left: left, right: right}
				return &rootNode
			}
		}
	}
	return nil
}

func travelTree(rootNode *TreeNode) []int {
	var ret []int
	if rootNode == nil {
		return ret
	}
	ret = append(ret, rootNode.value)
	leftRet := travelTree(rootNode.left)
	rightRet := travelTree(rootNode.right)

	for _, ele := range leftRet {
		ret = append(ret, ele)
	}
	for _, ele := range rightRet {
		ret = append(ret, ele)
	}
	return ret
}

type cQueue struct{
	stack1 *stack.Stack
    stack2  *stack.Stack
}

func (q *cQueue) appendTail(ele interface{}) {
	q.stack1.Push(ele)
}

func (q *cQueue) deleteHead() interface{} {
	var ret, ele interface{}
	for true {
		ele = q.stack1.Pop();
		if q.stack1.Len() == 0 {
			ret = ele
			break
		} else {
			q.stack2.Push(ele)
		}
	}

	for q.stack2.Len() > 0 {
		q.stack1.Push(q.stack2.Pop())
	}

	return ret
}

// 写一个函数，求两个整数的之和，要求在函数体内不得使用＋、－、×、÷。
func sum(a, b int) int {
	if b == 0 {
		return a
	}

	s := a ^ b
	add := (a & b) << 1
	return sum(s, add)
}

// 求1+2+…+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字以及条件判断语句（A?B:C）
func accumulativeumS(n int) int {
	return (1 + n) *n / 2
}


var (
	count = 0
)

// 在8×8的国际象棋上摆放八个皇后，使其不能相互攻击，即任意两个皇后不得处在同一行、同一列或者同一对角斜线上。下图中的每个黑色格子表示一个皇后，这就是一种符合条件的摆放方法。请求出总共有多少种摆法。
func queens(chessBoard [][]bool, queensNumber int, startX, startY int) {
	printChessBoard(chessBoard)
	if queensNumber <= 0 {
		count++
		return
	}

	for i := startX; i < len(chessBoard); i++ {
		for j := startY; j < len(chessBoard[i]); j++ {
			chessBoard[i][j] = true
			if validate(chessBoard, i, j) {
				queens(chessBoard, queensNumber - 1, i, j)
			}
			chessBoard[i][j] = false
		}
	}

}

func validate(chessBoard [][]bool, x, y int) bool {
	for i := 0; i <= x; i++ {
		for j := 0; j < len(chessBoard[i]); j++ {
			if i == x {
				if j >= y {
					return true
				}
			}
			if chessBoard[i][j] == true {
				if i == x {
					return false
				}
				if j == y {
					return false
				}
				if math.Abs(float64(x - i)) == math.Abs(float64(y - j)) {
					return false
				}
			}
		}
	}
	return true
}

func printChessBoard(chessBoard [][]bool) {
	for i := 0; i <= len(chessBoard); i++ {
		for j := 0; j < len(chessBoard[i]); j++ {
			fmt.Print(chessBoard[i][j], ",")
		}
		fmt.Println("")
	}
}