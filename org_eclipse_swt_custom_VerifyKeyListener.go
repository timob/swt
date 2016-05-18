package swt

import "github.com/timob/javabind"

type CustomVerifyKeyListenerInterface interface {

	// public abstract void verifyKey(org.eclipse.swt.events.VerifyEvent)
	VerifyKey(a EventsVerifyEventInterface) 
}

type CustomVerifyKeyListener struct {
	*javabind.Callable
}

// public abstract void verifyKey(org.eclipse.swt.events.VerifyEvent)
func (jbobject *CustomVerifyKeyListener) VerifyKey(a EventsVerifyEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "verifyKey", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/VerifyEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


