package own

import "testing"

func Test_climbingStairsBacktrack(t *testing.T) {
	t.Logf("爬楼梯:%d\n", climbingStairsBacktrack(5))
	t.Logf("爬楼梯:%d\n", climbingStairsDFS(5))
	t.Logf("爬楼梯:%d\n", climbingStairsDPMemo(5))
	t.Logf("爬楼梯:%d\n", climbingStairsDPWithoutMemo(5))
}
