package swt

import "github.com/timob/javabind"

type LayoutFormLayoutInterface interface {
	WidgetsLayoutInterface
}

type LayoutFormLayout struct {
	WidgetsLayout
}

// public org.eclipse.swt.layout.FormLayout()
func NewLayoutFormLayout() (*LayoutFormLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormLayout")
	if err != nil {
		panic(err)
	}
	x := &LayoutFormLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutFormLayout) ToString() string {
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

func (jbobject *LayoutFormLayout) MarginWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormLayout) SetFieldMarginWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormLayout) MarginHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormLayout) SetFieldMarginHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormLayout) MarginLeft() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginLeft", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormLayout) SetFieldMarginLeft(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginLeft", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormLayout) MarginTop() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginTop", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormLayout) SetFieldMarginTop(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginTop", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormLayout) MarginRight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginRight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormLayout) SetFieldMarginRight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginRight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormLayout) MarginBottom() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginBottom", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormLayout) SetFieldMarginBottom(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginBottom", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormLayout) Spacing() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "spacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormLayout) SetFieldSpacing(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "spacing", val)
	if err != nil {
		panic(err)
	}

}


