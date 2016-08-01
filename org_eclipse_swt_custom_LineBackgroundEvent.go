package swt

import "github.com/timob/javabind"

type CustomLineBackgroundEventInterface interface {
	EventsTypedEventInterface
}

type CustomLineBackgroundEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.custom.LineBackgroundEvent(org.eclipse.swt.custom.StyledTextEvent)
func NewCustomLineBackgroundEvent(a CustomStyledTextEventInterface) (*CustomLineBackgroundEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/LineBackgroundEvent", conv_a.Value().Cast("org/eclipse/swt/custom/StyledTextEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomLineBackgroundEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *CustomLineBackgroundEvent) LineOffset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomLineBackgroundEvent) SetFieldLineOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "lineOffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineBackgroundEvent) LineText() string {
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

func (jbobject *CustomLineBackgroundEvent) SetFieldLineText(val string) {
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

func (jbobject *CustomLineBackgroundEvent) LineBackground() *GraphicsColor {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineBackground", "org/eclipse/swt/graphics/Color")
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
	unique_x := &GraphicsColor{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomLineBackgroundEvent) SetFieldLineBackground(val GraphicsColorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "lineBackground", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


