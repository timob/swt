package swt

import "github.com/timob/javabind"

type EventsControlEventInterface interface {
	EventsTypedEventInterface
}

type EventsControlEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.ControlEvent(org.eclipse.swt.widgets.Event)
func NewEventsControlEvent(a WidgetsEventInterface) (*EventsControlEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ControlEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsControlEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


