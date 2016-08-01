package swt

import "github.com/timob/javabind"

type CustomMovementEventInterface interface {
	EventsTypedEventInterface
}

type CustomMovementEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.custom.MovementEvent(org.eclipse.swt.custom.StyledTextEvent)
func NewCustomMovementEvent(a CustomStyledTextEventInterface) (*CustomMovementEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/MovementEvent", conv_a.Value().Cast("org/eclipse/swt/custom/StyledTextEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomMovementEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *CustomMovementEvent) LineOffset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomMovementEvent) SetFieldLineOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "lineOffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomMovementEvent) LineText() string {
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

func (jbobject *CustomMovementEvent) SetFieldLineText(val string) {
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

func (jbobject *CustomMovementEvent) Offset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "offset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomMovementEvent) SetFieldOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "offset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomMovementEvent) NewOffset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "newOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomMovementEvent) SetFieldNewOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "newOffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomMovementEvent) Movement() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "movement", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomMovementEvent) SetFieldMovement(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "movement", val)
	if err != nil {
		panic(err)
	}

}


