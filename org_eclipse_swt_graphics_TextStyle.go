package swt

import "github.com/timob/javabind"

type GraphicsTextStyleInterface interface {
	JavaLangObjectInterface
}

type GraphicsTextStyle struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.TextStyle()
func NewGraphicsTextStyle() (*GraphicsTextStyle) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/TextStyle")
	if err != nil {
		panic(err)
	}
	x := &GraphicsTextStyle{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.TextStyle(org.eclipse.swt.graphics.Font, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color)
func NewGraphicsTextStyle3(a GraphicsFontInterface, b GraphicsColorInterface, c GraphicsColorInterface) (*GraphicsTextStyle) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/TextStyle", conv_a.Value().Cast("org/eclipse/swt/graphics/Font"), conv_b.Value().Cast("org/eclipse/swt/graphics/Color"), conv_c.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &GraphicsTextStyle{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.TextStyle(org.eclipse.swt.graphics.TextStyle)
func NewGraphicsTextStyle2(a GraphicsTextStyleInterface) (*GraphicsTextStyle) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/TextStyle", conv_a.Value().Cast("org/eclipse/swt/graphics/TextStyle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsTextStyle{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsTextStyle) Equals(a interface{}) bool {
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
func (jbobject *GraphicsTextStyle) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *GraphicsTextStyle) ToString() string {
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

func (jbobject *GraphicsTextStyle) Font() *GraphicsFont {
	jret, err := jbobject.GetField(javabind.GetEnv(), "font", "org/eclipse/swt/graphics/Font")
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

func (jbobject *GraphicsTextStyle) SetFieldFont(val GraphicsFontInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "font", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsTextStyle) Foreground() *GraphicsColor {
	jret, err := jbobject.GetField(javabind.GetEnv(), "foreground", "org/eclipse/swt/graphics/Color")
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

func (jbobject *GraphicsTextStyle) SetFieldForeground(val GraphicsColorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "foreground", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsTextStyle) Background() *GraphicsColor {
	jret, err := jbobject.GetField(javabind.GetEnv(), "background", "org/eclipse/swt/graphics/Color")
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

func (jbobject *GraphicsTextStyle) SetFieldBackground(val GraphicsColorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "background", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsTextStyle) Underline() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "underline", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsTextStyle) SetFieldUnderline(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "underline", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsTextStyle) UnderlineColor() *GraphicsColor {
	jret, err := jbobject.GetField(javabind.GetEnv(), "underlineColor", "org/eclipse/swt/graphics/Color")
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

func (jbobject *GraphicsTextStyle) SetFieldUnderlineColor(val GraphicsColorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "underlineColor", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsTextStyle) UnderlineStyle() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "underlineStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsTextStyle) SetFieldUnderlineStyle(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "underlineStyle", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsTextStyle) Strikeout() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "strikeout", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsTextStyle) SetFieldStrikeout(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "strikeout", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsTextStyle) StrikeoutColor() *GraphicsColor {
	jret, err := jbobject.GetField(javabind.GetEnv(), "strikeoutColor", "org/eclipse/swt/graphics/Color")
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

func (jbobject *GraphicsTextStyle) SetFieldStrikeoutColor(val GraphicsColorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "strikeoutColor", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsTextStyle) BorderStyle() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "borderStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsTextStyle) SetFieldBorderStyle(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "borderStyle", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsTextStyle) BorderColor() *GraphicsColor {
	jret, err := jbobject.GetField(javabind.GetEnv(), "borderColor", "org/eclipse/swt/graphics/Color")
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

func (jbobject *GraphicsTextStyle) SetFieldBorderColor(val GraphicsColorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "borderColor", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsTextStyle) Metrics() *GraphicsGlyphMetrics {
	jret, err := jbobject.GetField(javabind.GetEnv(), "metrics", "org/eclipse/swt/graphics/GlyphMetrics")
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
	unique_x := &GraphicsGlyphMetrics{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsTextStyle) SetFieldMetrics(val GraphicsGlyphMetricsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "metrics", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsTextStyle) Rise() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "rise", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsTextStyle) SetFieldRise(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "rise", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsTextStyle) Data() *JavaLangObject {
	jret, err := jbobject.GetField(javabind.GetEnv(), "data", "java/lang/Object")
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

func (jbobject *GraphicsTextStyle) SetFieldData(val JavaLangObjectInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "data", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


