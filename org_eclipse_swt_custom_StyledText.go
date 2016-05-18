package swt

import "github.com/timob/javabind"

type CustomStyledTextInterface interface {
	WidgetsCanvasInterface

	// public void addExtendedModifyListener(org.eclipse.swt.custom.ExtendedModifyListener)
	AddExtendedModifyListener(a CustomExtendedModifyListenerInterface) 

	// public void addBidiSegmentListener(org.eclipse.swt.custom.BidiSegmentListener)
	AddBidiSegmentListener(a CustomBidiSegmentListenerInterface) 

	// public void addCaretListener(org.eclipse.swt.custom.CaretListener)
	AddCaretListener(a CustomCaretListenerInterface) 

	// public void addLineBackgroundListener(org.eclipse.swt.custom.LineBackgroundListener)
	AddLineBackgroundListener(a CustomLineBackgroundListenerInterface) 

	// public void addLineStyleListener(org.eclipse.swt.custom.LineStyleListener)
	AddLineStyleListener(a CustomLineStyleListenerInterface) 

	// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
	AddModifyListener(a EventsModifyListenerInterface) 

	// public void addPaintObjectListener(org.eclipse.swt.custom.PaintObjectListener)
	AddPaintObjectListener(a CustomPaintObjectListenerInterface) 

	// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
	AddSelectionListener(a EventsSelectionListenerInterface) 

	// public void addVerifyKeyListener(org.eclipse.swt.custom.VerifyKeyListener)
	AddVerifyKeyListener(a CustomVerifyKeyListenerInterface) 

	// public void addVerifyListener(org.eclipse.swt.events.VerifyListener)
	AddVerifyListener(a EventsVerifyListenerInterface) 

	// public void addWordMovementListener(org.eclipse.swt.custom.MovementListener)
	AddWordMovementListener(a CustomMovementListenerInterface) 

	// public void append(java.lang.String)
	Append(a string) 

	// public void copy()
	Copy() 

	// public void copy(int)
	Copy2(a int) 

	// public int getAlignment()
	GetAlignment() int

	// public boolean getAlwaysShowScrollBars()
	GetAlwaysShowScrollBars() bool

	// public org.eclipse.swt.graphics.Color getMarginColor()
	GetMarginColor() *GraphicsColor

	// public void cut()
	Cut() 

	// public int getBaseline()
	GetBaseline() int

	// public int getBaseline(int)
	GetBaseline2(a int) int

	// public boolean getBidiColoring()
	GetBidiColoring() bool

	// public boolean getBlockSelection()
	GetBlockSelection() bool

	// public org.eclipse.swt.graphics.Rectangle getBlockSelectionBounds()
	GetBlockSelectionBounds() *GraphicsRectangle

	// public int getBottomMargin()
	GetBottomMargin() int

	// public int getCaretOffset()
	GetCaretOffset() int

	// public org.eclipse.swt.custom.StyledTextContent getContent()
	GetContent() *CustomStyledTextContent

	// public boolean getDoubleClickEnabled()
	GetDoubleClickEnabled() bool

	// public boolean getEditable()
	GetEditable() bool

	// public int getHorizontalIndex()
	GetHorizontalIndex() int

	// public int getHorizontalPixel()
	GetHorizontalPixel() int

	// public int getIndent()
	GetIndent() int

	// public boolean getJustify()
	GetJustify() bool

	// public int getKeyBinding(int)
	GetKeyBinding(a int) int

	// public int getCharCount()
	GetCharCount() int

	// public java.lang.String getLine(int)
	GetLine(a int) string

	// public int getLineAlignment(int)
	GetLineAlignment(a int) int

	// public int getLineAtOffset(int)
	GetLineAtOffset(a int) int

	// public org.eclipse.swt.graphics.Color getLineBackground(int)
	GetLineBackground(a int) *GraphicsColor

	// public org.eclipse.swt.custom.Bullet getLineBullet(int)
	GetLineBullet(a int) *CustomBullet

	// public int getLineCount()
	GetLineCount() int

	// public java.lang.String getLineDelimiter()
	GetLineDelimiter() string

	// public int getLineHeight()
	GetLineHeight() int

	// public int getLineHeight(int)
	GetLineHeight2(a int) int

	// public int getLineIndent(int)
	GetLineIndent(a int) int

	// public boolean getLineJustify(int)
	GetLineJustify(a int) bool

	// public int getLineSpacing()
	GetLineSpacing() int

	// public int getLinePixel(int)
	GetLinePixel(a int) int

	// public int getLineIndex(int)
	GetLineIndex(a int) int

	// public int[] getLineTabStops(int)
	GetLineTabStops(a int) []int

	// public int getLineWrapIndent(int)
	GetLineWrapIndent(a int) int

	// public int getLeftMargin()
	GetLeftMargin() int

	// public org.eclipse.swt.graphics.Point getLocationAtOffset(int)
	GetLocationAtOffset(a int) *GraphicsPoint

	// public int getOffsetAtLine(int)
	GetOffsetAtLine(a int) int

	// public int getOffsetAtLocation(org.eclipse.swt.graphics.Point)
	GetOffsetAtLocation(a GraphicsPointInterface) int

	// public int[] getRanges()
	GetRanges() []int

	// public int[] getRanges(int, int)
	GetRanges2(a int, b int) []int

	// public int getRightMargin()
	GetRightMargin() int

	// public org.eclipse.swt.graphics.Point getSelection()
	GetSelection() *GraphicsPoint

	// public org.eclipse.swt.graphics.Point getSelectionRange()
	GetSelectionRange() *GraphicsPoint

	// public int[] getSelectionRanges()
	GetSelectionRanges() []int

	// public org.eclipse.swt.graphics.Color getSelectionBackground()
	GetSelectionBackground() *GraphicsColor

	// public int getSelectionCount()
	GetSelectionCount() int

	// public org.eclipse.swt.graphics.Color getSelectionForeground()
	GetSelectionForeground() *GraphicsColor

	// public java.lang.String getSelectionText()
	GetSelectionText() string

	// public org.eclipse.swt.custom.StyleRange getStyleRangeAtOffset(int)
	GetStyleRangeAtOffset(a int) *CustomStyleRange

	// public org.eclipse.swt.custom.StyleRange[] getStyleRanges()
	GetStyleRanges() []*CustomStyleRange

	// public org.eclipse.swt.custom.StyleRange[] getStyleRanges(boolean)
	GetStyleRanges2(a bool) []*CustomStyleRange

	// public org.eclipse.swt.custom.StyleRange[] getStyleRanges(int, int)
	GetStyleRanges3(a int, b int) []*CustomStyleRange

	// public org.eclipse.swt.custom.StyleRange[] getStyleRanges(int, int, boolean)
	GetStyleRanges4(a int, b int, c bool) []*CustomStyleRange

	// public int getTabs()
	GetTabs() int

	// public int[] getTabStops()
	GetTabStops() []int

	// public java.lang.String getText()
	GetText() string

	// public java.lang.String getText(int, int)
	GetText2(a int, b int) string

	// public org.eclipse.swt.graphics.Rectangle getTextBounds(int, int)
	GetTextBounds(a int, b int) *GraphicsRectangle

	// public java.lang.String getTextRange(int, int)
	GetTextRange(a int, b int) string

	// public int getTextLimit()
	GetTextLimit() int

	// public int getTopIndex()
	GetTopIndex() int

	// public int getTopMargin()
	GetTopMargin() int

	// public int getTopPixel()
	GetTopPixel() int

	// public boolean getWordWrap()
	GetWordWrap() bool

	// public int getWrapIndent()
	GetWrapIndent() int

	// public void insert(java.lang.String)
	Insert(a string) 

	// public void invokeAction(int)
	InvokeAction(a int) 

	// public void paste()
	Paste() 

	// public void print()
	Print2() 

	// public void redrawRange(int, int, boolean)
	RedrawRange(a int, b int, c bool) 

	// public void removeBidiSegmentListener(org.eclipse.swt.custom.BidiSegmentListener)
	RemoveBidiSegmentListener(a CustomBidiSegmentListenerInterface) 

	// public void removeCaretListener(org.eclipse.swt.custom.CaretListener)
	RemoveCaretListener(a CustomCaretListenerInterface) 

	// public void removeExtendedModifyListener(org.eclipse.swt.custom.ExtendedModifyListener)
	RemoveExtendedModifyListener(a CustomExtendedModifyListenerInterface) 

	// public void removeLineBackgroundListener(org.eclipse.swt.custom.LineBackgroundListener)
	RemoveLineBackgroundListener(a CustomLineBackgroundListenerInterface) 

	// public void removeLineStyleListener(org.eclipse.swt.custom.LineStyleListener)
	RemoveLineStyleListener(a CustomLineStyleListenerInterface) 

	// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
	RemoveModifyListener(a EventsModifyListenerInterface) 

	// public void removePaintObjectListener(org.eclipse.swt.custom.PaintObjectListener)
	RemovePaintObjectListener(a CustomPaintObjectListenerInterface) 

	// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
	RemoveSelectionListener(a EventsSelectionListenerInterface) 

	// public void removeVerifyListener(org.eclipse.swt.events.VerifyListener)
	RemoveVerifyListener(a EventsVerifyListenerInterface) 

	// public void removeVerifyKeyListener(org.eclipse.swt.custom.VerifyKeyListener)
	RemoveVerifyKeyListener(a CustomVerifyKeyListenerInterface) 

	// public void removeWordMovementListener(org.eclipse.swt.custom.MovementListener)
	RemoveWordMovementListener(a CustomMovementListenerInterface) 

	// public void replaceStyleRanges(int, int, org.eclipse.swt.custom.StyleRange[])
	ReplaceStyleRanges(a int, b int, c []*CustomStyleRange) 

	// public void replaceTextRange(int, int, java.lang.String)
	ReplaceTextRange(a int, b int, c string) 

	// public void selectAll()
	SelectAll() 

	// public void setAlignment(int)
	SetAlignment(a int) 

	// public void setAlwaysShowScrollBars(boolean)
	SetAlwaysShowScrollBars(a bool) 

	// public void setBlockSelection(boolean)
	SetBlockSelection(a bool) 

	// public void setBlockSelectionBounds(org.eclipse.swt.graphics.Rectangle)
	SetBlockSelectionBounds2(a GraphicsRectangleInterface) 

	// public void setBlockSelectionBounds(int, int, int, int)
	SetBlockSelectionBounds(a int, b int, c int, d int) 

	// public void setBidiColoring(boolean)
	SetBidiColoring(a bool) 

	// public void setBottomMargin(int)
	SetBottomMargin(a int) 

	// public void setCaretOffset(int)
	SetCaretOffset(a int) 

	// public void setContent(org.eclipse.swt.custom.StyledTextContent)
	SetContent(a CustomStyledTextContentInterface) 

	// public void setDoubleClickEnabled(boolean)
	SetDoubleClickEnabled(a bool) 

	// public void setEditable(boolean)
	SetEditable(a bool) 

	// public void setHorizontalIndex(int)
	SetHorizontalIndex(a int) 

	// public void setHorizontalPixel(int)
	SetHorizontalPixel(a int) 

	// public void setIndent(int)
	SetIndent(a int) 

	// public void setJustify(boolean)
	SetJustify(a bool) 

	// public void setKeyBinding(int, int)
	SetKeyBinding(a int, b int) 

	// public void setLeftMargin(int)
	SetLeftMargin(a int) 

	// public void setLineAlignment(int, int, int)
	SetLineAlignment(a int, b int, c int) 

	// public void setLineBackground(int, int, org.eclipse.swt.graphics.Color)
	SetLineBackground(a int, b int, c GraphicsColorInterface) 

	// public void setLineBullet(int, int, org.eclipse.swt.custom.Bullet)
	SetLineBullet(a int, b int, c CustomBulletInterface) 

	// public void setLineIndent(int, int, int)
	SetLineIndent(a int, b int, c int) 

	// public void setLineJustify(int, int, boolean)
	SetLineJustify(a int, b int, c bool) 

	// public void setLineSpacing(int)
	SetLineSpacing(a int) 

	// public void setLineTabStops(int, int, int[])
	SetLineTabStops(a int, b int, c []int) 

	// public void setLineWrapIndent(int, int, int)
	SetLineWrapIndent(a int, b int, c int) 

	// public void setMarginColor(org.eclipse.swt.graphics.Color)
	SetMarginColor(a GraphicsColorInterface) 

	// public void setMargins(int, int, int, int)
	SetMargins(a int, b int, c int, d int) 

	// public void setRightMargin(int)
	SetRightMargin(a int) 

	// public void setSelection(int)
	SetSelection(a int) 

	// public void setSelection(org.eclipse.swt.graphics.Point)
	SetSelection3(a GraphicsPointInterface) 

	// public void setSelectionBackground(org.eclipse.swt.graphics.Color)
	SetSelectionBackground(a GraphicsColorInterface) 

	// public void setSelectionForeground(org.eclipse.swt.graphics.Color)
	SetSelectionForeground(a GraphicsColorInterface) 

	// public void setSelection(int, int)
	SetSelection2(a int, b int) 

	// public void setSelectionRange(int, int)
	SetSelectionRange(a int, b int) 

	// public void setStyleRange(org.eclipse.swt.custom.StyleRange)
	SetStyleRange(a CustomStyleRangeInterface) 

	// public void setStyleRanges(int, int, int[], org.eclipse.swt.custom.StyleRange[])
	SetStyleRanges3(a int, b int, c []int, d []*CustomStyleRange) 

	// public void setStyleRanges(int[], org.eclipse.swt.custom.StyleRange[])
	SetStyleRanges2(a []int, b []*CustomStyleRange) 

	// public void setStyleRanges(org.eclipse.swt.custom.StyleRange[])
	SetStyleRanges(a []*CustomStyleRange) 

	// public void setTabs(int)
	SetTabs(a int) 

	// public void setTabStops(int[])
	SetTabStops(a []int) 

	// public void setText(java.lang.String)
	SetText(a string) 

	// public void setTextLimit(int)
	SetTextLimit(a int) 

	// public void setTopIndex(int)
	SetTopIndex(a int) 

	// public void setTopMargin(int)
	SetTopMargin(a int) 

	// public void setTopPixel(int)
	SetTopPixel(a int) 

	// public void setWordWrap(boolean)
	SetWordWrap(a bool) 

	// public void setWrapIndent(int)
	SetWrapIndent(a int) 

	// public void showSelection()
	ShowSelection() 
}

type CustomStyledText struct {
	WidgetsCanvas
}

// public org.eclipse.swt.custom.StyledText(org.eclipse.swt.widgets.Composite, int)
func NewCustomStyledText(a WidgetsCompositeInterface, b int) (*CustomStyledText) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StyledText", conv_a.Value().Cast("org/eclipse/swt/widgets/Composite"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomStyledText{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addExtendedModifyListener(org.eclipse.swt.custom.ExtendedModifyListener)
func (jbobject *CustomStyledText) AddExtendedModifyListener(a CustomExtendedModifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addExtendedModifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/ExtendedModifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addBidiSegmentListener(org.eclipse.swt.custom.BidiSegmentListener)
func (jbobject *CustomStyledText) AddBidiSegmentListener(a CustomBidiSegmentListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addBidiSegmentListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/BidiSegmentListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addCaretListener(org.eclipse.swt.custom.CaretListener)
func (jbobject *CustomStyledText) AddCaretListener(a CustomCaretListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addCaretListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CaretListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addLineBackgroundListener(org.eclipse.swt.custom.LineBackgroundListener)
func (jbobject *CustomStyledText) AddLineBackgroundListener(a CustomLineBackgroundListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addLineBackgroundListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/LineBackgroundListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addLineStyleListener(org.eclipse.swt.custom.LineStyleListener)
func (jbobject *CustomStyledText) AddLineStyleListener(a CustomLineStyleListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addLineStyleListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/LineStyleListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *CustomStyledText) AddModifyListener(a EventsModifyListenerInterface)  {
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

// public void addPaintObjectListener(org.eclipse.swt.custom.PaintObjectListener)
func (jbobject *CustomStyledText) AddPaintObjectListener(a CustomPaintObjectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addPaintObjectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/PaintObjectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomStyledText) AddSelectionListener(a EventsSelectionListenerInterface)  {
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

// public void addVerifyKeyListener(org.eclipse.swt.custom.VerifyKeyListener)
func (jbobject *CustomStyledText) AddVerifyKeyListener(a CustomVerifyKeyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addVerifyKeyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/VerifyKeyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addVerifyListener(org.eclipse.swt.events.VerifyListener)
func (jbobject *CustomStyledText) AddVerifyListener(a EventsVerifyListenerInterface)  {
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

// public void addWordMovementListener(org.eclipse.swt.custom.MovementListener)
func (jbobject *CustomStyledText) AddWordMovementListener(a CustomMovementListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addWordMovementListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/MovementListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void append(java.lang.String)
func (jbobject *CustomStyledText) Append(a string)  {
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

// public org.eclipse.swt.graphics.Point computeSize(int, int, boolean)
func (jbobject *CustomStyledText) ComputeSize2(a int, b int, c bool) *GraphicsPoint {
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

// public void copy()
func (jbobject *CustomStyledText) Copy()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copy", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void copy(int)
func (jbobject *CustomStyledText) Copy2(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copy", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public int getAlignment()
func (jbobject *CustomStyledText) GetAlignment() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getAlwaysShowScrollBars()
func (jbobject *CustomStyledText) GetAlwaysShowScrollBars() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlwaysShowScrollBars", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Color getMarginColor()
func (jbobject *CustomStyledText) GetMarginColor() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMarginColor", "org/eclipse/swt/graphics/Color")
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

// public void cut()
func (jbobject *CustomStyledText) Cut()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cut", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *CustomStyledText) GetBackground() *GraphicsColor {
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

// public int getBaseline()
func (jbobject *CustomStyledText) GetBaseline() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBaseline", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getBaseline(int)
func (jbobject *CustomStyledText) GetBaseline2(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBaseline", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getBidiColoring()
func (jbobject *CustomStyledText) GetBidiColoring() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBidiColoring", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getBlockSelection()
func (jbobject *CustomStyledText) GetBlockSelection() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockSelection", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Rectangle getBlockSelectionBounds()
func (jbobject *CustomStyledText) GetBlockSelectionBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockSelectionBounds", "org/eclipse/swt/graphics/Rectangle")
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

// public int getBottomMargin()
func (jbobject *CustomStyledText) GetBottomMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBottomMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getCaretOffset()
func (jbobject *CustomStyledText) GetCaretOffset() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCaretOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.custom.StyledTextContent getContent()
func (jbobject *CustomStyledText) GetContent() *CustomStyledTextContent {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContent", "org/eclipse/swt/custom/StyledTextContent")
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
	unique_x := &CustomStyledTextContent{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean getDragDetect()
func (jbobject *CustomStyledText) GetDragDetect() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDragDetect", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getDoubleClickEnabled()
func (jbobject *CustomStyledText) GetDoubleClickEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDoubleClickEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean getEditable()
func (jbobject *CustomStyledText) GetEditable() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEditable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.eclipse.swt.graphics.Color getForeground()
func (jbobject *CustomStyledText) GetForeground() *GraphicsColor {
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

// public int getHorizontalIndex()
func (jbobject *CustomStyledText) GetHorizontalIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHorizontalIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getHorizontalPixel()
func (jbobject *CustomStyledText) GetHorizontalPixel() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHorizontalPixel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getIndent()
func (jbobject *CustomStyledText) GetIndent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getJustify()
func (jbobject *CustomStyledText) GetJustify() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJustify", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getKeyBinding(int)
func (jbobject *CustomStyledText) GetKeyBinding(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKeyBinding", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getCharCount()
func (jbobject *CustomStyledText) GetCharCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCharCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getLine(int)
func (jbobject *CustomStyledText) GetLine(a int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLine", "java/lang/String", a)
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

// public int getLineAlignment(int)
func (jbobject *CustomStyledText) GetLineAlignment(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineAlignment", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineAtOffset(int)
func (jbobject *CustomStyledText) GetLineAtOffset(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineAtOffset", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Color getLineBackground(int)
func (jbobject *CustomStyledText) GetLineBackground(a int) *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineBackground", "org/eclipse/swt/graphics/Color", a)
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

// public org.eclipse.swt.custom.Bullet getLineBullet(int)
func (jbobject *CustomStyledText) GetLineBullet(a int) *CustomBullet {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineBullet", "org/eclipse/swt/custom/Bullet", a)
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
	unique_x := &CustomBullet{}
	unique_x.Callable = dst
	return unique_x
}

// public int getLineCount()
func (jbobject *CustomStyledText) GetLineCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getLineDelimiter()
func (jbobject *CustomStyledText) GetLineDelimiter() string {
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
func (jbobject *CustomStyledText) GetLineHeight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineHeight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineHeight(int)
func (jbobject *CustomStyledText) GetLineHeight2(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineHeight", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineIndent(int)
func (jbobject *CustomStyledText) GetLineIndent(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineIndent", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getLineJustify(int)
func (jbobject *CustomStyledText) GetLineJustify(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineJustify", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getLineSpacing()
func (jbobject *CustomStyledText) GetLineSpacing() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineSpacing", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLinePixel(int)
func (jbobject *CustomStyledText) GetLinePixel(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLinePixel", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineIndex(int)
func (jbobject *CustomStyledText) GetLineIndex(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineIndex", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getLineTabStops(int)
func (jbobject *CustomStyledText) GetLineTabStops(a int) []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineTabStops", javabind.Int | javabind.Array, a)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int getLineWrapIndent(int)
func (jbobject *CustomStyledText) GetLineWrapIndent(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineWrapIndent", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLeftMargin()
func (jbobject *CustomStyledText) GetLeftMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLeftMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point getLocationAtOffset(int)
func (jbobject *CustomStyledText) GetLocationAtOffset(a int) *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocationAtOffset", "org/eclipse/swt/graphics/Point", a)
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

// public int getOffsetAtLine(int)
func (jbobject *CustomStyledText) GetOffsetAtLine(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOffsetAtLine", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getOffsetAtLocation(org.eclipse.swt.graphics.Point)
func (jbobject *CustomStyledText) GetOffsetAtLocation(a GraphicsPointInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOffsetAtLocation", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int getOrientation()
func (jbobject *CustomStyledText) GetOrientation() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrientation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getRanges()
func (jbobject *CustomStyledText) GetRanges() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRanges", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int[] getRanges(int, int)
func (jbobject *CustomStyledText) GetRanges2(a int, b int) []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRanges", javabind.Int | javabind.Array, a, b)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public int getRightMargin()
func (jbobject *CustomStyledText) GetRightMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRightMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Point getSelection()
func (jbobject *CustomStyledText) GetSelection() *GraphicsPoint {
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

// public org.eclipse.swt.graphics.Point getSelectionRange()
func (jbobject *CustomStyledText) GetSelectionRange() *GraphicsPoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionRange", "org/eclipse/swt/graphics/Point")
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

// public int[] getSelectionRanges()
func (jbobject *CustomStyledText) GetSelectionRanges() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionRanges", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public org.eclipse.swt.graphics.Color getSelectionBackground()
func (jbobject *CustomStyledText) GetSelectionBackground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionBackground", "org/eclipse/swt/graphics/Color")
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

// public int getSelectionCount()
func (jbobject *CustomStyledText) GetSelectionCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.Color getSelectionForeground()
func (jbobject *CustomStyledText) GetSelectionForeground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelectionForeground", "org/eclipse/swt/graphics/Color")
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

// public java.lang.String getSelectionText()
func (jbobject *CustomStyledText) GetSelectionText() string {
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

// public org.eclipse.swt.custom.StyleRange getStyleRangeAtOffset(int)
func (jbobject *CustomStyledText) GetStyleRangeAtOffset(a int) *CustomStyleRange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyleRangeAtOffset", "org/eclipse/swt/custom/StyleRange", a)
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
	unique_x := &CustomStyleRange{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.custom.StyleRange[] getStyleRanges()
func (jbobject *CustomStyledText) GetStyleRanges() []*CustomStyleRange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyleRanges", javabind.ObjectArrayType("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/StyleRange")
	dst := new([]*CustomStyleRange)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.custom.StyleRange[] getStyleRanges(boolean)
func (jbobject *CustomStyledText) GetStyleRanges2(a bool) []*CustomStyleRange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyleRanges", javabind.ObjectArrayType("org/eclipse/swt/custom/StyleRange"), a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/StyleRange")
	dst := new([]*CustomStyleRange)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.custom.StyleRange[] getStyleRanges(int, int)
func (jbobject *CustomStyledText) GetStyleRanges3(a int, b int) []*CustomStyleRange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyleRanges", javabind.ObjectArrayType("org/eclipse/swt/custom/StyleRange"), a, b)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/StyleRange")
	dst := new([]*CustomStyleRange)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.custom.StyleRange[] getStyleRanges(int, int, boolean)
func (jbobject *CustomStyledText) GetStyleRanges4(a int, b int, c bool) []*CustomStyleRange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyleRanges", javabind.ObjectArrayType("org/eclipse/swt/custom/StyleRange"), a, b, c)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/custom/StyleRange")
	dst := new([]*CustomStyleRange)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getTabs()
func (jbobject *CustomStyledText) GetTabs() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTabs", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int[] getTabStops()
func (jbobject *CustomStyledText) GetTabStops() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTabStops", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

// public java.lang.String getText()
func (jbobject *CustomStyledText) GetText() string {
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
func (jbobject *CustomStyledText) GetText2(a int, b int) string {
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

// public org.eclipse.swt.graphics.Rectangle getTextBounds(int, int)
func (jbobject *CustomStyledText) GetTextBounds(a int, b int) *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextBounds", "org/eclipse/swt/graphics/Rectangle", a, b)
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

// public java.lang.String getTextRange(int, int)
func (jbobject *CustomStyledText) GetTextRange(a int, b int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextRange", "java/lang/String", a, b)
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

// public int getTextLimit()
func (jbobject *CustomStyledText) GetTextLimit() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextLimit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTopIndex()
func (jbobject *CustomStyledText) GetTopIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTopMargin()
func (jbobject *CustomStyledText) GetTopMargin() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopMargin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getTopPixel()
func (jbobject *CustomStyledText) GetTopPixel() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTopPixel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean getWordWrap()
func (jbobject *CustomStyledText) GetWordWrap() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWordWrap", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int getWrapIndent()
func (jbobject *CustomStyledText) GetWrapIndent() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWrapIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void insert(java.lang.String)
func (jbobject *CustomStyledText) Insert(a string)  {
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

// public void invokeAction(int)
func (jbobject *CustomStyledText) InvokeAction(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "invokeAction", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void paste()
func (jbobject *CustomStyledText) Paste()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "paste", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void print()
func (jbobject *CustomStyledText) Print2()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void redraw()
func (jbobject *CustomStyledText) Redraw()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "redraw", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void redraw(int, int, int, int, boolean)
func (jbobject *CustomStyledText) Redraw2(a int, b int, c int, d int, e bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "redraw", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public void redrawRange(int, int, boolean)
func (jbobject *CustomStyledText) RedrawRange(a int, b int, c bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "redrawRange", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void removeBidiSegmentListener(org.eclipse.swt.custom.BidiSegmentListener)
func (jbobject *CustomStyledText) RemoveBidiSegmentListener(a CustomBidiSegmentListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeBidiSegmentListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/BidiSegmentListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeCaretListener(org.eclipse.swt.custom.CaretListener)
func (jbobject *CustomStyledText) RemoveCaretListener(a CustomCaretListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeCaretListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/CaretListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeExtendedModifyListener(org.eclipse.swt.custom.ExtendedModifyListener)
func (jbobject *CustomStyledText) RemoveExtendedModifyListener(a CustomExtendedModifyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeExtendedModifyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/ExtendedModifyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeLineBackgroundListener(org.eclipse.swt.custom.LineBackgroundListener)
func (jbobject *CustomStyledText) RemoveLineBackgroundListener(a CustomLineBackgroundListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeLineBackgroundListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/LineBackgroundListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeLineStyleListener(org.eclipse.swt.custom.LineStyleListener)
func (jbobject *CustomStyledText) RemoveLineStyleListener(a CustomLineStyleListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeLineStyleListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/LineStyleListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeModifyListener(org.eclipse.swt.events.ModifyListener)
func (jbobject *CustomStyledText) RemoveModifyListener(a EventsModifyListenerInterface)  {
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

// public void removePaintObjectListener(org.eclipse.swt.custom.PaintObjectListener)
func (jbobject *CustomStyledText) RemovePaintObjectListener(a CustomPaintObjectListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removePaintObjectListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/PaintObjectListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSelectionListener(org.eclipse.swt.events.SelectionListener)
func (jbobject *CustomStyledText) RemoveSelectionListener(a EventsSelectionListenerInterface)  {
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
func (jbobject *CustomStyledText) RemoveVerifyListener(a EventsVerifyListenerInterface)  {
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

// public void removeVerifyKeyListener(org.eclipse.swt.custom.VerifyKeyListener)
func (jbobject *CustomStyledText) RemoveVerifyKeyListener(a CustomVerifyKeyListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeVerifyKeyListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/VerifyKeyListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeWordMovementListener(org.eclipse.swt.custom.MovementListener)
func (jbobject *CustomStyledText) RemoveWordMovementListener(a CustomMovementListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeWordMovementListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/MovementListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void replaceStyleRanges(int, int, org.eclipse.swt.custom.StyleRange[])
func (jbobject *CustomStyledText) ReplaceStyleRanges(a int, b int, c []*CustomStyleRange)  {
	conv_c := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/custom/StyleRange")
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "replaceStyleRanges", javabind.Void, a, b, conv_c.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void replaceTextRange(int, int, java.lang.String)
func (jbobject *CustomStyledText) ReplaceTextRange(a int, b int, c string)  {
	conv_c := javabind.NewGoToJavaString()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "replaceTextRange", javabind.Void, a, b, conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void scroll(int, int, int, int, int, int, boolean)
func (jbobject *CustomStyledText) Scroll(a int, b int, c int, d int, e int, f int, g bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "scroll", javabind.Void, a, b, c, d, e, f, g)
	if err != nil {
		panic(err)
	}

}

// public void selectAll()
func (jbobject *CustomStyledText) SelectAll()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "selectAll", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setAlignment(int)
func (jbobject *CustomStyledText) SetAlignment(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlignment", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setAlwaysShowScrollBars(boolean)
func (jbobject *CustomStyledText) SetAlwaysShowScrollBars(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlwaysShowScrollBars", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomStyledText) SetBackground(a GraphicsColorInterface)  {
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

// public void setBlockSelection(boolean)
func (jbobject *CustomStyledText) SetBlockSelection(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlockSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBlockSelectionBounds(org.eclipse.swt.graphics.Rectangle)
func (jbobject *CustomStyledText) SetBlockSelectionBounds2(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlockSelectionBounds", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBlockSelectionBounds(int, int, int, int)
func (jbobject *CustomStyledText) SetBlockSelectionBounds(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlockSelectionBounds", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void setCaret(org.eclipse.swt.widgets.Caret)
func (jbobject *CustomStyledText) SetCaret(a WidgetsCaretInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCaret", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Caret"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setBidiColoring(boolean)
func (jbobject *CustomStyledText) SetBidiColoring(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBidiColoring", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBottomMargin(int)
func (jbobject *CustomStyledText) SetBottomMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBottomMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setCaretOffset(int)
func (jbobject *CustomStyledText) SetCaretOffset(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCaretOffset", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setContent(org.eclipse.swt.custom.StyledTextContent)
func (jbobject *CustomStyledText) SetContent(a CustomStyledTextContentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setContent", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/StyledTextContent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setCursor(org.eclipse.swt.graphics.Cursor)
func (jbobject *CustomStyledText) SetCursor(a GraphicsCursorInterface)  {
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

// public void setDoubleClickEnabled(boolean)
func (jbobject *CustomStyledText) SetDoubleClickEnabled(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDoubleClickEnabled", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setDragDetect(boolean)
func (jbobject *CustomStyledText) SetDragDetect(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDragDetect", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEditable(boolean)
func (jbobject *CustomStyledText) SetEditable(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEditable", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setFont(org.eclipse.swt.graphics.Font)
func (jbobject *CustomStyledText) SetFont(a GraphicsFontInterface)  {
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
func (jbobject *CustomStyledText) SetForeground(a GraphicsColorInterface)  {
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

// public void setHorizontalIndex(int)
func (jbobject *CustomStyledText) SetHorizontalIndex(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHorizontalIndex", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setHorizontalPixel(int)
func (jbobject *CustomStyledText) SetHorizontalPixel(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHorizontalPixel", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setIndent(int)
func (jbobject *CustomStyledText) SetIndent(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIndent", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setJustify(boolean)
func (jbobject *CustomStyledText) SetJustify(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJustify", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setKeyBinding(int, int)
func (jbobject *CustomStyledText) SetKeyBinding(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKeyBinding", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setLeftMargin(int)
func (jbobject *CustomStyledText) SetLeftMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLeftMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLineAlignment(int, int, int)
func (jbobject *CustomStyledText) SetLineAlignment(a int, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineAlignment", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setLineBackground(int, int, org.eclipse.swt.graphics.Color)
func (jbobject *CustomStyledText) SetLineBackground(a int, b int, c GraphicsColorInterface)  {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineBackground", javabind.Void, a, b, conv_c.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void setLineBullet(int, int, org.eclipse.swt.custom.Bullet)
func (jbobject *CustomStyledText) SetLineBullet(a int, b int, c CustomBulletInterface)  {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineBullet", javabind.Void, a, b, conv_c.Value().Cast("org/eclipse/swt/custom/Bullet"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void setLineIndent(int, int, int)
func (jbobject *CustomStyledText) SetLineIndent(a int, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineIndent", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setLineJustify(int, int, boolean)
func (jbobject *CustomStyledText) SetLineJustify(a int, b int, c bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineJustify", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setLineSpacing(int)
func (jbobject *CustomStyledText) SetLineSpacing(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineSpacing", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setLineTabStops(int, int, int[])
func (jbobject *CustomStyledText) SetLineTabStops(a int, b int, c []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineTabStops", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setLineWrapIndent(int, int, int)
func (jbobject *CustomStyledText) SetLineWrapIndent(a int, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLineWrapIndent", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setMarginColor(org.eclipse.swt.graphics.Color)
func (jbobject *CustomStyledText) SetMarginColor(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMarginColor", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMargins(int, int, int, int)
func (jbobject *CustomStyledText) SetMargins(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMargins", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void setOrientation(int)
func (jbobject *CustomStyledText) SetOrientation(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrientation", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setRightMargin(int)
func (jbobject *CustomStyledText) SetRightMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRightMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(int)
func (jbobject *CustomStyledText) SetSelection(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setSelection(org.eclipse.swt.graphics.Point)
func (jbobject *CustomStyledText) SetSelection3(a GraphicsPointInterface)  {
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

// public void setSelectionBackground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomStyledText) SetSelectionBackground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelectionForeground(org.eclipse.swt.graphics.Color)
func (jbobject *CustomStyledText) SetSelectionForeground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionForeground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelection(int, int)
func (jbobject *CustomStyledText) SetSelection2(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelection", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setSelectionRange(int, int)
func (jbobject *CustomStyledText) SetSelectionRange(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelectionRange", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void setStyleRange(org.eclipse.swt.custom.StyleRange)
func (jbobject *CustomStyledText) SetStyleRange(a CustomStyleRangeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStyleRange", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setStyleRanges(int, int, int[], org.eclipse.swt.custom.StyleRange[])
func (jbobject *CustomStyledText) SetStyleRanges3(a int, b int, c []int, d []*CustomStyleRange)  {
	conv_d := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/custom/StyleRange")
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStyleRanges", javabind.Void, a, b, c, conv_d.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_d.CleanUp()

}

// public void setStyleRanges(int[], org.eclipse.swt.custom.StyleRange[])
func (jbobject *CustomStyledText) SetStyleRanges2(a []int, b []*CustomStyleRange)  {
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/custom/StyleRange")
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStyleRanges", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void setStyleRanges(org.eclipse.swt.custom.StyleRange[])
func (jbobject *CustomStyledText) SetStyleRanges(a []*CustomStyleRange)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/custom/StyleRange")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStyleRanges", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTabs(int)
func (jbobject *CustomStyledText) SetTabs(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTabs", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTabStops(int[])
func (jbobject *CustomStyledText) SetTabStops(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTabStops", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setText(java.lang.String)
func (jbobject *CustomStyledText) SetText(a string)  {
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

// public void setTextLimit(int)
func (jbobject *CustomStyledText) SetTextLimit(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTextLimit", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTopIndex(int)
func (jbobject *CustomStyledText) SetTopIndex(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopIndex", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTopMargin(int)
func (jbobject *CustomStyledText) SetTopMargin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopMargin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTopPixel(int)
func (jbobject *CustomStyledText) SetTopPixel(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTopPixel", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setWordWrap(boolean)
func (jbobject *CustomStyledText) SetWordWrap(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWordWrap", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setWrapIndent(int)
func (jbobject *CustomStyledText) SetWrapIndent(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWrapIndent", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void showSelection()
func (jbobject *CustomStyledText) ShowSelection()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "showSelection", javabind.Void)
	if err != nil {
		panic(err)
	}

}


