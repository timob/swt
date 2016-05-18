package swt

import "github.com/timob/javabind"

type EventsMenuDetectEventInterface interface {
	EventsTypedEventInterface
}

type EventsMenuDetectEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.MenuDetectEvent(org.eclipse.swt.widgets.Event)
func NewEventsMenuDetectEvent(a WidgetsEventInterface) (*EventsMenuDetectEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MenuDetectEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsMenuDetectEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsMenuDetectEvent) ToString() string {
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

func (jbobject *EventsMenuDetectEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMenuDetectEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsMenuDetectEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMenuDetectEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsMenuDetectEvent) Doit() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "doit", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *EventsMenuDetectEvent) SetFieldDoit(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "doit", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsMenuDetectEvent) Detail() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "detail", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsMenuDetectEvent) SetFieldDetail(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "detail", val)
	if err != nil {
		panic(err)
	}

}


