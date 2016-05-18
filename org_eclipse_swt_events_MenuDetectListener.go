package swt

import "github.com/timob/javabind"

type EventsMenuDetectListenerInterface interface {

	// public abstract void menuDetected(org.eclipse.swt.events.MenuDetectEvent)
	MenuDetected(a EventsMenuDetectEventInterface) 
}

type EventsMenuDetectListener struct {
	*javabind.Callable
}

// public abstract void menuDetected(org.eclipse.swt.events.MenuDetectEvent)
func (jbobject *EventsMenuDetectListener) MenuDetected(a EventsMenuDetectEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "menuDetected", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuDetectEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


