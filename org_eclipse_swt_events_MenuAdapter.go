package swt

import "github.com/timob/javabind"

type EventsMenuAdapterInterface interface {
	JavaLangObjectInterface

	// public void menuHidden(org.eclipse.swt.events.MenuEvent)
	MenuHidden(a EventsMenuEventInterface) 

	// public void menuShown(org.eclipse.swt.events.MenuEvent)
	MenuShown(a EventsMenuEventInterface) 
}

type EventsMenuAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.MenuAdapter()
func NewEventsMenuAdapter() (*EventsMenuAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MenuAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsMenuAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void menuHidden(org.eclipse.swt.events.MenuEvent)
func (jbobject *EventsMenuAdapter) MenuHidden(a EventsMenuEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "menuHidden", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void menuShown(org.eclipse.swt.events.MenuEvent)
func (jbobject *EventsMenuAdapter) MenuShown(a EventsMenuEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "menuShown", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


