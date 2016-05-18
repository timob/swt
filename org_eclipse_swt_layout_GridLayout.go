package swt

import "github.com/timob/javabind"

type LayoutGridLayoutInterface interface {
	WidgetsLayoutInterface
}

type LayoutGridLayout struct {
	WidgetsLayout
}

// public org.eclipse.swt.layout.GridLayout()
func NewLayoutGridLayout() (*LayoutGridLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/GridLayout")
	if err != nil {
		panic(err)
	}
	x := &LayoutGridLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.GridLayout(int, boolean)
func NewLayoutGridLayout2(a int, b bool) (*LayoutGridLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/GridLayout", a, b)
	if err != nil {
		panic(err)
	}
	x := &LayoutGridLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutGridLayout) ToString() string {
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

func (jbobject *LayoutGridLayout) NumColumns() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "numColumns", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldNumColumns(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "numColumns", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) MakeColumnsEqualWidth() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "makeColumnsEqualWidth", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutGridLayout) SetFieldMakeColumnsEqualWidth(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "makeColumnsEqualWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) MarginWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldMarginWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) MarginHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldMarginHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) MarginLeft() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginLeft", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldMarginLeft(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginLeft", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) MarginTop() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginTop", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldMarginTop(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginTop", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) MarginRight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginRight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldMarginRight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginRight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) MarginBottom() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginBottom", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldMarginBottom(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginBottom", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) HorizontalSpacing() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "horizontalSpacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldHorizontalSpacing(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "horizontalSpacing", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridLayout) VerticalSpacing() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "verticalSpacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridLayout) SetFieldVerticalSpacing(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "verticalSpacing", val)
	if err != nil {
		panic(err)
	}

}


