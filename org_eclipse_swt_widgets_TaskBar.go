package swt

import "github.com/timob/javabind"

type WidgetsTaskBarInterface interface {
	WidgetsWidgetInterface

	// public org.eclipse.swt.widgets.TaskItem getItem(int)
	GetItem(a int) *WidgetsTaskItem

	// public int getItemCount()
	GetItemCount() int

	// public org.eclipse.swt.widgets.TaskItem getItem(org.eclipse.swt.widgets.Shell)
	GetItem2(a WidgetsShellInterface) *WidgetsTaskItem

	// public org.eclipse.swt.widgets.TaskItem[] getItems()
	GetItems() []*WidgetsTaskItem
}

type WidgetsTaskBar struct {
	WidgetsWidget
}

// public org.eclipse.swt.widgets.TaskItem getItem(int)
func (jbobject *WidgetsTaskBar) GetItem(a int) *WidgetsTaskItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TaskItem", a)
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
	unique_x := &WidgetsTaskItem{}
	unique_x.Callable = dst
	return unique_x
}

// public int getItemCount()
func (jbobject *WidgetsTaskBar) GetItemCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItemCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.widgets.TaskItem getItem(org.eclipse.swt.widgets.Shell)
func (jbobject *WidgetsTaskBar) GetItem2(a WidgetsShellInterface) *WidgetsTaskItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItem", "org/eclipse/swt/widgets/TaskItem", conv_a.Value().Cast("org/eclipse/swt/widgets/Shell"))
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
	unique_x := &WidgetsTaskItem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.widgets.TaskItem[] getItems()
func (jbobject *WidgetsTaskBar) GetItems() []*WidgetsTaskItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getItems", javabind.ObjectArrayType("org/eclipse/swt/widgets/TaskItem"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/widgets/TaskItem")
	dst := new([]*WidgetsTaskItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


