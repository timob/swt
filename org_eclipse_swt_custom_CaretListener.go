package swt

import "github.com/timob/javabind"

type CustomCaretListenerInterface interface {

	// public abstract void caretMoved(org.eclipse.swt.custom.CaretEvent)
	CaretMoved(a CustomCaretEventInterface) 
}

type CustomCaretListener struct {
	*javabind.Callable
}

// public abstract void caretMoved(org.eclipse.swt.custom.CaretEvent)
func (jbobject *CustomCaretListener) CaretMoved(a CustomCaretEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "caretMoved", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CaretEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


