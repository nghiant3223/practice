package sumroot2leafnumbers

import (
	"testing"

	"github.com/nghiant3223/gopractice/alg/tree/ds"
	"github.com/stretchr/testify/assert"
)

/*
Input: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
Output: 1026
*/

func TestHappy1(t *testing.T) {
	node4 := ds.NewNode(4)
	node9 := ds.NewNode(9)
	node0 := ds.NewNode(0)
	node5 := ds.NewNode(5)
	node1 := ds.NewNode(1)
	node4.Left = node9
	node4.Right = node0
	node9.Left = node5
	node9.Right = node1

	ans := sumNumbers(node4)
	ans2 := sumNumbers2(node4)
	assert.Equal(t, 1026, ans)
	assert.Equal(t, 1026, ans2)
}

/*
Input: [1,2,3]
    1
   / \
  2   3
Output: 25
*/

func TestHappy2(t *testing.T) {
	node1 := ds.NewNode(1)
	node2 := ds.NewNode(2)
	node3 := ds.NewNode(3)
	node1.Left = node2
	node1.Right = node3

	ans := sumNumbers(node1)
	ans2 := sumNumbers2(node1)
	assert.Equal(t, 25, ans)
	assert.Equal(t, 25, ans2)
}

/*
Input: [1,2,null,3]
    1
   /
  2
 /
3
Output: 123
*/

func TestHappy3(t *testing.T) {
	node1 := ds.NewNode(1)
	node2 := ds.NewNode(2)
	node3 := ds.NewNode(3)
	node1.Left = node2
	node2.Left = node3

	ans := sumNumbers(node1)
	ans2 := sumNumbers2(node1)
	assert.Equal(t, 123, ans)
	assert.Equal(t, 123, ans2)
}
