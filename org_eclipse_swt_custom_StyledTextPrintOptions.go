package swt

import "github.com/timob/javabind"

type CustomStyledTextPrintOptionsInterface interface {
	JavaLangObjectInterface
}

type CustomStyledTextPrintOptions struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.StyledTextPrintOptions()
func NewCustomStyledTextPrintOptions() (*CustomStyledTextPrintOptions) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/StyledTextPrintOptions")
	if err != nil {
		panic(err)
	}
	x := &CustomStyledTextPrintOptions{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *CustomStyledTextPrintOptions) PAGE_TAG() string {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/StyledTextPrintOptions", "PAGE_TAG", "java/lang/String")
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

func (jbobject *CustomStyledTextPrintOptions) SetFieldPAGE_TAG(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/StyledTextPrintOptions", "PAGE_TAG", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomStyledTextPrintOptions) SEPARATOR() string {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/StyledTextPrintOptions", "SEPARATOR", "java/lang/String")
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

func (jbobject *CustomStyledTextPrintOptions) SetFieldSEPARATOR(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/StyledTextPrintOptions", "SEPARATOR", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomStyledTextPrintOptions) Header() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "header", "java/lang/String")
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

func (jbobject *CustomStyledTextPrintOptions) SetFieldHeader(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "header", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomStyledTextPrintOptions) Footer() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "footer", "java/lang/String")
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

func (jbobject *CustomStyledTextPrintOptions) SetFieldFooter(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "footer", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomStyledTextPrintOptions) JobName() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "jobName", "java/lang/String")
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

func (jbobject *CustomStyledTextPrintOptions) SetFieldJobName(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "jobName", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *CustomStyledTextPrintOptions) PrintTextForeground() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "printTextForeground", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomStyledTextPrintOptions) SetFieldPrintTextForeground(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "printTextForeground", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStyledTextPrintOptions) PrintTextBackground() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "printTextBackground", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomStyledTextPrintOptions) SetFieldPrintTextBackground(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "printTextBackground", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStyledTextPrintOptions) PrintTextFontStyle() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "printTextFontStyle", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomStyledTextPrintOptions) SetFieldPrintTextFontStyle(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "printTextFontStyle", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStyledTextPrintOptions) PrintLineBackground() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "printLineBackground", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomStyledTextPrintOptions) SetFieldPrintLineBackground(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "printLineBackground", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStyledTextPrintOptions) PrintLineNumbers() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "printLineNumbers", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *CustomStyledTextPrintOptions) SetFieldPrintLineNumbers(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "printLineNumbers", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomStyledTextPrintOptions) LineLabels() []string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "lineLabels", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *CustomStyledTextPrintOptions) SetFieldLineLabels(val []string) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "lineLabels", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


