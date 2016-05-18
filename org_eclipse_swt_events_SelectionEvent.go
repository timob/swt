package swt

import "github.com/timob/javabind"

type EventsSelectionEventInterface interface {
	EventsTypedEventInterface
}

type EventsSelectionEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.SelectionEvent(org.eclipse.swt.widgets.Event)
func NewEventsSelectionEvent(a WidgetsEventInterface) (*EventsSelectionEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/SelectionEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsSelectionEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsSelectionEvent) ToString() string {
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

func (jbobject *EventsSelectionEvent) Item() *WidgetsWidget {
	jret, err := jbobject.GetField(javabind.GetEnv(), "item", "org/eclipse/swt/widgets/Widget")
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

func (jbobject *EventsSelectionEvent) SetFieldItem(val WidgetsWidgetInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "item", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventsSelectionEvent) Detail() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "detail", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsSelectionEvent) SetFieldDetail(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "detail", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSelectionEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsSelectionEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSelectionEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsSelectionEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSelectionEvent) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsSelectionEvent) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSelectionEvent) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsSelectionEvent) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSelectionEvent) StateMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stateMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsSelectionEvent) SetFieldStateMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stateMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSelectionEvent) Text() string {
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

func (jbobject *EventsSelectionEvent) SetFieldText(val string) {
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

func (jbobject *EventsSelectionEvent) Doit() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "doit", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *EventsSelectionEvent) SetFieldDoit(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "doit", val)
	if err != nil {
		panic(err)
	}

}


