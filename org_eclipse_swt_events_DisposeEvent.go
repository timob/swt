package swt

import "github.com/timob/javabind"

type EventsDisposeEventInterface interface {
	EventsTypedEventInterface
}

type EventsDisposeEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.DisposeEvent(org.eclipse.swt.widgets.Event)
func NewEventsDisposeEvent(a WidgetsEventInterface) (*EventsDisposeEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/DisposeEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsDisposeEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


