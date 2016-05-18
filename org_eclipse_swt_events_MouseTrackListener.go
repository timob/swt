package swt

import "github.com/timob/javabind"

type EventsMouseTrackListenerInterface interface {

	// public abstract void mouseEnter(org.eclipse.swt.events.MouseEvent)
	MouseEnter(a EventsMouseEventInterface) 

	// public abstract void mouseExit(org.eclipse.swt.events.MouseEvent)
	MouseExit(a EventsMouseEventInterface) 

	// public abstract void mouseHover(org.eclipse.swt.events.MouseEvent)
	MouseHover(a EventsMouseEventInterface) 
}

type EventsMouseTrackListener struct {
	*javabind.Callable
}

// public abstract void mouseEnter(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseTrackListener) MouseEnter(a EventsMouseEventInterface)  {
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

// public abstract void mouseExit(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseTrackListener) MouseExit(a EventsMouseEventInterface)  {
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

// public abstract void mouseHover(org.eclipse.swt.events.MouseEvent)
func (jbobject *EventsMouseTrackListener) MouseHover(a EventsMouseEventInterface)  {
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


