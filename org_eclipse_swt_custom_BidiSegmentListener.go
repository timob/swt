package swt

import "github.com/timob/javabind"

type CustomBidiSegmentListenerInterface interface {

	// public abstract void lineGetSegments(org.eclipse.swt.custom.BidiSegmentEvent)
	LineGetSegments(a CustomBidiSegmentEventInterface) 
}

type CustomBidiSegmentListener struct {
	*javabind.Callable
}

// public abstract void lineGetSegments(org.eclipse.swt.custom.BidiSegmentEvent)
func (jbobject *CustomBidiSegmentListener) LineGetSegments(a CustomBidiSegmentEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "lineGetSegments", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/BidiSegmentEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


