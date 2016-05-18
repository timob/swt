package swt

import "github.com/timob/javabind"

type EventsArmEventInterface interface {
	EventsTypedEventInterface
}

type EventsArmEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.ArmEvent(org.eclipse.swt.widgets.Event)
func NewEventsArmEvent(a WidgetsEventInterface) (*EventsArmEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ArmEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsArmEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


