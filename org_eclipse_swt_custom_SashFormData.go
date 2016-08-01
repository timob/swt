package swt

import "github.com/timob/javabind"

type CustomSashFormDataInterface interface {
	JavaLangObjectInterface
}

type CustomSashFormData struct {
	JavaLangObject
}

// public java.lang.String toString()
func (jbobject *CustomSashFormData) ToString() string {
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


