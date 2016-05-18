package swt

import "github.com/timob/javabind"

type WidgetsTouchSourceInterface interface {
	JavaLangObjectInterface

	// public boolean isDirect()
	IsDirect() bool

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle
}

type WidgetsTouchSource struct {
	JavaLangObject
}

// public boolean isDirect()
func (jbobject *WidgetsTouchSource) IsDirect() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDirect", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsTouchSource) GetBounds() *GraphicsRectangle {
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

// public java.lang.String toString()
func (jbobject *WidgetsTouchSource) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


