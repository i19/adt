package booking

import (
	"fmt"
	"testing"
)

func Test_carParkingDFS(t *testing.T) {
	parted := []bool{true, false, true, false, true, true, true, true}
	for k := 0; k < len(parted)+1; k++ {
		fmt.Println(carParkingForce(parted, k))
		fmt.Println(carParkingSlidingWindow(parted, k))
		fmt.Println("======")
	}
}
