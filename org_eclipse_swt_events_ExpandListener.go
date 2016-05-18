package swt

import "github.com/timob/javabind"

type EventsExpandListenerInterface interface {

	// public abstract void itemCollapsed(org.eclipse.swt.events.ExpandEvent)
	ItemCollapsed(a EventsExpandEventInterface) 

	// public abstract void itemExpanded(org.eclipse.swt.events.ExpandEvent)
	ItemExpanded(a EventsExpandEventInterface) 
}

type EventsExpandListener struct {
	*javabind.Callable
}

// public abstract void itemCollapsed(org.eclipse.swt.events.ExpandEvent)
func (jbobject *EventsExpandListener) ItemCollapsed(a EventsExpandEventInterface)  {
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

// public abstract void itemExpanded(org.eclipse.swt.events.ExpandEvent)
func (jbobject *EventsExpandListener) ItemExpanded(a EventsExpandEventInterface)  {
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


