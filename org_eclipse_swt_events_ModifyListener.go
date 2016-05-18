package swt

import "github.com/timob/javabind"

type EventsModifyListenerInterface interface {

	// public abstract void modifyText(org.eclipse.swt.events.ModifyEvent)
	ModifyText(a EventsModifyEventInterface) 
}

type EventsModifyListener struct {
	*javabind.Callable
}

// public abstract void modifyText(org.eclipse.swt.events.ModifyEvent)
func (jbobject *EventsModifyListener) ModifyText(a EventsModifyEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyText", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ModifyEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


