package swt

import "github.com/timob/javabind"

type CustomTextChangingEventInterface interface {
	EventsTypedEventInterface
}

type CustomTextChangingEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.custom.TextChangingEvent(org.eclipse.swt.custom.StyledTextContent)
func NewCustomTextChangingEvent(a CustomStyledTextContentInterface) (*CustomTextChangingEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TextChangingEvent", conv_a.Value().Cast("org/eclipse/swt/custom/StyledTextContent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTextChangingEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *CustomTextChangingEvent) Start() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "start", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomTextChangingEvent) SetFieldStart(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "start", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomTextChangingEvent) NewText() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "newText", "java/lang/String")
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

func (jbobject *CustomTextChangingEvent) SetFieldNewText(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "newText", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomTextChangingEvent) ReplaceCharCount() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "replaceCharCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomTextChangingEvent) SetFieldReplaceCharCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "replaceCharCount", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomTextChangingEvent) NewCharCount() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "newCharCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomTextChangingEvent) SetFieldNewCharCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "newCharCount", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomTextChangingEvent) ReplaceLineCount() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "replaceLineCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomTextChangingEvent) SetFieldReplaceLineCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "replaceLineCount", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomTextChangingEvent) NewLineCount() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "newLineCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomTextChangingEvent) SetFieldNewLineCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "newLineCount", val)
	if err != nil {
		panic(err)
	}

}


