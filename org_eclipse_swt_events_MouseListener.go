package swt

import "github.com/timob/javabind"

type EventsMouseListenerInterface interface {

	// public abstract void mouseDoubleClick(org.eclipse.swt.events.MouseEvent)
	MouseDoubleClick(a EventsMouseEventInterface) 

	// public abstract void mouseDown(org.eclipse.swt.events.MouseEvent)
	MouseDown(a EventsMouseEventInterface) 

	// public abstract void mouseUp(org.eclipse.swt.events.MouseEvent)
	MouseUp(a EventsMouseEventInterface) 
}

type EventsMouseListener struct {
	*javabind.Callable
}

// public abstract void mouseDoubleClick(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseListener) MouseDoubleClick(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseDoubleClick", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void mouseDown(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseListener) MouseDown(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseDown", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void mouseUp(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseListener) MouseUp(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseUp", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


