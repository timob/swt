package swt

import "github.com/timob/javabind"

type EventsPaintListenerInterface interface {

	// public abstract void paintControl(org.eclipse.swt.events.PaintEvent)
	PaintControl(a EventsPaintEventInterface) 
}

type EventsPaintListener struct {
	*javabind.Callable
}

// public abstract void paintControl(org.eclipse.swt.events.PaintEvent)
func (jbobject *EventsPaintListener) PaintControl(a EventsPaintEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "paintControl", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/PaintEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


