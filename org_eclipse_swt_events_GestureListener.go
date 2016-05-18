package swt

import "github.com/timob/javabind"

type EventsGestureListenerInterface interface {

	// public abstract void gesture(org.eclipse.swt.events.GestureEvent)
	Gesture(a EventsGestureEventInterface) 
}

type EventsGestureListener struct {
	*javabind.Callable
}

// public abstract void gesture(org.eclipse.swt.events.GestureEvent)
func (jbobject *EventsGestureListener) Gesture(a EventsGestureEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "gesture", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/GestureEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


