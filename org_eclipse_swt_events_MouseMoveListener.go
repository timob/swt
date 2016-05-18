package swt

import "github.com/timob/javabind"

type EventsMouseMoveListenerInterface interface {

	// public abstract void mouseMove(org.eclipse.swt.events.MouseEvent)
	MouseMove(a EventsMouseEventInterface) 
}

type EventsMouseMoveListener struct {
	*javabind.Callable
}

// public abstract void mouseMove(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseMoveListener) MouseMove(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseMove", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


