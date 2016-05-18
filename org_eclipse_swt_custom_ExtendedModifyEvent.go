package swt

import "github.com/timob/javabind"

type CustomExtendedModifyEventInterface interface {
	EventsTypedEventInterface
}

type CustomExtendedModifyEvent struct {
	EventsTypedEvent
}

func (jbobject *CustomExtendedModifyEvent) Start() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "start", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomExtendedModifyEvent) SetFieldStart(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "start", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomExtendedModifyEvent) Length() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "length", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomExtendedModifyEvent) SetFieldLength(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "length", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomExtendedModifyEvent) ReplacedText() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "replacedText", "java/lang/String")
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

func (jbobject *CustomExtendedModifyEvent) SetFieldReplacedText(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "replacedText", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


