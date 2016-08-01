package swt

import "github.com/timob/javabind"

type WidgetsEventInterface interface {
	JavaLangObjectInterface

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public void setBounds(org.eclipse.swt.graphics.Rectangle)
	SetBounds(a GraphicsRectangleInterface) 
}

type WidgetsEvent struct {
	JavaLangObject
}

// public org.eclipse.swt.widgets.Event()
func NewWidgetsEvent() (*WidgetsEvent) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Event")
	if err != nil {
		panic(err)
	}
	x := &WidgetsEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *WidgetsEvent) GetBounds() *GraphicsRectangle {
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
func (jbobject *WidgetsEvent) SetBounds(a GraphicsRectangleInterface)  {
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

// public java.lang.String toString()
func (jbobject *WidgetsEvent) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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

func (jbobject *WidgetsEvent) Display() *WidgetsDisplay {
	jret, err := jbobject.GetField(javabind.GetEnv(), "display", "org/eclipse/swt/widgets/Display")
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
	unique_x := &WidgetsDisplay{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *WidgetsEvent) SetFieldDisplay(val WidgetsDisplayInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "display", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsEvent) Widget() *WidgetsWidget {
	jret, err := jbobject.GetField(javabind.GetEnv(), "widget", "org/eclipse/swt/widgets/Widget")
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
	unique_x := &WidgetsWidget{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *WidgetsEvent) SetFieldWidget(val WidgetsWidgetInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "widget", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsEvent) Type() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "type", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldType(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "type", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Detail() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "detail", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldDetail(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "detail", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Item() *WidgetsWidget {
	jret, err := jbobject.GetField(javabind.GetEnv(), "item", "org/eclipse/swt/widgets/Widget")
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
	unique_x := &WidgetsWidget{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *WidgetsEvent) SetFieldItem(val WidgetsWidgetInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "item", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsEvent) Index() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "index", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldIndex(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "index", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Gc() *GraphicsGC {
	jret, err := jbobject.GetField(javabind.GetEnv(), "gc", "org/eclipse/swt/graphics/GC")
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
	unique_x := &GraphicsGC{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *WidgetsEvent) SetFieldGc(val GraphicsGCInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "gc", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Count() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "count", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldCount(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "count", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Time() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "time", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldTime(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "time", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Button() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "button", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldButton(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "button", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Character() uint16 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "character", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *WidgetsEvent) SetFieldCharacter(val uint16) {
	err := jbobject.SetField(javabind.GetEnv(), "character", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) KeyCode() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "keyCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldKeyCode(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "keyCode", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) KeyLocation() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "keyLocation", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldKeyLocation(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "keyLocation", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) StateMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "stateMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldStateMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "stateMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Start() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "start", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldStart(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "start", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) End() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "end", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldEnd(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "end", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Text() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "text", "java/lang/String")
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

func (jbobject *WidgetsEvent) SetFieldText(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "text", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsEvent) Segments() []int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "segments", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

func (jbobject *WidgetsEvent) SetFieldSegments(val []int) {
	err := jbobject.SetField(javabind.GetEnv(), "segments", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) SegmentsChars() []uint16 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "segmentsChars", javabind.Char | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]uint16)
}

func (jbobject *WidgetsEvent) SetFieldSegmentsChars(val []uint16) {
	err := jbobject.SetField(javabind.GetEnv(), "segmentsChars", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Doit() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "doit", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *WidgetsEvent) SetFieldDoit(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "doit", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Data() *JavaLangObject {
	jret, err := jbobject.GetField(javabind.GetEnv(), "data", "java/lang/Object")
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

func (jbobject *WidgetsEvent) SetFieldData(val JavaLangObjectInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "data", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsEvent) Touches() []*WidgetsTouch {
	jret, err := jbobject.GetField(javabind.GetEnv(), "touches", javabind.ObjectArrayType("org/eclipse/swt/widgets/Touch"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/Touch")
	dst := new([]*WidgetsTouch)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *WidgetsEvent) SetFieldTouches(val []*WidgetsTouch) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/widgets/Touch")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "touches", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *WidgetsEvent) XDirection() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "xDirection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldXDirection(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "xDirection", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) YDirection() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "yDirection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *WidgetsEvent) SetFieldYDirection(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "yDirection", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Magnification() float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "magnification", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

func (jbobject *WidgetsEvent) SetFieldMagnification(val float64) {
	err := jbobject.SetField(javabind.GetEnv(), "magnification", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *WidgetsEvent) Rotation() float64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "rotation", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

func (jbobject *WidgetsEvent) SetFieldRotation(val float64) {
	err := jbobject.SetField(javabind.GetEnv(), "rotation", val)
	if err != nil {
		panic(err)
	}

}


