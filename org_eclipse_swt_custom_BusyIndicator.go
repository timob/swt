package swt

import "github.com/timob/javabind"

type CustomBusyIndicatorInterface interface {
	JavaLangObjectInterface
}

type CustomBusyIndicator struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.BusyIndicator()
func NewCustomBusyIndicator() (*CustomBusyIndicator) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/BusyIndicator")
	if err != nil {
		panic(err)
	}
	x := &CustomBusyIndicator{}
	x.Callable = &javabind.Callable{obj}
	return x
}


