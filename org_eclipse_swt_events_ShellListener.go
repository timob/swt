package swt

import "github.com/timob/javabind"

type EventsShellListenerInterface interface {

	// public abstract void shellActivated(org.eclipse.swt.events.ShellEvent)
	ShellActivated(a EventsShellEventInterface) 

	// public abstract void shellClosed(org.eclipse.swt.events.ShellEvent)
	ShellClosed(a EventsShellEventInterface) 

	// public abstract void shellDeactivated(org.eclipse.swt.events.ShellEvent)
	ShellDeactivated(a EventsShellEventInterface) 

	// public abstract void shellDeiconified(org.eclipse.swt.events.ShellEvent)
	ShellDeiconified(a EventsShellEventInterface) 

	// public abstract void shellIconified(org.eclipse.swt.events.ShellEvent)
	ShellIconified(a EventsShellEventInterface) 
}

type EventsShellListener struct {
	*javabind.Callable
}

// public abstract void shellActivated(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellListener) ShellActivated(a EventsShellEventInterface)  {
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

// public abstract void shellClosed(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellListener) ShellClosed(a EventsShellEventInterface)  {
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

// public abstract void shellDeactivated(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellListener) ShellDeactivated(a EventsShellEventInterface)  {
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

// public abstract void shellDeiconified(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellListener) ShellDeiconified(a EventsShellEventInterface)  {
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

// public abstract void shellIconified(org.eclipse.swt.events.ShellEvent)
func (jbobject *EventsShellListener) ShellIconified(a EventsShellEventInterface)  {
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


