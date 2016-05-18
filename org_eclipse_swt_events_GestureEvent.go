package swt

import "github.com/timob/javabind"

type EventsGestureEventInterface interface {
	EventsTypedEventInterface
}

type EventsGestureEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.GestureEvent(org.eclipse.swt.widgets.Event)
func NewEventsGestureEvent(a WidgetsEventInterface) (*EventsGestureEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/GestureEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsGestureEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsGestureEvent) ToString() string {
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

func (jbobject *EventsGestureEvent) StateMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stateMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsGestureEvent) SetFieldStateMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stateMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) Detail() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "detail", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsGestureEvent) SetFieldDetail(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "detail", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsGestureEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsGestureEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) Rotation() float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "rotation", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

func (jbobject *EventsGestureEvent) SetFieldRotation(val float64) {
	err := jbobject.SetField(javabind.GetEnv(), "rotation", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) XDirection() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "xDirection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsGestureEvent) SetFieldXDirection(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "xDirection", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) YDirection() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "yDirection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsGestureEvent) SetFieldYDirection(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "yDirection", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) Magnification() float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "magnification", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

func (jbobject *EventsGestureEvent) SetFieldMagnification(val float64) {
	err := jbobject.SetField(javabind.GetEnv(), "magnification", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsGestureEvent) Doit() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "doit", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *EventsGestureEvent) SetFieldDoit(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "doit", val)
	if err != nil {
		panic(err)
	}

}


