package swt

import "github.com/timob/javabind"

type WidgetsTreeItemInterface interface {
	WidgetsItemInterface

	// public void clear(int, boolean)
	Clear(a int, b bool) 

	// public void clearAll(boolean)
	ClearAll(a bool) 

	// public org.eclipse.swt.graphics.Color getBackground()
	GetBackground() *GraphicsColor

	// public org.eclipse.swt.graphics.Color getBackground(int)
	GetBackground2(a int) *GraphicsColor

	// public org.eclipse.swt.graphics.Rectangle getBounds(int)
	GetBounds2(a int) *GraphicsRectangle

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public boolean getChecked()
	GetChecked() bool

	// public boolean getExpanded()
	GetExpanded() bool

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

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.TreeItem getItem(int)
	GetItem(a int) *WidgetsTreeItem

	// public org.eclipse.swt.widgets.TreeItem[] getItems()
	GetItems() []*WidgetsTreeItem

	// public org.eclipse.swt.widgets.Tree getParent()
	GetParent() *WidgetsTree

	// public org.eclipse.swt.widgets.TreeItem getParentItem()
	GetParentItem() *WidgetsTreeItem

	// public java.lang.String getText(int)
	GetText2(a int) string

	// public org.eclipse.swt.graphics.Rectangle getTextBounds(int)
	GetTextBounds(a int) *GraphicsRectangle

	// public int indexOf(org.eclipse.swt.widgets.TreeItem)
	IndexOf(a WidgetsTreeItemInterface) int

	// public void removeAll()
	RemoveAll() 

	// public void setBackground(org.eclipse.swt.graphics.Color)
	SetBackground(a GraphicsColorInterface) 

	// public void setBackground(int, org.eclipse.swt.graphics.Color)
	SetBackground2(a int, b GraphicsColorInterface) 

	// public void setChecked(boolean)
	SetChecked(a bool) 

	// public void setExpanded(boolean)
	SetExpanded(a bool) 

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

	// public void setItemCount(int)
	SetItemCount(a int) 

	// public void setText(int, java.lang.String)
	SetText3(a int, b string) 

	// public void setText(java.lang.String[])
	SetText2(a []string) 
}

type WidgetsTreeItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.TreeItem(org.eclipse.swt.widgets.Tree, int)
func NewWidgetsTreeItem(a WidgetsTreeInterface, b int) (*WidgetsTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TreeItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Tree"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.TreeItem(org.eclipse.swt.widgets.Tree, int, int)
func NewWidgetsTreeItem3(a WidgetsTreeInterface, b int, c int) (*WidgetsTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TreeItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Tree"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.TreeItem(org.eclipse.swt.widgets.TreeItem, int)
func NewWidgetsTreeItem2(a WidgetsTreeItemInterface, b int) (*WidgetsTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TreeItem", conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.TreeItem(org.eclipse.swt.widgets.TreeItem, int, int)
func NewWidgetsTreeItem4(a WidgetsTreeItemInterface, b int, c int) (*WidgetsTreeItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/TreeItem", conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTreeItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void clear(int, boolean)
func (jbobject *WidgetsTreeItem) Clear(a int, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void clearAll(boolean)
func (jbobject *WidgetsTreeItem) ClearAll(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearAll", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *WidgetsTreeItem) GetBackground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Color getBackground(int)
func (jbobject *WidgetsTreeItem) GetBackground2(a int) *GraphicsColor {
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
func (jbobject *WidgetsTreeItem) GetBounds2(a int) *GraphicsRectangle {
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

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsTreeItem) GetBounds() *GraphicsRectangle {
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

// public boolean getChecked()
func (jbobject *WidgetsTreeItem) GetChecked() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getChecked", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getExpanded()
func (jbobject *WidgetsTreeItem) GetExpanded() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpanded", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *WidgetsTreeItem) GetFont() *GraphicsFont {
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
func (jbobject *WidgetsTreeItem) GetFont2(a int) *GraphicsFont {
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
func (jbobject *WidgetsTreeItem) GetForeground() *GraphicsColor {
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
func (jbobject *WidgetsTreeItem) GetForeground2(a int) *GraphicsColor {
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
func (jbobject *WidgetsTreeItem) GetGrayed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGrayed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Image getImage()
func (jbobject *WidgetsTreeItem) GetImage() *GraphicsImage {
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
func (jbobject *WidgetsTreeItem) GetImage2(a int) *GraphicsImage {
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
func (jbobject *WidgetsTreeItem) GetImageBounds(a int) *GraphicsRectangle {
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

// public int getItemCount()
func (jbobject *WidgetsTreeItem) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TreeItem getItem(int)
func (jbobject *WidgetsTreeItem) GetItem(a int) *WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TreeItem", a)
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
	unique_x := &WidgetsTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.TreeItem[] getItems()
func (jbobject *WidgetsTreeItem) GetItems() []*WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TreeItem")
	dst := new([]*WidgetsTreeItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.widgets.Tree getParent()
func (jbobject *WidgetsTreeItem) GetParent() *WidgetsTree {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Tree")
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
	unique_x := &WidgetsTree{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.TreeItem getParentItem()
func (jbobject *WidgetsTreeItem) GetParentItem() *WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParentItem", "org/eclipse/swt/widgets/TreeItem")
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
	unique_x := &WidgetsTreeItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getText()
func (jbobject *WidgetsTreeItem) GetText() string {
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
func (jbobject *WidgetsTreeItem) GetText2(a int) string {
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
func (jbobject *WidgetsTreeItem) GetTextBounds(a int) *GraphicsRectangle {
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

// public int indexOf(org.eclipse.swt.widgets.TreeItem)
func (jbobject *WidgetsTreeItem) IndexOf(a WidgetsTreeItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public void removeAll()
func (jbobject *WidgetsTreeItem) RemoveAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *WidgetsTreeItem) SetBackground(a GraphicsColorInterface)  {
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
func (jbobject *WidgetsTreeItem) SetBackground2(a int, b GraphicsColorInterface)  {
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
func (jbobject *WidgetsTreeItem) SetChecked(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setChecked", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setExpanded(boolean)
func (jbobject *WidgetsTreeItem) SetExpanded(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpanded", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *WidgetsTreeItem) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *WidgetsTreeItem) SetFont2(a int, b GraphicsFontInterface)  {
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
func (jbobject *WidgetsTreeItem) SetForeground(a GraphicsColorInterface)  {
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
func (jbobject *WidgetsTreeItem) SetForeground2(a int, b GraphicsColorInterface)  {
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
func (jbobject *WidgetsTreeItem) SetGrayed(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGrayed", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(int, org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTreeItem) SetImage3(a int, b GraphicsImageInterface)  {
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
func (jbobject *WidgetsTreeItem) SetImage(a GraphicsImageInterface)  {
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
func (jbobject *WidgetsTreeItem) SetImage2(a []*GraphicsImage)  {
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

// public void setItemCount(int)
func (jbobject *WidgetsTreeItem) SetItemCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItemCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(int, java.lang.String)
func (jbobject *WidgetsTreeItem) SetText3(a int, b string)  {
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
func (jbobject *WidgetsTreeItem) SetText(a string)  {
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
func (jbobject *WidgetsTreeItem) SetText2(a []string)  {
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


