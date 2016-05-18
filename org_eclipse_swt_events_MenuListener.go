package swt

import "github.com/timob/javabind"

type EventsMenuListenerInterface interface {

	// public abstract void menuHidden(org.eclipse.swt.events.MenuEvent)
	MenuHidden(a EventsMenuEventInterface) 

	// public abstract void menuShown(org.eclipse.swt.events.MenuEvent)
	MenuShown(a EventsMenuEventInterface) 
}

type EventsMenuListener struct {
	*javabind.Callable
}

// public abstract void menuHidden(org.eclipse.swt.events.MenuEvent)
func (jbobject *EventsMenuListener) MenuHidden(a EventsMenuEventInterface)  {
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

// public abstract void menuShown(org.eclipse.swt.events.MenuEvent)
func (jbobject *EventsMenuListener) MenuShown(a EventsMenuEventInterface)  {
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


