package swt

import "github.com/timob/javabind"

type CustomStyledTextListenerInterface interface {
	WidgetsTypedListenerInterface
}

type CustomStyledTextListener struct {
	WidgetsTypedListener
}

// public void handleEvent(org.eclipse.swt.widgets.Event)
func (jbobject *CustomStyledTextListener) HandleEvent(a WidgetsEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "handleEvent", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/widgets/Event"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


