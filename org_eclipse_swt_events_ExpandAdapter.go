package swt

import "github.com/timob/javabind"

type EventsExpandAdapterInterface interface {
	JavaLangObjectInterface

	// public void itemCollapsed(org.eclipse.swt.events.ExpandEvent)
	ItemCollapsed(a EventsExpandEventInterface) 

	// public void itemExpanded(org.eclipse.swt.events.ExpandEvent)
	ItemExpanded(a EventsExpandEventInterface) 
}

type EventsExpandAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.ExpandAdapter()
func NewEventsExpandAdapter() (*EventsExpandAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ExpandAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsExpandAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void itemCollapsed(org.eclipse.swt.events.ExpandEvent)
func (jbobject *EventsExpandAdapter) ItemCollapsed(a EventsExpandEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "itemCollapsed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ExpandEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void itemExpanded(org.eclipse.swt.events.ExpandEvent)
func (jbobject *EventsExpandAdapter) ItemExpanded(a EventsExpandEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "itemExpanded", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ExpandEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


