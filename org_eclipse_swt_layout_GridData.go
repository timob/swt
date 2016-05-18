package swt

import "github.com/timob/javabind"

type LayoutGridDataInterface interface {
	JavaLangObjectInterface
}

type LayoutGridData struct {
	JavaLangObject
}

// public org.eclipse.swt.layout.GridData()
func NewLayoutGridData() (*LayoutGridData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/GridData")
	if err != nil {
		panic(err)
	}
	x := &LayoutGridData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.GridData(int)
func NewLayoutGridData2(a int) (*LayoutGridData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/GridData", a)
	if err != nil {
		panic(err)
	}
	x := &LayoutGridData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.GridData(int, int, boolean, boolean)
func NewLayoutGridData4(a int, b int, c bool, d bool) (*LayoutGridData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/GridData", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &LayoutGridData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.GridData(int, int, boolean, boolean, int, int)
func NewLayoutGridData5(a int, b int, c bool, d bool, e int, f int) (*LayoutGridData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/GridData", a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}
	x := &LayoutGridData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.GridData(int, int)
func NewLayoutGridData3(a int, b int) (*LayoutGridData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/GridData", a, b)
	if err != nil {
		panic(err)
	}
	x := &LayoutGridData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutGridData) ToString() string {
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

func (jbobject *LayoutGridData) VerticalAlignment() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "verticalAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldVerticalAlignment(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "verticalAlignment", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HorizontalAlignment() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "horizontalAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHorizontalAlignment(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "horizontalAlignment", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) WidthHint() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "widthHint", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldWidthHint(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "widthHint", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HeightHint() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "heightHint", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHeightHint(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "heightHint", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HorizontalIndent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "horizontalIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHorizontalIndent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "horizontalIndent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) VerticalIndent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "verticalIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldVerticalIndent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "verticalIndent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HorizontalSpan() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "horizontalSpan", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHorizontalSpan(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "horizontalSpan", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) VerticalSpan() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "verticalSpan", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldVerticalSpan(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "verticalSpan", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) GrabExcessHorizontalSpace() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "grabExcessHorizontalSpace", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutGridData) SetFieldGrabExcessHorizontalSpace(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "grabExcessHorizontalSpace", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) GrabExcessVerticalSpace() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "grabExcessVerticalSpace", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutGridData) SetFieldGrabExcessVerticalSpace(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "grabExcessVerticalSpace", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) MinimumWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "minimumWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldMinimumWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "minimumWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) MinimumHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "minimumHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldMinimumHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "minimumHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) Exclude() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "exclude", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *LayoutGridData) SetFieldExclude(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "exclude", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) BEGINNING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "BEGINNING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldBEGINNING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "BEGINNING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) CENTER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "CENTER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldCENTER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "CENTER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldEND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) FILL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "FILL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldFILL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "FILL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) VERTICAL_ALIGN_BEGINNING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_BEGINNING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldVERTICAL_ALIGN_BEGINNING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_BEGINNING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) VERTICAL_ALIGN_CENTER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_CENTER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldVERTICAL_ALIGN_CENTER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_CENTER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) VERTICAL_ALIGN_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldVERTICAL_ALIGN_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) VERTICAL_ALIGN_FILL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_FILL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldVERTICAL_ALIGN_FILL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "VERTICAL_ALIGN_FILL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HORIZONTAL_ALIGN_BEGINNING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_BEGINNING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHORIZONTAL_ALIGN_BEGINNING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_BEGINNING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HORIZONTAL_ALIGN_CENTER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_CENTER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHORIZONTAL_ALIGN_CENTER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_CENTER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HORIZONTAL_ALIGN_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHORIZONTAL_ALIGN_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) HORIZONTAL_ALIGN_FILL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_FILL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldHORIZONTAL_ALIGN_FILL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "HORIZONTAL_ALIGN_FILL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) GRAB_HORIZONTAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "GRAB_HORIZONTAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldGRAB_HORIZONTAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "GRAB_HORIZONTAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) GRAB_VERTICAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "GRAB_VERTICAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldGRAB_VERTICAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "GRAB_VERTICAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) FILL_VERTICAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "FILL_VERTICAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldFILL_VERTICAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "FILL_VERTICAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) FILL_HORIZONTAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "FILL_HORIZONTAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldFILL_HORIZONTAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "FILL_HORIZONTAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutGridData) FILL_BOTH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/layout/GridData", "FILL_BOTH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutGridData) SetFieldFILL_BOTH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/layout/GridData", "FILL_BOTH", val)
	if err != nil {
		panic(err)
	}

}


