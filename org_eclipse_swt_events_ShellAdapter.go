package swt

import "github.com/timob/javabind"

type EventsShellAdapterInterface interface {
	JavaLangObjectInterface

	// public void shellActivated(org.eclipse.swt.events.ShellEvent)
	ShellActivated(a EventsShellEventInterface) 

	// public void shellClosed(org.eclipse.swt.events.ShellEvent)
	ShellClosed(a EventsShellEventInterface) 

	// public void shellDeactivated(org.eclipse.swt.events.ShellEvent)
	ShellDeactivated(a EventsShellEventInterface) 

	// public void shellDeiconified(org.eclipse.swt.events.ShellEvent)
	ShellDeiconified(a EventsShellEventInterface) 

	// public void shellIconified(org.eclipse.swt.events.ShellEvent)
	ShellIconified(a EventsShellEventInterface) 
}

type EventsShellAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.ShellAdapter()
func NewEventsShellAdapter() (*EventsShellAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ShellAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsShellAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void shellActivated(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellAdapter) ShellActivated(a EventsShellEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shellActivated", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ShellEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void shellClosed(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellAdapter) ShellClosed(a EventsShellEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shellClosed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ShellEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void shellDeactivated(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellAdapter) ShellDeactivated(a EventsShellEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shellDeactivated", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ShellEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void shellDeiconified(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellAdapter) ShellDeiconified(a EventsShellEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shellDeiconified", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ShellEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void shellIconified(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellAdapter) ShellIconified(a EventsShellEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shellIconified", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ShellEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


