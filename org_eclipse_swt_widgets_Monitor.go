package swt

import "github.com/timob/javabind"

type WidgetsMonitorInterface interface {
	JavaLangObjectInterface

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.graphics.Rectangle getClientArea()
	GetClientArea() *GraphicsRectangle
}

type WidgetsMonitor struct {
	JavaLangObject
}

// public boolean equals(java.lang.Object)
func (jbobject *WidgetsMonitor) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsMonitor) GetBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", "org/eclipse/swt/graphics/Rectangle")
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
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *WidgetsMonitor) GetClientArea() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientArea", "org/eclipse/swt/graphics/Rectangle")
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
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public int hashCode()
func (jbobject *WidgetsMonitor) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


