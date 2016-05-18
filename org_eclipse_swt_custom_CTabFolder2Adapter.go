package swt

import "github.com/timob/javabind"

type CustomCTabFolder2AdapterInterface interface {
	JavaLangObjectInterface

	// public void close(org.eclipse.swt.custom.CTabFolderEvent)
	Close(a CustomCTabFolderEventInterface) 

	// public void minimize(org.eclipse.swt.custom.CTabFolderEvent)
	Minimize(a CustomCTabFolderEventInterface) 

	// public void maximize(org.eclipse.swt.custom.CTabFolderEvent)
	Maximize(a CustomCTabFolderEventInterface) 

	// public void restore(org.eclipse.swt.custom.CTabFolderEvent)
	Restore(a CustomCTabFolderEventInterface) 

	// public void showList(org.eclipse.swt.custom.CTabFolderEvent)
	ShowList(a CustomCTabFolderEventInterface) 
}

type CustomCTabFolder2Adapter struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.CTabFolder2Adapter()
func NewCustomCTabFolder2Adapter() (*CustomCTabFolder2Adapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CTabFolder2Adapter")
	if err != nil {
		panic(err)
	}
	x := &CustomCTabFolder2Adapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void close(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Adapter) Close(a CustomCTabFolderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void minimize(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Adapter) Minimize(a CustomCTabFolderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "minimize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void maximize(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Adapter) Maximize(a CustomCTabFolderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "maximize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void restore(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Adapter) Restore(a CustomCTabFolderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "restore", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void showList(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Adapter) ShowList(a CustomCTabFolderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showList", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


