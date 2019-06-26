package mediator

type Mediator interface {
	Register(Colleague)
	Operate() error
}

type ConcreteMediator struct {
	Colleagues 	[]Colleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		Colleagues: make([]Colleague, 0),
	}
}

func (cm *ConcreteMediator) Register(co Colleague) {
	cm.Colleagues = append(cm.Colleagues, co)
}

func (cm *ConcreteMediator) Operate() error {

	return nil
}

// colleague
type Colleague interface {
	Mediator
}

type ConcreteColleague struct {

}
