package swt

import "github.com/timob/javabind"

type WidgetsLayoutInterface interface {
	JavaLangObjectInterface
}

type WidgetsLayout struct {
	JavaLangObject
}

// public org.eclipse.swt.widgets.Layout()
func NewWidgetsLayout() (*WidgetsLayout) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Layout")
	if err != nil {
		panic(err)
	}
	x := &WidgetsLayout{}
	x.Callable = &javabind.Callable{obj}
	return x
}


