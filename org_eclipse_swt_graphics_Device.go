package swt

import "github.com/timob/javabind"

type GraphicsDeviceInterface interface {
	JavaLangObjectInterface

	// public void dispose()
	Dispose() 

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.graphics.DeviceData getDeviceData()
	GetDeviceData() *GraphicsDeviceData

	// public org.eclipse.swt.graphics.Rectangle getClientArea()
	GetClientArea() *GraphicsRectangle

	// public int getDepth()
	GetDepth() int

	// public org.eclipse.swt.graphics.Point getDPI()
	GetDPI() *GraphicsPoint

	// public org.eclipse.swt.graphics.FontData[] getFontList(java.lang.String, boolean)
	GetFontList(a string, b bool) []*GraphicsFontData

	// public org.eclipse.swt.graphics.Color getSystemColor(int)
	GetSystemColor(a int) *GraphicsColor

	// public org.eclipse.swt.graphics.Font getSystemFont()
	GetSystemFont() *GraphicsFont

	// public boolean getWarnings()
	GetWarnings() bool

	// public abstract long internal_new_GC(org.eclipse.swt.graphics.GCData)
	Internal_new_GC(a GraphicsGCDataInterface) int64

	// public abstract void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
	Internal_dispose_GC(a int64, b GraphicsGCDataInterface) 

	// public boolean isDisposed()
	IsDisposed() bool

	// public boolean loadFont(java.lang.String)
	LoadFont(a string) bool

	// public void setWarnings(boolean)
	SetWarnings(a bool) 
}

type GraphicsDevice struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.Device()
func NewGraphicsDevice() (*GraphicsDevice) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Device")
	if err != nil {
		panic(err)
	}
	x := &GraphicsDevice{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Device(org.eclipse.swt.graphics.DeviceData)
func NewGraphicsDevice2(a GraphicsDeviceDataInterface) (*GraphicsDevice) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Device", conv_a.Value().Cast("org/eclipse/swt/graphics/DeviceData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsDevice{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dispose()
func (jbobject *GraphicsDevice) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *GraphicsDevice) GetBounds() *GraphicsRectangle {
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

// public org.eclipse.swt.graphics.DeviceData getDeviceData()
func (jbobject *GraphicsDevice) GetDeviceData() *GraphicsDeviceData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeviceData", "org/eclipse/swt/graphics/DeviceData")
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
	unique_x := &GraphicsDeviceData{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Rectangle getClientArea()
func (jbobject *GraphicsDevice) GetClientArea() *GraphicsRectangle {
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

// public int getDepth()
func (jbobject *GraphicsDevice) GetDepth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDepth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point getDPI()
func (jbobject *GraphicsDevice) GetDPI() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDPI", "org/eclipse/swt/graphics/Point")
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

// public org.eclipse.swt.graphics.FontData[] getFontList(java.lang.String, boolean)
func (jbobject *GraphicsDevice) GetFontList(a string, b bool) []*GraphicsFontData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFontList", javabind.ObjectArrayType("org/eclipse/swt/graphics/FontData"), conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/FontData")
	dst := new([]*GraphicsFontData)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.graphics.Color getSystemColor(int)
func (jbobject *GraphicsDevice) GetSystemColor(a int) *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemColor", "org/eclipse/swt/graphics/Color", a)
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
	unique_x := &GraphicsColor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Font getSystemFont()
func (jbobject *GraphicsDevice) GetSystemFont() *GraphicsFont {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemFont", "org/eclipse/swt/graphics/Font")
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
	unique_x := &GraphicsFont{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getWarnings()
func (jbobject *GraphicsDevice) GetWarnings() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWarnings", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract long internal_new_GC(org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsDevice) Internal_new_GC(a GraphicsGCDataInterface) int64 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "internal_new_GC", javabind.Long, conv_a.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public abstract void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsDevice) Internal_dispose_GC(a int64, b GraphicsGCDataInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "internal_dispose_GC", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public boolean isDisposed()
func (jbobject *GraphicsDevice) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean loadFont(java.lang.String)
func (jbobject *GraphicsDevice) LoadFont(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "loadFont", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void setWarnings(boolean)
func (jbobject *GraphicsDevice) SetWarnings(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWarnings", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsDevice) DEBUG() bool {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/graphics/Device", "DEBUG", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsDevice) SetFieldDEBUG(val bool) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/graphics/Device", "DEBUG", val)
	if err != nil {
		panic(err)
	}

}


