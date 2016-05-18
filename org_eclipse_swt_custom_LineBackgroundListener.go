package swt

import "github.com/timob/javabind"

type CustomLineBackgroundListenerInterface interface {

	// public abstract void lineGetBackground(org.eclipse.swt.custom.LineBackgroundEvent)
	LineGetBackground(a CustomLineBackgroundEventInterface) 
}

type CustomLineBackgroundListener struct {
	*javabind.Callable
}

// public abstract void lineGetBackground(org.eclipse.swt.custom.LineBackgroundEvent)
func (jbobject *CustomLineBackgroundListener) LineGetBackground(a CustomLineBackgroundEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "lineGetBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/LineBackgroundEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


