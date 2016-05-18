package swt

import "github.com/timob/javabind"

type EventsVerifyEventInterface interface {
	EventsKeyEventInterface
}

type EventsVerifyEvent struct {
	EventsKeyEvent
}

// public org.eclipse.swt.events.VerifyEvent(org.eclipse.swt.widgets.Event)
func NewEventsVerifyEvent(a WidgetsEventInterface) (*EventsVerifyEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/VerifyEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsVerifyEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsVerifyEvent) ToString() string {
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

func (jbobject *EventsVerifyEvent) Start() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "start", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsVerifyEvent) SetFieldStart(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "start", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsVerifyEvent) End() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "end", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsVerifyEvent) SetFieldEnd(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "end", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsVerifyEvent) Text() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "text", "java/lang/String")
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

func (jbobject *EventsVerifyEvent) SetFieldText(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "text", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


