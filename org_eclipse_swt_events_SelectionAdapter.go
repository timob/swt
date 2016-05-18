package swt

import "github.com/timob/javabind"

type EventsSelectionAdapterInterface interface {
	JavaLangObjectInterface

	// public void widgetSelected(org.eclipse.swt.events.SelectionEvent)
	WidgetSelected(a EventsSelectionEventInterface) 

	// public void widgetDefaultSelected(org.eclipse.swt.events.SelectionEvent)
	WidgetDefaultSelected(a EventsSelectionEventInterface) 
}

type EventsSelectionAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.SelectionAdapter()
func NewEventsSelectionAdapter() (*EventsSelectionAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/SelectionAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsSelectionAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void widgetSelected(org.eclipse.swt.events.SelectionEvent)
func (jbobject *EventsSelectionAdapter) WidgetSelected(a EventsSelectionEventInterface)  {
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

// public void widgetDefaultSelected(org.eclipse.swt.events.SelectionEvent)
func (jbobject *EventsSelectionAdapter) WidgetDefaultSelected(a EventsSelectionEventInterface)  {
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


