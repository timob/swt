package swt

import "github.com/timob/javabind"

type GraphicsTextLayoutInterface interface {
	GraphicsResourceInterface

	// public void draw(org.eclipse.swt.graphics.GC, int, int)
	Draw(a GraphicsGCInterface, b int, c int) 

	// public void draw(org.eclipse.swt.graphics.GC, int, int, int, int, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color)
	Draw2(a GraphicsGCInterface, b int, c int, d int, e int, f GraphicsColorInterface, g GraphicsColorInterface) 

	// public void draw(org.eclipse.swt.graphics.GC, int, int, int, int, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color, int)
	Draw3(a GraphicsGCInterface, b int, c int, d int, e int, f GraphicsColorInterface, g GraphicsColorInterface, h int) 

	// public int getAlignment()
	GetAlignment() int

	// public int getAscent()
	GetAscent() int

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.graphics.Rectangle getBounds(int, int)
	GetBounds2(a int, b int) *GraphicsRectangle

	// public int getDescent()
	GetDescent() int

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public int getIndent()
	GetIndent() int

	// public boolean getJustify()
	GetJustify() bool

	// public int getLevel(int)
	GetLevel(a int) int

	// public org.eclipse.swt.graphics.Rectangle getLineBounds(int)
	GetLineBounds(a int) *GraphicsRectangle

	// public int getLineCount()
	GetLineCount() int

	// public int getLineIndex(int)
	GetLineIndex(a int) int

	// public org.eclipse.swt.graphics.FontMetrics getLineMetrics(int)
	GetLineMetrics(a int) *GraphicsFontMetrics

	// public int[] getLineOffsets()
	GetLineOffsets() []int

	// public org.eclipse.swt.graphics.Point getLocation(int, boolean)
	GetLocation(a int, b bool) *GraphicsPoint

	// public int getNextOffset(int, int)
	GetNextOffset(a int, b int) int

	// public int getOffset(org.eclipse.swt.graphics.Point, int[])
	GetOffset2(a GraphicsPointInterface, b []int) int

	// public int getOffset(int, int, int[])
	GetOffset(a int, b int, c []int) int

	// public int getOrientation()
	GetOrientation() int

	// public int getPreviousOffset(int, int)
	GetPreviousOffset(a int, b int) int

	// public int[] getRanges()
	GetRanges() []int

	// public int[] getSegments()
	GetSegments() []int

	// public char[] getSegmentsChars()
	GetSegmentsChars() []uint16

	// public int getSpacing()
	GetSpacing() int

	// public org.eclipse.swt.graphics.TextStyle getStyle(int)
	GetStyle(a int) *GraphicsTextStyle

	// public org.eclipse.swt.graphics.TextStyle[] getStyles()
	GetStyles() []*GraphicsTextStyle

	// public int[] getTabs()
	GetTabs() []int

	// public java.lang.String getText()
	GetText() string

	// public int getWidth()
	GetWidth() int

	// public int getWrapIndent()
	GetWrapIndent() int

	// public void setAlignment(int)
	SetAlignment(a int) 

	// public void setAscent(int)
	SetAscent(a int) 

	// public void setDescent(int)
	SetDescent(a int) 

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setIndent(int)
	SetIndent(a int) 

	// public void setJustify(boolean)
	SetJustify(a bool) 

	// public void setOrientation(int)
	SetOrientation(a int) 

	// public void setSpacing(int)
	SetSpacing(a int) 

	// public void setSegments(int[])
	SetSegments(a []int) 

	// public void setSegmentsChars(char[])
	SetSegmentsChars(a []uint16) 

	// public void setStyle(org.eclipse.swt.graphics.TextStyle, int, int)
	SetStyle(a GraphicsTextStyleInterface, b int, c int) 

	// public void setTabs(int[])
	SetTabs(a []int) 

	// public void setText(java.lang.String)
	SetText(a string) 

	// public void setWidth(int)
	SetWidth(a int) 

	// public void setWrapIndent(int)
	SetWrapIndent(a int) 
}

type GraphicsTextLayout struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.TextLayout(org.eclipse.swt.graphics.Device)
func NewGraphicsTextLayout(a GraphicsDeviceInterface) (*GraphicsTextLayout) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/TextLayout", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsTextLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void draw(org.eclipse.swt.graphics.GC, int, int)
func (jbobject *GraphicsTextLayout) Draw(a GraphicsGCInterface, b int, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "draw", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/GC"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void draw(org.eclipse.swt.graphics.GC, int, int, int, int, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color)
func (jbobject *GraphicsTextLayout) Draw2(a GraphicsGCInterface, b int, c int, d int, e int, f GraphicsColorInterface, g GraphicsColorInterface)  {
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
	_, err := jbobject.CallMethod(javabind.GetEnv(), "draw", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/GC"), b, c, d, e, conv_f.Value().Cast("org/eclipse/swt/graphics/Color"), conv_g.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()

}

// public void draw(org.eclipse.swt.graphics.GC, int, int, int, int, org.eclipse.swt.graphics.Color, org.eclipse.swt.graphics.Color, int)
func (jbobject *GraphicsTextLayout) Draw3(a GraphicsGCInterface, b int, c int, d int, e int, f GraphicsColorInterface, g GraphicsColorInterface, h int)  {
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
	_, err := jbobject.CallMethod(javabind.GetEnv(), "draw", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/GC"), b, c, d, e, conv_f.Value().Cast("org/eclipse/swt/graphics/Color"), conv_g.Value().Cast("org/eclipse/swt/graphics/Color"), h)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()

}

// public int getAlignment()
func (jbobject *GraphicsTextLayout) GetAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getAscent()
func (jbobject *GraphicsTextLayout) GetAscent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAscent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *GraphicsTextLayout) GetBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", "org/eclipse/swt/graphics/Rectangle")
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
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Rectangle getBounds(int, int)
func (jbobject *GraphicsTextLayout) GetBounds2(a int, b int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", "org/eclipse/swt/graphics/Rectangle", a, b)
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
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public int getDescent()
func (jbobject *GraphicsTextLayout) GetDescent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *GraphicsTextLayout) GetFont() *GraphicsFont {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFont", "org/eclipse/swt/graphics/Font")
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

// public int getIndent()
func (jbobject *GraphicsTextLayout) GetIndent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getJustify()
func (jbobject *GraphicsTextLayout) GetJustify() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJustify", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getLevel(int)
func (jbobject *GraphicsTextLayout) GetLevel(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLevel", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Rectangle getLineBounds(int)
func (jbobject *GraphicsTextLayout) GetLineBounds(a int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineBounds", "org/eclipse/swt/graphics/Rectangle", a)
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
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public int getLineCount()
func (jbobject *GraphicsTextLayout) GetLineCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineIndex(int)
func (jbobject *GraphicsTextLayout) GetLineIndex(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineIndex", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.FontMetrics getLineMetrics(int)
func (jbobject *GraphicsTextLayout) GetLineMetrics(a int) *GraphicsFontMetrics {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineMetrics", "org/eclipse/swt/graphics/FontMetrics", a)
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

// public int[] getLineOffsets()
func (jbobject *GraphicsTextLayout) GetLineOffsets() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineOffsets", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.eclipse.swt.graphics.Point getLocation(int, boolean)
func (jbobject *GraphicsTextLayout) GetLocation(a int, b bool) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocation", "org/eclipse/swt/graphics/Point", a, b)
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
	unique_x := &GraphicsPoint{}
	unique_x.Callable = dst
	return unique_x
}

// public int getNextOffset(int, int)
func (jbobject *GraphicsTextLayout) GetNextOffset(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextOffset", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getOffset(org.eclipse.swt.graphics.Point, int[])
func (jbobject *GraphicsTextLayout) GetOffset2(a GraphicsPointInterface, b []int) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOffset", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int getOffset(int, int, int[])
func (jbobject *GraphicsTextLayout) GetOffset(a int, b int, c []int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOffset", javabind.Int, a, b, c)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getOrientation()
func (jbobject *GraphicsTextLayout) GetOrientation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrientation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getPreviousOffset(int, int)
func (jbobject *GraphicsTextLayout) GetPreviousOffset(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPreviousOffset", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getRanges()
func (jbobject *GraphicsTextLayout) GetRanges() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRanges", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int[] getSegments()
func (jbobject *GraphicsTextLayout) GetSegments() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSegments", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public char[] getSegmentsChars()
func (jbobject *GraphicsTextLayout) GetSegmentsChars() []uint16 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSegmentsChars", javabind.Char | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]uint16)
}

// public int getSpacing()
func (jbobject *GraphicsTextLayout) GetSpacing() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.TextStyle getStyle(int)
func (jbobject *GraphicsTextLayout) GetStyle(a int) *GraphicsTextStyle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", "org/eclipse/swt/graphics/TextStyle", a)
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
	unique_x := &GraphicsTextStyle{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.TextStyle[] getStyles()
func (jbobject *GraphicsTextLayout) GetStyles() []*GraphicsTextStyle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyles", javabind.ObjectArrayType("org/eclipse/swt/graphics/TextStyle"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/TextStyle")
	dst := new([]*GraphicsTextStyle)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int[] getTabs()
func (jbobject *GraphicsTextLayout) GetTabs() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTabs", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public java.lang.String getText()
func (jbobject *GraphicsTextLayout) GetText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getText", "java/lang/String")
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

// public int getWidth()
func (jbobject *GraphicsTextLayout) GetWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getWrapIndent()
func (jbobject *GraphicsTextLayout) GetWrapIndent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWrapIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isDisposed()
func (jbobject *GraphicsTextLayout) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setAlignment(int)
func (jbobject *GraphicsTextLayout) SetAlignment(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlignment", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setAscent(int)
func (jbobject *GraphicsTextLayout) SetAscent(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAscent", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setDescent(int)
func (jbobject *GraphicsTextLayout) SetDescent(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescent", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *GraphicsTextLayout) SetFont(a GraphicsFontInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFont", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Font"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setIndent(int)
func (jbobject *GraphicsTextLayout) SetIndent(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIndent", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setJustify(boolean)
func (jbobject *GraphicsTextLayout) SetJustify(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJustify", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setOrientation(int)
func (jbobject *GraphicsTextLayout) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSpacing(int)
func (jbobject *GraphicsTextLayout) SetSpacing(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpacing", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSegments(int[])
func (jbobject *GraphicsTextLayout) SetSegments(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSegments", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSegmentsChars(char[])
func (jbobject *GraphicsTextLayout) SetSegmentsChars(a []uint16)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSegmentsChars", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setStyle(org.eclipse.swt.graphics.TextStyle, int, int)
func (jbobject *GraphicsTextLayout) SetStyle(a GraphicsTextStyleInterface, b int, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStyle", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/TextStyle"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTabs(int[])
func (jbobject *GraphicsTextLayout) SetTabs(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTabs", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *GraphicsTextLayout) SetText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setWidth(int)
func (jbobject *GraphicsTextLayout) SetWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setWrapIndent(int)
func (jbobject *GraphicsTextLayout) SetWrapIndent(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWrapIndent", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String toString()
func (jbobject *GraphicsTextLayout) ToString() string {
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


