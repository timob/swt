package swt

import "github.com/timob/javabind"

type EventsKeyEventInterface interface {
	EventsTypedEventInterface
}

type EventsKeyEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.KeyEvent(org.eclipse.swt.widgets.Event)
func NewEventsKeyEvent(a WidgetsEventInterface) (*EventsKeyEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/KeyEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsKeyEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsKeyEvent) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *EventsKeyEvent) KeyCode() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "keyCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsKeyEvent) SetFieldKeyCode(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "keyCode", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsKeyEvent) KeyLocation() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "keyLocation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsKeyEvent) SetFieldKeyLocation(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "keyLocation", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsKeyEvent) StateMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stateMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsKeyEvent) SetFieldStateMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stateMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsKeyEvent) Doit() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "doit", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *EventsKeyEvent) SetFieldDoit(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "doit", val)
	if err != nil {
		panic(err)
	}

}


