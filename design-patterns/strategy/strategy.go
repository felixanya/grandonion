package strategy

/*
策略模式
	NO if-elseif
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

//
type ArithmeticCalculator struct {
	Arithmetic      IArithmetic
}

func NewArithmeticCalculator(arithmetic IArithmetic) *ArithmeticCalculator {
	return &ArithmeticCalculator{
		Arithmetic: arithmetic,
	}
}

func (ae *ArithmeticCalculator) Calculate(num1, num2 int) int {
	return ae.Arithmetic.Compute(num1, num2)
}
