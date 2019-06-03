package Observer

type ISubject interface {
	Register(...IObserver)
	Dismiss(...IObserver)
	Notify(*Event)
}

type Subject struct {
	observers map[IObserver]struct{}
}

func (s *Subject) Register(observers ...IObserver) {
	for _, ob := range observers {
		s.observers[ob] = struct{}{}
	}
}

func (s *Subject) Dismiss(observers ...IObserver) {
	for _, ob := range observers {
		delete(s.observers, ob)
	}
}

func (s *Subject) Notify(event *Event) {
	event.wg.Add(len(s.observers))
	for ob := range s.observers {
		go ob.Update(event)
	}
	event.wg.Wait()
}
