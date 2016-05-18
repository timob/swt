package swt

import "github.com/timob/javabind"

type WidgetsControlInterface interface {
	WidgetsWidgetInterface

	// public int getOrientation()
	GetOrientation() int

	// public boolean print(org.eclipse.swt.graphics.GC)
	Print(a GraphicsGCInterface) bool

	// public org.eclipse.swt.graphics.Point computeSize(int, int)
	ComputeSize(a int, b int) *GraphicsPoint

	// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
	ComputeSize2(a int, b int, c bool) *GraphicsPoint

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public void setBounds(org.eclipse.swt.graphics.Rectangle)
	SetBounds2(a GraphicsRectangleInterface) 

	// public void setBounds(int, int, int, int)
	SetBounds(a int, b int, c int, d int) 

	// public org.eclipse.swt.graphics.Point getLocation()
	GetLocation() *GraphicsPoint

	// public void setLocation(org.eclipse.swt.graphics.Point)
	SetLocation2(a GraphicsPointInterface) 

	// public void setLocation(int, int)
	SetLocation(a int, b int) 

	// public org.eclipse.swt.graphics.Point getSize()
	GetSize() *GraphicsPoint

	// public void setSize(org.eclipse.swt.graphics.Point)
	SetSize2(a GraphicsPointInterface) 

	// public void setRegion(org.eclipse.swt.graphics.Region)
	SetRegion(a GraphicsRegionInterface) 

	// public void setSize(int, int)
	SetSize(a int, b int) 

	// public void moveAbove(org.eclipse.swt.widgets.Control)
	MoveAbove(a WidgetsControlInterface) 

	// public void moveBelow(org.eclipse.swt.widgets.Control)
	MoveBelow(a WidgetsControlInterface) 

	// public void pack()
	Pack() 

	// public void pack(boolean)
	Pack2(a bool) 

	// public void setLayoutData(java.lang.Object)
	SetLayoutData(a interface{}) 

	// public org.eclipse.swt.graphics.Point toControl(int, int)
	ToControl(a int, b int) *GraphicsPoint

	// public org.eclipse.swt.graphics.Point toControl(org.eclipse.swt.graphics.Point)
	ToControl2(a GraphicsPointInterface) *GraphicsPoint

	// public org.eclipse.swt.graphics.Point toDisplay(int, int)
	ToDisplay(a int, b int) *GraphicsPoint

	// public org.eclipse.swt.graphics.Point toDisplay(org.eclipse.swt.graphics.Point)
	ToDisplay2(a GraphicsPointInterface) *GraphicsPoint

	// public void addControlListener(org.eclipse.swt.events.ControlListener)
	AddControlListener(a EventsControlListenerInterface) 

	// public void addDragDetectListener(org.eclipse.swt.events.DragDetectListener)
	AddDragDetectListener(a EventsDragDetectListenerInterface) 

	// public void addFocusListener(org.eclipse.swt.events.FocusListener)
	AddFocusListener(a EventsFocusListenerInterface) 

	// public void addGestureListener(org.eclipse.swt.events.GestureListener)
	AddGestureListener(a EventsGestureListenerInterface) 

	// public void addHelpListener(org.eclipse.swt.events.HelpListener)
	AddHelpListener(a EventsHelpListenerInterface) 

	// public void addKeyListener(org.eclipse.swt.events.KeyListener)
	AddKeyListener(a EventsKeyListenerInterface) 

	// public void addMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
	AddMenuDetectListener(a EventsMenuDetectListenerInterface) 

	// public void addMouseListener(org.eclipse.swt.events.MouseListener)
	AddMouseListener(a EventsMouseListenerInterface) 

	// public void addMouseMoveListener(org.eclipse.swt.events.MouseMoveListener)
	AddMouseMoveListener(a EventsMouseMoveListenerInterface) 

	// public void addMouseTrackListener(org.eclipse.swt.events.MouseTrackListener)
	AddMouseTrackListener(a EventsMouseTrackListenerInterface) 

	// public void addMouseWheelListener(org.eclipse.swt.events.MouseWheelListener)
	AddMouseWheelListener(a EventsMouseWheelListenerInterface) 

	// public void addPaintListener(org.eclipse.swt.events.PaintListener)
	AddPaintListener(a EventsPaintListenerInterface) 

	// public void addTouchListener(org.eclipse.swt.events.TouchListener)
	AddTouchListener(a EventsTouchListenerInterface) 

	// public void addTraverseListener(org.eclipse.swt.events.TraverseListener)
	AddTraverseListener(a EventsTraverseListenerInterface) 

	// public void removeControlListener(org.eclipse.swt.events.ControlListener)
	RemoveControlListener(a EventsControlListenerInterface) 

	// public void removeDragDetectListener(org.eclipse.swt.events.DragDetectListener)
	RemoveDragDetectListener(a EventsDragDetectListenerInterface) 

	// public void removeFocusListener(org.eclipse.swt.events.FocusListener)
	RemoveFocusListener(a EventsFocusListenerInterface) 

	// public void removeGestureListener(org.eclipse.swt.events.GestureListener)
	RemoveGestureListener(a EventsGestureListenerInterface) 

	// public void removeHelpListener(org.eclipse.swt.events.HelpListener)
	RemoveHelpListener(a EventsHelpListenerInterface) 

	// public void removeKeyListener(org.eclipse.swt.events.KeyListener)
	RemoveKeyListener(a EventsKeyListenerInterface) 

	// public void removeMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
	RemoveMenuDetectListener(a EventsMenuDetectListenerInterface) 

	// public void removeMouseListener(org.eclipse.swt.events.MouseListener)
	RemoveMouseListener(a EventsMouseListenerInterface) 

	// public void removeMouseMoveListener(org.eclipse.swt.events.MouseMoveListener)
	RemoveMouseMoveListener(a EventsMouseMoveListenerInterface) 

	// public void removeMouseTrackListener(org.eclipse.swt.events.MouseTrackListener)
	RemoveMouseTrackListener(a EventsMouseTrackListenerInterface) 

	// public void removeMouseWheelListener(org.eclipse.swt.events.MouseWheelListener)
	RemoveMouseWheelListener(a EventsMouseWheelListenerInterface) 

	// public void removePaintListener(org.eclipse.swt.events.PaintListener)
	RemovePaintListener(a EventsPaintListenerInterface) 

	// public void removeTouchListener(org.eclipse.swt.events.TouchListener)
	RemoveTouchListener(a EventsTouchListenerInterface) 

	// public void removeTraverseListener(org.eclipse.swt.events.TraverseListener)
	RemoveTraverseListener(a EventsTraverseListenerInterface) 

	// public boolean dragDetect(org.eclipse.swt.widgets.Event)
	DragDetect2(a WidgetsEventInterface) bool

	// public boolean dragDetect(org.eclipse.swt.events.MouseEvent)
	DragDetect(a EventsMouseEventInterface) bool

	// public boolean forceFocus()
	ForceFocus() bool

	// public org.eclipse.swt.graphics.Color getBackground()
	GetBackground() *GraphicsColor

	// public org.eclipse.swt.graphics.Image getBackgroundImage()
	GetBackgroundImage() *GraphicsImage

	// public int getBorderWidth()
	GetBorderWidth() int

	// public org.eclipse.swt.graphics.Cursor getCursor()
	GetCursor() *GraphicsCursor

	// public boolean getDragDetect()
	GetDragDetect() bool

	// public boolean getEnabled()
	GetEnabled() bool

	// public org.eclipse.swt.graphics.Font getFont()
	GetFont() *GraphicsFont

	// public org.eclipse.swt.graphics.Color getForeground()
	GetForeground() *GraphicsColor

	// public java.lang.Object getLayoutData()
	GetLayoutData() *JavaLangObject

	// public org.eclipse.swt.widgets.Menu getMenu()
	GetMenu() *WidgetsMenu

	// public org.eclipse.swt.widgets.Monitor getMonitor()
	GetMonitor() *WidgetsMonitor

	// public org.eclipse.swt.widgets.Composite getParent()
	GetParent() *WidgetsComposite

	// public org.eclipse.swt.graphics.Region getRegion()
	GetRegion() *GraphicsRegion

	// public org.eclipse.swt.widgets.Shell getShell()
	GetShell() *WidgetsShell

	// public java.lang.String getToolTipText()
	GetToolTipText() string

	// public boolean getTouchEnabled()
	GetTouchEnabled() bool

	// public boolean getVisible()
	GetVisible() bool

	// public long internal_new_GC(org.eclipse.swt.graphics.GCData)
	Internal_new_GC(a GraphicsGCDataInterface) int64

	// public void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
	Internal_dispose_GC(a int64, b GraphicsGCDataInterface) 

	// public boolean isReparentable()
	IsReparentable() bool

	// public boolean isEnabled()
	IsEnabled() bool

	// public boolean isFocusControl()
	IsFocusControl() bool

	// public boolean isVisible()
	IsVisible() bool

	// public void redraw()
	Redraw() 

	// public void redraw(int, int, int, int, boolean)
	Redraw2(a int, b int, c int, d int, e bool) 

	// public void setBackground(org.eclipse.swt.graphics.Color)
	SetBackground(a GraphicsColorInterface) 

	// public void setBackgroundImage(org.eclipse.swt.graphics.Image)
	SetBackgroundImage(a GraphicsImageInterface) 

	// public void setCapture(boolean)
	SetCapture(a bool) 

	// public void setCursor(org.eclipse.swt.graphics.Cursor)
	SetCursor(a GraphicsCursorInterface) 

	// public void setDragDetect(boolean)
	SetDragDetect(a bool) 

	// public void setEnabled(boolean)
	SetEnabled(a bool) 

	// public boolean setFocus()
	SetFocus() bool

	// public void setFont(org.eclipse.swt.graphics.Font)
	SetFont(a GraphicsFontInterface) 

	// public void setForeground(org.eclipse.swt.graphics.Color)
	SetForeground(a GraphicsColorInterface) 

	// public void setMenu(org.eclipse.swt.widgets.Menu)
	SetMenu(a WidgetsMenuInterface) 

	// public void setOrientation(int)
	SetOrientation(a int) 

	// public boolean setParent(org.eclipse.swt.widgets.Composite)
	SetParent(a WidgetsCompositeInterface) bool

	// public void setRedraw(boolean)
	SetRedraw(a bool) 

	// public void setToolTipText(java.lang.String)
	SetToolTipText(a string) 

	// public void setTouchEnabled(boolean)
	SetTouchEnabled(a bool) 

	// public void setVisible(boolean)
	SetVisible(a bool) 

	// public boolean traverse(int)
	Traverse(a int) bool

	// public boolean traverse(int, org.eclipse.swt.widgets.Event)
	Traverse3(a int, b WidgetsEventInterface) bool

	// public boolean traverse(int, org.eclipse.swt.events.KeyEvent)
	Traverse2(a int, b EventsKeyEventInterface) bool

	// public void update()
	Update() 
}

type WidgetsControl struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.Control(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsControl(a WidgetsCompositeInterface, b int) (*WidgetsControl) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Control", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsControl{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getOrientation()
func (jbobject *WidgetsControl) GetOrientation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrientation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean print(org.eclipse.swt.graphics.GC)
func (jbobject *WidgetsControl) Print(a GraphicsGCInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/graphics/GC"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Point computeSize(int, int)
func (jbobject *WidgetsControl) ComputeSize(a int, b int) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeSize", "org/eclipse/swt/graphics/Point", a, b)
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

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsControl) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsControl) GetBounds() *GraphicsRectangle {
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

// public void setBounds(org.eclipse.swt.graphics.Rectangle)
func (jbobject *WidgetsControl) SetBounds2(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBounds", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBounds(int, int, int, int)
func (jbobject *WidgetsControl) SetBounds(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBounds", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point getLocation()
func (jbobject *WidgetsControl) GetLocation() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocation", "org/eclipse/swt/graphics/Point")
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

// public void setLocation(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsControl) SetLocation2(a GraphicsPointInterface)  {
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

// public void setLocation(int, int)
func (jbobject *WidgetsControl) SetLocation(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocation", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point getSize()
func (jbobject *WidgetsControl) GetSize() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", "org/eclipse/swt/graphics/Point")
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

// public void setSize(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsControl) SetSize2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setRegion(org.eclipse.swt.graphics.Region)
func (jbobject *WidgetsControl) SetRegion(a GraphicsRegionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRegion", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Region"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSize(int, int)
func (jbobject *WidgetsControl) SetSize(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void moveAbove(org.eclipse.swt.widgets.Control)
func (jbobject *WidgetsControl) MoveAbove(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "moveAbove", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void moveBelow(org.eclipse.swt.widgets.Control)
func (jbobject *WidgetsControl) MoveBelow(a WidgetsControlInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "moveBelow", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Control"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void pack()
func (jbobject *WidgetsControl) Pack()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "pack", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void pack(boolean)
func (jbobject *WidgetsControl) Pack2(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "pack", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLayoutData(java.lang.Object)
func (jbobject *WidgetsControl) SetLayoutData(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLayoutData", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.graphics.Point toControl(int, int)
func (jbobject *WidgetsControl) ToControl(a int, b int) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toControl", "org/eclipse/swt/graphics/Point", a, b)
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

// public org.eclipse.swt.graphics.Point toControl(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsControl) ToControl2(a GraphicsPointInterface) *GraphicsPoint {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toControl", "org/eclipse/swt/graphics/Point", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
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

// public org.eclipse.swt.graphics.Point toDisplay(int, int)
func (jbobject *WidgetsControl) ToDisplay(a int, b int) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toDisplay", "org/eclipse/swt/graphics/Point", a, b)
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

// public org.eclipse.swt.graphics.Point toDisplay(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsControl) ToDisplay2(a GraphicsPointInterface) *GraphicsPoint {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toDisplay", "org/eclipse/swt/graphics/Point", conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
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

// public void addControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsControl) AddControlListener(a EventsControlListenerInterface)  {
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

// public void addDragDetectListener(org.eclipse.swt.events.DragDetectListener)
func (jbobject *WidgetsControl) AddDragDetectListener(a EventsDragDetectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addDragDetectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/DragDetectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addFocusListener(org.eclipse.swt.events.FocusListener)
func (jbobject *WidgetsControl) AddFocusListener(a EventsFocusListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addFocusListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/FocusListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addGestureListener(org.eclipse.swt.events.GestureListener)
func (jbobject *WidgetsControl) AddGestureListener(a EventsGestureListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addGestureListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/GestureListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addHelpListener(org.eclipse.swt.events.HelpListener)
func (jbobject *WidgetsControl) AddHelpListener(a EventsHelpListenerInterface)  {
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

// public void addKeyListener(org.eclipse.swt.events.KeyListener)
func (jbobject *WidgetsControl) AddKeyListener(a EventsKeyListenerInterface)  {
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

// public void addMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
func (jbobject *WidgetsControl) AddMenuDetectListener(a EventsMenuDetectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addMenuDetectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuDetectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addMouseListener(org.eclipse.swt.events.MouseListener)
func (jbobject *WidgetsControl) AddMouseListener(a EventsMouseListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addMouseListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addMouseMoveListener(org.eclipse.swt.events.MouseMoveListener)
func (jbobject *WidgetsControl) AddMouseMoveListener(a EventsMouseMoveListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addMouseMoveListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseMoveListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addMouseTrackListener(org.eclipse.swt.events.MouseTrackListener)
func (jbobject *WidgetsControl) AddMouseTrackListener(a EventsMouseTrackListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addMouseTrackListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseTrackListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addMouseWheelListener(org.eclipse.swt.events.MouseWheelListener)
func (jbobject *WidgetsControl) AddMouseWheelListener(a EventsMouseWheelListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addMouseWheelListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseWheelListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addPaintListener(org.eclipse.swt.events.PaintListener)
func (jbobject *WidgetsControl) AddPaintListener(a EventsPaintListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addPaintListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/PaintListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addTouchListener(org.eclipse.swt.events.TouchListener)
func (jbobject *WidgetsControl) AddTouchListener(a EventsTouchListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addTouchListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TouchListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addTraverseListener(org.eclipse.swt.events.TraverseListener)
func (jbobject *WidgetsControl) AddTraverseListener(a EventsTraverseListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addTraverseListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TraverseListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeControlListener(org.eclipse.swt.events.ControlListener)
func (jbobject *WidgetsControl) RemoveControlListener(a EventsControlListenerInterface)  {
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

// public void removeDragDetectListener(org.eclipse.swt.events.DragDetectListener)
func (jbobject *WidgetsControl) RemoveDragDetectListener(a EventsDragDetectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeDragDetectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/DragDetectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeFocusListener(org.eclipse.swt.events.FocusListener)
func (jbobject *WidgetsControl) RemoveFocusListener(a EventsFocusListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeFocusListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/FocusListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeGestureListener(org.eclipse.swt.events.GestureListener)
func (jbobject *WidgetsControl) RemoveGestureListener(a EventsGestureListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeGestureListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/GestureListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeHelpListener(org.eclipse.swt.events.HelpListener)
func (jbobject *WidgetsControl) RemoveHelpListener(a EventsHelpListenerInterface)  {
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

// public void removeKeyListener(org.eclipse.swt.events.KeyListener)
func (jbobject *WidgetsControl) RemoveKeyListener(a EventsKeyListenerInterface)  {
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

// public void removeMenuDetectListener(org.eclipse.swt.events.MenuDetectListener)
func (jbobject *WidgetsControl) RemoveMenuDetectListener(a EventsMenuDetectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeMenuDetectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MenuDetectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeMouseListener(org.eclipse.swt.events.MouseListener)
func (jbobject *WidgetsControl) RemoveMouseListener(a EventsMouseListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeMouseListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeMouseMoveListener(org.eclipse.swt.events.MouseMoveListener)
func (jbobject *WidgetsControl) RemoveMouseMoveListener(a EventsMouseMoveListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeMouseMoveListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseMoveListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeMouseTrackListener(org.eclipse.swt.events.MouseTrackListener)
func (jbobject *WidgetsControl) RemoveMouseTrackListener(a EventsMouseTrackListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeMouseTrackListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseTrackListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeMouseWheelListener(org.eclipse.swt.events.MouseWheelListener)
func (jbobject *WidgetsControl) RemoveMouseWheelListener(a EventsMouseWheelListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeMouseWheelListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/MouseWheelListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removePaintListener(org.eclipse.swt.events.PaintListener)
func (jbobject *WidgetsControl) RemovePaintListener(a EventsPaintListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removePaintListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/PaintListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeTouchListener(org.eclipse.swt.events.TouchListener)
func (jbobject *WidgetsControl) RemoveTouchListener(a EventsTouchListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeTouchListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TouchListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeTraverseListener(org.eclipse.swt.events.TraverseListener)
func (jbobject *WidgetsControl) RemoveTraverseListener(a EventsTraverseListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeTraverseListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/TraverseListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean dragDetect(org.eclipse.swt.widgets.Event)
func (jbobject *WidgetsControl) DragDetect2(a WidgetsEventInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "dragDetect", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean dragDetect(org.eclipse.swt.events.MouseEvent)
func (jbobject *WidgetsControl) DragDetect(a EventsMouseEventInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "dragDetect", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/events/MouseEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean forceFocus()
func (jbobject *WidgetsControl) ForceFocus() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "forceFocus", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *WidgetsControl) GetBackground() *GraphicsColor {
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

// public org.eclipse.swt.graphics.Image getBackgroundImage()
func (jbobject *WidgetsControl) GetBackgroundImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBackgroundImage", "org/eclipse/swt/graphics/Image")
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

// public int getBorderWidth()
func (jbobject *WidgetsControl) GetBorderWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBorderWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Cursor getCursor()
func (jbobject *WidgetsControl) GetCursor() *GraphicsCursor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCursor", "org/eclipse/swt/graphics/Cursor")
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
	unique_x := &GraphicsCursor{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getDragDetect()
func (jbobject *WidgetsControl) GetDragDetect() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDragDetect", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getEnabled()
func (jbobject *WidgetsControl) GetEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Font getFont()
func (jbobject *WidgetsControl) GetFont() *GraphicsFont {
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
func (jbobject *WidgetsControl) GetForeground() *GraphicsColor {
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

// public java.lang.Object getLayoutData()
func (jbobject *WidgetsControl) GetLayoutData() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLayoutData", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Menu getMenu()
func (jbobject *WidgetsControl) GetMenu() *WidgetsMenu {
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

// public org.eclipse.swt.widgets.Monitor getMonitor()
func (jbobject *WidgetsControl) GetMonitor() *WidgetsMonitor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitor", "org/eclipse/swt/widgets/Monitor")
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
	unique_x := &WidgetsMonitor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Composite getParent()
func (jbobject *WidgetsControl) GetParent() *WidgetsComposite {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/Composite")
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
	unique_x := &WidgetsComposite{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Region getRegion()
func (jbobject *WidgetsControl) GetRegion() *GraphicsRegion {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRegion", "org/eclipse/swt/graphics/Region")
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
	unique_x := &GraphicsRegion{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Shell getShell()
func (jbobject *WidgetsControl) GetShell() *WidgetsShell {
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

// public java.lang.String getToolTipText()
func (jbobject *WidgetsControl) GetToolTipText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getToolTipText", "java/lang/String")
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

// public boolean getTouchEnabled()
func (jbobject *WidgetsControl) GetTouchEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTouchEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getVisible()
func (jbobject *WidgetsControl) GetVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long internal_new_GC(org.eclipse.swt.graphics.GCData)
func (jbobject *WidgetsControl) Internal_new_GC(a GraphicsGCDataInterface) int64 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "internal_new_GC", javabind.Long, conv_a.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
func (jbobject *WidgetsControl) Internal_dispose_GC(a int64, b GraphicsGCDataInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "internal_dispose_GC", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public boolean isReparentable()
func (jbobject *WidgetsControl) IsReparentable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isReparentable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isEnabled()
func (jbobject *WidgetsControl) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isFocusControl()
func (jbobject *WidgetsControl) IsFocusControl() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isFocusControl", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isVisible()
func (jbobject *WidgetsControl) IsVisible() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isVisible", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void redraw()
func (jbobject *WidgetsControl) Redraw()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "redraw", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void redraw(int, int, int, int, boolean)
func (jbobject *WidgetsControl) Redraw2(a int, b int, c int, d int, e bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "redraw", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *WidgetsControl) SetBackground(a GraphicsColorInterface)  {
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

// public void setBackgroundImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsControl) SetBackgroundImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackgroundImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setCapture(boolean)
func (jbobject *WidgetsControl) SetCapture(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCapture", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setCursor(org.eclipse.swt.graphics.Cursor)
func (jbobject *WidgetsControl) SetCursor(a GraphicsCursorInterface)  {
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

// public void setDragDetect(boolean)
func (jbobject *WidgetsControl) SetDragDetect(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDragDetect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEnabled(boolean)
func (jbobject *WidgetsControl) SetEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean setFocus()
func (jbobject *WidgetsControl) SetFocus() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setFocus", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *WidgetsControl) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *WidgetsControl) SetForeground(a GraphicsColorInterface)  {
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

// public void setMenu(org.eclipse.swt.widgets.Menu)
func (jbobject *WidgetsControl) SetMenu(a WidgetsMenuInterface)  {
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

// public void setOrientation(int)
func (jbobject *WidgetsControl) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean setParent(org.eclipse.swt.widgets.Composite)
func (jbobject *WidgetsControl) SetParent(a WidgetsCompositeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setParent", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void setRedraw(boolean)
func (jbobject *WidgetsControl) SetRedraw(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRedraw", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setToolTipText(java.lang.String)
func (jbobject *WidgetsControl) SetToolTipText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setToolTipText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTouchEnabled(boolean)
func (jbobject *WidgetsControl) SetTouchEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTouchEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setVisible(boolean)
func (jbobject *WidgetsControl) SetVisible(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVisible", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean traverse(int)
func (jbobject *WidgetsControl) Traverse(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "traverse", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean traverse(int, org.eclipse.swt.widgets.Event)
func (jbobject *WidgetsControl) Traverse3(a int, b WidgetsEventInterface) bool {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "traverse", javabind.Boolean, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	return jret.(bool)
}

// public boolean traverse(int, org.eclipse.swt.events.KeyEvent)
func (jbobject *WidgetsControl) Traverse2(a int, b EventsKeyEventInterface) bool {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "traverse", javabind.Boolean, a, conv_b.Value().Cast("org/eclipse/swt/events/KeyEvent"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	return jret.(bool)
}

// public void update()
func (jbobject *WidgetsControl) Update()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "update", javabind.Void)
	if err != nil {
		panic(err)
	}

}


