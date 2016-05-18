package swt

import "github.com/timob/javabind"

type GraphicsGCDataInterface interface {
	JavaLangObjectInterface
}

type GraphicsGCData struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.GCData()
func NewGraphicsGCData() (*GraphicsGCData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/GCData")
	if err != nil {
		panic(err)
	}
	x := &GraphicsGCData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *GraphicsGCData) Device() *GraphicsDevice {
	jret, err := jbobject.GetField(javabind.GetEnv(), "device", "org/eclipse/swt/graphics/Device")
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
	unique_x := &GraphicsDevice{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsGCData) SetFieldDevice(val GraphicsDeviceInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "device", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsGCData) Style() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "style", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldStyle(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "style", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) State() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "state", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldState(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "state", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Font() *GraphicsFont {
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

func (jbobject *GraphicsGCData) SetFieldFont(val GraphicsFontInterface) {
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

func (jbobject *GraphicsGCData) ForegroundPattern() *GraphicsPattern {
	jret, err := jbobject.GetField(javabind.GetEnv(), "foregroundPattern", "org/eclipse/swt/graphics/Pattern")
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
	unique_x := &GraphicsPattern{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsGCData) SetFieldForegroundPattern(val GraphicsPatternInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "foregroundPattern", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsGCData) BackgroundPattern() *GraphicsPattern {
	jret, err := jbobject.GetField(javabind.GetEnv(), "backgroundPattern", "org/eclipse/swt/graphics/Pattern")
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
	unique_x := &GraphicsPattern{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsGCData) SetFieldBackgroundPattern(val GraphicsPatternInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "backgroundPattern", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsGCData) ClipRgn() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "clipRgn", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsGCData) SetFieldClipRgn(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "clipRgn", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) LineWidth() float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineWidth", javabind.Float)
	if err != nil {
		panic(err)
	}
	return jret.(float32)
}

func (jbobject *GraphicsGCData) SetFieldLineWidth(val float32) {
	err := jbobject.SetField(javabind.GetEnv(), "lineWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) LineStyle() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldLineStyle(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "lineStyle", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) LineDashes() []float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineDashes", javabind.Float | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]float32)
}

func (jbobject *GraphicsGCData) SetFieldLineDashes(val []float32) {
	err := jbobject.SetField(javabind.GetEnv(), "lineDashes", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) LineDashesOffset() float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineDashesOffset", javabind.Float)
	if err != nil {
		panic(err)
	}
	return jret.(float32)
}

func (jbobject *GraphicsGCData) SetFieldLineDashesOffset(val float32) {
	err := jbobject.SetField(javabind.GetEnv(), "lineDashesOffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) LineMiterLimit() float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineMiterLimit", javabind.Float)
	if err != nil {
		panic(err)
	}
	return jret.(float32)
}

func (jbobject *GraphicsGCData) SetFieldLineMiterLimit(val float32) {
	err := jbobject.SetField(javabind.GetEnv(), "lineMiterLimit", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) LineCap() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineCap", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldLineCap(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "lineCap", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) LineJoin() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineJoin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldLineJoin(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "lineJoin", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) XorMode() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "xorMode", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsGCData) SetFieldXorMode(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "xorMode", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Alpha() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "alpha", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldAlpha(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "alpha", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Interpolation() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "interpolation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldInterpolation(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "interpolation", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Context() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "context", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsGCData) SetFieldContext(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "context", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Layout() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "layout", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsGCData) SetFieldLayout(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "layout", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) DamageRgn() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "damageRgn", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsGCData) SetFieldDamageRgn(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "damageRgn", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Image() *GraphicsImage {
	jret, err := jbobject.GetField(javabind.GetEnv(), "image", "org/eclipse/swt/graphics/Image")
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
	unique_x := &GraphicsImage{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsGCData) SetFieldImage(val GraphicsImageInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "image", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsGCData) Drawable() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "drawable", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsGCData) SetFieldDrawable(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "drawable", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Cairo() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "cairo", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsGCData) SetFieldCairo(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "cairo", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) CairoXoffset() float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "cairoXoffset", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

func (jbobject *GraphicsGCData) SetFieldCairoXoffset(val float64) {
	err := jbobject.SetField(javabind.GetEnv(), "cairoXoffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) CairoYoffset() float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "cairoYoffset", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

func (jbobject *GraphicsGCData) SetFieldCairoYoffset(val float64) {
	err := jbobject.SetField(javabind.GetEnv(), "cairoYoffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) DisposeCairo() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "disposeCairo", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsGCData) SetFieldDisposeCairo(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "disposeCairo", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Identity() []float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "identity", javabind.Double | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]float64)
}

func (jbobject *GraphicsGCData) SetFieldIdentity(val []float64) {
	err := jbobject.SetField(javabind.GetEnv(), "identity", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) ClippingTransform() []float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "clippingTransform", javabind.Double | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]float64)
}

func (jbobject *GraphicsGCData) SetFieldClippingTransform(val []float64) {
	err := jbobject.SetField(javabind.GetEnv(), "clippingTransform", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) String() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "string", "java/lang/String")
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

func (jbobject *GraphicsGCData) SetFieldString(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "string", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsGCData) StringWidth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stringWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldStringWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stringWidth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) StringHeight() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stringHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldStringHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stringHeight", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) DrawFlags() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "drawFlags", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldDrawFlags(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "drawFlags", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) RealDrawable() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "realDrawable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsGCData) SetFieldRealDrawable(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "realDrawable", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsGCData) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsGCData) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}


