package swt

import "github.com/timob/javabind"

type CustomPaintObjectEventInterface interface {
	EventsTypedEventInterface
}

type CustomPaintObjectEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.custom.PaintObjectEvent(org.eclipse.swt.custom.StyledTextEvent)
func NewCustomPaintObjectEvent(a CustomStyledTextEventInterface) (*CustomPaintObjectEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/PaintObjectEvent", conv_a.Value().Cast("org/eclipse/swt/custom/StyledTextEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomPaintObjectEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *CustomPaintObjectEvent) Gc() *GraphicsGC {
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

func (jbobject *CustomPaintObjectEvent) SetFieldGc(val GraphicsGCInterface) {
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

func (jbobject *CustomPaintObjectEvent) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomPaintObjectEvent) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomPaintObjectEvent) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomPaintObjectEvent) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomPaintObjectEvent) Ascent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "ascent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomPaintObjectEvent) SetFieldAscent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "ascent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomPaintObjectEvent) Descent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "descent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomPaintObjectEvent) SetFieldDescent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "descent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomPaintObjectEvent) Style() *CustomStyleRange {
	jret, err := jbobject.GetField(javabind.GetEnv(), "style", "org/eclipse/swt/custom/StyleRange")
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

func (jbobject *CustomPaintObjectEvent) SetFieldStyle(val CustomStyleRangeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "style", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomPaintObjectEvent) Bullet() *CustomBullet {
	jret, err := jbobject.GetField(javabind.GetEnv(), "bullet", "org/eclipse/swt/custom/Bullet")
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

func (jbobject *CustomPaintObjectEvent) SetFieldBullet(val CustomBulletInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "bullet", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomPaintObjectEvent) BulletIndex() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "bulletIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomPaintObjectEvent) SetFieldBulletIndex(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "bulletIndex", val)
	if err != nil {
		panic(err)
	}

}


