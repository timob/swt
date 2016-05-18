package swt

import "github.com/timob/javabind"

type CustomCTabFolderAdapterInterface interface {
	JavaLangObjectInterface

	// public void itemClosed(org.eclipse.swt.custom.CTabFolderEvent)
	ItemClosed(a CustomCTabFolderEventInterface) 
}

type CustomCTabFolderAdapter struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.CTabFolderAdapter()
func NewCustomCTabFolderAdapter() (*CustomCTabFolderAdapter) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CTabFolderAdapter")
	if err != nil {
		panic(err)
	}
	x := &CustomCTabFolderAdapter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void itemClosed(org.eclipse.swt.custom.CTabFolderEvent)
func (jbobject *CustomCTabFolderAdapter) ItemClosed(a CustomCTabFolderEventInterface)  {
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


