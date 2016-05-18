package swt

import "github.com/timob/javabind"

type CustomCTabFolder2ListenerInterface interface {

	// public abstract void close(org.eclipse.swt.custom.CTabFolderEvent)
	Close(a CustomCTabFolderEventInterface) 

	// public abstract void minimize(org.eclipse.swt.custom.CTabFolderEvent)
	Minimize(a CustomCTabFolderEventInterface) 

	// public abstract void maximize(org.eclipse.swt.custom.CTabFolderEvent)
	Maximize(a CustomCTabFolderEventInterface) 

	// public abstract void restore(org.eclipse.swt.custom.CTabFolderEvent)
	Restore(a CustomCTabFolderEventInterface) 

	// public abstract void showList(org.eclipse.swt.custom.CTabFolderEvent)
	ShowList(a CustomCTabFolderEventInterface) 
}

type CustomCTabFolder2Listener struct {
	*javabind.Callable
}

// public abstract void close(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Listener) Close(a CustomCTabFolderEventInterface)  {
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

// public abstract void minimize(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Listener) Minimize(a CustomCTabFolderEventInterface)  {
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

// public abstract void maximize(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Listener) Maximize(a CustomCTabFolderEventInterface)  {
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

// public abstract void restore(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Listener) Restore(a CustomCTabFolderEventInterface)  {
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

// public abstract void showList(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolder2Listener) ShowList(a CustomCTabFolderEventInterface)  {
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


