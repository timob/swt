package swt

import "github.com/timob/javabind"

type EventsVerifyListenerInterface interface {

	// public abstract void verifyText(org.eclipse.swt.events.VerifyEvent)
	VerifyText(a EventsVerifyEventInterface) 
}

type EventsVerifyListener struct {
	*javabind.Callable
}

// public abstract void verifyText(org.eclipse.swt.events.VerifyEvent)
func (jbobject *EventsVerifyListener) VerifyText(a EventsVerifyEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "verifyText", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/VerifyEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


