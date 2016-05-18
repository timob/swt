package swt

import "github.com/timob/javabind"

type GraphicsFontInterface interface {
	GraphicsResourceInterface

	// public org.eclipse.swt.graphics.FontData[] getFontData()
	GetFontData() []*GraphicsFontData

	// public static org.eclipse.swt.graphics.Font gtk_new(org.eclipse.swt.graphics.Device, long)
	Gtk_new(a GraphicsDeviceInterface, b int64) *GraphicsFont
}

type GraphicsFont struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Font(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.FontData)
func NewGraphicsFont(a GraphicsDeviceInterface, b GraphicsFontDataInterface) (*GraphicsFont) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Font", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/FontData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsFont{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Font(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.FontData[])
func NewGraphicsFont2(a GraphicsDeviceInterface, b []*GraphicsFontData) (*GraphicsFont) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/FontData")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Font", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/FontData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsFont{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Font(org.eclipse.swt.graphics.Device, java.lang.String, int, int)
func NewGraphicsFont3(a GraphicsDeviceInterface, b string, c int, d int) (*GraphicsFont) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Font", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("java/lang/String"), c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsFont{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsFont) Equals(a interface{}) bool {
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

// public org.eclipse.swt.graphics.FontData[] getFontData()
func (jbobject *GraphicsFont) GetFontData() []*GraphicsFontData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFontData", javabind.ObjectArrayType("org/eclipse/swt/graphics/FontData"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/FontData")
	dst := new([]*GraphicsFontData)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static org.eclipse.swt.graphics.Font gtk_new(org.eclipse.swt.graphics.Device, long)
func (jbobject *GraphicsFont) Gtk_new(a GraphicsDeviceInterface, b int64) *GraphicsFont {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/Font", "gtk_new", "org/eclipse/swt/graphics/Font", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b)
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
	unique_x := &GraphicsFont{}
	unique_x.Callable = dst
	return unique_x
}

// public int hashCode()
func (jbobject *GraphicsFont) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isDisposed()
func (jbobject *GraphicsFont) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *GraphicsFont) ToString() string {
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

func (jbobject *GraphicsFont) Handle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsFont) SetFieldHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


