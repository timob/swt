package swt

import "github.com/timob/javabind"

type WidgetsExpandBarInterface interface {
	WidgetsCompositeInterface

	// public void addExpandListener(org.eclipse.swt.events.ExpandListener)
	AddExpandListener(a EventsExpandListenerInterface) 

	// public org.eclipse.swt.widgets.ExpandItem getItem(int)
	GetItem(a int) *WidgetsExpandItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.ExpandItem[] getItems()
	GetItems() []*WidgetsExpandItem

	// public int getSpacing()
	GetSpacing() int

	// public int indexOf(org.eclipse.swt.widgets.ExpandItem)
	IndexOf(a WidgetsExpandItemInterface) int

	// public void removeExpandListener(org.eclipse.swt.events.ExpandListener)
	RemoveExpandListener(a EventsExpandListenerInterface) 

	// public void setSpacing(int)
	SetSpacing(a int) 
}

type WidgetsExpandBar struct {
	WidgetsComposite
}

// public org.eclipse.swt.widgets.ExpandBar(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsExpandBar(a WidgetsCompositeInterface, b int) (*WidgetsExpandBar) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ExpandBar", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsExpandBar{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addExpandListener(org.eclipse.swt.events.ExpandListener)
func (jbobject *WidgetsExpandBar) AddExpandListener(a EventsExpandListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addExpandListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ExpandListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsExpandBar) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeSize", "org/eclipse/swt/graphics/Point", a, b, c)
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
	unique_x := &GraphicsPoint{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.ExpandItem getItem(int)
func (jbobject *WidgetsExpandBar) GetItem(a int) *WidgetsExpandItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/ExpandItem", a)
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
	unique_x := &WidgetsExpandItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsExpandBar) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.ExpandItem[] getItems()
func (jbobject *WidgetsExpandBar) GetItems() []*WidgetsExpandItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/ExpandItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/ExpandItem")
	dst := new([]*WidgetsExpandItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getSpacing()
func (jbobject *WidgetsExpandBar) GetSpacing() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(org.eclipse.swt.widgets.ExpandItem)
func (jbobject *WidgetsExpandBar) IndexOf(a WidgetsExpandItemInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/widgets/ExpandItem"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public void removeExpandListener(org.eclipse.swt.events.ExpandListener)
func (jbobject *WidgetsExpandBar) RemoveExpandListener(a EventsExpandListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeExpandListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ExpandListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSpacing(int)
func (jbobject *WidgetsExpandBar) SetSpacing(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpacing", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


