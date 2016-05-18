package swt

import "github.com/timob/javabind"

type GraphicsColorInterface interface {
	GraphicsResourceInterface

	// public int getBlue()
	GetBlue() int

	// public int getGreen()
	GetGreen() int

	// public int getRed()
	GetRed() int

	// public org.eclipse.swt.graphics.RGB getRGB()
	GetRGB() *GraphicsRGB
}

type GraphicsColor struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Color(org.eclipse.swt.graphics.Device, int, int, int)
func NewGraphicsColor(a GraphicsDeviceInterface, b int, c int, d int) (*GraphicsColor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Color", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsColor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Color(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.RGB)
func NewGraphicsColor2(a GraphicsDeviceInterface, b GraphicsRGBInterface) (*GraphicsColor) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Color", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/RGB"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsColor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsColor) Equals(a interface{}) bool {
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

// public int getBlue()
func (jbobject *GraphicsColor) GetBlue() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlue", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getGreen()
func (jbobject *GraphicsColor) GetGreen() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGreen", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getRed()
func (jbobject *GraphicsColor) GetRed() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRed", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int hashCode()
func (jbobject *GraphicsColor) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.RGB getRGB()
func (jbobject *GraphicsColor) GetRGB() *GraphicsRGB {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRGB", "org/eclipse/swt/graphics/RGB")
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isDisposed()
func (jbobject *GraphicsColor) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *GraphicsColor) ToString() string {
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


