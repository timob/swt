package swt

import "github.com/timob/javabind"

type EventsMenuEventInterface interface {
	EventsTypedEventInterface
}

type EventsMenuEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.MenuEvent(org.eclipse.swt.widgets.Event)
func NewEventsMenuEvent(a WidgetsEventInterface) (*EventsMenuEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MenuEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsMenuEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


