package swt

import "github.com/timob/javabind"

type EventsDisposeListenerInterface interface {

	// public abstract void widgetDisposed(org.eclipse.swt.events.DisposeEvent)
	WidgetDisposed(a EventsDisposeEventInterface) 
}

type EventsDisposeListener struct {
	*javabind.Callable
}

// public abstract void widgetDisposed(org.eclipse.swt.events.DisposeEvent)
func (jbobject *EventsDisposeListener) WidgetDisposed(a EventsDisposeEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "widgetDisposed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/DisposeEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


