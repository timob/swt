package swt

import "github.com/timob/javabind"

type EventsControlListenerInterface interface {

	// public abstract void controlMoved(org.eclipse.swt.events.ControlEvent)
	ControlMoved(a EventsControlEventInterface) 

	// public abstract void controlResized(org.eclipse.swt.events.ControlEvent)
	ControlResized(a EventsControlEventInterface) 
}

type EventsControlListener struct {
	*javabind.Callable
}

// public abstract void controlMoved(org.eclipse.swt.events.ControlEvent)
func (jbobject *EventsControlListener) ControlMoved(a EventsControlEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "controlMoved", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ControlEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void controlResized(org.eclipse.swt.events.ControlEvent)
func (jbobject *EventsControlListener) ControlResized(a EventsControlEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "controlResized", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ControlEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


