package swt

import "github.com/timob/javabind"

type GraphicsFontMetricsInterface interface {
	JavaLangObjectInterface

	// public int getAscent()
	GetAscent() int

	// public int getAverageCharWidth()
	GetAverageCharWidth() int

	// public int getDescent()
	GetDescent() int

	// public int getHeight()
	GetHeight() int

	// public int getLeading()
	GetLeading() int

	// public static org.eclipse.swt.graphics.FontMetrics gtk_new(int, int, int, int, int)
	Gtk_new(a int, b int, c int, d int, e int) *GraphicsFontMetrics
}

type GraphicsFontMetrics struct {
	JavaLangObject
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsFontMetrics) Equals(a interface{}) bool {
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

// public int getAscent()
func (jbobject *GraphicsFontMetrics) GetAscent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAscent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getAverageCharWidth()
func (jbobject *GraphicsFontMetrics) GetAverageCharWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAverageCharWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getDescent()
func (jbobject *GraphicsFontMetrics) GetDescent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getHeight()
func (jbobject *GraphicsFontMetrics) GetHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLeading()
func (jbobject *GraphicsFontMetrics) GetLeading() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLeading", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public static org.eclipse.swt.graphics.FontMetrics gtk_new(int, int, int, int, int)
func (jbobject *GraphicsFontMetrics) Gtk_new(a int, b int, c int, d int, e int) *GraphicsFontMetrics {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/FontMetrics", "gtk_new", "org/eclipse/swt/graphics/FontMetrics", a, b, c, d, e)
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
	unique_x := &GraphicsFontMetrics{}
	unique_x.Callable = dst
	return unique_x
}

// public int hashCode()
func (jbobject *GraphicsFontMetrics) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


