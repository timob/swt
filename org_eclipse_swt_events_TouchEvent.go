package swt

import "github.com/timob/javabind"

type EventsTouchEventInterface interface {
	EventsTypedEventInterface
}

type EventsTouchEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.TouchEvent(org.eclipse.swt.widgets.Event)
func NewEventsTouchEvent(a WidgetsEventInterface) (*EventsTouchEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TouchEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsTouchEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsTouchEvent) ToString() string {
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

func (jbobject *EventsTouchEvent) Touches() []*WidgetsTouch {
	jret, err := jbobject.GetField(javabind.GetEnv(), "touches", javabind.ObjectArrayType("org/eclipse/swt/widgets/Touch"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/Touch")
	dst := new([]*WidgetsTouch)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *EventsTouchEvent) SetFieldTouches(val []*WidgetsTouch) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/Touch")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "touches", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventsTouchEvent) StateMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stateMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsTouchEvent) SetFieldStateMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stateMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsTouchEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsTouchEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsTouchEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsTouchEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}


