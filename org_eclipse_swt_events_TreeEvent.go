package swt

import "github.com/timob/javabind"

type EventsTreeEventInterface interface {
	EventsSelectionEventInterface
}

type EventsTreeEvent struct {
	EventsSelectionEvent
}

// public org.eclipse.swt.events.TreeEvent(org.eclipse.swt.widgets.Event)
func NewEventsTreeEvent(a WidgetsEventInterface) (*EventsTreeEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TreeEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsTreeEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


