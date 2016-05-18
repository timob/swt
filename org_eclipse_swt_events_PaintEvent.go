package swt

import "github.com/timob/javabind"

type EventsPaintEventInterface interface {
	EventsTypedEventInterface
}

type EventsPaintEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.PaintEvent(org.eclipse.swt.widgets.Event)
func NewEventsPaintEvent(a WidgetsEventInterface) (*EventsPaintEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/PaintEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsPaintEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *EventsPaintEvent) ToString() string {
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

func (jbobject *EventsPaintEvent) Gc() *GraphicsGC {
	jret, err := jbobject.GetField(javabind.GetEnv(), "gc", "org/eclipse/swt/graphics/GC")
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
	unique_x := &GraphicsGC{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventsPaintEvent) SetFieldGc(val GraphicsGCInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "gc", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventsPaintEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsPaintEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsPaintEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsPaintEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsPaintEvent) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsPaintEvent) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsPaintEvent) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsPaintEvent) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsPaintEvent) Count() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "count", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsPaintEvent) SetFieldCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "count", val)
	if err != nil {
		panic(err)
	}

}


