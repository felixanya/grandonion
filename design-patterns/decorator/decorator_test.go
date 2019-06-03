package decorator

import (
	"fmt"
	"testing"
)

func TestSum1(t *testing.T) {
	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)

	fmt.Printf("%d, %d \n", sum1(-1000000, 1000000), sum2(-1000000, 1000000))
}
