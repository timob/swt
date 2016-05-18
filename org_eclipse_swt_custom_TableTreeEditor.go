package swt

import "github.com/timob/javabind"

type CustomTableTreeEditorInterface interface {
	CustomControlEditorInterface

	// public int getColumn()
	GetColumn() int

	// public org.eclipse.swt.custom.TableTreeItem getItem()
	GetItem() *CustomTableTreeItem

	// public void setColumn(int)
	SetColumn(a int) 

	// public void setItem(org.eclipse.swt.custom.TableTreeItem)
	SetItem(a CustomTableTreeItemInterface) 

	// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.custom.TableTreeItem, int)
	SetEditor2(a WidgetsControlInterface, b CustomTableTreeItemInterface, c int) 
}

type CustomTableTreeEditor struct {
	CustomControlEditor
}

// public org.eclipse.swt.custom.TableTreeEditor(org.eclipse.swt.custom.TableTree)
func NewCustomTableTreeEditor(a CustomTableTreeInterface) (*CustomTableTreeEditor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableTreeEditor", conv_a.Value().Cast("org/eclipse/swt/custom/TableTree"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableTreeEditor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dispose()
func (jbobject *CustomTableTreeEditor) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public int getColumn()
func (jbobject *CustomTableTreeEditor) GetColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.custom.TableTreeItem getItem()
func (jbobject *CustomTableTreeEditor) GetItem() *CustomTableTreeItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/custom/TableTreeItem")
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

// public void setColumn(int)
func (jbobject *CustomTableTreeEditor) SetColumn(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setColumn", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setItem(org.eclipse.swt.custom.TableTreeItem)
func (jbobject *CustomTableTreeEditor) SetItem(a CustomTableTreeItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TableTreeItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.custom.TableTreeItem, int)
func (jbobject *CustomTableTreeEditor) SetEditor2(a WidgetsControlInterface, b CustomTableTreeItemInterface, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/custom/TableTreeItem"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void layout()
func (jbobject *CustomTableTreeEditor) Layout()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void)
	if err != nil {
		panic(err)
	}

}


