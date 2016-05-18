package swt

import "github.com/timob/javabind"

type WidgetsMenuItemInterface interface {
	WidgetsItemInterface

	// public void addArmListener(org.eclipse.swt.events.ArmListener)
	AddArmListener(a EventsArmListenerInterface) 

	// public void addHelpListener(org.eclipse.swt.events.HelpListener)
	AddHelpListener(a EventsHelpListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public int getAccelerator()
	GetAccelerator() int

	// public boolean getEnabled()
	GetEnabled() bool

	// public int getID()
	GetID() int

	// public org.eclipse.swt.widgets.Menu getMenu()
	GetMenu() *WidgetsMenu

	// public org.eclipse.swt.widgets.Menu getParent()
	GetParent() *WidgetsMenu

	// public boolean getSelection()
	GetSelection() bool

	// public boolean isEnabled()
	IsEnabled() bool

	// public void removeArmListener(org.eclipse.swt.events.ArmListener)
	RemoveArmListener(a EventsArmListenerInterface) 

	// public void removeHelpListener(org.eclipse.swt.events.HelpListener)
	RemoveHelpListener(a EventsHelpListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void setAccelerator(int)
	SetAccelerator(a int) 

	// public void setEnabled(boolean)
	SetEnabled(a bool) 

	// public void setID(int)
	SetID(a int) 

	// public void setMenu(org.eclipse.swt.widgets.Menu)
	SetMenu(a WidgetsMenuInterface) 

	// public void setSelection(boolean)
	SetSelection(a bool) 
}

type WidgetsMenuItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.MenuItem(org.eclipse.swt.widgets.Menu, int)
func NewWidgetsMenuItem(a WidgetsMenuInterface, b int) (*WidgetsMenuItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/MenuItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Menu"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMenuItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.MenuItem(org.eclipse.swt.widgets.Menu, int, int)
func NewWidgetsMenuItem2(a WidgetsMenuInterface, b int, c int) (*WidgetsMenuItem) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/MenuItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Menu"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsMenuItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addArmListener(org.eclipse.swt.events.ArmListener)
func (jbobject *WidgetsMenuItem) AddArmListener(a EventsArmListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addArmListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ArmListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addHelpListener(org.eclipse.swt.events.HelpListener)
func (jbobject *WidgetsMenuItem) AddHelpListener(a EventsHelpListenerInterface)  {
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

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsMenuItem) AddSelectionListener(a EventsSelectionListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addSelectionListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public int getAccelerator()
func (jbobject *WidgetsMenuItem) GetAccelerator() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAccelerator", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getEnabled()
func (jbobject *WidgetsMenuItem) GetEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getID()
func (jbobject *WidgetsMenuItem) GetID() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getID", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.Menu getMenu()
func (jbobject *WidgetsMenuItem) GetMenu() *WidgetsMenu {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMenu", "org/eclipse/swt/widgets/Menu")
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

// public org.eclipse.swt.widgets.Menu getParent()
func (jbobject *WidgetsMenuItem) GetParent() *WidgetsMenu {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Menu")
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

// public boolean getSelection()
func (jbobject *WidgetsMenuItem) GetSelection() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isEnabled()
func (jbobject *WidgetsMenuItem) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeArmListener(org.eclipse.swt.events.ArmListener)
func (jbobject *WidgetsMenuItem) RemoveArmListener(a EventsArmListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeArmListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ArmListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeHelpListener(org.eclipse.swt.events.HelpListener)
func (jbobject *WidgetsMenuItem) RemoveHelpListener(a EventsHelpListenerInterface)  {
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

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsMenuItem) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeSelectionListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SelectionListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setAccelerator(int)
func (jbobject *WidgetsMenuItem) SetAccelerator(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAccelerator", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEnabled(boolean)
func (jbobject *WidgetsMenuItem) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setID(int)
func (jbobject *WidgetsMenuItem) SetID(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setID", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsMenuItem) SetImage(a GraphicsImageInterface)  {
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

// public void setMenu(org.eclipse.swt.widgets.Menu)
func (jbobject *WidgetsMenuItem) SetMenu(a WidgetsMenuInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMenu", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Menu"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(boolean)
func (jbobject *WidgetsMenuItem) SetSelection(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *WidgetsMenuItem) SetText(a string)  {
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


