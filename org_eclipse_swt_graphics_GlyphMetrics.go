package swt

import "github.com/timob/javabind"

type GraphicsGlyphMetricsInterface interface {
	JavaLangObjectInterface
}

type GraphicsGlyphMetrics struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.GlyphMetrics(int, int, int)
func NewGraphicsGlyphMetrics(a int, b int, c int) (*GraphicsGlyphMetrics) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/GlyphMetrics", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &GraphicsGlyphMetrics{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsGlyphMetrics) Equals(a interface{}) bool {
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
func (jbobject *GraphicsGlyphMetrics) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *GraphicsGlyphMetrics) ToString() string {
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

func (jbobject *GraphicsGlyphMetrics) Ascent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "ascent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGlyphMetrics) SetFieldAscent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "ascent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGlyphMetrics) Descent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "descent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGlyphMetrics) SetFieldDescent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "descent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGlyphMetrics) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGlyphMetrics) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}


