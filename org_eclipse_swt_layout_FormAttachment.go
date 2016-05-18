package swt

import "github.com/timob/javabind"

type LayoutFormAttachmentInterface interface {
	JavaLangObjectInterface
}

type LayoutFormAttachment struct {
	JavaLangObject
}

// public org.eclipse.swt.layout.FormAttachment()
func NewLayoutFormAttachment() (*LayoutFormAttachment) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormAttachment")
	if err != nil {
		panic(err)
	}
	x := &LayoutFormAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FormAttachment(int)
func NewLayoutFormAttachment2(a int) (*LayoutFormAttachment) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormAttachment", a)
	if err != nil {
		panic(err)
	}
	x := &LayoutFormAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FormAttachment(int, int)
func NewLayoutFormAttachment3(a int, b int) (*LayoutFormAttachment) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormAttachment", a, b)
	if err != nil {
		panic(err)
	}
	x := &LayoutFormAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FormAttachment(int, int, int)
func NewLayoutFormAttachment4(a int, b int, c int) (*LayoutFormAttachment) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormAttachment", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &LayoutFormAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FormAttachment(org.eclipse.swt.widgets.Control)
func NewLayoutFormAttachment5(a WidgetsControlInterface) (*LayoutFormAttachment) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormAttachment", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &LayoutFormAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FormAttachment(org.eclipse.swt.widgets.Control, int)
func NewLayoutFormAttachment6(a WidgetsControlInterface, b int) (*LayoutFormAttachment) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormAttachment", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &LayoutFormAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.layout.FormAttachment(org.eclipse.swt.widgets.Control, int, int)
func NewLayoutFormAttachment7(a WidgetsControlInterface, b int, c int) (*LayoutFormAttachment) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/layout/FormAttachment", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &LayoutFormAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String toString()
func (jbobject *LayoutFormAttachment) ToString() string {
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

func (jbobject *LayoutFormAttachment) Numerator() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "numerator", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormAttachment) SetFieldNumerator(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "numerator", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormAttachment) Denominator() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "denominator", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormAttachment) SetFieldDenominator(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "denominator", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormAttachment) Offset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "offset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormAttachment) SetFieldOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "offset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *LayoutFormAttachment) Control() *WidgetsControl {
	jret, err := jbobject.GetField(javabind.GetEnv(), "control", "org/eclipse/swt/widgets/Control")
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

func (jbobject *LayoutFormAttachment) SetFieldControl(val WidgetsControlInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "control", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *LayoutFormAttachment) Alignment() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "alignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *LayoutFormAttachment) SetFieldAlignment(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "alignment", val)
	if err != nil {
		panic(err)
	}

}


