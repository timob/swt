package swt

import "github.com/timob/javabind"

type EventsControlAdapterInterface interface {
	JavaLangObjectInterface

	// public void controlMoved(org.eclipse.swt.events.ControlEvent)
	ControlMoved(a EventsControlEventInterface) 

	// public void controlResized(org.eclipse.swt.events.ControlEvent)
	ControlResized(a EventsControlEventInterface) 
}

type EventsControlAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.ControlAdapter()
func NewEventsControlAdapter() (*EventsControlAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ControlAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsControlAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void controlMoved(org.eclipse.swt.events.ControlEvent)
func (jbobject *EventsControlAdapter) ControlMoved(a EventsControlEventInterface)  {
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

// public void controlResized(org.eclipse.swt.events.ControlEvent)
func (jbobject *EventsControlAdapter) ControlResized(a EventsControlEventInterface)  {
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


