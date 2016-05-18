package swt

import "github.com/timob/javabind"

type EventsFocusListenerInterface interface {

	// public abstract void focusGained(org.eclipse.swt.events.FocusEvent)
	FocusGained(a EventsFocusEventInterface) 

	// public abstract void focusLost(org.eclipse.swt.events.FocusEvent)
	FocusLost(a EventsFocusEventInterface) 
}

type EventsFocusListener struct {
	*javabind.Callable
}

// public abstract void focusGained(org.eclipse.swt.events.FocusEvent)
func (jbobject *EventsFocusListener) FocusGained(a EventsFocusEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "focusGained", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/FocusEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void focusLost(org.eclipse.swt.events.FocusEvent)
func (jbobject *EventsFocusListener) FocusLost(a EventsFocusEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "focusLost", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/FocusEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


