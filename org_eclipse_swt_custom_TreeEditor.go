package swt

import "github.com/timob/javabind"

type CustomTreeEditorInterface interface {
	CustomControlEditorInterface

	// public int getColumn()
	GetColumn() int

	// public org.eclipse.swt.widgets.TreeItem getItem()
	GetItem() *WidgetsTreeItem

	// public void setColumn(int)
	SetColumn(a int) 

	// public void setItem(org.eclipse.swt.widgets.TreeItem)
	SetItem(a WidgetsTreeItemInterface) 

	// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.TreeItem, int)
	SetEditor4(a WidgetsControlInterface, b WidgetsTreeItemInterface, c int) 

	// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.TreeItem)
	SetEditor3(a WidgetsControlInterface, b WidgetsTreeItemInterface) 
}

type CustomTreeEditor struct {
	CustomControlEditor
}

// public org.eclipse.swt.custom.TreeEditor(org.eclipse.swt.widgets.Tree)
func NewCustomTreeEditor(a WidgetsTreeInterface) (*CustomTreeEditor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TreeEditor", conv_a.Value().Cast("org/eclipse/swt/widgets/Tree"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTreeEditor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dispose()
func (jbobject *CustomTreeEditor) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public int getColumn()
func (jbobject *CustomTreeEditor) GetColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TreeItem getItem()
func (jbobject *CustomTreeEditor) GetItem() *WidgetsTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TreeItem")
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

// public void setColumn(int)
func (jbobject *CustomTreeEditor) SetColumn(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setColumn", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setItem(org.eclipse.swt.widgets.TreeItem)
func (jbobject *CustomTreeEditor) SetItem(a WidgetsTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.TreeItem, int)
func (jbobject *CustomTreeEditor) SetEditor4(a WidgetsControlInterface, b WidgetsTreeItemInterface, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/widgets/TreeItem"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void setEditor(org.eclipse.swt.widgets.Control)
func (jbobject *CustomTreeEditor) SetEditor2(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.TreeItem)
func (jbobject *CustomTreeEditor) SetEditor3(a WidgetsControlInterface, b WidgetsTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/widgets/TreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void layout()
func (jbobject *CustomTreeEditor) Layout()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void)
	if err != nil {
		panic(err)
	}

}


