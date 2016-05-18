package swt

import "github.com/timob/javabind"

type WidgetsMessageBoxInterface interface {
	WidgetsDialogInterface

	// public java.lang.String getMessage()
	GetMessage() string

	// public void setMessage(java.lang.String)
	SetMessage(a string) 

	// public int open()
	Open() int
}

type WidgetsMessageBox struct {
	WidgetsDialog
}

// public org.eclipse.swt.widgets.MessageBox(org.eclipse.swt.widgets.Shell)
func NewWidgetsMessageBox(a WidgetsShellInterface) (*WidgetsMessageBox) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/MessageBox", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMessageBox{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.MessageBox(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsMessageBox2(a WidgetsShellInterface, b int) (*WidgetsMessageBox) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/MessageBox", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMessageBox{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getMessage()
func (jbobject *WidgetsMessageBox) GetMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessage", "java/lang/String")
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

// public void setMessage(java.lang.String)
func (jbobject *WidgetsMessageBox) SetMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public int open()
func (jbobject *WidgetsMessageBox) Open() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "open", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


