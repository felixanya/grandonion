package strategy

/*
策略模式
 */

type IArithmetic interface {
	Compute(int, int) int
}


type Addition struct {
}

func (a *Addition) Compute(num1, num2 int) int {
	return num1 + num2
}

type Subtraction struct{}

func (s *Subtraction) Compute(num1, num2 int) int {
	return num1 - num2
}

type Multiplication struct {
}

func (m *Multiplication) Compute(num1, num2 int) int {
	return num1 * num2
}

type Division struct{}

func (d *Division) Compute(num1, num2 int) int {
	return num1 / num2
}
