package swt

import "github.com/timob/javabind"

type CustomBulletInterface interface {
	JavaLangObjectInterface
}

type CustomBullet struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.Bullet(org.eclipse.swt.custom.StyleRange)
func NewCustomBullet(a CustomStyleRangeInterface) (*CustomBullet) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/Bullet", conv_a.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomBullet{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.custom.Bullet(int, org.eclipse.swt.custom.StyleRange)
func NewCustomBullet2(a int, b CustomStyleRangeInterface) (*CustomBullet) {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/Bullet", a, conv_b.Value().Cast("org/eclipse/swt/custom/StyleRange"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &CustomBullet{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int hashCode()
func (jbobject *CustomBullet) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomBullet) Type() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "type", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomBullet) SetFieldType(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "type", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomBullet) Style() *CustomStyleRange {
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

func (jbobject *CustomBullet) SetFieldStyle(val CustomStyleRangeInterface) {
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

func (jbobject *CustomBullet) Text() string {
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

func (jbobject *CustomBullet) SetFieldText(val string) {
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


