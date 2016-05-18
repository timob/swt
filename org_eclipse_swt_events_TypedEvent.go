package swt

import "github.com/timob/javabind"

type EventsTypedEventInterface interface {

	// public java.lang.String toString()
	ToString() string
}

type EventsTypedEvent struct {
	*javabind.Callable
}

// public org.eclipse.swt.events.TypedEvent(java.lang.Object)
func NewEventsTypedEvent(a interface{}) (*EventsTypedEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TypedEvent", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsTypedEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.events.TypedEvent(org.eclipse.swt.widgets.Event)
func NewEventsTypedEvent2(a WidgetsEventInterface) (*EventsTypedEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TypedEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsTypedEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsTypedEvent) ToString() string {
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

func (jbobject *EventsTypedEvent) Display() *WidgetsDisplay {
	jret, err := jbobject.GetField(javabind.GetEnv(), "display", "org/eclipse/swt/widgets/Display")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsDisplay{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventsTypedEvent) SetFieldDisplay(val WidgetsDisplayInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "display", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventsTypedEvent) Widget() *WidgetsWidget {
	jret, err := jbobject.GetField(javabind.GetEnv(), "widget", "org/eclipse/swt/widgets/Widget")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &WidgetsWidget{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventsTypedEvent) SetFieldWidget(val WidgetsWidgetInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "widget", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventsTypedEvent) Time() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "time", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsTypedEvent) SetFieldTime(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "time", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsTypedEvent) Data() *JavaLangObject {
	jret, err := jbobject.GetField(javabind.GetEnv(), "data", "java/lang/Object")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventsTypedEvent) SetFieldData(val JavaLangObjectInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "data", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


