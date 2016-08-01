package swt

import "github.com/timob/javabind"

type WidgetsTextInterface interface {
	WidgetsScrollableInterface

	// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
	AddModifyListener(a EventsModifyListenerInterface) 

	// public void addSegmentListener(org.eclipse.swt.events.SegmentListener)
	AddSegmentListener(a EventsSegmentListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void addVerifyListener(org.eclipse.swt.events.VerifyListener)
	AddVerifyListener(a EventsVerifyListenerInterface) 

	// public void append(java.lang.String)
	Append(a string) 

	// public void clearSelection()
	ClearSelection() 

	// public void copy()
	Copy() 

	// public void cut()
	Cut() 

	// public int getCaretLineNumber()
	GetCaretLineNumber() int

	// public org.eclipse.swt.graphics.Point getCaretLocation()
	GetCaretLocation() *GraphicsPoint

	// public int getCaretPosition()
	GetCaretPosition() int

	// public int getCharCount()
	GetCharCount() int

	// public boolean getDoubleClickEnabled()
	GetDoubleClickEnabled() bool

	// public char getEchoChar()
	GetEchoChar() uint16

	// public boolean getEditable()
	GetEditable() bool

	// public int getLineCount()
	GetLineCount() int

	// public java.lang.String getLineDelimiter()
	GetLineDelimiter() string

	// public int getLineHeight()
	GetLineHeight() int

	// public java.lang.String getMessage()
	GetMessage() string

	// public org.eclipse.swt.graphics.Point getSelection()
	GetSelection() *GraphicsPoint

	// public int getSelectionCount()
	GetSelectionCount() int

	// public java.lang.String getSelectionText()
	GetSelectionText() string

	// public int getTabs()
	GetTabs() int

	// public java.lang.String getText()
	GetText() string

	// public java.lang.String getText(int, int)
	GetText2(a int, b int) string

	// public char[] getTextChars()
	GetTextChars() []uint16

	// public int getTextLimit()
	GetTextLimit() int

	// public int getTopIndex()
	GetTopIndex() int

	// public int getTopPixel()
	GetTopPixel() int

	// public void insert(java.lang.String)
	Insert(a string) 

	// public void paste()
	Paste() 

	// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
	RemoveModifyListener(a EventsModifyListenerInterface) 

	// public void removeSegmentListener(org.eclipse.swt.events.SegmentListener)
	RemoveSegmentListener(a EventsSegmentListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void removeVerifyListener(org.eclipse.swt.events.VerifyListener)
	RemoveVerifyListener(a EventsVerifyListenerInterface) 

	// public void selectAll()
	SelectAll() 

	// public void setDoubleClickEnabled(boolean)
	SetDoubleClickEnabled(a bool) 

	// public void setEchoChar(char)
	SetEchoChar(a uint16) 

	// public void setEditable(boolean)
	SetEditable(a bool) 

	// public void setMessage(java.lang.String)
	SetMessage(a string) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setSelection(int, int)
	SetSelection2(a int, b int) 

	// public void setSelection(org.eclipse.swt.graphics.Point)
	SetSelection3(a GraphicsPointInterface) 

	// public void setTabs(int)
	SetTabs(a int) 

	// public void setText(java.lang.String)
	SetText(a string) 

	// public void setTextChars(char[])
	SetTextChars(a []uint16) 

	// public void setTextLimit(int)
	SetTextLimit(a int) 

	// public void setTopIndex(int)
	SetTopIndex(a int) 

	// public void showSelection()
	ShowSelection() 
}

type WidgetsText struct {
	WidgetsScrollable
}

// public org.eclipse.swt.widgets.Text(org.eclipse.swt.widgets.Composite, int)
func NewWidgetsText(a WidgetsCompositeInterface, b int) (*WidgetsText) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Text", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsText{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *WidgetsText) AddModifyListener(a EventsModifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addModifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ModifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSegmentListener(org.eclipse.swt.events.SegmentListener)
func (jbobject *WidgetsText) AddSegmentListener(a EventsSegmentListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addSegmentListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SegmentListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsText) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void addVerifyListener(org.eclipse.swt.events.VerifyListener)
func (jbobject *WidgetsText) AddVerifyListener(a EventsVerifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addVerifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/VerifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void append(java.lang.String)
func (jbobject *WidgetsText) Append(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "append", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void clearSelection()
func (jbobject *WidgetsText) ClearSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *WidgetsText) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public org.eclipse.swt.graphics.Rectangle computeTrim(int, int, int, int)
func (jbobject *WidgetsText) ComputeTrim(a int, b int, c int, d int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "computeTrim", "org/eclipse/swt/graphics/Rectangle", a, b, c, d)
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

// public void copy()
func (jbobject *WidgetsText) Copy()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copy", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void cut()
func (jbobject *WidgetsText) Cut()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cut", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public int getBorderWidth()
func (jbobject *WidgetsText) GetBorderWidth() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBorderWidth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getCaretLineNumber()
func (jbobject *WidgetsText) GetCaretLineNumber() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaretLineNumber", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point getCaretLocation()
func (jbobject *WidgetsText) GetCaretLocation() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaretLocation", "org/eclipse/swt/graphics/Point")
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

// public int getCaretPosition()
func (jbobject *WidgetsText) GetCaretPosition() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaretPosition", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getCharCount()
func (jbobject *WidgetsText) GetCharCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCharCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getDoubleClickEnabled()
func (jbobject *WidgetsText) GetDoubleClickEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDoubleClickEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public char getEchoChar()
func (jbobject *WidgetsText) GetEchoChar() uint16 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEchoChar", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

// public boolean getEditable()
func (jbobject *WidgetsText) GetEditable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEditable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getLineCount()
func (jbobject *WidgetsText) GetLineCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getLineDelimiter()
func (jbobject *WidgetsText) GetLineDelimiter() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineDelimiter", "java/lang/String")
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

// public int getLineHeight()
func (jbobject *WidgetsText) GetLineHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getMessage()
func (jbobject *WidgetsText) GetMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessage", "java/lang/String")
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

// public int getOrientation()
func (jbobject *WidgetsText) GetOrientation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrientation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point getSelection()
func (jbobject *WidgetsText) GetSelection() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelection", "org/eclipse/swt/graphics/Point")
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

// public int getSelectionCount()
func (jbobject *WidgetsText) GetSelectionCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getSelectionText()
func (jbobject *WidgetsText) GetSelectionText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionText", "java/lang/String")
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

// public int getTabs()
func (jbobject *WidgetsText) GetTabs() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTabs", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getText()
func (jbobject *WidgetsText) GetText() string {
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

// public java.lang.String getText(int, int)
func (jbobject *WidgetsText) GetText2(a int, b int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getText", "java/lang/String", a, b)
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

// public char[] getTextChars()
func (jbobject *WidgetsText) GetTextChars() []uint16 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextChars", javabind.Char | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]uint16)
}

// public int getTextLimit()
func (jbobject *WidgetsText) GetTextLimit() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextLimit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTopIndex()
func (jbobject *WidgetsText) GetTopIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTopPixel()
func (jbobject *WidgetsText) GetTopPixel() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopPixel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void insert(java.lang.String)
func (jbobject *WidgetsText) Insert(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "insert", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void paste()
func (jbobject *WidgetsText) Paste()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "paste", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *WidgetsText) RemoveModifyListener(a EventsModifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeModifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ModifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSegmentListener(org.eclipse.swt.events.SegmentListener)
func (jbobject *WidgetsText) RemoveSegmentListener(a EventsSegmentListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeSegmentListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/SegmentListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *WidgetsText) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void removeVerifyListener(org.eclipse.swt.events.VerifyListener)
func (jbobject *WidgetsText) RemoveVerifyListener(a EventsVerifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeVerifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/VerifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void selectAll()
func (jbobject *WidgetsText) SelectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "selectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setDoubleClickEnabled(boolean)
func (jbobject *WidgetsText) SetDoubleClickEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDoubleClickEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEchoChar(char)
func (jbobject *WidgetsText) SetEchoChar(a uint16)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEchoChar", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEditable(boolean)
func (jbobject *WidgetsText) SetEditable(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditable", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setMessage(java.lang.String)
func (jbobject *WidgetsText) SetMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setOrientation(int)
func (jbobject *WidgetsText) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *WidgetsText) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int, int)
func (jbobject *WidgetsText) SetSelection2(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.graphics.Point)
func (jbobject *WidgetsText) SetSelection3(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTabs(int)
func (jbobject *WidgetsText) SetTabs(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTabs", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *WidgetsText) SetText(a string)  {
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

// public void setTextChars(char[])
func (jbobject *WidgetsText) SetTextChars(a []uint16)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTextChars", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTextLimit(int)
func (jbobject *WidgetsText) SetTextLimit(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTextLimit", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTopIndex(int)
func (jbobject *WidgetsText) SetTopIndex(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopIndex", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void showSelection()
func (jbobject *WidgetsText) ShowSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsText) LIMIT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/widgets/Text", "LIMIT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsText) SetFieldLIMIT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/widgets/Text", "LIMIT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsText) DELIMITER() string {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/widgets/Text", "DELIMITER", "java/lang/String")
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

func (jbobject *WidgetsText) SetFieldDELIMITER(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/widgets/Text", "DELIMITER", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


