package swt

import "github.com/timob/javabind"

type CustomAnimatedProgressInterface interface {
	WidgetsCanvasInterface

	// public synchronized void clear()
	Clear() 

	// public synchronized void start()
	Start() 

	// public synchronized void stop()
	Stop() 
}

type CustomAnimatedProgress struct {
	WidgetsCanvas
}

// public org.eclipse.swt.custom.AnimatedProgress(org.eclipse.swt.widgets.Composite, int)
func NewCustomAnimatedProgress(a WidgetsCompositeInterface, b int) (*CustomAnimatedProgress) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/AnimatedProgress", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomAnimatedProgress{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public synchronized void clear()
func (jbobject *CustomAnimatedProgress) Clear()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *CustomAnimatedProgress) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeSize", "org/eclipse/swt/graphics/Point", a, b, c)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsPoint{}
	unique_x.Callable = dst
	return unique_x
}

// public synchronized void start()
func (jbobject *CustomAnimatedProgress) Start()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "start", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public synchronized void stop()
func (jbobject *CustomAnimatedProgress) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


