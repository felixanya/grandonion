package Observer

type ISubject interface {
	AddObserver(observers ...IObserver)
	NotifyObservers()
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) AddObservers(observers ...IObserver) {
	s.observers = append(s.observers, observers...)
}

func (s *Subject) NotifyObservers() {
	for k := range s.observers {
		s.observers[k].Notify()
	}
}
