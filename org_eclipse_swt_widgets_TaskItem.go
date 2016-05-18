package swt

import "github.com/timob/javabind"

type WidgetsTaskItemInterface interface {
	WidgetsItemInterface

	// public org.eclipse.swt.widgets.Menu getMenu()
	GetMenu() *WidgetsMenu

	// public org.eclipse.swt.graphics.Image getOverlayImage()
	GetOverlayImage() *GraphicsImage

	// public java.lang.String getOverlayText()
	GetOverlayText() string

	// public org.eclipse.swt.widgets.TaskBar getParent()
	GetParent() *WidgetsTaskBar

	// public int getProgress()
	GetProgress() int

	// public int getProgressState()
	GetProgressState() int

	// public void setMenu(org.eclipse.swt.widgets.Menu)
	SetMenu(a WidgetsMenuInterface) 

	// public void setOverlayImage(org.eclipse.swt.graphics.Image)
	SetOverlayImage(a GraphicsImageInterface) 

	// public void setOverlayText(java.lang.String)
	SetOverlayText(a string) 

	// public void setProgressState(int)
	SetProgressState(a int) 

	// public void setProgress(int)
	SetProgress(a int) 
}

type WidgetsTaskItem struct {
	WidgetsItem
}

// public org.eclipse.swt.widgets.Menu getMenu()
func (jbobject *WidgetsTaskItem) GetMenu() *WidgetsMenu {
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

// public org.eclipse.swt.graphics.Image getOverlayImage()
func (jbobject *WidgetsTaskItem) GetOverlayImage() *GraphicsImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOverlayImage", "org/eclipse/swt/graphics/Image")
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

// public java.lang.String getOverlayText()
func (jbobject *WidgetsTaskItem) GetOverlayText() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOverlayText", "java/lang/String")
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

// public org.eclipse.swt.widgets.TaskBar getParent()
func (jbobject *WidgetsTaskItem) GetParent() *WidgetsTaskBar {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "org/eclipse/swt/widgets/TaskBar")
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
	unique_x := &WidgetsTaskBar{}
	unique_x.Callable = dst
	return unique_x
}

// public int getProgress()
func (jbobject *WidgetsTaskItem) GetProgress() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProgress", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getProgressState()
func (jbobject *WidgetsTaskItem) GetProgressState() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProgressState", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setMenu(org.eclipse.swt.widgets.Menu)
func (jbobject *WidgetsTaskItem) SetMenu(a WidgetsMenuInterface)  {
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

// public void setOverlayImage(org.eclipse.swt.graphics.Image)
func (jbobject *WidgetsTaskItem) SetOverlayImage(a GraphicsImageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOverlayImage", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Image"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setOverlayText(java.lang.String)
func (jbobject *WidgetsTaskItem) SetOverlayText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOverlayText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setProgressState(int)
func (jbobject *WidgetsTaskItem) SetProgressState(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProgressState", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setProgress(int)
func (jbobject *WidgetsTaskItem) SetProgress(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProgress", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


