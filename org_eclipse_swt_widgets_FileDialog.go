package swt

import "github.com/timob/javabind"

type WidgetsFileDialogInterface interface {
	WidgetsDialogInterface

	// public java.lang.String getFileName()
	GetFileName() string

	// public java.lang.String[] getFileNames()
	GetFileNames() []string

	// public java.lang.String[] getFilterExtensions()
	GetFilterExtensions() []string

	// public int getFilterIndex()
	GetFilterIndex() int

	// public java.lang.String[] getFilterNames()
	GetFilterNames() []string

	// public java.lang.String getFilterPath()
	GetFilterPath() string

	// public boolean getOverwrite()
	GetOverwrite() bool

	// public java.lang.String open()
	Open() string

	// public void setFileName(java.lang.String)
	SetFileName(a string) 

	// public void setFilterExtensions(java.lang.String[])
	SetFilterExtensions(a []string) 

	// public void setFilterIndex(int)
	SetFilterIndex(a int) 

	// public void setFilterNames(java.lang.String[])
	SetFilterNames(a []string) 

	// public void setFilterPath(java.lang.String)
	SetFilterPath(a string) 

	// public void setOverwrite(boolean)
	SetOverwrite(a bool) 
}

type WidgetsFileDialog struct {
	WidgetsDialog
}

// public org.eclipse.swt.widgets.FileDialog(org.eclipse.swt.widgets.Shell)
func NewWidgetsFileDialog(a WidgetsShellInterface) (*WidgetsFileDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/FileDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsFileDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.FileDialog(org.eclipse.swt.widgets.Shell, int)
func NewWidgetsFileDialog2(a WidgetsShellInterface, b int) (*WidgetsFileDialog) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/FileDialog", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsFileDialog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getFileName()
func (jbobject *WidgetsFileDialog) GetFileName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFileName", "java/lang/String")
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

// public java.lang.String[] getFileNames()
func (jbobject *WidgetsFileDialog) GetFileNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFileNames", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String[] getFilterExtensions()
func (jbobject *WidgetsFileDialog) GetFilterExtensions() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFilterExtensions", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getFilterIndex()
func (jbobject *WidgetsFileDialog) GetFilterIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFilterIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String[] getFilterNames()
func (jbobject *WidgetsFileDialog) GetFilterNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFilterNames", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String getFilterPath()
func (jbobject *WidgetsFileDialog) GetFilterPath() string {
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

// public boolean getOverwrite()
func (jbobject *WidgetsFileDialog) GetOverwrite() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOverwrite", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String open()
func (jbobject *WidgetsFileDialog) Open() string {
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

// public void setFileName(java.lang.String)
func (jbobject *WidgetsFileDialog) SetFileName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFileName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFilterExtensions(java.lang.String[])
func (jbobject *WidgetsFileDialog) SetFilterExtensions(a []string)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFilterExtensions", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFilterIndex(int)
func (jbobject *WidgetsFileDialog) SetFilterIndex(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFilterIndex", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFilterNames(java.lang.String[])
func (jbobject *WidgetsFileDialog) SetFilterNames(a []string)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFilterNames", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFilterPath(java.lang.String)
func (jbobject *WidgetsFileDialog) SetFilterPath(a string)  {
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

// public void setOverwrite(boolean)
func (jbobject *WidgetsFileDialog) SetOverwrite(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOverwrite", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


