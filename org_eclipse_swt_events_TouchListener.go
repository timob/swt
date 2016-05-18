package swt

import "github.com/timob/javabind"

type EventsTouchListenerInterface interface {

	// public abstract void touch(org.eclipse.swt.events.TouchEvent)
	Touch(a EventsTouchEventInterface) 
}

type EventsTouchListener struct {
	*javabind.Callable
}

// public abstract void touch(org.eclipse.swt.events.TouchEvent)
func (jbobject *EventsTouchListener) Touch(a EventsTouchEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "touch", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TouchEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


