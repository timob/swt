package swt

import "github.com/timob/javabind"

type CustomTableEditorInterface interface {
	CustomControlEditorInterface

	// public int getColumn()
	GetColumn() int

	// public org.eclipse.swt.widgets.TableItem getItem()
	GetItem() *WidgetsTableItem

	// public void setColumn(int)
	SetColumn(a int) 

	// public void setItem(org.eclipse.swt.widgets.TableItem)
	SetItem(a WidgetsTableItemInterface) 

	// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.TableItem, int)
	SetEditor3(a WidgetsControlInterface, b WidgetsTableItemInterface, c int) 
}

type CustomTableEditor struct {
	CustomControlEditor
}

// public org.eclipse.swt.custom.TableEditor(org.eclipse.swt.widgets.Table)
func NewCustomTableEditor(a WidgetsTableInterface) (*CustomTableEditor) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TableEditor", conv_a.Value().Cast("org/eclipse/swt/widgets/Table"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTableEditor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void dispose()
func (jbobject *CustomTableEditor) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public int getColumn()
func (jbobject *CustomTableEditor) GetColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TableItem getItem()
func (jbobject *CustomTableEditor) GetItem() *WidgetsTableItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TableItem")
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
	unique_x := &WidgetsTableItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setColumn(int)
func (jbobject *CustomTableEditor) SetColumn(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setColumn", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setItem(org.eclipse.swt.widgets.TableItem)
func (jbobject *CustomTableEditor) SetItem(a WidgetsTableItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/TableItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setEditor(org.eclipse.swt.widgets.Control)
func (jbobject *CustomTableEditor) SetEditor2(a WidgetsControlInterface)  {
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

// public void setEditor(org.eclipse.swt.widgets.Control, org.eclipse.swt.widgets.TableItem, int)
func (jbobject *CustomTableEditor) SetEditor3(a WidgetsControlInterface, b WidgetsTableItemInterface, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"), conv_b.Value().Cast("org/eclipse/swt/widgets/TableItem"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void layout()
func (jbobject *CustomTableEditor) Layout()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "layout", javabind.Void)
	if err != nil {
		panic(err)
	}

}


