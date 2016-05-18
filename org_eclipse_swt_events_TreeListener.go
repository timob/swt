package swt

import "github.com/timob/javabind"

type EventsTreeListenerInterface interface {

	// public abstract void treeCollapsed(org.eclipse.swt.events.TreeEvent)
	TreeCollapsed(a EventsTreeEventInterface) 

	// public abstract void treeExpanded(org.eclipse.swt.events.TreeEvent)
	TreeExpanded(a EventsTreeEventInterface) 
}

type EventsTreeListener struct {
	*javabind.Callable
}

// public abstract void treeCollapsed(org.eclipse.swt.events.TreeEvent)
func (jbobject *EventsTreeListener) TreeCollapsed(a EventsTreeEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "treeCollapsed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TreeEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void treeExpanded(org.eclipse.swt.events.TreeEvent)
func (jbobject *EventsTreeListener) TreeExpanded(a EventsTreeEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "treeExpanded", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TreeEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


