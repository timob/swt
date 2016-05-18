package swt

import "github.com/timob/javabind"

type EventsTreeAdapterInterface interface {
	JavaLangObjectInterface

	// public void treeCollapsed(org.eclipse.swt.events.TreeEvent)
	TreeCollapsed(a EventsTreeEventInterface) 

	// public void treeExpanded(org.eclipse.swt.events.TreeEvent)
	TreeExpanded(a EventsTreeEventInterface) 
}

type EventsTreeAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.TreeAdapter()
func NewEventsTreeAdapter() (*EventsTreeAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TreeAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsTreeAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void treeCollapsed(org.eclipse.swt.events.TreeEvent)
func (jbobject *EventsTreeAdapter) TreeCollapsed(a EventsTreeEventInterface)  {
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

// public void treeExpanded(org.eclipse.swt.events.TreeEvent)
func (jbobject *EventsTreeAdapter) TreeExpanded(a EventsTreeEventInterface)  {
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


