package swt

import "github.com/timob/javabind"

type WidgetsTrayInterface interface {
	WidgetsWidgetInterface

	// public org.eclipse.swt.widgets.TrayItem getItem(int)
	GetItem(a int) *WidgetsTrayItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.TrayItem[] getItems()
	GetItems() []*WidgetsTrayItem
}

type WidgetsTray struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.TrayItem getItem(int)
func (jbobject *WidgetsTray) GetItem(a int) *WidgetsTrayItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TrayItem", a)
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
	unique_x := &WidgetsTrayItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsTray) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TrayItem[] getItems()
func (jbobject *WidgetsTray) GetItems() []*WidgetsTrayItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/TrayItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TrayItem")
	dst := new([]*WidgetsTrayItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


