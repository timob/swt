package swt

import "github.com/timob/javabind"

type EventsMouseWheelListenerInterface interface {

	// public abstract void mouseScrolled(org.eclipse.swt.events.MouseEvent)
	MouseScrolled(a EventsMouseEventInterface) 
}

type EventsMouseWheelListener struct {
	*javabind.Callable
}

// public abstract void mouseScrolled(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseWheelListener) MouseScrolled(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseScrolled", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


