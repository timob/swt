package swt

import "github.com/timob/javabind"

type GraphicsGCInterface interface {
	GraphicsResourceInterface

	// public static org.eclipse.swt.graphics.GC gtk_new(long, org.eclipse.swt.graphics.GCData)
	Gtk_new(a int64, b GraphicsGCDataInterface) *GraphicsGC

	// public static org.eclipse.swt.graphics.GC gtk_new(org.eclipse.swt.graphics.Drawable, org.eclipse.swt.graphics.GCData)
	Gtk_new2(a GraphicsDrawableInterface, b GraphicsGCDataInterface) *GraphicsGC

	// public void copyArea(org.eclipse.swt.graphics.Image, int, int)
	CopyArea3(a GraphicsImageInterface, b int, c int) 

	// public void copyArea(int, int, int, int, int, int)
	CopyArea(a int, b int, c int, d int, e int, f int) 

	// public void copyArea(int, int, int, int, int, int, boolean)
	CopyArea2(a int, b int, c int, d int, e int, f int, g bool) 

	// public void drawArc(int, int, int, int, int, int)
	DrawArc(a int, b int, c int, d int, e int, f int) 

	// public void drawFocus(int, int, int, int)
	DrawFocus(a int, b int, c int, d int) 

	// public void drawImage(org.eclipse.swt.graphics.Image, int, int)
	DrawImage(a GraphicsImageInterface, b int, c int) 

	// public void drawImage(org.eclipse.swt.graphics.Image, int, int, int, int, int, int, int, int)
	DrawImage2(a GraphicsImageInterface, b int, c int, d int, e int, f int, g int, h int, i int) 

	// public void drawLine(int, int, int, int)
	DrawLine(a int, b int, c int, d int) 

	// public void drawOval(int, int, int, int)
	DrawOval(a int, b int, c int, d int) 

	// public void drawPath(org.eclipse.swt.graphics.Path)
	DrawPath(a GraphicsPathInterface) 

	// public void drawPoint(int, int)
	DrawPoint(a int, b int) 

	// public void drawPolygon(int[])
	DrawPolygon(a []int) 

	// public void drawPolyline(int[])
	DrawPolyline(a []int) 

	// public void drawRectangle(int, int, int, int)
	DrawRectangle(a int, b int, c int, d int) 

	// public void drawRectangle(org.eclipse.swt.graphics.Rectangle)
	DrawRectangle2(a GraphicsRectangleInterface) 

	// public void drawRoundRectangle(int, int, int, int, int, int)
	DrawRoundRectangle(a int, b int, c int, d int, e int, f int) 

	// public void drawString(java.lang.String, int, int)
	DrawString(a string, b int, c int) 

	// public void drawString(java.lang.String, int, int, boolean)
	DrawString2(a string, b int, c int, d bool) 

	// public void drawText(java.lang.String, int, int)
	DrawText(a string, b int, c int) 

	// public void drawText(java.lang.String, int, int, boolean)
	DrawText2(a string, b int, c int, d bool) 

	// public void drawText(java.lang.String, int, int, int)
	DrawText3(a string, b int, c int, d int) 

	// public void fillArc(int, int, int, int, int, int)
	FillArc(a int, b int, c int, d int, e int, f int) 

	// public void fillGradientRectangle(int, int, int, int, boolean)
	FillGradientRectangle(a int, b int, c int, d int, e bool) 

	// public void fillOval(int, int, int, int)
	FillOval(a int, b int, c int, d int) 

	// public void fillPath(org.eclipse.swt.graphics.Path)
	FillPath(a GraphicsPathInterface) 

	// public void fillPolygon(int[])
	FillPolygon(a []int) 

	// public void fillRectangle(int, int, int, int)
	FillRectangle(a int, b int, c int, d int) 

	// public void fillRectangle(org.eclipse.swt.graphics.Rectangle)
	FillRectangle2(a GraphicsRectangleInterface) 

	// public void fillRoundRectangle(int, int, int, int, int, int)
	FillRoundRectangle(a int, b int, c int, d int, e int, f int) 

	// public int getAdvanceWidth(char)
	GetAdvanceWidth(a uint16) int

	// public boolean getAdvanced()
	GetAdvanced() bool

	// public int getAlpha()
	GetAlpha() int

	// public int getAntialias()
	GetAntialias() int

	// public org.eclipse.swt.graphics.Color getBackground()
	GetBackground() *GraphicsColor

	// public org.eclipse.swt.graphics.Pattern getBackgroundPattern()
	GetBackgroundPattern() *GraphicsPattern

	// public int getCharWidth(char)
	GetCharWidth(a uint16) int

	// public org.eclipse.swt.graphics.Rectangle getClipping()
	GetClipping() *GraphicsRectangle

	// public void getClipping(org.eclipse.swt.graphics.Region)
	GetClipping2(a GraphicsRegionInterface) 

	// public int getFillRule()
	GetFillRule() int

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public org.eclipse.swt.graphics.FontMetrics getFontMetrics()
	GetFontMetrics() *GraphicsFontMetrics

	// public org.eclipse.swt.graphics.Color getForeground()
	GetForeground() *GraphicsColor

	// public org.eclipse.swt.graphics.Pattern getForegroundPattern()
	GetForegroundPattern() *GraphicsPattern

	// public org.eclipse.swt.graphics.GCData getGCData()
	GetGCData() *GraphicsGCData

	// public int getInterpolation()
	GetInterpolation() int

	// public org.eclipse.swt.graphics.LineAttributes getLineAttributes()
	GetLineAttributes() *GraphicsLineAttributes

	// public int getLineCap()
	GetLineCap() int

	// public int[] getLineDash()
	GetLineDash() []int

	// public int getLineJoin()
	GetLineJoin() int

	// public int getLineStyle()
	GetLineStyle() int

	// public int getLineWidth()
	GetLineWidth() int

	// public int getStyle()
	GetStyle() int

	// public int getTextAntialias()
	GetTextAntialias() int

	// public void getTransform(org.eclipse.swt.graphics.Transform)
	GetTransform(a GraphicsTransformInterface) 

	// public boolean getXORMode()
	GetXORMode() bool

	// public boolean isClipped()
	IsClipped() bool

	// public void setAdvanced(boolean)
	SetAdvanced(a bool) 

	// public void setAlpha(int)
	SetAlpha(a int) 

	// public void setAntialias(int)
	SetAntialias(a int) 

	// public void setBackground(org.eclipse.swt.graphics.Color)
	SetBackground(a GraphicsColorInterface) 

	// public void setBackgroundPattern(org.eclipse.swt.graphics.Pattern)
	SetBackgroundPattern(a GraphicsPatternInterface) 

	// public void setClipping(int, int, int, int)
	SetClipping(a int, b int, c int, d int) 

	// public void setClipping(org.eclipse.swt.graphics.Path)
	SetClipping2(a GraphicsPathInterface) 

	// public void setClipping(org.eclipse.swt.graphics.Rectangle)
	SetClipping3(a GraphicsRectangleInterface) 

	// public void setClipping(org.eclipse.swt.graphics.Region)
	SetClipping4(a GraphicsRegionInterface) 

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setFillRule(int)
	SetFillRule(a int) 

	// public void setForeground(org.eclipse.swt.graphics.Color)
	SetForeground(a GraphicsColorInterface) 

	// public void setForegroundPattern(org.eclipse.swt.graphics.Pattern)
	SetForegroundPattern(a GraphicsPatternInterface) 

	// public void setInterpolation(int)
	SetInterpolation(a int) 

	// public void setLineAttributes(org.eclipse.swt.graphics.LineAttributes)
	SetLineAttributes(a GraphicsLineAttributesInterface) 

	// public void setLineCap(int)
	SetLineCap(a int) 

	// public void setLineDash(int[])
	SetLineDash(a []int) 

	// public void setLineJoin(int)
	SetLineJoin(a int) 

	// public void setLineStyle(int)
	SetLineStyle(a int) 

	// public void setLineWidth(int)
	SetLineWidth(a int) 

	// public void setTextAntialias(int)
	SetTextAntialias(a int) 

	// public void setTransform(org.eclipse.swt.graphics.Transform)
	SetTransform(a GraphicsTransformInterface) 

	// public void setXORMode(boolean)
	SetXORMode(a bool) 

	// public org.eclipse.swt.graphics.Point stringExtent(java.lang.String)
	StringExtent(a string) *GraphicsPoint

	// public org.eclipse.swt.graphics.Point textExtent(java.lang.String)
	TextExtent(a string) *GraphicsPoint

	// public org.eclipse.swt.graphics.Point textExtent(java.lang.String, int)
	TextExtent2(a string, b int) *GraphicsPoint
}

type GraphicsGC struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.GC(org.eclipse.swt.graphics.Drawable)
func NewGraphicsGC(a GraphicsDrawableInterface) (*GraphicsGC) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/GC", conv_a.Value().Cast("org/eclipse/swt/graphics/Drawable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsGC{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.GC(org.eclipse.swt.graphics.Drawable, int)
func NewGraphicsGC2(a GraphicsDrawableInterface, b int) (*GraphicsGC) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/GC", conv_a.Value().Cast("org/eclipse/swt/graphics/Drawable"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsGC{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.eclipse.swt.graphics.GC gtk_new(long, org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsGC) Gtk_new(a int64, b GraphicsGCDataInterface) *GraphicsGC {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/GC", "gtk_new", "org/eclipse/swt/graphics/GC", a, conv_b.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsGC{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.eclipse.swt.graphics.GC gtk_new(org.eclipse.swt.graphics.Drawable, org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsGC) Gtk_new2(a GraphicsDrawableInterface, b GraphicsGCDataInterface) *GraphicsGC {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/GC", "gtk_new", "org/eclipse/swt/graphics/GC", conv_a.Value().Cast("org/eclipse/swt/graphics/Drawable"), conv_b.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsGC{}
	unique_x.Callable = dst
	return unique_x
}

// public void copyArea(org.eclipse.swt.graphics.Image, int, int)
func (jbobject *GraphicsGC) CopyArea3(a GraphicsImageInterface, b int, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copyArea", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void copyArea(int, int, int, int, int, int)
func (jbobject *GraphicsGC) CopyArea(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copyArea", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void copyArea(int, int, int, int, int, int, boolean)
func (jbobject *GraphicsGC) CopyArea2(a int, b int, c int, d int, e int, f int, g bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copyArea", javabind.Void, a, b, c, d, e, f, g)
	if err != nil {
		panic(err)
	}

}

// public void drawArc(int, int, int, int, int, int)
func (jbobject *GraphicsGC) DrawArc(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawArc", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void drawFocus(int, int, int, int)
func (jbobject *GraphicsGC) DrawFocus(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawFocus", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void drawImage(org.eclipse.swt.graphics.Image, int, int)
func (jbobject *GraphicsGC) DrawImage(a GraphicsImageInterface, b int, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawImage(org.eclipse.swt.graphics.Image, int, int, int, int, int, int, int, int)
func (jbobject *GraphicsGC) DrawImage2(a GraphicsImageInterface, b int, c int, d int, e int, f int, g int, h int, i int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"), b, c, d, e, f, g, h, i)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawLine(int, int, int, int)
func (jbobject *GraphicsGC) DrawLine(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawLine", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void drawOval(int, int, int, int)
func (jbobject *GraphicsGC) DrawOval(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawOval", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void drawPath(org.eclipse.swt.graphics.Path)
func (jbobject *GraphicsGC) DrawPath(a GraphicsPathInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawPath", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Path"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawPoint(int, int)
func (jbobject *GraphicsGC) DrawPoint(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawPoint", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void drawPolygon(int[])
func (jbobject *GraphicsGC) DrawPolygon(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawPolygon", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void drawPolyline(int[])
func (jbobject *GraphicsGC) DrawPolyline(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawPolyline", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void drawRectangle(int, int, int, int)
func (jbobject *GraphicsGC) DrawRectangle(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawRectangle", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void drawRectangle(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsGC) DrawRectangle2(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawRectangle", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawRoundRectangle(int, int, int, int, int, int)
func (jbobject *GraphicsGC) DrawRoundRectangle(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawRoundRectangle", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void drawString(java.lang.String, int, int)
func (jbobject *GraphicsGC) DrawString(a string, b int, c int)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawString", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawString(java.lang.String, int, int, boolean)
func (jbobject *GraphicsGC) DrawString2(a string, b int, c int, d bool)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawString", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawText(java.lang.String, int, int)
func (jbobject *GraphicsGC) DrawText(a string, b int, c int)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawText", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawText(java.lang.String, int, int, boolean)
func (jbobject *GraphicsGC) DrawText2(a string, b int, c int, d bool)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawText", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void drawText(java.lang.String, int, int, int)
func (jbobject *GraphicsGC) DrawText3(a string, b int, c int, d int)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "drawText", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsGC) Equals(a interface{}) bool {
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

// public void fillArc(int, int, int, int, int, int)
func (jbobject *GraphicsGC) FillArc(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillArc", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public void fillGradientRectangle(int, int, int, int, boolean)
func (jbobject *GraphicsGC) FillGradientRectangle(a int, b int, c int, d int, e bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillGradientRectangle", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public void fillOval(int, int, int, int)
func (jbobject *GraphicsGC) FillOval(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillOval", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void fillPath(org.eclipse.swt.graphics.Path)
func (jbobject *GraphicsGC) FillPath(a GraphicsPathInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillPath", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Path"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void fillPolygon(int[])
func (jbobject *GraphicsGC) FillPolygon(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillPolygon", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void fillRectangle(int, int, int, int)
func (jbobject *GraphicsGC) FillRectangle(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillRectangle", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void fillRectangle(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsGC) FillRectangle2(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillRectangle", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void fillRoundRectangle(int, int, int, int, int, int)
func (jbobject *GraphicsGC) FillRoundRectangle(a int, b int, c int, d int, e int, f int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "fillRoundRectangle", javabind.Void, a, b, c, d, e, f)
	if err != nil {
		panic(err)
	}

}

// public int getAdvanceWidth(char)
func (jbobject *GraphicsGC) GetAdvanceWidth(a uint16) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAdvanceWidth", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getAdvanced()
func (jbobject *GraphicsGC) GetAdvanced() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAdvanced", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getAlpha()
func (jbobject *GraphicsGC) GetAlpha() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlpha", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getAntialias()
func (jbobject *GraphicsGC) GetAntialias() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAntialias", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *GraphicsGC) GetBackground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBackground", "org/eclipse/swt/graphics/Color")
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

// public org.eclipse.swt.graphics.Pattern getBackgroundPattern()
func (jbobject *GraphicsGC) GetBackgroundPattern() *GraphicsPattern {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBackgroundPattern", "org/eclipse/swt/graphics/Pattern")
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

// public int getCharWidth(char)
func (jbobject *GraphicsGC) GetCharWidth(a uint16) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCharWidth", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Rectangle getClipping()
func (jbobject *GraphicsGC) GetClipping() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClipping", "org/eclipse/swt/graphics/Rectangle")
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

// public void getClipping(org.eclipse.swt.graphics.Region)
func (jbobject *GraphicsGC) GetClipping2(a GraphicsRegionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getClipping", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Region"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public int getFillRule()
func (jbobject *GraphicsGC) GetFillRule() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFillRule", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *GraphicsGC) GetFont() *GraphicsFont {
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

// public org.eclipse.swt.graphics.FontMetrics getFontMetrics()
func (jbobject *GraphicsGC) GetFontMetrics() *GraphicsFontMetrics {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFontMetrics", "org/eclipse/swt/graphics/FontMetrics")
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

// public org.eclipse.swt.graphics.Color getForeground()
func (jbobject *GraphicsGC) GetForeground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getForeground", "org/eclipse/swt/graphics/Color")
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

// public org.eclipse.swt.graphics.Pattern getForegroundPattern()
func (jbobject *GraphicsGC) GetForegroundPattern() *GraphicsPattern {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getForegroundPattern", "org/eclipse/swt/graphics/Pattern")
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

// public org.eclipse.swt.graphics.GCData getGCData()
func (jbobject *GraphicsGC) GetGCData() *GraphicsGCData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGCData", "org/eclipse/swt/graphics/GCData")
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
	unique_x := &GraphicsGCData{}
	unique_x.Callable = dst
	return unique_x
}

// public int getInterpolation()
func (jbobject *GraphicsGC) GetInterpolation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInterpolation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.LineAttributes getLineAttributes()
func (jbobject *GraphicsGC) GetLineAttributes() *GraphicsLineAttributes {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineAttributes", "org/eclipse/swt/graphics/LineAttributes")
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
	unique_x := &GraphicsLineAttributes{}
	unique_x.Callable = dst
	return unique_x
}

// public int getLineCap()
func (jbobject *GraphicsGC) GetLineCap() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineCap", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getLineDash()
func (jbobject *GraphicsGC) GetLineDash() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineDash", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int getLineJoin()
func (jbobject *GraphicsGC) GetLineJoin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineJoin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineStyle()
func (jbobject *GraphicsGC) GetLineStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineWidth()
func (jbobject *GraphicsGC) GetLineWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getStyle()
func (jbobject *GraphicsGC) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTextAntialias()
func (jbobject *GraphicsGC) GetTextAntialias() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextAntialias", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void getTransform(org.eclipse.swt.graphics.Transform)
func (jbobject *GraphicsGC) GetTransform(a GraphicsTransformInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getTransform", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Transform"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean getXORMode()
func (jbobject *GraphicsGC) GetXORMode() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getXORMode", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int hashCode()
func (jbobject *GraphicsGC) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isClipped()
func (jbobject *GraphicsGC) IsClipped() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isClipped", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isDisposed()
func (jbobject *GraphicsGC) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setAdvanced(boolean)
func (jbobject *GraphicsGC) SetAdvanced(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAdvanced", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setAlpha(int)
func (jbobject *GraphicsGC) SetAlpha(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlpha", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setAntialias(int)
func (jbobject *GraphicsGC) SetAntialias(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAntialias", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *GraphicsGC) SetBackground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBackgroundPattern(org.eclipse.swt.graphics.Pattern)
func (jbobject *GraphicsGC) SetBackgroundPattern(a GraphicsPatternInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackgroundPattern", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Pattern"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setClipping(int, int, int, int)
func (jbobject *GraphicsGC) SetClipping(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClipping", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void setClipping(org.eclipse.swt.graphics.Path)
func (jbobject *GraphicsGC) SetClipping2(a GraphicsPathInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClipping", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Path"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setClipping(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsGC) SetClipping3(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClipping", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setClipping(org.eclipse.swt.graphics.Region)
func (jbobject *GraphicsGC) SetClipping4(a GraphicsRegionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClipping", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Region"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *GraphicsGC) SetFont(a GraphicsFontInterface)  {
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

// public void setFillRule(int)
func (jbobject *GraphicsGC) SetFillRule(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFillRule", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setForeground(org.eclipse.swt.graphics.Color)
func (jbobject *GraphicsGC) SetForeground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setForeground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setForegroundPattern(org.eclipse.swt.graphics.Pattern)
func (jbobject *GraphicsGC) SetForegroundPattern(a GraphicsPatternInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setForegroundPattern", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Pattern"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setInterpolation(int)
func (jbobject *GraphicsGC) SetInterpolation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInterpolation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLineAttributes(org.eclipse.swt.graphics.LineAttributes)
func (jbobject *GraphicsGC) SetLineAttributes(a GraphicsLineAttributesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineAttributes", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/LineAttributes"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setLineCap(int)
func (jbobject *GraphicsGC) SetLineCap(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineCap", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLineDash(int[])
func (jbobject *GraphicsGC) SetLineDash(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineDash", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLineJoin(int)
func (jbobject *GraphicsGC) SetLineJoin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineJoin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLineStyle(int)
func (jbobject *GraphicsGC) SetLineStyle(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineStyle", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLineWidth(int)
func (jbobject *GraphicsGC) SetLineWidth(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineWidth", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTextAntialias(int)
func (jbobject *GraphicsGC) SetTextAntialias(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTextAntialias", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTransform(org.eclipse.swt.graphics.Transform)
func (jbobject *GraphicsGC) SetTransform(a GraphicsTransformInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTransform", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Transform"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setXORMode(boolean)
func (jbobject *GraphicsGC) SetXORMode(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setXORMode", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point stringExtent(java.lang.String)
func (jbobject *GraphicsGC) StringExtent(a string) *GraphicsPoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stringExtent", "org/eclipse/swt/graphics/Point", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
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

// public org.eclipse.swt.graphics.Point textExtent(java.lang.String)
func (jbobject *GraphicsGC) TextExtent(a string) *GraphicsPoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "textExtent", "org/eclipse/swt/graphics/Point", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
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

// public org.eclipse.swt.graphics.Point textExtent(java.lang.String, int)
func (jbobject *GraphicsGC) TextExtent2(a string, b int) *GraphicsPoint {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "textExtent", "org/eclipse/swt/graphics/Point", conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
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

// public java.lang.String toString()
func (jbobject *GraphicsGC) ToString() string {
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

func (jbobject *GraphicsGC) Handle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsGC) SetFieldHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


