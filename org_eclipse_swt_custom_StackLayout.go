package swt

import "github.com/timob/javabind"

type CustomStackLayoutInterface interface {
	WidgetsLayoutInterface
}

type CustomStackLayout struct {
	WidgetsLayout
}

// public org.eclipse.swt.custom.StackLayout()
func NewCustomStackLayout() (*CustomStackLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StackLayout")
	if err != nil {
		panic(err)
	}
	x := &CustomStackLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *CustomStackLayout) ToString() string {
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

func (jbobject *CustomStackLayout) MarginWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomStackLayout) SetFieldMarginWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStackLayout) MarginHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomStackLayout) SetFieldMarginHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStackLayout) TopControl() *WidgetsControl {
	jret, err := jbobject.GetField(javabind.GetEnv(), "topControl", "org/eclipse/swt/widgets/Control")
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
	unique_x := &WidgetsControl{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *CustomStackLayout) SetFieldTopControl(val WidgetsControlInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "topControl", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


