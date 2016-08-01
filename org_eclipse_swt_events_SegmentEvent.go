package swt

import "github.com/timob/javabind"

type EventsSegmentEventInterface interface {
	EventsTypedEventInterface
}

type EventsSegmentEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.events.SegmentEvent(org.eclipse.swt.widgets.Event)
func NewEventsSegmentEvent(a WidgetsEventInterface) (*EventsSegmentEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/SegmentEvent", conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventsSegmentEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *EventsSegmentEvent) LineOffset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventsSegmentEvent) SetFieldLineOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "lineOffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSegmentEvent) LineText() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineText", "java/lang/String")
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

func (jbobject *EventsSegmentEvent) SetFieldLineText(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "lineText", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventsSegmentEvent) Segments() []int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "segments", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

func (jbobject *EventsSegmentEvent) SetFieldSegments(val []int) {
	err := jbobject.SetField(javabind.GetEnv(), "segments", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventsSegmentEvent) SegmentsChars() []uint16 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "segmentsChars", javabind.Char | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]uint16)
}

func (jbobject *EventsSegmentEvent) SetFieldSegmentsChars(val []uint16) {
	err := jbobject.SetField(javabind.GetEnv(), "segmentsChars", val)
	if err != nil {
		panic(err)
	}

}


