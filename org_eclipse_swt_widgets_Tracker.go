package swt

import "github.com/timob/javabind"

type WidgetsTrackerInterface interface {
	WidgetsWidgetInterface

	// public void addControlListener(org.eclipse.swt.events.ControlListener)
	AddControlListener(a EventsControlListenerInterface) 

	// public void addKeyListener(org.eclipse.swt.events.KeyListener)
	AddKeyListener(a EventsKeyListenerInterface) 

	// public void close()
	Close() 

	// public org.eclipse.swt.graphics.Rectangle[] getRectangles()
	GetRectangles() []*GraphicsRectangle

	// public boolean getStippled()
	GetStippled() bool

	// public boolean open()
	Open() bool

	// public void removeControlListener(org.eclipse.swt.events.ControlListener)
	RemoveControlListener(a EventsControlListenerInterface) 

	// public void removeKeyListener(org.eclipse.swt.events.KeyListener)
	RemoveKeyListener(a EventsKeyListenerInterface) 

	// public void setCursor(org.eclipse.swt.graphics.Cursor)
	SetCursor(a GraphicsCursorInterface) 

	// public void setRectangles(org.eclipse.swt.graphics.Rectangle[])
	SetRectangles(a []*GraphicsRectangle) 

	// public void setStippled(boolean)
	SetStippled(a bool) 
}

type WidgetsTracker struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.Tracker(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsTracker(a WidgetsCompositeInterface, b int) (*WidgetsTracker) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Tracker", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTracker{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.widgets.Tracker(org.eclipse.swt.widgets.Display, int)
func NewWidgetsTracker2(a WidgetsDisplayInterface, b int) (*WidgetsTracker) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Tracker", conv_a.Value().Cast("org/eclipse/swt/widgets/Display"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsTracker{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsTracker) AddControlListener(a EventsControlListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addControlListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ControlListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addKeyListener(org.eclipse.swt.events.KeyListener)
func (jbobject *WidgetsTracker) AddKeyListener(a EventsKeyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addKeyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/KeyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void close()
func (jbobject *WidgetsTracker) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Rectangle[] getRectangles()
func (jbobject *WidgetsTracker) GetRectangles() []*GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRectangles", javabind.ObjectArrayType("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/Rectangle")
	dst := new([]*GraphicsRectangle)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean getStippled()
func (jbobject *WidgetsTracker) GetStippled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStippled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean open()
func (jbobject *WidgetsTracker) Open() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "open", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void removeControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsTracker) RemoveControlListener(a EventsControlListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeControlListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ControlListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeKeyListener(org.eclipse.swt.events.KeyListener)
func (jbobject *WidgetsTracker) RemoveKeyListener(a EventsKeyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeKeyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/KeyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setCursor(org.eclipse.swt.graphics.Cursor)
func (jbobject *WidgetsTracker) SetCursor(a GraphicsCursorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCursor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Cursor"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setRectangles(org.eclipse.swt.graphics.Rectangle[])
func (jbobject *WidgetsTracker) SetRectangles(a []*GraphicsRectangle)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/Rectangle")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRectangles", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setStippled(boolean)
func (jbobject *WidgetsTracker) SetStippled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStippled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


