package swt

import "github.com/timob/javabind"

type EventsHelpListenerInterface interface {

	// public abstract void helpRequested(org.eclipse.swt.events.HelpEvent)
	HelpRequested(a EventsHelpEventInterface) 
}

type EventsHelpListener struct {
	*javabind.Callable
}

// public abstract void helpRequested(org.eclipse.swt.events.HelpEvent)
func (jbobject *EventsHelpListener) HelpRequested(a EventsHelpEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "helpRequested", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/HelpEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


