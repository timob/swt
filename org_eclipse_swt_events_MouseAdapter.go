package swt

import "github.com/timob/javabind"

type EventsMouseAdapterInterface interface {
	JavaLangObjectInterface

	// public void mouseDoubleClick(org.eclipse.swt.events.MouseEvent)
	MouseDoubleClick(a EventsMouseEventInterface) 

	// public void mouseDown(org.eclipse.swt.events.MouseEvent)
	MouseDown(a EventsMouseEventInterface) 

	// public void mouseUp(org.eclipse.swt.events.MouseEvent)
	MouseUp(a EventsMouseEventInterface) 
}

type EventsMouseAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.MouseAdapter()
func NewEventsMouseAdapter() (*EventsMouseAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MouseAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsMouseAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void mouseDoubleClick(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseAdapter) MouseDoubleClick(a EventsMouseEventInterface)  {
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

// public void mouseDown(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseAdapter) MouseDown(a EventsMouseEventInterface)  {
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

// public void mouseUp(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseAdapter) MouseUp(a EventsMouseEventInterface)  {
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


