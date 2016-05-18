package swt

import "github.com/timob/javabind"

type EventsSelectionListenerInterface interface {

	// public abstract void widgetSelected(org.eclipse.swt.events.SelectionEvent)
	WidgetSelected(a EventsSelectionEventInterface) 

	// public abstract void widgetDefaultSelected(org.eclipse.swt.events.SelectionEvent)
	WidgetDefaultSelected(a EventsSelectionEventInterface) 
}

type EventsSelectionListener struct {
	*javabind.Callable
}

// public abstract void widgetSelected(org.eclipse.swt.events.SelectionEvent)
func (jbobject *EventsSelectionListener) WidgetSelected(a EventsSelectionEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "widgetSelected", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void widgetDefaultSelected(org.eclipse.swt.events.SelectionEvent)
func (jbobject *EventsSelectionListener) WidgetDefaultSelected(a EventsSelectionEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "widgetDefaultSelected", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


