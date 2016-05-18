package swt

import "github.com/timob/javabind"

type EventsKeyAdapterInterface interface {
	JavaLangObjectInterface

	// public void keyPressed(org.eclipse.swt.events.KeyEvent)
	KeyPressed(a EventsKeyEventInterface) 

	// public void keyReleased(org.eclipse.swt.events.KeyEvent)
	KeyReleased(a EventsKeyEventInterface) 
}

type EventsKeyAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.events.KeyAdapter()
func NewEventsKeyAdapter() (*EventsKeyAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/KeyAdapter")
	if err != nil {
		panic(err)
	}
	x := &EventsKeyAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void keyPressed(org.eclipse.swt.events.KeyEvent)
func (jbobject *EventsKeyAdapter) KeyPressed(a EventsKeyEventInterface)  {
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

// public void keyReleased(org.eclipse.swt.events.KeyEvent)
func (jbobject *EventsKeyAdapter) KeyReleased(a EventsKeyEventInterface)  {
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


