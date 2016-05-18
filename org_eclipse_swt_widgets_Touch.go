package swt

import "github.com/timob/javabind"

type WidgetsTouchInterface interface {
	JavaLangObjectInterface
}

type WidgetsTouch struct {
	JavaLangObject
}

// public java.lang.String toString()
func (jbobject *WidgetsTouch) ToString() string {
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

func (jbobject *WidgetsTouch) Id() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "id", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *WidgetsTouch) SetFieldId(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "id", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsTouch) Source() *WidgetsTouchSource {
	jret, err := jbobject.GetField(javabind.GetEnv(), "source", "org/eclipse/swt/widgets/TouchSource")
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
	unique_x := &WidgetsTouchSource{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *WidgetsTouch) SetFieldSource(val WidgetsTouchSourceInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "source", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsTouch) State() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "state", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsTouch) SetFieldState(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "state", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsTouch) Primary() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "primary", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *WidgetsTouch) SetFieldPrimary(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "primary", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsTouch) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsTouch) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsTouch) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsTouch) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}


