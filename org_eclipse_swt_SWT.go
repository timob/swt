package swt

import "github.com/timob/javabind"

type SWTInterface interface {
	JavaLangObjectInterface

	// public static boolean isLoadable()
	IsLoadable() bool

	// public static java.lang.String getMessage(java.lang.String)
	GetMessage(a string) string

	// public static java.lang.String getMessage(java.lang.String, java.lang.Object[])
	GetMessage2(a string, b []*JavaLangObject) string

	// public static java.lang.String getPlatform()
	GetPlatform() string

	// public static int getVersion()
	GetVersion() int

	// public static void error(int)
	Error(a int) 
}

type SWT struct {
	JavaLangObject
}

// public org.eclipse.swt.SWT()
func NewSWT() (*SWT) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/SWT")
	if err != nil {
		panic(err)
	}
	x := &SWT{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static boolean isLoadable()
func (jbobject *SWT) IsLoadable() bool {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/SWT", "isLoadable", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static java.lang.String getMessage(java.lang.String)
func (jbobject *SWT) GetMessage(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/SWT", "getMessage", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String getMessage(java.lang.String, java.lang.Object[])
func (jbobject *SWT) GetMessage2(a string, b []*JavaLangObject) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/SWT", "getMessage", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String getPlatform()
func (jbobject *SWT) GetPlatform() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/SWT", "getPlatform", "java/lang/String")
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

// public static int getVersion()
func (jbobject *SWT) GetVersion() int {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/SWT", "getVersion", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public static void error(int)
func (jbobject *SWT) Error(a int)  {
	_, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/SWT", "error", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) None() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "None", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNone(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "None", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KeyDown() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KeyDown", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKeyDown(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KeyDown", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KeyUp() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KeyUp", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKeyUp(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KeyUp", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseDown() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseDown", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseDown(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseDown", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseUp() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseUp", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseUp(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseUp", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseMove() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseMove", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseMove(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseMove", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseEnter() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseEnter", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseEnter(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseEnter", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseExit() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseExit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseExit(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseExit", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseDoubleClick() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseDoubleClick", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseDoubleClick(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseDoubleClick", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Paint() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Paint", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPaint(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Paint", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Move() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Move", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMove(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Move", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Resize() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Resize", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldResize(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Resize", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Dispose() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Dispose", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDispose(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Dispose", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Selection() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Selection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSelection(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Selection", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DefaultSelection() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DefaultSelection", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDefaultSelection(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DefaultSelection", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FocusIn() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FocusIn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFocusIn(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FocusIn", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FocusOut() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FocusOut", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFocusOut(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FocusOut", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Expand() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Expand", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldExpand(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Expand", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Collapse() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Collapse", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCollapse(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Collapse", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Iconify() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Iconify", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIconify(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Iconify", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Deiconify() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Deiconify", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDeiconify(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Deiconify", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Close() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Close", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldClose(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Close", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Show() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Show", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldShow(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Show", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Hide() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Hide", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHide(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Hide", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Modify() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Modify", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldModify(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Modify", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Verify() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Verify", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldVerify(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Verify", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Activate() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Activate", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldActivate(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Activate", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Deactivate() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Deactivate", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDeactivate(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Deactivate", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Help() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Help", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHelp(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Help", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DragDetect() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DragDetect", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDragDetect(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DragDetect", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Arm() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Arm", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldArm(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Arm", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Traverse() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Traverse", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTraverse(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Traverse", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseHover() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseHover", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseHover(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseHover", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HardKeyDown() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HardKeyDown", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHardKeyDown(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HardKeyDown", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HardKeyUp() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HardKeyUp", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHardKeyUp(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HardKeyUp", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MenuDetect() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MenuDetect", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMenuDetect(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MenuDetect", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SetData() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SetData", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSetData(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SetData", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseVerticalWheel() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseVerticalWheel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseVerticalWheel(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseVerticalWheel", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseHorizontalWheel() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseHorizontalWheel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseHorizontalWheel(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseHorizontalWheel", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MouseWheel() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MouseWheel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMouseWheel(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MouseWheel", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Settings() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Settings", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSettings(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Settings", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) EraseItem() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "EraseItem", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldEraseItem(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "EraseItem", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MeasureItem() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MeasureItem", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMeasureItem(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MeasureItem", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PaintItem() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PaintItem", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPaintItem(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PaintItem", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ImeComposition() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ImeComposition", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldImeComposition(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ImeComposition", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) OrientationChange() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "OrientationChange", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldOrientationChange(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "OrientationChange", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Skin() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Skin", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSkin(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Skin", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) OpenDocument() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "OpenDocument", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldOpenDocument(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "OpenDocument", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Touch() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Touch", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTouch(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Touch", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Gesture() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Gesture", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldGesture(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Gesture", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) Segments() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "Segments", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSegments(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "Segments", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COMPOSITION_CHANGED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COMPOSITION_CHANGED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOMPOSITION_CHANGED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COMPOSITION_CHANGED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COMPOSITION_OFFSET() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COMPOSITION_OFFSET", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOMPOSITION_OFFSET(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COMPOSITION_OFFSET", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COMPOSITION_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COMPOSITION_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOMPOSITION_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COMPOSITION_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DRAG() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DRAG", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDRAG(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DRAG", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SELECTED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SELECTED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSELECTED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SELECTED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FOCUSED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FOCUSED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFOCUSED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FOCUSED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FOREGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FOREGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFOREGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FOREGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HOT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HOT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHOT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HOT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_NONE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_NONE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_NONE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_NONE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_ESCAPE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_ESCAPE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_ESCAPE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_ESCAPE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_RETURN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_RETURN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_RETURN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_RETURN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_TAB_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_TAB_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_TAB_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_TAB_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_TAB_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_TAB_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_TAB_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_TAB_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_ARROW_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_ARROW_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_ARROW_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_ARROW_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_ARROW_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_ARROW_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_ARROW_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_ARROW_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_MNEMONIC() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_MNEMONIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_MNEMONIC(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_MNEMONIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_PAGE_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_PAGE_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_PAGE_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_PAGE_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAVERSE_PAGE_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAVERSE_PAGE_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAVERSE_PAGE_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAVERSE_PAGE_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) GESTURE_BEGIN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "GESTURE_BEGIN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldGESTURE_BEGIN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "GESTURE_BEGIN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) GESTURE_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "GESTURE_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldGESTURE_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "GESTURE_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) GESTURE_ROTATE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "GESTURE_ROTATE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldGESTURE_ROTATE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "GESTURE_ROTATE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) GESTURE_SWIPE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "GESTURE_SWIPE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldGESTURE_SWIPE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "GESTURE_SWIPE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) GESTURE_MAGNIFY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "GESTURE_MAGNIFY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldGESTURE_MAGNIFY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "GESTURE_MAGNIFY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) GESTURE_PAN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "GESTURE_PAN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldGESTURE_PAN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "GESTURE_PAN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TOUCHSTATE_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TOUCHSTATE_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTOUCHSTATE_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TOUCHSTATE_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TOUCHSTATE_MOVE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TOUCHSTATE_MOVE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTOUCHSTATE_MOVE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TOUCHSTATE_MOVE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TOUCHSTATE_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TOUCHSTATE_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTOUCHSTATE_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TOUCHSTATE_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MENU_MOUSE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MENU_MOUSE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMENU_MOUSE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MENU_MOUSE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MENU_KEYBOARD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MENU_KEYBOARD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMENU_KEYBOARD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MENU_KEYBOARD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CHANGED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CHANGED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCHANGED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CHANGED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DEFER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DEFER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDEFER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DEFER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NONE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NONE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNONE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NONE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NULL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NULL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNULL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NULL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DEFAULT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DEFAULT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDEFAULT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DEFAULT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) OFF() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "OFF", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldOFF(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "OFF", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ON() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldON(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LOW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LOW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLOW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LOW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HIGH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HIGH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHIGH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HIGH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BAR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BAR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBAR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BAR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DROP_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DROP_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDROP_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DROP_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) POP_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "POP_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPOP_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "POP_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SEPARATOR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SEPARATOR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSEPARATOR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SEPARATOR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SEPARATOR_FILL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SEPARATOR_FILL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSEPARATOR_FILL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SEPARATOR_FILL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TOGGLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TOGGLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTOGGLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TOGGLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ARROW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ARROW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldARROW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ARROW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PUSH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PUSH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPUSH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PUSH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) RADIO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "RADIO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldRADIO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "RADIO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CHECK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CHECK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCHECK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CHECK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CASCADE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CASCADE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCASCADE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CASCADE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MULTI() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MULTI", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMULTI(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MULTI", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SINGLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SINGLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSINGLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SINGLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) READ_ONLY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "READ_ONLY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldREAD_ONLY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "READ_ONLY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) WRAP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "WRAP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldWRAP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "WRAP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SEARCH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SEARCH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSEARCH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SEARCH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SIMPLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SIMPLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSIMPLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SIMPLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PASSWORD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PASSWORD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPASSWORD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PASSWORD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHADOW_IN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHADOW_IN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHADOW_IN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHADOW_IN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHADOW_OUT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHADOW_OUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHADOW_OUT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHADOW_OUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHADOW_ETCHED_IN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHADOW_ETCHED_IN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHADOW_ETCHED_IN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHADOW_ETCHED_IN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHADOW_ETCHED_OUT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHADOW_ETCHED_OUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHADOW_ETCHED_OUT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHADOW_ETCHED_OUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHADOW_NONE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHADOW_NONE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHADOW_NONE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHADOW_NONE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) INDETERMINATE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "INDETERMINATE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldINDETERMINATE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "INDETERMINATE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TOOL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TOOL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTOOL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TOOL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO_TRIM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO_TRIM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO_TRIM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO_TRIM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) RESIZE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "RESIZE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldRESIZE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "RESIZE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TITLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TITLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTITLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TITLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CLOSE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CLOSE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCLOSE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CLOSE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MENU() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MENU", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMENU(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MENU", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MIN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MIN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMIN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MIN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MAX() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MAX", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMAX(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MAX", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) H_SCROLL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "H_SCROLL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldH_SCROLL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "H_SCROLL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) V_SCROLL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "V_SCROLL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldV_SCROLL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "V_SCROLL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO_SCROLL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO_SCROLL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO_SCROLL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO_SCROLL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BORDER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BORDER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBORDER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BORDER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CLIP_CHILDREN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CLIP_CHILDREN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCLIP_CHILDREN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CLIP_CHILDREN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CLIP_SIBLINGS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CLIP_SIBLINGS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCLIP_SIBLINGS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CLIP_SIBLINGS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ON_TOP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ON_TOP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldON_TOP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ON_TOP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHEET() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHEET", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHEET(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHEET", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHELL_TRIM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHELL_TRIM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHELL_TRIM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHELL_TRIM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DIALOG_TRIM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DIALOG_TRIM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDIALOG_TRIM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DIALOG_TRIM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MODELESS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MODELESS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMODELESS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MODELESS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PRIMARY_MODAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PRIMARY_MODAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPRIMARY_MODAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PRIMARY_MODAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) APPLICATION_MODAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "APPLICATION_MODAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldAPPLICATION_MODAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "APPLICATION_MODAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SYSTEM_MODAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SYSTEM_MODAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSYSTEM_MODAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SYSTEM_MODAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HIDE_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HIDE_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHIDE_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HIDE_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FULL_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FULL_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFULL_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FULL_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FLAT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FLAT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFLAT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FLAT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SMOOTH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SMOOTH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSMOOTH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SMOOTH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO_FOCUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO_FOCUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO_FOCUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO_FOCUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO_REDRAW_RESIZE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO_REDRAW_RESIZE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO_REDRAW_RESIZE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO_REDRAW_RESIZE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO_MERGE_PAINTS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO_MERGE_PAINTS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO_MERGE_PAINTS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO_MERGE_PAINTS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO_RADIO_GROUP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO_RADIO_GROUP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO_RADIO_GROUP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO_RADIO_GROUP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LEFT_TO_RIGHT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LEFT_TO_RIGHT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLEFT_TO_RIGHT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LEFT_TO_RIGHT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) RIGHT_TO_LEFT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "RIGHT_TO_LEFT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldRIGHT_TO_LEFT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "RIGHT_TO_LEFT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MIRRORED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MIRRORED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMIRRORED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MIRRORED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) EMBEDDED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "EMBEDDED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldEMBEDDED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "EMBEDDED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) VIRTUAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "VIRTUAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldVIRTUAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "VIRTUAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DOUBLE_BUFFERED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DOUBLE_BUFFERED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDOUBLE_BUFFERED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DOUBLE_BUFFERED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRANSPARENT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRANSPARENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRANSPARENT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRANSPARENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldUP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) UNDERLINE_SINGLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "UNDERLINE_SINGLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldUNDERLINE_SINGLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "UNDERLINE_SINGLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) UNDERLINE_DOUBLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "UNDERLINE_DOUBLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldUNDERLINE_DOUBLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "UNDERLINE_DOUBLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) UNDERLINE_ERROR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "UNDERLINE_ERROR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldUNDERLINE_ERROR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "UNDERLINE_ERROR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) UNDERLINE_SQUIGGLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "UNDERLINE_SQUIGGLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldUNDERLINE_SQUIGGLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "UNDERLINE_SQUIGGLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) UNDERLINE_LINK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "UNDERLINE_LINK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldUNDERLINE_LINK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "UNDERLINE_LINK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BORDER_SOLID() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BORDER_SOLID", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBORDER_SOLID(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BORDER_SOLID", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BORDER_DASH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BORDER_DASH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBORDER_DASH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BORDER_DASH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BORDER_DOT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BORDER_DOT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBORDER_DOT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BORDER_DOT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TOP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TOP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTOP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TOP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BOTTOM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BOTTOM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBOTTOM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BOTTOM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LEAD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LEAD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLEAD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LEAD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LEFT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LEFT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLEFT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LEFT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRAIL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRAIL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRAIL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRAIL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) RIGHT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "RIGHT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldRIGHT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "RIGHT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CENTER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CENTER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCENTER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CENTER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HORIZONTAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HORIZONTAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHORIZONTAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HORIZONTAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) VERTICAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "VERTICAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldVERTICAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "VERTICAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DATE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DATE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDATE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DATE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TIME() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TIME", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTIME(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TIME", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CALENDAR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CALENDAR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCALENDAR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CALENDAR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHORT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHORT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHORT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHORT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MEDIUM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MEDIUM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMEDIUM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MEDIUM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LONG() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LONG", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLONG(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LONG", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOZILLA() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOZILLA", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOZILLA(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOZILLA", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) WEBKIT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "WEBKIT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldWEBKIT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "WEBKIT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BALLOON() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BALLOON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBALLOON(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BALLOON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BEGINNING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BEGINNING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBEGINNING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BEGINNING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FILL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FILL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFILL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FILL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DBCS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DBCS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDBCS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DBCS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ALPHA() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ALPHA", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldALPHA(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ALPHA", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NATIVE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NATIVE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNATIVE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NATIVE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PHONETIC() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PHONETIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPHONETIC(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PHONETIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ROMAN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ROMAN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldROMAN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ROMAN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BS() uint16 {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BS", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *SWT) SetFieldBS(val uint16) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CR() uint16 {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CR", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *SWT) SetFieldCR(val uint16) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DEL() uint16 {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DEL", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *SWT) SetFieldDEL(val uint16) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DEL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ESC() uint16 {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ESC", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *SWT) SetFieldESC(val uint16) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ESC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LF() uint16 {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LF", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *SWT) SetFieldLF(val uint16) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LF", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TAB() uint16 {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TAB", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *SWT) SetFieldTAB(val uint16) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TAB", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SPACE() uint16 {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SPACE", javabind.Char)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

func (jbobject *SWT) SetFieldSPACE(val uint16) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SPACE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ALT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ALT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldALT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ALT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SHIFT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SHIFT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSHIFT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SHIFT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CTRL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CTRL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCTRL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CTRL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CONTROL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CONTROL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCONTROL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CONTROL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COMMAND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COMMAND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOMMAND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COMMAND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MODIFIER_MASK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MODIFIER_MASK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMODIFIER_MASK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MODIFIER_MASK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BUTTON1() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BUTTON1", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBUTTON1(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BUTTON1", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BUTTON2() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BUTTON2", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBUTTON2(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BUTTON2", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BUTTON3() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BUTTON3", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBUTTON3(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BUTTON3", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BUTTON4() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BUTTON4", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBUTTON4(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BUTTON4", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BUTTON5() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BUTTON5", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBUTTON5(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BUTTON5", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BUTTON_MASK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BUTTON_MASK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBUTTON_MASK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BUTTON_MASK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOD1() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOD1", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOD1(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOD1", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOD2() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOD2", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOD2(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOD2", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOD3() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOD3", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOD3(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOD3", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOD4() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOD4", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOD4(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOD4", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SCROLL_LINE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SCROLL_LINE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSCROLL_LINE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SCROLL_LINE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SCROLL_PAGE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SCROLL_PAGE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSCROLL_PAGE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SCROLL_PAGE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYCODE_BIT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYCODE_BIT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYCODE_BIT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYCODE_BIT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEY_MASK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEY_MASK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEY_MASK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEY_MASK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ARROW_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ARROW_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldARROW_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ARROW_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ARROW_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ARROW_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldARROW_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ARROW_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ARROW_LEFT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ARROW_LEFT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldARROW_LEFT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ARROW_LEFT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ARROW_RIGHT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ARROW_RIGHT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldARROW_RIGHT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ARROW_RIGHT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PAGE_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PAGE_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPAGE_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PAGE_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PAGE_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PAGE_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPAGE_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PAGE_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HOME() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HOME", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHOME(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HOME", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldEND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) INSERT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "INSERT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldINSERT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "INSERT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F1() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F1", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF1(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F1", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F2() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F2", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF2(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F2", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F3() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F3", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF3(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F3", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F4() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F4", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF4(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F4", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F5() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F5", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF5(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F5", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F6() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F6", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF6(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F6", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F7() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F7", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF7(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F7", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F8() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F8", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF8(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F8", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F9() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F9", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF9(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F9", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F10() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F10", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF10(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F10", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F11() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F11", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF11(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F11", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F12() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F12", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF12(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F12", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F13() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F13", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF13(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F13", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F14() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F14", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF14(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F14", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F15() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F15", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF15(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F15", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F16() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F16", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF16(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F16", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F17() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F17", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF17(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F17", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F18() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F18", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF18(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F18", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F19() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F19", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF19(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F19", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) F20() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "F20", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldF20(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "F20", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_MULTIPLY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_MULTIPLY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_MULTIPLY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_MULTIPLY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_ADD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_ADD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_ADD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_ADD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_SUBTRACT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_SUBTRACT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_SUBTRACT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_SUBTRACT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_DECIMAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_DECIMAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_DECIMAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_DECIMAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_DIVIDE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_DIVIDE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_DIVIDE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_DIVIDE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_0() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_0", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_0(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_0", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_1() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_1", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_1(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_1", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_2() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_2", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_2(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_2", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_3() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_3", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_3(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_3", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_4() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_4", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_4(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_4", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_5() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_5", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_5(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_5", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_6() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_6", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_6(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_6", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_7() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_7", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_7(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_7", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_8() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_8", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_8(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_8", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_9() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_9", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_9(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_9", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_EQUAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_EQUAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_EQUAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_EQUAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) KEYPAD_CR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "KEYPAD_CR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldKEYPAD_CR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "KEYPAD_CR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) HELP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "HELP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldHELP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "HELP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CAPS_LOCK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CAPS_LOCK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCAPS_LOCK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CAPS_LOCK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NUM_LOCK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NUM_LOCK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNUM_LOCK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NUM_LOCK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SCROLL_LOCK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SCROLL_LOCK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSCROLL_LOCK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SCROLL_LOCK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PAUSE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PAUSE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPAUSE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PAUSE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BREAK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BREAK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBREAK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BREAK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PRINT_SCREEN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PRINT_SCREEN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPRINT_SCREEN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PRINT_SCREEN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON_ERROR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON_ERROR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON_ERROR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON_ERROR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON_INFORMATION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON_INFORMATION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON_INFORMATION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON_INFORMATION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON_QUESTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON_QUESTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON_QUESTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON_QUESTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON_WARNING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON_WARNING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON_WARNING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON_WARNING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON_WORKING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON_WORKING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON_WORKING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON_WORKING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON_SEARCH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON_SEARCH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON_SEARCH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON_SEARCH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON_CANCEL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON_CANCEL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON_CANCEL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON_CANCEL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) OK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "OK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldOK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "OK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) YES() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "YES", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldYES(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "YES", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CANCEL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CANCEL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCANCEL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CANCEL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ABORT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ABORT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldABORT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ABORT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) RETRY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "RETRY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldRETRY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "RETRY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IGNORE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IGNORE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIGNORE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IGNORE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) OPEN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "OPEN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldOPEN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "OPEN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SAVE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SAVE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSAVE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SAVE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) INHERIT_NONE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "INHERIT_NONE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldINHERIT_NONE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "INHERIT_NONE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) INHERIT_DEFAULT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "INHERIT_DEFAULT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldINHERIT_DEFAULT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "INHERIT_DEFAULT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) INHERIT_FORCE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "INHERIT_FORCE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldINHERIT_FORCE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "INHERIT_FORCE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WHITE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WHITE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WHITE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WHITE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_BLACK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_BLACK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_BLACK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_BLACK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_RED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_RED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_RED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_RED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_DARK_RED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_RED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_DARK_RED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_RED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_GREEN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_GREEN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_GREEN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_GREEN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_DARK_GREEN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_GREEN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_DARK_GREEN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_GREEN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_YELLOW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_YELLOW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_YELLOW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_YELLOW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_DARK_YELLOW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_YELLOW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_DARK_YELLOW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_YELLOW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_BLUE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_BLUE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_BLUE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_BLUE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_DARK_BLUE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_BLUE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_DARK_BLUE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_BLUE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_MAGENTA() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_MAGENTA", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_MAGENTA(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_MAGENTA", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_DARK_MAGENTA() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_MAGENTA", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_DARK_MAGENTA(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_MAGENTA", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_CYAN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_CYAN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_CYAN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_CYAN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_DARK_CYAN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_CYAN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_DARK_CYAN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_CYAN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_GRAY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_GRAY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_GRAY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_GRAY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_DARK_GRAY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_GRAY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_DARK_GRAY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_DARK_GRAY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WIDGET_DARK_SHADOW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_DARK_SHADOW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WIDGET_DARK_SHADOW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_DARK_SHADOW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WIDGET_NORMAL_SHADOW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_NORMAL_SHADOW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WIDGET_NORMAL_SHADOW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_NORMAL_SHADOW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WIDGET_LIGHT_SHADOW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_LIGHT_SHADOW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WIDGET_LIGHT_SHADOW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_LIGHT_SHADOW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WIDGET_HIGHLIGHT_SHADOW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_HIGHLIGHT_SHADOW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WIDGET_HIGHLIGHT_SHADOW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_HIGHLIGHT_SHADOW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WIDGET_FOREGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_FOREGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WIDGET_FOREGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_FOREGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WIDGET_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WIDGET_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_WIDGET_BORDER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_BORDER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_WIDGET_BORDER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_WIDGET_BORDER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_LIST_FOREGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_FOREGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_LIST_FOREGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_FOREGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_LIST_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_LIST_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_LIST_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_LIST_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_LIST_SELECTION_TEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_SELECTION_TEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_LIST_SELECTION_TEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_LIST_SELECTION_TEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_INFO_FOREGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_INFO_FOREGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_INFO_FOREGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_INFO_FOREGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_INFO_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_INFO_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_INFO_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_INFO_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_TITLE_FOREGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_FOREGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_TITLE_FOREGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_FOREGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_TITLE_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_TITLE_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_TITLE_BACKGROUND_GRADIENT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_BACKGROUND_GRADIENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_TITLE_BACKGROUND_GRADIENT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_BACKGROUND_GRADIENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_TITLE_INACTIVE_FOREGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_INACTIVE_FOREGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_TITLE_INACTIVE_FOREGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_INACTIVE_FOREGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_TITLE_INACTIVE_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_INACTIVE_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_TITLE_INACTIVE_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_INACTIVE_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) COLOR_TITLE_INACTIVE_BACKGROUND_GRADIENT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_INACTIVE_BACKGROUND_GRADIENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCOLOR_TITLE_INACTIVE_BACKGROUND_GRADIENT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "COLOR_TITLE_INACTIVE_BACKGROUND_GRADIENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DRAW_TRANSPARENT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DRAW_TRANSPARENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDRAW_TRANSPARENT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DRAW_TRANSPARENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DRAW_DELIMITER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DRAW_DELIMITER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDRAW_DELIMITER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DRAW_DELIMITER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DRAW_TAB() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DRAW_TAB", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDRAW_TAB(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DRAW_TAB", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DRAW_MNEMONIC() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DRAW_MNEMONIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDRAW_MNEMONIC(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DRAW_MNEMONIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DELIMITER_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DELIMITER_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDELIMITER_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DELIMITER_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LAST_LINE_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LAST_LINE_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLAST_LINE_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LAST_LINE_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_UNSPECIFIED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_UNSPECIFIED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_UNSPECIFIED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_UNSPECIFIED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_NO_HANDLES() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_NO_HANDLES", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_NO_HANDLES(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_NO_HANDLES", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_NO_MORE_CALLBACKS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_NO_MORE_CALLBACKS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_NO_MORE_CALLBACKS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_NO_MORE_CALLBACKS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_NULL_ARGUMENT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_NULL_ARGUMENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_NULL_ARGUMENT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_NULL_ARGUMENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_INVALID_ARGUMENT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_ARGUMENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_INVALID_ARGUMENT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_ARGUMENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_INVALID_RANGE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_RANGE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_INVALID_RANGE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_RANGE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_BE_ZERO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_BE_ZERO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_BE_ZERO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_BE_ZERO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_GET_ITEM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_ITEM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_GET_ITEM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_ITEM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_GET_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_GET_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_INVERT_MATRIX() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_INVERT_MATRIX", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_INVERT_MATRIX(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_INVERT_MATRIX", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_GET_ITEM_HEIGHT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_ITEM_HEIGHT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_GET_ITEM_HEIGHT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_ITEM_HEIGHT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_GET_TEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_TEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_GET_TEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_TEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_SET_TEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_TEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_SET_TEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_TEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_ITEM_NOT_ADDED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_ITEM_NOT_ADDED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_ITEM_NOT_ADDED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_ITEM_NOT_ADDED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_ITEM_NOT_REMOVED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_ITEM_NOT_REMOVED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_ITEM_NOT_REMOVED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_ITEM_NOT_REMOVED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_NO_GRAPHICS_LIBRARY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_NO_GRAPHICS_LIBRARY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_NO_GRAPHICS_LIBRARY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_NO_GRAPHICS_LIBRARY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_NOT_IMPLEMENTED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_NOT_IMPLEMENTED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_NOT_IMPLEMENTED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_NOT_IMPLEMENTED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_MENU_NOT_DROP_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_MENU_NOT_DROP_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_MENU_NOT_DROP_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_MENU_NOT_DROP_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_THREAD_INVALID_ACCESS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_THREAD_INVALID_ACCESS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_THREAD_INVALID_ACCESS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_THREAD_INVALID_ACCESS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_WIDGET_DISPOSED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_WIDGET_DISPOSED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_WIDGET_DISPOSED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_WIDGET_DISPOSED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_MENUITEM_NOT_CASCADE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_MENUITEM_NOT_CASCADE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_MENUITEM_NOT_CASCADE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_MENUITEM_NOT_CASCADE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_SET_SELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_SELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_SET_SELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_SELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_SET_MENU() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_MENU", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_SET_MENU(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_MENU", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_SET_ENABLED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_ENABLED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_SET_ENABLED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_SET_ENABLED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_GET_ENABLED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_ENABLED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_GET_ENABLED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_ENABLED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_INVALID_PARENT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_PARENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_INVALID_PARENT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_PARENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_MENU_NOT_BAR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_MENU_NOT_BAR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_MENU_NOT_BAR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_MENU_NOT_BAR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_CANNOT_GET_COUNT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_COUNT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_CANNOT_GET_COUNT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_CANNOT_GET_COUNT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_MENU_NOT_POP_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_MENU_NOT_POP_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_MENU_NOT_POP_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_MENU_NOT_POP_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_UNSUPPORTED_DEPTH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_UNSUPPORTED_DEPTH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_UNSUPPORTED_DEPTH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_UNSUPPORTED_DEPTH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_IO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_IO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_IO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_IO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_INVALID_IMAGE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_IMAGE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_INVALID_IMAGE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_IMAGE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_UNSUPPORTED_FORMAT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_UNSUPPORTED_FORMAT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_UNSUPPORTED_FORMAT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_UNSUPPORTED_FORMAT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_INVALID_SUBCLASS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_SUBCLASS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_INVALID_SUBCLASS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_SUBCLASS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_GRAPHIC_DISPOSED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_GRAPHIC_DISPOSED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_GRAPHIC_DISPOSED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_GRAPHIC_DISPOSED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_DEVICE_DISPOSED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_DEVICE_DISPOSED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_DEVICE_DISPOSED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_DEVICE_DISPOSED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_FAILED_EXEC() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_FAILED_EXEC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_FAILED_EXEC(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_FAILED_EXEC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_FAILED_LOAD_LIBRARY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_FAILED_LOAD_LIBRARY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_FAILED_LOAD_LIBRARY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_FAILED_LOAD_LIBRARY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_INVALID_FONT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_FONT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_INVALID_FONT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_FONT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_FUNCTION_DISPOSED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_FUNCTION_DISPOSED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_FUNCTION_DISPOSED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_FUNCTION_DISPOSED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_FAILED_EVALUATE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_FAILED_EVALUATE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_FAILED_EVALUATE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_FAILED_EVALUATE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR_INVALID_RETURN_VALUE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_RETURN_VALUE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR_INVALID_RETURN_VALUE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR_INVALID_RETURN_VALUE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BITMAP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BITMAP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBITMAP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BITMAP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ICON() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ICON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldICON(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ICON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_COPY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_COPY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_COPY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_COPY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_DISABLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_DISABLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_DISABLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_DISABLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_GRAY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_GRAY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_GRAY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_GRAY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ERROR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ERROR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldERROR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ERROR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PAUSED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PAUSED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPAUSED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PAUSED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) NORMAL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "NORMAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldNORMAL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "NORMAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) BOLD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "BOLD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldBOLD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "BOLD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ITALIC() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ITALIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldITALIC(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ITALIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_ARROW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_ARROW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_ARROW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_ARROW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_WAIT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_WAIT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_WAIT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_WAIT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_CROSS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_CROSS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_CROSS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_CROSS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_APPSTARTING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_APPSTARTING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_APPSTARTING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_APPSTARTING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_HELP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_HELP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_HELP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_HELP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZEALL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEALL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZEALL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEALL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZENESW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENESW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZENESW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENESW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZENS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZENS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZENWSE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENWSE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZENWSE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENWSE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZEWE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEWE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZEWE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEWE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZEN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZEN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZES() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZES", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZES(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZES", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZEE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZEE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZEW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZEW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZEW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZENE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZENE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZESE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZESE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZESE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZESE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZESW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZESW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZESW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZESW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_SIZENW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_SIZENW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_SIZENW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_UPARROW() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_UPARROW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_UPARROW(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_UPARROW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_IBEAM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_IBEAM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_IBEAM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_IBEAM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_NO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_NO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_NO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_NO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CURSOR_HAND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CURSOR_HAND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCURSOR_HAND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CURSOR_HAND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CAP_FLAT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CAP_FLAT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCAP_FLAT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CAP_FLAT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CAP_ROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CAP_ROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCAP_ROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CAP_ROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) CAP_SQUARE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "CAP_SQUARE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldCAP_SQUARE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "CAP_SQUARE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) JOIN_MITER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "JOIN_MITER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldJOIN_MITER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "JOIN_MITER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) JOIN_ROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "JOIN_ROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldJOIN_ROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "JOIN_ROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) JOIN_BEVEL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "JOIN_BEVEL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldJOIN_BEVEL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "JOIN_BEVEL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LINE_SOLID() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LINE_SOLID", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLINE_SOLID(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LINE_SOLID", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LINE_DASH() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LINE_DASH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLINE_DASH(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LINE_DASH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LINE_DOT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LINE_DOT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLINE_DOT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LINE_DOT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LINE_DASHDOT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LINE_DASHDOT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLINE_DASHDOT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LINE_DASHDOT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LINE_DASHDOTDOT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LINE_DASHDOTDOT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLINE_DASHDOTDOT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LINE_DASHDOTDOT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) LINE_CUSTOM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "LINE_CUSTOM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldLINE_CUSTOM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "LINE_CUSTOM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PATH_MOVE_TO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PATH_MOVE_TO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPATH_MOVE_TO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PATH_MOVE_TO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PATH_LINE_TO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PATH_LINE_TO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPATH_LINE_TO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PATH_LINE_TO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PATH_QUAD_TO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PATH_QUAD_TO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPATH_QUAD_TO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PATH_QUAD_TO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PATH_CUBIC_TO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PATH_CUBIC_TO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPATH_CUBIC_TO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PATH_CUBIC_TO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) PATH_CLOSE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "PATH_CLOSE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldPATH_CLOSE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "PATH_CLOSE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FILL_EVEN_ODD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FILL_EVEN_ODD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFILL_EVEN_ODD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FILL_EVEN_ODD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) FILL_WINDING() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "FILL_WINDING", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldFILL_WINDING(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "FILL_WINDING", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_UNDEFINED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_UNDEFINED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_UNDEFINED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_UNDEFINED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_BMP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_BMP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_BMP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_BMP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_BMP_RLE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_BMP_RLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_BMP_RLE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_BMP_RLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_GIF() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_GIF", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_GIF(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_GIF", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_ICO() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_ICO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_ICO(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_ICO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_JPEG() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_JPEG", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_JPEG(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_JPEG", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_PNG() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_PNG", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_PNG(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_PNG", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_TIFF() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_TIFF", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_TIFF(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_TIFF", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) IMAGE_OS2_BMP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "IMAGE_OS2_BMP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldIMAGE_OS2_BMP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "IMAGE_OS2_BMP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DM_UNSPECIFIED() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DM_UNSPECIFIED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDM_UNSPECIFIED(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DM_UNSPECIFIED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DM_FILL_NONE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DM_FILL_NONE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDM_FILL_NONE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DM_FILL_NONE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DM_FILL_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DM_FILL_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDM_FILL_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DM_FILL_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) DM_FILL_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "DM_FILL_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldDM_FILL_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "DM_FILL_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRANSPARENCY_NONE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_NONE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRANSPARENCY_NONE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_NONE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRANSPARENCY_ALPHA() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_ALPHA", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRANSPARENCY_ALPHA(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_ALPHA", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRANSPARENCY_MASK() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_MASK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRANSPARENCY_MASK(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_MASK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) TRANSPARENCY_PIXEL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_PIXEL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldTRANSPARENCY_PIXEL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "TRANSPARENCY_PIXEL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOVEMENT_CHAR() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOVEMENT_CHAR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOVEMENT_CHAR(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOVEMENT_CHAR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOVEMENT_CLUSTER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOVEMENT_CLUSTER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOVEMENT_CLUSTER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOVEMENT_CLUSTER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOVEMENT_WORD() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOVEMENT_WORD", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOVEMENT_WORD(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOVEMENT_WORD", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOVEMENT_WORD_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOVEMENT_WORD_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOVEMENT_WORD_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOVEMENT_WORD_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) MOVEMENT_WORD_START() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "MOVEMENT_WORD_START", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldMOVEMENT_WORD_START(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "MOVEMENT_WORD_START", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ALL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ALL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldALL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ALL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ID_ABOUT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ID_ABOUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldID_ABOUT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ID_ABOUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ID_PREFERENCES() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ID_PREFERENCES", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldID_PREFERENCES(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ID_PREFERENCES", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ID_HIDE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ID_HIDE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldID_HIDE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ID_HIDE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ID_HIDE_OTHERS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ID_HIDE_OTHERS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldID_HIDE_OTHERS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ID_HIDE_OTHERS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ID_SHOW_ALL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ID_SHOW_ALL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldID_SHOW_ALL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ID_SHOW_ALL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) ID_QUIT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "ID_QUIT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldID_QUIT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "ID_QUIT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *SWT) SKIN_CLASS() string {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SKIN_CLASS", "java/lang/String")
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

func (jbobject *SWT) SetFieldSKIN_CLASS(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SKIN_CLASS", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *SWT) SKIN_ID() string {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SKIN_ID", "java/lang/String")
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

func (jbobject *SWT) SetFieldSKIN_ID(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SKIN_ID", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *SWT) SCROLLBAR_OVERLAY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/SWT", "SCROLLBAR_OVERLAY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *SWT) SetFieldSCROLLBAR_OVERLAY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/SWT", "SCROLLBAR_OVERLAY", val)
	if err != nil {
		panic(err)
	}

}


