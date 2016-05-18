package swt

import "github.com/timob/javabind"

type LayoutFillLayoutInterface interface {
	WidgetsLayoutInterface
}

type LayoutFillLayout struct {
	WidgetsLayout
}

// public org.eclipse.swt.layout.FillLayout()
func NewLayoutFillLayout() (*LayoutFillLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FillLayout")
	if err != nil {
		panic(err)
	}
	x := &LayoutFillLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FillLayout(int)
func NewLayoutFillLayout2(a int) (*LayoutFillLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FillLayout", a)
	if err != nil {
		panic(err)
	}
	x := &LayoutFillLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutFillLayout) ToString() string {
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

func (jbobject *LayoutFillLayout) Type() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "type", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFillLayout) SetFieldType(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "type", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFillLayout) MarginWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFillLayout) SetFieldMarginWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFillLayout) MarginHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFillLayout) SetFieldMarginHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFillLayout) Spacing() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "spacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFillLayout) SetFieldSpacing(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "spacing", val)
	if err != nil {
		panic(err)
	}

}


