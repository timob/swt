package swt

import "github.com/timob/javabind"

type EventsDragDetectListenerInterface interface {

	// public abstract void dragDetected(org.eclipse.swt.events.DragDetectEvent)
	DragDetected(a EventsDragDetectEventInterface) 
}

type EventsDragDetectListener struct {
	*javabind.Callable
}

// public abstract void dragDetected(org.eclipse.swt.events.DragDetectEvent)
func (jbobject *EventsDragDetectListener) DragDetected(a EventsDragDetectEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dragDetected", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/DragDetectEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


