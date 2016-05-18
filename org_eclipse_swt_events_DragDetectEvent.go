package swt

import "github.com/timob/javabind"

type EventsDragDetectEventInterface interface {
	EventsMouseEventInterface
}

type EventsDragDetectEvent struct {
	EventsMouseEvent
}

// public org.eclipse.swt.events.DragDetectEvent(org.eclipse.swt.widgets.Event)
func NewEventsDragDetectEvent(a WidgetsEventInterface) (*EventsDragDetectEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/DragDetectEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsDragDetectEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


