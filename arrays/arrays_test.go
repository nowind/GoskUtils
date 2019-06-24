package arrays

import "testing"

func TestIn(t *testing.T) {
	t.Log(In([]int{1,2,4},4))
	t.Log(In([]int{1,2,4},3))
}
