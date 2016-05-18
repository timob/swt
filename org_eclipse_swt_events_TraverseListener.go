package swt

import "github.com/timob/javabind"

type EventsTraverseListenerInterface interface {

	// public abstract void keyTraversed(org.eclipse.swt.events.TraverseEvent)
	KeyTraversed(a EventsTraverseEventInterface) 
}

type EventsTraverseListener struct {
	*javabind.Callable
}

// public abstract void keyTraversed(org.eclipse.swt.events.TraverseEvent)
func (jbobject *EventsTraverseListener) KeyTraversed(a EventsTraverseEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "keyTraversed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TraverseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


