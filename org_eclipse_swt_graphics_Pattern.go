package swt

import "github.com/timob/javabind"

type GraphicsPatternInterface interface {
	GraphicsResourceInterface
}

type GraphicsPattern struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Pattern(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.Image)
func NewGraphicsPattern(a GraphicsDeviceInterface, b GraphicsImageInterface) (*GraphicsPattern) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Pattern", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsPattern{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Pattern(org.eclipse.swt.graphics.Device, float, float, float, float, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color)
func NewGraphicsPattern2(a GraphicsDeviceInterface, b float32, c float32, d float32, e float32, f GraphicsColorInterface, g GraphicsColorInterface) (*GraphicsPattern) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Pattern", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b, c, d, e, conv_f.Value().Cast("org/eclipse/swt/graphics/Color"), conv_g.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &GraphicsPattern{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Pattern(org.eclipse.swt.graphics.Device, float, float, float, float, org.eclipse.swt.graphics.Color, int, org.eclipse.swt.graphics.Color, int)
func NewGraphicsPattern3(a GraphicsDeviceInterface, b float32, c float32, d float32, e float32, f GraphicsColorInterface, g int, h GraphicsColorInterface, i int) (*GraphicsPattern) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Pattern", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b, c, d, e, conv_f.Value().Cast("org/eclipse/swt/graphics/Color"), g, conv_h.Value().Cast("org/eclipse/swt/graphics/Color"), i)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_f.CleanUp()
	conv_h.CleanUp()
	x := &GraphicsPattern{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean isDisposed()
func (jbobject *GraphicsPattern) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *GraphicsPattern) ToString() string {
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

func (jbobject *GraphicsPattern) Handle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsPattern) SetFieldHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


