package swt

import "github.com/timob/javabind"

type CustomExtendedModifyListenerInterface interface {

	// public abstract void modifyText(org.eclipse.swt.custom.ExtendedModifyEvent)
	ModifyText(a CustomExtendedModifyEventInterface) 
}

type CustomExtendedModifyListener struct {
	*javabind.Callable
}

// public abstract void modifyText(org.eclipse.swt.custom.ExtendedModifyEvent)
func (jbobject *CustomExtendedModifyListener) ModifyText(a CustomExtendedModifyEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyText", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/ExtendedModifyEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


