package swt

import "github.com/timob/javabind"

type EventsMouseTrackAdapterInterface interface {
	JavaLangObjectInterface

	// public void mouseEnter(org.eclipse.swt.events.MouseEvent)
	MouseEnter(a EventsMouseEventInterface) 

	// public void mouseExit(org.eclipse.swt.events.MouseEvent)
	MouseExit(a EventsMouseEventInterface) 

	// public void mouseHover(org.eclipse.swt.events.MouseEvent)
	MouseHover(a EventsMouseEventInterface) 
}

type EventsMouseTrackAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.MouseTrackAdapter()
func NewEventsMouseTrackAdapter() (*EventsMouseTrackAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MouseTrackAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsMouseTrackAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void mouseEnter(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseTrackAdapter) MouseEnter(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseEnter", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void mouseExit(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseTrackAdapter) MouseExit(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseExit", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void mouseHover(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseTrackAdapter) MouseHover(a EventsMouseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mouseHover", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


