package swt

import "github.com/timob/javabind"

type CustomTableTreeItemInterface interface {
	WidgetsItemInterface

	// public org.eclipse.swt.graphics.Color getBackground()
	GetBackground() *GraphicsColor

	// public org.eclipse.swt.graphics.Rectangle getBounds(int)
	GetBounds(a int) *GraphicsRectangle

	// public boolean getChecked()
	GetChecked() bool

	// public boolean getGrayed()
	GetGrayed() bool

	// public boolean getExpanded()
	GetExpanded() bool

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public org.eclipse.swt.graphics.Color getForeground()
	GetForeground() *GraphicsColor

	// public org.eclipse.swt.graphics.Image getImage(int)
	GetImage3(a int) *GraphicsImage

	// public org.eclipse.swt.custom.TableTreeItem getItem(int)
	GetItem(a int) *CustomTableTreeItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.custom.TableTreeItem[] getItems()
	GetItems() []*CustomTableTreeItem

	// public org.eclipse.swt.custom.TableTree getParent()
	GetParent() *CustomTableTree

	// public org.eclipse.swt.custom.TableTreeItem getParentItem()
	GetParentItem() *CustomTableTreeItem

	// public java.lang.String getText(int)
	GetText3(a int) string

	// public int indexOf(org.eclipse.swt.custom.TableTreeItem)
	IndexOf(a CustomTableTreeItemInterface) int

	// public void setBackground(org.eclipse.swt.graphics.Color)
	SetBackground(a GraphicsColorInterface) 

	// public void setChecked(boolean)
	SetChecked(a bool) 

	// public void setGrayed(boolean)
	SetGrayed(a bool) 

	// public void setExpanded(boolean)
	SetExpanded(a bool) 

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setForeground(org.eclipse.swt.graphics.Color)
	SetForeground(a GraphicsColorInterface) 

	// public void setImage(int, org.eclipse.swt.graphics.Image)
	SetImage3(a int, b GraphicsImageInterface) 

	// public void setText(int, java.lang.String)
	SetText3(a int, b string) 
}

type CustomTableTreeItem struct {
	WidgetsItem
}

// public org.eclipse.swt.custom.TableTreeItem(org.eclipse.swt.custom.TableTree, int)
func NewCustomTableTreeItem(a CustomTableTreeInterface, b int) (*CustomTableTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableTreeItem", conv_a.Value().Cast("org/eclipse/swt/custom/TableTree"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.TableTreeItem(org.eclipse.swt.custom.TableTree, int, int)
func NewCustomTableTreeItem3(a CustomTableTreeInterface, b int, c int) (*CustomTableTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableTreeItem", conv_a.Value().Cast("org/eclipse/swt/custom/TableTree"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.TableTreeItem(org.eclipse.swt.custom.TableTreeItem, int)
func NewCustomTableTreeItem2(a CustomTableTreeItemInterface, b int) (*CustomTableTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableTreeItem", conv_a.Value().Cast("org/eclipse/swt/custom/TableTreeItem"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.TableTreeItem(org.eclipse.swt.custom.TableTreeItem, int, int)
func NewCustomTableTreeItem4(a CustomTableTreeItemInterface, b int, c int) (*CustomTableTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableTreeItem", conv_a.Value().Cast("org/eclipse/swt/custom/TableTreeItem"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *CustomTableTreeItem) GetBackground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Rectangle getBounds(int)
func (jbobject *CustomTableTreeItem) GetBounds(a int) *GraphicsRectangle {
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
func (jbobject *CustomTableTreeItem) GetChecked() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getChecked", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getGrayed()
func (jbobject *CustomTableTreeItem) GetGrayed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGrayed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getExpanded()
func (jbobject *CustomTableTreeItem) GetExpanded() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpanded", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *CustomTableTreeItem) GetFont() *GraphicsFont {
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

// public org.eclipse.swt.graphics.Color getForeground()
func (jbobject *CustomTableTreeItem) GetForeground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *CustomTableTreeItem) GetImage2() *GraphicsImage {
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
func (jbobject *CustomTableTreeItem) GetImage3(a int) *GraphicsImage {
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

// public org.eclipse.swt.custom.TableTreeItem getItem(int)
func (jbobject *CustomTableTreeItem) GetItem(a int) *CustomTableTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/custom/TableTreeItem", a)
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
	unique_x := &CustomTableTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *CustomTableTreeItem) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.custom.TableTreeItem[] getItems()
func (jbobject *CustomTableTreeItem) GetItems() []*CustomTableTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/TableTreeItem")
	dst := new([]*CustomTableTreeItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.custom.TableTree getParent()
func (jbobject *CustomTableTreeItem) GetParent() *CustomTableTree {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/custom/TableTree")
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
	unique_x := &CustomTableTree{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.custom.TableTreeItem getParentItem()
func (jbobject *CustomTableTreeItem) GetParentItem() *CustomTableTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParentItem", "org/eclipse/swt/custom/TableTreeItem")
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
	unique_x := &CustomTableTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getText()
func (jbobject *CustomTableTreeItem) GetText2() string {
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
func (jbobject *CustomTableTreeItem) GetText3(a int) string {
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

// public int indexOf(org.eclipse.swt.custom.TableTreeItem)
func (jbobject *CustomTableTreeItem) IndexOf(a CustomTableTreeItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public void dispose()
func (jbobject *CustomTableTreeItem) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomTableTreeItem) SetBackground(a GraphicsColorInterface)  {
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

// public void setChecked(boolean)
func (jbobject *CustomTableTreeItem) SetChecked(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setChecked", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setGrayed(boolean)
func (jbobject *CustomTableTreeItem) SetGrayed(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGrayed", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setExpanded(boolean)
func (jbobject *CustomTableTreeItem) SetExpanded(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpanded", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomTableTreeItem) SetFont(a GraphicsFontInterface)  {
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

// public void setForeground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomTableTreeItem) SetForeground(a GraphicsColorInterface)  {
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

// public void setImage(int, org.eclipse.swt.graphics.Image)
func (jbobject *CustomTableTreeItem) SetImage3(a int, b GraphicsImageInterface)  {
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
func (jbobject *CustomTableTreeItem) SetImage2(a GraphicsImageInterface)  {
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

// public void setText(int, java.lang.String)
func (jbobject *CustomTableTreeItem) SetText3(a int, b string)  {
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
func (jbobject *CustomTableTreeItem) SetText2(a string)  {
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


