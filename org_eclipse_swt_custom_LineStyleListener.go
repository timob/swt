package swt

import "github.com/timob/javabind"

type CustomLineStyleListenerInterface interface {

	// public abstract void lineGetStyle(org.eclipse.swt.custom.LineStyleEvent)
	LineGetStyle(a CustomLineStyleEventInterface) 
}

type CustomLineStyleListener struct {
	*javabind.Callable
}

// public abstract void lineGetStyle(org.eclipse.swt.custom.LineStyleEvent)
func (jbobject *CustomLineStyleListener) LineGetStyle(a CustomLineStyleEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "lineGetStyle", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/LineStyleEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


