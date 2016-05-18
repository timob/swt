package swt

import "github.com/timob/javabind"

type WidgetsWidgetInterface interface {
	JavaLangObjectInterface

	// public void addListener(int, org.eclipse.swt.widgets.Listener)
	AddListener(a int, b WidgetsListenerInterface) 

	// public void addDisposeListener(org.eclipse.swt.events.DisposeListener)
	AddDisposeListener(a EventsDisposeListenerInterface) 

	// public void dispose()
	Dispose() 

	// public java.lang.Object getData()
	GetData() *JavaLangObject

	// public java.lang.Object getData(java.lang.String)
	GetData2(a string) *JavaLangObject

	// public org.eclipse.swt.widgets.Display getDisplay()
	GetDisplay() *WidgetsDisplay

	// public org.eclipse.swt.widgets.Listener[] getListeners(int)
	GetListeners(a int) []*WidgetsListener

	// public int getStyle()
	GetStyle() int

	// public boolean isDisposed()
	IsDisposed() bool

	// public boolean isListening(int)
	IsListening(a int) bool

	// public void notifyListeners(int, org.eclipse.swt.widgets.Event)
	NotifyListeners(a int, b WidgetsEventInterface) 

	// public void removeListener(int, org.eclipse.swt.widgets.Listener)
	RemoveListener(a int, b WidgetsListenerInterface) 

	// public void reskin(int)
	Reskin(a int) 

	// public void removeDisposeListener(org.eclipse.swt.events.DisposeListener)
	RemoveDisposeListener(a EventsDisposeListenerInterface) 

	// public void setData(java.lang.Object)
	SetData(a interface{}) 

	// public void setData(java.lang.String, java.lang.Object)
	SetData2(a string, b interface{}) 
}

type WidgetsWidget struct {
	JavaLangObject
}

// public org.eclipse.swt.widgets.Widget(org.eclipse.swt.widgets.Widget, int)
func NewWidgetsWidget(a WidgetsWidgetInterface, b int) (*WidgetsWidget) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Widget", conv_a.Value().Cast("org/eclipse/swt/widgets/Widget"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsWidget{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void addListener(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsWidget) AddListener(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addListener", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void addDisposeListener(org.eclipse.swt.events.DisposeListener)
func (jbobject *WidgetsWidget) AddDisposeListener(a EventsDisposeListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addDisposeListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/DisposeListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void dispose()
func (jbobject *WidgetsWidget) Dispose()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "dispose", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public java.lang.Object getData()
func (jbobject *WidgetsWidget) GetData() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getData", "java/lang/Object")
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

// public java.lang.Object getData(java.lang.String)
func (jbobject *WidgetsWidget) GetData2(a string) *JavaLangObject {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getData", "java/lang/Object", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.Display getDisplay()
func (jbobject *WidgetsWidget) GetDisplay() *WidgetsDisplay {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplay", "org/eclipse/swt/widgets/Display")
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

// public org.eclipse.swt.widgets.Listener[] getListeners(int)
func (jbobject *WidgetsWidget) GetListeners(a int) []*WidgetsListener {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getListeners", javabind.ObjectArrayType("org/eclipse/swt/widgets/Listener"), a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/Listener")
	dst := new([]*WidgetsListener)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getStyle()
func (jbobject *WidgetsWidget) GetStyle() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isDisposed()
func (jbobject *WidgetsWidget) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isListening(int)
func (jbobject *WidgetsWidget) IsListening(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isListening", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void notifyListeners(int, org.eclipse.swt.widgets.Event)
func (jbobject *WidgetsWidget) NotifyListeners(a int, b WidgetsEventInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "notifyListeners", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void removeListener(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsWidget) RemoveListener(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeListener", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public void reskin(int)
func (jbobject *WidgetsWidget) Reskin(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reskin", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void removeDisposeListener(org.eclipse.swt.events.DisposeListener)
func (jbobject *WidgetsWidget) RemoveDisposeListener(a EventsDisposeListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeDisposeListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/DisposeListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setData(java.lang.Object)
func (jbobject *WidgetsWidget) SetData(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setData", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setData(java.lang.String, java.lang.Object)
func (jbobject *WidgetsWidget) SetData2(a string, b interface{})  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setData", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public java.lang.String toString()
func (jbobject *WidgetsWidget) ToString() string {
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

func (jbobject *WidgetsWidget) Handle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *WidgetsWidget) SetFieldHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


