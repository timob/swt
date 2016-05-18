package swt

import "github.com/timob/javabind"

type GraphicsCursorInterface interface {
	GraphicsResourceInterface

	// public static org.eclipse.swt.graphics.Cursor gtk_new(org.eclipse.swt.graphics.Device, long)
	Gtk_new(a GraphicsDeviceInterface, b int64) *GraphicsCursor
}

type GraphicsCursor struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Cursor(org.eclipse.swt.graphics.Device, int)
func NewGraphicsCursor(a GraphicsDeviceInterface, b int) (*GraphicsCursor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Cursor", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsCursor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Cursor(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.ImageData, org.eclipse.swt.graphics.ImageData, int, int)
func NewGraphicsCursor3(a GraphicsDeviceInterface, b GraphicsImageDataInterface, c GraphicsImageDataInterface, d int, e int) (*GraphicsCursor) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Cursor", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/ImageData"), conv_c.Value().Cast("org/eclipse/swt/graphics/ImageData"), d, e)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &GraphicsCursor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Cursor(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.ImageData, int, int)
func NewGraphicsCursor2(a GraphicsDeviceInterface, b GraphicsImageDataInterface, c int, d int) (*GraphicsCursor) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Cursor", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/ImageData"), c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsCursor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsCursor) Equals(a interface{}) bool {
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

// public static org.eclipse.swt.graphics.Cursor gtk_new(org.eclipse.swt.graphics.Device, long)
func (jbobject *GraphicsCursor) Gtk_new(a GraphicsDeviceInterface, b int64) *GraphicsCursor {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/Cursor", "gtk_new", "org/eclipse/swt/graphics/Cursor", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsCursor{}
	unique_x.Callable = dst
	return unique_x
}

// public int hashCode()
func (jbobject *GraphicsCursor) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isDisposed()
func (jbobject *GraphicsCursor) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *GraphicsCursor) ToString() string {
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

func (jbobject *GraphicsCursor) Handle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsCursor) SetFieldHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


