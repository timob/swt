package swt

import "github.com/timob/javabind"

type CustomStyledTextContentInterface interface {

	// public abstract void addTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
	AddTextChangeListener(a CustomTextChangeListenerInterface) 

	// public abstract int getCharCount()
	GetCharCount() int

	// public abstract java.lang.String getLine(int)
	GetLine(a int) string

	// public abstract int getLineAtOffset(int)
	GetLineAtOffset(a int) int

	// public abstract int getLineCount()
	GetLineCount() int

	// public abstract java.lang.String getLineDelimiter()
	GetLineDelimiter() string

	// public abstract int getOffsetAtLine(int)
	GetOffsetAtLine(a int) int

	// public abstract java.lang.String getTextRange(int, int)
	GetTextRange(a int, b int) string

	// public abstract void removeTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
	RemoveTextChangeListener(a CustomTextChangeListenerInterface) 

	// public abstract void replaceTextRange(int, int, java.lang.String)
	ReplaceTextRange(a int, b int, c string) 

	// public abstract void setText(java.lang.String)
	SetText(a string) 
}

type CustomStyledTextContent struct {
	JavaLangObject
}

// public abstract void addTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
func (jbobject *CustomStyledTextContent) AddTextChangeListener(a CustomTextChangeListenerInterface)  {
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

// public abstract int getCharCount()
func (jbobject *CustomStyledTextContent) GetCharCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCharCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.String getLine(int)
func (jbobject *CustomStyledTextContent) GetLine(a int) string {
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

// public abstract int getLineAtOffset(int)
func (jbobject *CustomStyledTextContent) GetLineAtOffset(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineAtOffset", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int getLineCount()
func (jbobject *CustomStyledTextContent) GetLineCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.String getLineDelimiter()
func (jbobject *CustomStyledTextContent) GetLineDelimiter() string {
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

// public abstract int getOffsetAtLine(int)
func (jbobject *CustomStyledTextContent) GetOffsetAtLine(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOffsetAtLine", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.String getTextRange(int, int)
func (jbobject *CustomStyledTextContent) GetTextRange(a int, b int) string {
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

// public abstract void removeTextChangeListener(org.eclipse.swt.custom.TextChangeListener)
func (jbobject *CustomStyledTextContent) RemoveTextChangeListener(a CustomTextChangeListenerInterface)  {
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

// public abstract void replaceTextRange(int, int, java.lang.String)
func (jbobject *CustomStyledTextContent) ReplaceTextRange(a int, b int, c string)  {
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

// public abstract void setText(java.lang.String)
func (jbobject *CustomStyledTextContent) SetText(a string)  {
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


