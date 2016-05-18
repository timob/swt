package swt

import "github.com/timob/javabind"

type ProgramProgramInterface interface {
	JavaLangObjectInterface

	// public static org.eclipse.swt.program.Program findProgram(java.lang.String)
	FindProgram(a string) *ProgramProgram

	// public static java.lang.String[] getExtensions()
	GetExtensions() []string

	// public static org.eclipse.swt.program.Program[] getPrograms()
	GetPrograms() []*ProgramProgram

	// public static boolean launch(java.lang.String)
	Launch(a string) bool

	// public static boolean launch(java.lang.String, java.lang.String)
	Launch2(a string, b string) bool

	// public boolean execute(java.lang.String)
	Execute(a string) bool

	// public org.eclipse.swt.graphics.ImageData getImageData()
	GetImageData() *GraphicsImageData

	// public java.lang.String getName()
	GetName() string
}

type ProgramProgram struct {
	JavaLangObject
}

// public static org.eclipse.swt.program.Program findProgram(java.lang.String)
func (jbobject *ProgramProgram) FindProgram(a string) *ProgramProgram {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/program/Program", "findProgram", "org/eclipse/swt/program/Program", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ProgramProgram{}
	unique_x.Callable = dst
	return unique_x
}

// public static java.lang.String[] getExtensions()
func (jbobject *ProgramProgram) GetExtensions() []string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/program/Program", "getExtensions", javabind.ObjectArrayType("java/lang/String"))
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

// public static org.eclipse.swt.program.Program[] getPrograms()
func (jbobject *ProgramProgram) GetPrograms() []*ProgramProgram {
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/program/Program", "getPrograms", javabind.ObjectArrayType("org/eclipse/swt/program/Program"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/program/Program")
	dst := new([]*ProgramProgram)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static boolean launch(java.lang.String)
func (jbobject *ProgramProgram) Launch(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/program/Program", "launch", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public static boolean launch(java.lang.String, java.lang.String)
func (jbobject *ProgramProgram) Launch2(a string, b string) bool {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/program/Program", "launch", javabind.Boolean, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(bool)
}

// public boolean equals(java.lang.Object)
func (jbobject *ProgramProgram) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean execute(java.lang.String)
func (jbobject *ProgramProgram) Execute(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "execute", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public org.eclipse.swt.graphics.ImageData getImageData()
func (jbobject *ProgramProgram) GetImageData() *GraphicsImageData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageData", "org/eclipse/swt/graphics/ImageData")
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
	unique_x := &GraphicsImageData{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getName()
func (jbobject *ProgramProgram) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public int hashCode()
func (jbobject *ProgramProgram) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *ProgramProgram) ToString() string {
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


