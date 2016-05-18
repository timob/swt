package swt

import "github.com/timob/javabind"

type CustomPaintObjectListenerInterface interface {

	// public abstract void paintObject(org.eclipse.swt.custom.PaintObjectEvent)
	PaintObject(a CustomPaintObjectEventInterface) 
}

type CustomPaintObjectListener struct {
	*javabind.Callable
}

// public abstract void paintObject(org.eclipse.swt.custom.PaintObjectEvent)
func (jbobject *CustomPaintObjectListener) PaintObject(a CustomPaintObjectEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "paintObject", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/PaintObjectEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


