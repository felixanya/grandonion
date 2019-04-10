package Observer

type ISubject interface {
	Register(...IObserver)
	Dismiss(...IObserver)
	Notify()
}

type Subject struct {
	observers map[IObserver]struct{}
}

func (s *Subject) AddObservers(observers ...IObserver) {
	s.observers = append(s.observers, observers...)
}

func (s *Subject) NotifyObservers() {
	for k := range s.observers {
		s.observers[k].Notify()
	}
}
