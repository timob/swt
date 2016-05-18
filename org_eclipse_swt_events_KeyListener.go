package swt

import "github.com/timob/javabind"

type EventsKeyListenerInterface interface {

	// public abstract void keyPressed(org.eclipse.swt.events.KeyEvent)
	KeyPressed(a EventsKeyEventInterface) 

	// public abstract void keyReleased(org.eclipse.swt.events.KeyEvent)
	KeyReleased(a EventsKeyEventInterface) 
}

type EventsKeyListener struct {
	*javabind.Callable
}

// public abstract void keyPressed(org.eclipse.swt.events.KeyEvent)
func (jbobject *EventsKeyListener) KeyPressed(a EventsKeyEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "keyPressed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/KeyEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void keyReleased(org.eclipse.swt.events.KeyEvent)
func (jbobject *EventsKeyListener) KeyReleased(a EventsKeyEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "keyReleased", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/KeyEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


