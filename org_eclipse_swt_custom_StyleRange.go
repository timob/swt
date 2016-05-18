package swt

import "github.com/timob/javabind"

type CustomStyleRangeInterface interface {
	GraphicsTextStyleInterface

	// public boolean isUnstyled()
	IsUnstyled() bool

	// public boolean similarTo(org.eclipse.swt.custom.StyleRange)
	SimilarTo(a CustomStyleRangeInterface) bool

	// public java.lang.Object clone()
	Clone() *JavaLangObject
}

type CustomStyleRange struct {
	GraphicsTextStyle
}

// public org.eclipse.swt.custom.StyleRange()
func NewCustomStyleRange() (*CustomStyleRange) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StyleRange")
	if err != nil {
		panic(err)
	}
	x := &CustomStyleRange{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.StyleRange(org.eclipse.swt.graphics.TextStyle)
func NewCustomStyleRange2(a GraphicsTextStyleInterface) (*CustomStyleRange) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StyleRange", conv_a.Value().Cast("org/eclipse/swt/graphics/TextStyle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomStyleRange{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.StyleRange(int, int, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color)
func NewCustomStyleRange3(a int, b int, c GraphicsColorInterface, d GraphicsColorInterface) (*CustomStyleRange) {
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StyleRange", a, b, conv_c.Value().Cast("org/eclipse/swt/graphics/Color"), conv_d.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &CustomStyleRange{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.StyleRange(int, int, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color, int)
func NewCustomStyleRange4(a int, b int, c GraphicsColorInterface, d GraphicsColorInterface, e int) (*CustomStyleRange) {
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StyleRange", a, b, conv_c.Value().Cast("org/eclipse/swt/graphics/Color"), conv_d.Value().Cast("org/eclipse/swt/graphics/Color"), e)
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &CustomStyleRange{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *CustomStyleRange) Equals(a interface{}) bool {
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
func (jbobject *CustomStyleRange) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isUnstyled()
func (jbobject *CustomStyleRange) IsUnstyled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isUnstyled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean similarTo(org.eclipse.swt.custom.StyleRange)
func (jbobject *CustomStyleRange) SimilarTo(a CustomStyleRangeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "similarTo", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.lang.Object clone()
func (jbobject *CustomStyleRange) Clone() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *CustomStyleRange) ToString() string {
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

func (jbobject *CustomStyleRange) Start() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "start", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomStyleRange) SetFieldStart(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "start", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStyleRange) Length() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "length", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomStyleRange) SetFieldLength(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "length", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStyleRange) FontStyle() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "fontStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomStyleRange) SetFieldFontStyle(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "fontStyle", val)
	if err != nil {
		panic(err)
	}

}


