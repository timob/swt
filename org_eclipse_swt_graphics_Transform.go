package swt

import "github.com/timob/javabind"

type GraphicsTransformInterface interface {
	GraphicsResourceInterface

	// public void getElements(float[])
	GetElements(a []float32) 

	// public void identity()
	Identity() 

	// public void invert()
	Invert() 

	// public boolean isIdentity()
	IsIdentity() bool

	// public void multiply(org.eclipse.swt.graphics.Transform)
	Multiply(a GraphicsTransformInterface) 

	// public void rotate(float)
	Rotate(a float32) 

	// public void scale(float, float)
	Scale(a float32, b float32) 

	// public void setElements(float, float, float, float, float, float)
	SetElements(a float32, b float32, c float32, d float32, e float32, f float32) 

	// public void shear(float, float)
	Shear(a float32, b float32) 

	// public void transform(float[])
	Transform(a []float32) 

	// public void translate(float, float)
	Translate(a float32, b float32) 
}

type GraphicsTransform struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Transform(org.eclipse.swt.graphics.Device)
func NewGraphicsTransform(a GraphicsDeviceInterface) (*GraphicsTransform) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Transform", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsTransform{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Transform(org.eclipse.swt.graphics.Device, float[])
func NewGraphicsTransform2(a GraphicsDeviceInterface, b []float32) (*GraphicsTransform) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Transform", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsTransform{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Transform(org.eclipse.swt.graphics.Device, float, float, float, float, float, float)
func NewGraphicsTransform3(a GraphicsDeviceInterface, b float32, c float32, d float32, e float32, f float32, g float32) (*GraphicsTransform) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Transform", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b, c, d, e, f, g)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsTransform{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void getElements(float[])
func (jbobject *GraphicsTransform) GetElements(a []float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getElements", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void identity()
func (jbobject *GraphicsTransform) Identity()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "identity", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void invert()
func (jbobject *GraphicsTransform) Invert()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "invert", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean isDisposed()
func (jbobject *GraphicsTransform) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isIdentity()
func (jbobject *GraphicsTransform) IsIdentity() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isIdentity", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void multiply(org.eclipse.swt.graphics.Transform)
func (jbobject *GraphicsTransform) Multiply(a GraphicsTransformInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "multiply", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Transform"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void rotate(float)
func (jbobject *GraphicsTransform) Rotate(a float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "rotate", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void scale(float, float)
func (jbobject *GraphicsTransform) Scale(a float32, b float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "scale", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setElements(float, float, float, float, float, float)
func (jbobject *GraphicsTransform) SetElements(a float32, b float32, c float32, d float32, e float32, f float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setElements", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void shear(float, float)
func (jbobject *GraphicsTransform) Shear(a float32, b float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "shear", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void transform(float[])
func (jbobject *GraphicsTransform) Transform(a []float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "transform", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void translate(float, float)
func (jbobject *GraphicsTransform) Translate(a float32, b float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "translate", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String toString()
func (jbobject *GraphicsTransform) ToString() string {
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

func (jbobject *GraphicsTransform) Handle() []float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Double | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]float64)
}

func (jbobject *GraphicsTransform) SetFieldHandle(val []float64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


