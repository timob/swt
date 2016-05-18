package swt

import "github.com/timob/javabind"

type GraphicsLineAttributesInterface interface {
	JavaLangObjectInterface
}

type GraphicsLineAttributes struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.LineAttributes(float)
func NewGraphicsLineAttributes(a float32) (*GraphicsLineAttributes) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/LineAttributes", a)
	if err != nil {
		panic(err)
	}
	x := &GraphicsLineAttributes{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.LineAttributes(float, int, int)
func NewGraphicsLineAttributes2(a float32, b int, c int) (*GraphicsLineAttributes) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/LineAttributes", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &GraphicsLineAttributes{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.LineAttributes(float, int, int, int, float[], float, float)
func NewGraphicsLineAttributes3(a float32, b int, c int, d int, e []float32, f float32, g float32) (*GraphicsLineAttributes) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/LineAttributes", a, b, c, d, e, f, g)
	if err != nil {
		panic(err)
	}
	x := &GraphicsLineAttributes{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsLineAttributes) Equals(a interface{}) bool {
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

// public int hashCode()
func (jbobject *GraphicsLineAttributes) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsLineAttributes) Width() float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Float)
	if err != nil {
		panic(err)
	}
	return jret.(float32)
}

func (jbobject *GraphicsLineAttributes) SetFieldWidth(val float32) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsLineAttributes) Style() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "style", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsLineAttributes) SetFieldStyle(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "style", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsLineAttributes) Cap() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "cap", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsLineAttributes) SetFieldCap(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "cap", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsLineAttributes) Join() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "join", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsLineAttributes) SetFieldJoin(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "join", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsLineAttributes) Dash() []float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "dash", javabind.Float | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]float32)
}

func (jbobject *GraphicsLineAttributes) SetFieldDash(val []float32) {
	err := jbobject.SetField(javabind.GetEnv(), "dash", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsLineAttributes) DashOffset() float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "dashOffset", javabind.Float)
	if err != nil {
		panic(err)
	}
	return jret.(float32)
}

func (jbobject *GraphicsLineAttributes) SetFieldDashOffset(val float32) {
	err := jbobject.SetField(javabind.GetEnv(), "dashOffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsLineAttributes) MiterLimit() float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "miterLimit", javabind.Float)
	if err != nil {
		panic(err)
	}
	return jret.(float32)
}

func (jbobject *GraphicsLineAttributes) SetFieldMiterLimit(val float32) {
	err := jbobject.SetField(javabind.GetEnv(), "miterLimit", val)
	if err != nil {
		panic(err)
	}

}


