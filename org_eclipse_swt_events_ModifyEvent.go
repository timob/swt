package swt

import "github.com/timob/javabind"

type EventsModifyEventInterface interface {
	EventsTypedEventInterface
}

type EventsModifyEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.ModifyEvent(org.eclipse.swt.widgets.Event)
func NewEventsModifyEvent(a WidgetsEventInterface) (*EventsModifyEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ModifyEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsModifyEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


