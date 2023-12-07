package leetcode

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Printf("GET 1 expect 1: %d\n", lru.Get(1))
	lru.Put(3, 3)
	fmt.Printf("GET 2 expect -1: %d\n", lru.Get(2))
	lru.Put(4, 4)
	fmt.Printf("GET 1 expect -1: %d\n", lru.Get(1))
	fmt.Printf("GET 3 expect 3: %d\n", lru.Get(3))
	fmt.Printf("GET 4 expect 4: %d\n", lru.Get(4))
}
