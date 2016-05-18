package swt

import "github.com/timob/javabind"

type LayoutRowLayoutInterface interface {
	WidgetsLayoutInterface
}

type LayoutRowLayout struct {
	WidgetsLayout
}

// public org.eclipse.swt.layout.RowLayout()
func NewLayoutRowLayout() (*LayoutRowLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/RowLayout")
	if err != nil {
		panic(err)
	}
	x := &LayoutRowLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.RowLayout(int)
func NewLayoutRowLayout2(a int) (*LayoutRowLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/RowLayout", a)
	if err != nil {
		panic(err)
	}
	x := &LayoutRowLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutRowLayout) ToString() string {
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

func (jbobject *LayoutRowLayout) Type() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "type", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldType(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "type", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) MarginWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldMarginWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) MarginHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldMarginHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) Spacing() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "spacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldSpacing(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "spacing", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) Wrap() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "wrap", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutRowLayout) SetFieldWrap(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "wrap", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) Pack() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "pack", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutRowLayout) SetFieldPack(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "pack", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) Fill() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "fill", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutRowLayout) SetFieldFill(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "fill", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) Center() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "center", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutRowLayout) SetFieldCenter(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "center", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) Justify() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "justify", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutRowLayout) SetFieldJustify(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "justify", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) MarginLeft() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginLeft", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldMarginLeft(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginLeft", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) MarginTop() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginTop", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldMarginTop(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginTop", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) MarginRight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginRight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldMarginRight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginRight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutRowLayout) MarginBottom() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "marginBottom", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutRowLayout) SetFieldMarginBottom(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "marginBottom", val)
	if err != nil {
		panic(err)
	}

}


