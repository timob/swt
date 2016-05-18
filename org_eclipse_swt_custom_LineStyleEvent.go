package swt

import "github.com/timob/javabind"

type CustomLineStyleEventInterface interface {
	EventsTypedEventInterface
}

type CustomLineStyleEvent struct {
	EventsTypedEvent
}

func (jbobject *CustomLineStyleEvent) LineOffset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomLineStyleEvent) SetFieldLineOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "lineOffset", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineStyleEvent) LineText() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineText", "java/lang/String")
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

func (jbobject *CustomLineStyleEvent) SetFieldLineText(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "lineText", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomLineStyleEvent) Ranges() []int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "ranges", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

func (jbobject *CustomLineStyleEvent) SetFieldRanges(val []int) {
	err := jbobject.SetField(javabind.GetEnv(), "ranges", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineStyleEvent) Styles() []*CustomStyleRange {
	jret, err := jbobject.GetField(javabind.GetEnv(), "styles", javabind.ObjectArrayType("org/eclipse/swt/custom/StyleRange"))
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

func (jbobject *CustomLineStyleEvent) SetFieldStyles(val []*CustomStyleRange) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/custom/StyleRange")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "styles", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomLineStyleEvent) Alignment() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "alignment", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomLineStyleEvent) SetFieldAlignment(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "alignment", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineStyleEvent) Indent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "indent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomLineStyleEvent) SetFieldIndent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "indent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineStyleEvent) WrapIndent() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "wrapIndent", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomLineStyleEvent) SetFieldWrapIndent(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "wrapIndent", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineStyleEvent) Justify() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "justify", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomLineStyleEvent) SetFieldJustify(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "justify", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineStyleEvent) Bullet() *CustomBullet {
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

func (jbobject *CustomLineStyleEvent) SetFieldBullet(val CustomBulletInterface) {
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

func (jbobject *CustomLineStyleEvent) BulletIndex() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "bulletIndex", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomLineStyleEvent) SetFieldBulletIndex(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "bulletIndex", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomLineStyleEvent) TabStops() []int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "tabStops", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

func (jbobject *CustomLineStyleEvent) SetFieldTabStops(val []int) {
	err := jbobject.SetField(javabind.GetEnv(), "tabStops", val)
	if err != nil {
		panic(err)
	}

}


