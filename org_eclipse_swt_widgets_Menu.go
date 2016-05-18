package swt

import "github.com/timob/javabind"

type WidgetsMenuInterface interface {
	WidgetsWidgetInterface

	// public void addMenuListener(org.eclipse.swt.events.MenuListener)
	AddMenuListener(a EventsMenuListenerInterface) 

	// public void addHelpListener(org.eclipse.swt.events.HelpListener)
	AddHelpListener(a EventsHelpListenerInterface) 

	// public org.eclipse.swt.widgets.MenuItem getDefaultItem()
	GetDefaultItem() *WidgetsMenuItem

	// public boolean getEnabled()
	GetEnabled() bool

	// public org.eclipse.swt.widgets.MenuItem getItem(int)
	GetItem(a int) *WidgetsMenuItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.MenuItem[] getItems()
	GetItems() []*WidgetsMenuItem

	// public int getOrientation()
	GetOrientation() int

	// public org.eclipse.swt.widgets.Decorations getParent()
	GetParent() *WidgetsDecorations

	// public org.eclipse.swt.widgets.MenuItem getParentItem()
	GetParentItem() *WidgetsMenuItem

	// public org.eclipse.swt.widgets.Menu getParentMenu()
	GetParentMenu() *WidgetsMenu

	// public org.eclipse.swt.widgets.Shell getShell()
	GetShell() *WidgetsShell

	// public boolean getVisible()
	GetVisible() bool

	// public int indexOf(org.eclipse.swt.widgets.MenuItem)
	IndexOf(a WidgetsMenuItemInterface) int

	// public boolean isEnabled()
	IsEnabled() bool

	// public boolean isVisible()
	IsVisible() bool

	// public void removeMenuListener(org.eclipse.swt.events.MenuListener)
	RemoveMenuListener(a EventsMenuListenerInterface) 

	// public void removeHelpListener(org.eclipse.swt.events.HelpListener)
	RemoveHelpListener(a EventsHelpListenerInterface) 

	// public void setDefaultItem(org.eclipse.swt.widgets.MenuItem)
	SetDefaultItem(a WidgetsMenuItemInterface) 

	// public void setEnabled(boolean)
	SetEnabled(a bool) 

	// public void setLocation(int, int)
	SetLocation(a int, b int) 

	// public void setLocation(org.eclipse.swt.graphics.Point)
	SetLocation2(a GraphicsPointInterface) 

	// public void setOrientation(int)
	SetOrientation(a int) 

	// public void setVisible(boolean)
	SetVisible(a bool) 
}

type WidgetsMenu struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.Menu(org.eclipse.swt.widgets.Control)
func NewWidgetsMenu(a WidgetsControlInterface) (*WidgetsMenu) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Menu", conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMenu{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Menu(org.eclipse.swt.widgets.Decorations, int)
func NewWidgetsMenu4(a WidgetsDecorationsInterface, b int) (*WidgetsMenu) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Menu", conv_a.Value().Cast("org/eclipse/swt/widgets/Decorations"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMenu{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Menu(org.eclipse.swt.widgets.Menu)
func NewWidgetsMenu2(a WidgetsMenuInterface) (*WidgetsMenu) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Menu", conv_a.Value().Cast("org/eclipse/swt/widgets/Menu"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMenu{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Menu(org.eclipse.swt.widgets.MenuItem)
func NewWidgetsMenu3(a WidgetsMenuItemInterface) (*WidgetsMenu) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Menu", conv_a.Value().Cast("org/eclipse/swt/widgets/MenuItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMenu{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addMenuListener(org.eclipse.swt.events.MenuListener)
func (jbobject *WidgetsMenu) AddMenuListener(a EventsMenuListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addMenuListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addHelpListener(org.eclipse.swt.events.HelpListener)
func (jbobject *WidgetsMenu) AddHelpListener(a EventsHelpListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addHelpListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/HelpListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.widgets.MenuItem getDefaultItem()
func (jbobject *WidgetsMenu) GetDefaultItem() *WidgetsMenuItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDefaultItem", "org/eclipse/swt/widgets/MenuItem")
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
	unique_x := &WidgetsMenuItem{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getEnabled()
func (jbobject *WidgetsMenu) GetEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.widgets.MenuItem getItem(int)
func (jbobject *WidgetsMenu) GetItem(a int) *WidgetsMenuItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/MenuItem", a)
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
	unique_x := &WidgetsMenuItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsMenu) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.MenuItem[] getItems()
func (jbobject *WidgetsMenu) GetItems() []*WidgetsMenuItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/MenuItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/MenuItem")
	dst := new([]*WidgetsMenuItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getOrientation()
func (jbobject *WidgetsMenu) GetOrientation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrientation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Decorations getParent()
func (jbobject *WidgetsMenu) GetParent() *WidgetsDecorations {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Decorations")
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
	unique_x := &WidgetsDecorations{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.MenuItem getParentItem()
func (jbobject *WidgetsMenu) GetParentItem() *WidgetsMenuItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParentItem", "org/eclipse/swt/widgets/MenuItem")
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
	unique_x := &WidgetsMenuItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Menu getParentMenu()
func (jbobject *WidgetsMenu) GetParentMenu() *WidgetsMenu {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParentMenu", "org/eclipse/swt/widgets/Menu")
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
	unique_x := &WidgetsMenu{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Shell getShell()
func (jbobject *WidgetsMenu) GetShell() *WidgetsShell {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getShell", "org/eclipse/swt/widgets/Shell")
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
	unique_x := &WidgetsShell{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getVisible()
func (jbobject *WidgetsMenu) GetVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int indexOf(org.eclipse.swt.widgets.MenuItem)
func (jbobject *WidgetsMenu) IndexOf(a WidgetsMenuItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/MenuItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public boolean isEnabled()
func (jbobject *WidgetsMenu) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isVisible()
func (jbobject *WidgetsMenu) IsVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeMenuListener(org.eclipse.swt.events.MenuListener)
func (jbobject *WidgetsMenu) RemoveMenuListener(a EventsMenuListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeMenuListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeHelpListener(org.eclipse.swt.events.HelpListener)
func (jbobject *WidgetsMenu) RemoveHelpListener(a EventsHelpListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeHelpListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/HelpListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setDefaultItem(org.eclipse.swt.widgets.MenuItem)
func (jbobject *WidgetsMenu) SetDefaultItem(a WidgetsMenuItemInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDefaultItem", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/MenuItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setEnabled(boolean)
func (jbobject *WidgetsMenu) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLocation(int, int)
func (jbobject *WidgetsMenu) SetLocation(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocation", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setLocation(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsMenu) SetLocation2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocation", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setOrientation(int)
func (jbobject *WidgetsMenu) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setVisible(boolean)
func (jbobject *WidgetsMenu) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


