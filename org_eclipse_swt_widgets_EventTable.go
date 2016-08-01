package swt

import "github.com/timob/javabind"

type WidgetsEventTableInterface interface {
	JavaLangObjectInterface

	// public org.eclipse.swt.widgets.Listener[] getListeners(int)
	GetListeners(a int) []*WidgetsListener

	// public void hook(int, org.eclipse.swt.widgets.Listener)
	Hook(a int, b WidgetsListenerInterface) 

	// public boolean hooks(int)
	Hooks(a int) bool

	// public void sendEvent(org.eclipse.swt.widgets.Event)
	SendEvent(a WidgetsEventInterface) 

	// public int size()
	Size() int

	// public void unhook(int, org.eclipse.swt.widgets.Listener)
	Unhook(a int, b WidgetsListenerInterface) 
}

type WidgetsEventTable struct {
	JavaLangObject
}

// public org.eclipse.swt.widgets.Listener[] getListeners(int)
func (jbobject *WidgetsEventTable) GetListeners(a int) []*WidgetsListener {
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

// public void hook(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsEventTable) Hook(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "hook", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public boolean hooks(int)
func (jbobject *WidgetsEventTable) Hooks(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hooks", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void sendEvent(org.eclipse.swt.widgets.Event)
func (jbobject *WidgetsEventTable) SendEvent(a WidgetsEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "sendEvent", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public int size()
func (jbobject *WidgetsEventTable) Size() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void unhook(int, org.eclipse.swt.widgets.Listener)
func (jbobject *WidgetsEventTable) Unhook(a int, b WidgetsListenerInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unhook", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/widgets/Listener"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}


