package swt

import "github.com/timob/javabind"

type EventsSegmentListenerInterface interface {

	// public abstract void getSegments(org.eclipse.swt.events.SegmentEvent)
	GetSegments(a EventsSegmentEventInterface) 
}

type EventsSegmentListener struct {
	*javabind.Callable
}

// public abstract void getSegments(org.eclipse.swt.events.SegmentEvent)
func (jbobject *EventsSegmentListener) GetSegments(a EventsSegmentEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getSegments", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SegmentEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


