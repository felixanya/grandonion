package strategy

import (
    "log"
    "testing"
)

func TestNewArithmeticExecutor(t *testing.T) {
    arithmetic := NewArithmeticCalculator(&Addition{})
    ret := arithmetic.Calculate(10, 23)
    log.Println(">>>addition result>>>", ret)

    arithmetic2 := NewArithmeticCalculator(&Division{})
    ret2 := arithmetic2.Calculate(10, 5)
    log.Println(">>>division result>>>", ret2)
}
