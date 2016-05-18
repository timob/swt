package swt

import "github.com/timob/javabind"

type EventsFocusEventInterface interface {
	EventsTypedEventInterface
}

type EventsFocusEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.FocusEvent(org.eclipse.swt.widgets.Event)
func NewEventsFocusEvent(a WidgetsEventInterface) (*EventsFocusEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/FocusEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsFocusEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


