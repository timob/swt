package swt

import "github.com/timob/javabind"

type CustomCTabFolderListenerInterface interface {

	// public abstract void itemClosed(org.eclipse.swt.custom.CTabFolderEvent)
	ItemClosed(a CustomCTabFolderEventInterface) 
}

type CustomCTabFolderListener struct {
	*javabind.Callable
}

// public abstract void itemClosed(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolderListener) ItemClosed(a CustomCTabFolderEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "itemClosed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CTabFolderEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


