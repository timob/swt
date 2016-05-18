package swt

import "github.com/timob/javabind"

type EventsFocusAdapterInterface interface {
	JavaLangObjectInterface

	// public void focusGained(org.eclipse.swt.events.FocusEvent)
	FocusGained(a EventsFocusEventInterface) 

	// public void focusLost(org.eclipse.swt.events.FocusEvent)
	FocusLost(a EventsFocusEventInterface) 
}

type EventsFocusAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.FocusAdapter()
func NewEventsFocusAdapter() (*EventsFocusAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/FocusAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsFocusAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void focusGained(org.eclipse.swt.events.FocusEvent)
func (jbobject *EventsFocusAdapter) FocusGained(a EventsFocusEventInterface)  {
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

// public void focusLost(org.eclipse.swt.events.FocusEvent)
func (jbobject *EventsFocusAdapter) FocusLost(a EventsFocusEventInterface)  {
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


