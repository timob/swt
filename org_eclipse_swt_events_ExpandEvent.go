package swt

import "github.com/timob/javabind"

type EventsExpandEventInterface interface {
	EventsSelectionEventInterface
}

type EventsExpandEvent struct {
	EventsSelectionEvent
}

// public org.eclipse.swt.events.ExpandEvent(org.eclipse.swt.widgets.Event)
func NewEventsExpandEvent(a WidgetsEventInterface) (*EventsExpandEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ExpandEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsExpandEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


