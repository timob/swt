package swt

import "github.com/timob/javabind"

type CustomDefaultContentInterface interface {
	JavaLangObjectInterface

	// public void addTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
	AddTextChangeListener(a CustomTextChangeListenerInterface) 

	// public int getCharCount()
	GetCharCount() int

	// public java.lang.String getLine(int)
	GetLine(a int) string

	// public java.lang.String getLineDelimiter()
	GetLineDelimiter() string

	// public int getLineCount()
	GetLineCount() int

	// public int getLineAtOffset(int)
	GetLineAtOffset(a int) int

	// public int getOffsetAtLine(int)
	GetOffsetAtLine(a int) int

	// public java.lang.String getTextRange(int, int)
	GetTextRange(a int, b int) string

	// public void removeTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
	RemoveTextChangeListener(a CustomTextChangeListenerInterface) 

	// public void replaceTextRange(int, int, java.lang.String)
	ReplaceTextRange(a int, b int, c string) 

	// public void setText(java.lang.String)
	SetText(a string) 
}

type CustomDefaultContent struct {
	JavaLangObject
}

// public void addTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
func (jbobject *CustomDefaultContent) AddTextChangeListener(a CustomTextChangeListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addTextChangeListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TextChangeListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public int getCharCount()
func (jbobject *CustomDefaultContent) GetCharCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCharCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getLine(int)
func (jbobject *CustomDefaultContent) GetLine(a int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLine", "java/lang/String", a)
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

// public java.lang.String getLineDelimiter()
func (jbobject *CustomDefaultContent) GetLineDelimiter() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineDelimiter", "java/lang/String")
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

// public int getLineCount()
func (jbobject *CustomDefaultContent) GetLineCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLineAtOffset(int)
func (jbobject *CustomDefaultContent) GetLineAtOffset(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineAtOffset", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getOffsetAtLine(int)
func (jbobject *CustomDefaultContent) GetOffsetAtLine(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOffsetAtLine", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getTextRange(int, int)
func (jbobject *CustomDefaultContent) GetTextRange(a int, b int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTextRange", "java/lang/String", a, b)
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

// public void removeTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
func (jbobject *CustomDefaultContent) RemoveTextChangeListener(a CustomTextChangeListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeTextChangeListener", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TextChangeListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void replaceTextRange(int, int, java.lang.String)
func (jbobject *CustomDefaultContent) ReplaceTextRange(a int, b int, c string)  {
	conv_c := javabind.NewGoToJavaString()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "replaceTextRange", javabind.Void, a, b, conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public void setText(java.lang.String)
func (jbobject *CustomDefaultContent) SetText(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setText", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


