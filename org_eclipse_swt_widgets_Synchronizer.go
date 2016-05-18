package swt

import "github.com/timob/javabind"

type WidgetsSynchronizerInterface interface {
	JavaLangObjectInterface
}

type WidgetsSynchronizer struct {
	JavaLangObject
}

// public org.eclipse.swt.widgets.Synchronizer(org.eclipse.swt.widgets.Display)
func NewWidgetsSynchronizer(a WidgetsDisplayInterface) (*WidgetsSynchronizer) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/Synchronizer", conv_a.Value().Cast("org/eclipse/swt/widgets/Display"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &WidgetsSynchronizer{}
	x.Callable = &javabind.Callable{obj}
	return x
}


