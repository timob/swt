package swt

import "github.com/timob/javabind"

type GraphicsPathInterface interface {
	GraphicsResourceInterface

	// public void addArc(float, float, float, float, float, float)
	AddArc(a float32, b float32, c float32, d float32, e float32, f float32) 

	// public void addPath(org.eclipse.swt.graphics.Path)
	AddPath(a GraphicsPathInterface) 

	// public void addRectangle(float, float, float, float)
	AddRectangle(a float32, b float32, c float32, d float32) 

	// public void addString(java.lang.String, float, float, org.eclipse.swt.graphics.Font)
	AddString(a string, b float32, c float32, d GraphicsFontInterface) 

	// public void close()
	Close() 

	// public boolean contains(float, float, org.eclipse.swt.graphics.GC, boolean)
	Contains(a float32, b float32, c GraphicsGCInterface, d bool) bool

	// public void cubicTo(float, float, float, float, float, float)
	CubicTo(a float32, b float32, c float32, d float32, e float32, f float32) 

	// public void getBounds(float[])
	GetBounds(a []float32) 

	// public void getCurrentPoint(float[])
	GetCurrentPoint(a []float32) 

	// public org.eclipse.swt.graphics.PathData getPathData()
	GetPathData() *GraphicsPathData

	// public void lineTo(float, float)
	LineTo(a float32, b float32) 

	// public void moveTo(float, float)
	MoveTo(a float32, b float32) 

	// public void quadTo(float, float, float, float)
	QuadTo(a float32, b float32, c float32, d float32) 
}

type GraphicsPath struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Path(org.eclipse.swt.graphics.Device)
func NewGraphicsPath(a GraphicsDeviceInterface) (*GraphicsPath) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Path", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsPath{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Path(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.Path, float)
func NewGraphicsPath3(a GraphicsDeviceInterface, b GraphicsPathInterface, c float32) (*GraphicsPath) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Path", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/Path"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsPath{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Path(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.PathData)
func NewGraphicsPath2(a GraphicsDeviceInterface, b GraphicsPathDataInterface) (*GraphicsPath) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Path", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/PathData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsPath{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addArc(float, float, float, float, float, float)
func (jbobject *GraphicsPath) AddArc(a float32, b float32, c float32, d float32, e float32, f float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addArc", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void addPath(org.eclipse.swt.graphics.Path)
func (jbobject *GraphicsPath) AddPath(a GraphicsPathInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addPath", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Path"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addRectangle(float, float, float, float)
func (jbobject *GraphicsPath) AddRectangle(a float32, b float32, c float32, d float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addRectangle", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void addString(java.lang.String, float, float, org.eclipse.swt.graphics.Font)
func (jbobject *GraphicsPath) AddString(a string, b float32, c float32, d GraphicsFontInterface)  {
	conv_a := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addString", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c, conv_d.Value().Cast("org/eclipse/swt/graphics/Font"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_d.CleanUp()

}

// public void close()
func (jbobject *GraphicsPath) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean contains(float, float, org.eclipse.swt.graphics.GC, boolean)
func (jbobject *GraphicsPath) Contains(a float32, b float32, c GraphicsGCInterface, d bool) bool {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, a, b, conv_c.Value().Cast("org/eclipse/swt/graphics/GC"), d)
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	return jret.(bool)
}

// public void cubicTo(float, float, float, float, float, float)
func (jbobject *GraphicsPath) CubicTo(a float32, b float32, c float32, d float32, e float32, f float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cubicTo", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void getBounds(float[])
func (jbobject *GraphicsPath) GetBounds(a []float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void getCurrentPoint(float[])
func (jbobject *GraphicsPath) GetCurrentPoint(a []float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getCurrentPoint", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.PathData getPathData()
func (jbobject *GraphicsPath) GetPathData() *GraphicsPathData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPathData", "org/eclipse/swt/graphics/PathData")
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
	unique_x := &GraphicsPathData{}
	unique_x.Callable = dst
	return unique_x
}

// public void lineTo(float, float)
func (jbobject *GraphicsPath) LineTo(a float32, b float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "lineTo", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void moveTo(float, float)
func (jbobject *GraphicsPath) MoveTo(a float32, b float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "moveTo", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void quadTo(float, float, float, float)
func (jbobject *GraphicsPath) QuadTo(a float32, b float32, c float32, d float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "quadTo", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public boolean isDisposed()
func (jbobject *GraphicsPath) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *GraphicsPath) ToString() string {
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

func (jbobject *GraphicsPath) Handle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsPath) SetFieldHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


