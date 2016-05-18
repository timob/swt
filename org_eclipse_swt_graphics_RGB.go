package swt

import "github.com/timob/javabind"

type GraphicsRGBInterface interface {
	JavaLangObjectInterface

	// public float[] getHSB()
	GetHSB() []float32
}

type GraphicsRGB struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.RGB(int, int, int)
func NewGraphicsRGB2(a int, b int, c int) (*GraphicsRGB) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/RGB", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &GraphicsRGB{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.RGB(float, float, float)
func NewGraphicsRGB(a float32, b float32, c float32) (*GraphicsRGB) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/RGB", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &GraphicsRGB{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public float[] getHSB()
func (jbobject *GraphicsRGB) GetHSB() []float32 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHSB", javabind.Float | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]float32)
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsRGB) Equals(a interface{}) bool {
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
func (jbobject *GraphicsRGB) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *GraphicsRGB) ToString() string {
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

func (jbobject *GraphicsRGB) Red() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "red", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsRGB) SetFieldRed(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "red", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsRGB) Green() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "green", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsRGB) SetFieldGreen(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "green", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsRGB) Blue() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "blue", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsRGB) SetFieldBlue(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "blue", val)
	if err != nil {
		panic(err)
	}

}


