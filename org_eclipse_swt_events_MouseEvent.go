package swt

import "github.com/timob/javabind"

type EventsMouseEventInterface interface {
	EventsTypedEventInterface
}

type EventsMouseEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.MouseEvent(org.eclipse.swt.widgets.Event)
func NewEventsMouseEvent(a WidgetsEventInterface) (*EventsMouseEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MouseEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsMouseEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsMouseEvent) ToString() string {
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

func (jbobject *EventsMouseEvent) Button() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "button", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMouseEvent) SetFieldButton(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "button", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsMouseEvent) StateMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stateMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMouseEvent) SetFieldStateMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stateMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsMouseEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMouseEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsMouseEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMouseEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsMouseEvent) Count() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "count", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMouseEvent) SetFieldCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "count", val)
	if err != nil {
		panic(err)
	}

}


