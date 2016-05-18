package swt

import "github.com/timob/javabind"

type WidgetsDirectoryDialogInterface interface {
	WidgetsDialogInterface

	// public java.lang.String getFilterPath()
	GetFilterPath() string

	// public java.lang.String getMessage()
	GetMessage() string

	// public java.lang.String open()
	Open() string

	// public void setFilterPath(java.lang.String)
	SetFilterPath(a string) 

	// public void setMessage(java.lang.String)
	SetMessage(a string) 
}

type WidgetsDirectoryDialog struct {
	WidgetsDialog
}

// public org.eclipse.swt.widgets.DirectoryDialog(org.eclipse.swt.widgets.Shell)
func NewWidgetsDirectoryDialog(a WidgetsShellInterface) (*WidgetsDirectoryDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/DirectoryDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsDirectoryDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.DirectoryDialog(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsDirectoryDialog2(a WidgetsShellInterface, b int) (*WidgetsDirectoryDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/DirectoryDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsDirectoryDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getFilterPath()
func (jbobject *WidgetsDirectoryDialog) GetFilterPath() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFilterPath", "java/lang/String")
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

// public java.lang.String getMessage()
func (jbobject *WidgetsDirectoryDialog) GetMessage() string {
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

// public java.lang.String open()
func (jbobject *WidgetsDirectoryDialog) Open() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "open", "java/lang/String")
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

// public void setFilterPath(java.lang.String)
func (jbobject *WidgetsDirectoryDialog) SetFilterPath(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFilterPath", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMessage(java.lang.String)
func (jbobject *WidgetsDirectoryDialog) SetMessage(a string)  {
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


