package swt

import "github.com/timob/javabind"

type WidgetsDialogInterface interface {
	JavaLangObjectInterface

	// public org.eclipse.swt.widgets.Shell getParent()
	GetParent() *WidgetsShell

	// public int getStyle()
	GetStyle() int

	// public java.lang.String getText()
	GetText() string

	// public void setText(java.lang.String)
	SetText(a string) 
}

type WidgetsDialog struct {
	JavaLangObject
}

// public org.eclipse.swt.widgets.Dialog(org.eclipse.swt.widgets.Shell)
func NewWidgetsDialog(a WidgetsShellInterface) (*WidgetsDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Dialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Dialog(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsDialog2(a WidgetsShellInterface, b int) (*WidgetsDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Dialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Shell getParent()
func (jbobject *WidgetsDialog) GetParent() *WidgetsShell {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Shell")
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
	unique_x := &WidgetsShell{}
	unique_x.Callable = dst
	return unique_x
}

// public int getStyle()
func (jbobject *WidgetsDialog) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getText()
func (jbobject *WidgetsDialog) GetText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getText", "java/lang/String")
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

// public void setText(java.lang.String)
func (jbobject *WidgetsDialog) SetText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


