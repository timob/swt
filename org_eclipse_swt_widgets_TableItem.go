package swt

import "github.com/timob/javabind"

type WidgetsTableItemInterface interface {
	WidgetsItemInterface

	// public org.eclipse.swt.graphics.Color getBackground()
	GetBackground() *GraphicsColor

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.graphics.Color getBackground(int)
	GetBackground2(a int) *GraphicsColor

	// public org.eclipse.swt.graphics.Rectangle getBounds(int)
	GetBounds2(a int) *GraphicsRectangle

	// public boolean getChecked()
	GetChecked() bool

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public org.eclipse.swt.graphics.Font getFont(int)
	GetFont2(a int) *GraphicsFont

	// public org.eclipse.swt.graphics.Color getForeground()
	GetForeground() *GraphicsColor

	// public org.eclipse.swt.graphics.Color getForeground(int)
	GetForeground2(a int) *GraphicsColor

	// public boolean getGrayed()
	GetGrayed() bool

	// public org.eclipse.swt.graphics.Image getImage(int)
	GetImage2(a int) *GraphicsImage

	// public org.eclipse.swt.graphics.Rectangle getImageBounds(int)
	GetImageBounds(a int) *GraphicsRectangle

	// public int getImageIndent()
	GetImageIndent() int

	// public org.eclipse.swt.widgets.Table getParent()
	GetParent() *WidgetsTable

	// public java.lang.String getText(int)
	GetText2(a int) string

	// public org.eclipse.swt.graphics.Rectangle getTextBounds(int)
	GetTextBounds(a int) *GraphicsRectangle

	// public void setBackground(org.eclipse.swt.graphics.Color)
	SetBackground(a GraphicsColorInterface) 

	// public void setBackground(int, org.eclipse.swt.graphics.Color)
	SetBackground2(a int, b GraphicsColorInterface) 

	// public void setChecked(boolean)
	SetChecked(a bool) 

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setFont(int, org.eclipse.swt.graphics.Font)
	SetFont2(a int, b GraphicsFontInterface) 

	// public void setForeground(org.eclipse.swt.graphics.Color)
	SetForeground(a GraphicsColorInterface) 

	// public void setForeground(int, org.eclipse.swt.graphics.Color)
	SetForeground2(a int, b GraphicsColorInterface) 

	// public void setGrayed(boolean)
	SetGrayed(a bool) 

	// public void setImage(int, org.eclipse.swt.graphics.Image)
	SetImage3(a int, b GraphicsImageInterface) 

	// public void setImage(org.eclipse.swt.graphics.Image[])
	SetImage2(a []*GraphicsImage) 

	// public void setImageIndent(int)
	SetImageIndent(a int) 

	// public void setText(int, java.lang.String)
	SetText3(a int, b string) 

	// public void setText(java.lang.String[])
	SetText2(a []string) 
}

type WidgetsTableItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.TableItem(org.eclipse.swt.widgets.Table, int, int)
func NewWidgetsTableItem2(a WidgetsTableInterface, b int, c int) (*WidgetsTableItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TableItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Table"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTableItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.TableItem(org.eclipse.swt.widgets.Table, int)
func NewWidgetsTableItem(a WidgetsTableInterface, b int) (*WidgetsTableItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TableItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Table"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTableItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *WidgetsTableItem) GetBackground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsTableItem) GetBounds() *GraphicsRectangle {
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

// public org.eclipse.swt.graphics.Color getBackground(int)
func (jbobject *WidgetsTableItem) GetBackground2(a int) *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBackground", "org/eclipse/swt/graphics/Color", a)
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

// public org.eclipse.swt.graphics.Rectangle getBounds(int)
func (jbobject *WidgetsTableItem) GetBounds2(a int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", "org/eclipse/swt/graphics/Rectangle", a)
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

// public boolean getChecked()
func (jbobject *WidgetsTableItem) GetChecked() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getChecked", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *WidgetsTableItem) GetFont() *GraphicsFont {
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

// public org.eclipse.swt.graphics.Font getFont(int)
func (jbobject *WidgetsTableItem) GetFont2(a int) *GraphicsFont {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFont", "org/eclipse/swt/graphics/Font", a)
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

// public org.eclipse.swt.graphics.Color getForeground()
func (jbobject *WidgetsTableItem) GetForeground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Color getForeground(int)
func (jbobject *WidgetsTableItem) GetForeground2(a int) *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getForeground", "org/eclipse/swt/graphics/Color", a)
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

// public boolean getGrayed()
func (jbobject *WidgetsTableItem) GetGrayed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGrayed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *WidgetsTableItem) GetImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImage", "org/eclipse/swt/graphics/Image")
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

// public org.eclipse.swt.graphics.Image getImage(int)
func (jbobject *WidgetsTableItem) GetImage2(a int) *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImage", "org/eclipse/swt/graphics/Image", a)
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

// public org.eclipse.swt.graphics.Rectangle getImageBounds(int)
func (jbobject *WidgetsTableItem) GetImageBounds(a int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageBounds", "org/eclipse/swt/graphics/Rectangle", a)
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

// public int getImageIndent()
func (jbobject *WidgetsTableItem) GetImageIndent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Table getParent()
func (jbobject *WidgetsTableItem) GetParent() *WidgetsTable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Table")
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
	unique_x := &WidgetsTable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getText()
func (jbobject *WidgetsTableItem) GetText() string {
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

// public java.lang.String getText(int)
func (jbobject *WidgetsTableItem) GetText2(a int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getText", "java/lang/String", a)
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

// public org.eclipse.swt.graphics.Rectangle getTextBounds(int)
func (jbobject *WidgetsTableItem) GetTextBounds(a int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextBounds", "org/eclipse/swt/graphics/Rectangle", a)
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

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *WidgetsTableItem) SetBackground(a GraphicsColorInterface)  {
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

// public void setBackground(int, org.eclipse.swt.graphics.Color)
func (jbobject *WidgetsTableItem) SetBackground2(a int, b GraphicsColorInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void setChecked(boolean)
func (jbobject *WidgetsTableItem) SetChecked(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setChecked", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *WidgetsTableItem) SetFont(a GraphicsFontInterface)  {
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

// public void setFont(int, org.eclipse.swt.graphics.Font)
func (jbobject *WidgetsTableItem) SetFont2(a int, b GraphicsFontInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFont", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/Font"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void setForeground(org.eclipse.swt.graphics.Color)
func (jbobject *WidgetsTableItem) SetForeground(a GraphicsColorInterface)  {
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

// public void setForeground(int, org.eclipse.swt.graphics.Color)
func (jbobject *WidgetsTableItem) SetForeground2(a int, b GraphicsColorInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setForeground", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void setGrayed(boolean)
func (jbobject *WidgetsTableItem) SetGrayed(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGrayed", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(int, org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTableItem) SetImage3(a int, b GraphicsImageInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTableItem) SetImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImage(org.eclipse.swt.graphics.Image[])
func (jbobject *WidgetsTableItem) SetImage2(a []*GraphicsImage)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Image")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImageIndent(int)
func (jbobject *WidgetsTableItem) SetImageIndent(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageIndent", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(int, java.lang.String)
func (jbobject *WidgetsTableItem) SetText3(a int, b string)  {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setText", javabind.Void, a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void setText(java.lang.String)
func (jbobject *WidgetsTableItem) SetText(a string)  {
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

// public void setText(java.lang.String[])
func (jbobject *WidgetsTableItem) SetText2(a []string)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


