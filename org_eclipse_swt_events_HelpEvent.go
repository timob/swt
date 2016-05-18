package swt

import "github.com/timob/javabind"

type EventsHelpEventInterface interface {
	EventsTypedEventInterface
}

type EventsHelpEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.HelpEvent(org.eclipse.swt.widgets.Event)
func NewEventsHelpEvent(a WidgetsEventInterface) (*EventsHelpEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/HelpEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsHelpEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


